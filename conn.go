package redis

import (
	"bufio"
	"errors"
	"net"
	"time"
)

type client2 struct {
	// Normalized node address in use. This field is read-only.
	Addr string

	// noCopy noCopy

	// sticky AUTH(entication)
	// password atomic.Value

	// sticky database SELECT
	db int64

	// optional execution expiry
	commandTimeout time.Duration

	// network establishment expiry
	dialTimeout time.Duration

	// The connection semaphore is used as a write lock.
	connSem chan *conn

	// The buffering reader from conn is used as a read lock.
	// Command submission holds the write lock [connSem] when sending
	// to readQueue.
	readQueue chan chan<- *bufio.Reader

	// The read routine stops on receive: no more readQueue receives
	// nor network use. The idle state is not set/restored.
	readInterrupt chan struct{}
}

type conn struct {
	net.Conn       // nil when offline
	offline  error // reason for connection absence

	// The token is nil when a read routine is using it.
	idle *bufio.Reader
}

// // connectOrClosed populates the connection semaphore.
// func (c *client2) connectOrClosed() {
// 	var retryDelay time.Duration
// 	for {
// 		config := connConfig{
// 			BufferSize:     conservativeMSS,
// 			Addr:           c.Addr,
// 			DB:             atomic.LoadInt64(&c.db),
// 			CommandTimeout: c.commandTimeout,
// 			DialTimeout:    c.dialTimeout,
// 		}
// 		config.Password, _ = c.password.Load().([]byte)
// 		conn, reader, err := connect(config)
// 		if err != nil {
// 			retry := time.NewTimer(retryDelay)

// 			// remove previous connect error unless closed
// 			if retryDelay != 0 {
// 				current := <-c.connSem
// 				if current.offline == ErrClosed {
// 					c.connSem <- current // restore
// 					retry.Stop()         // cleanup
// 					return               // abandon
// 				}
// 			}
// 			// propagate current connect error
// 			c.connSem <- &redisConn{offline: fmt.Errorf("redis: offline due %w", err)}

// 			retryDelay = 2*retryDelay + time.Millisecond
// 			if retryDelay > DialDelayMax {
// 				retryDelay = DialDelayMax
// 			}
// 			<-retry.C
// 			continue
// 		}

// 		// remove previous connect error unless closed
// 		if retryDelay != 0 {
// 			current := <-c.connSem
// 			if current.offline == ErrClosed {
// 				c.connSem <- current // restore
// 				conn.Close()         // discard
// 				return               // abandon
// 			}
// 		}

// 		// release
// 		c.connSem <- &redisConn{Conn: conn, idle: reader}
// 		return
// 	}
// }

// CancelQueue signals connection loss to all pending commands.
func (c *client2) cancelQueue() {
	for n := len(c.readQueue); n > 0; n-- {
		(<-c.readQueue) <- (*bufio.Reader)(nil)
	}
}

// Submit sends a request, and deals with response ordering.
func (c *client2) submit(req *request) (*bufio.Reader, error) {
	conn := <-c.connSem

	// validate connection state
	if err := conn.offline; err != nil {
		c.connSem <- conn // restore
		return nil, err
	}

	// apply timeout if set
	var deadline time.Time
	if c.commandTimeout != 0 {
		deadline = time.Now().Add(c.commandTimeout)
		conn.SetWriteDeadline(deadline)
	}

	// send command
	if _, err := conn.Write(req.buf); err != nil {
		// write remains locked
		go func() {
			c.haltReceive(conn)
			c.cancelQueue()
			conn.Close()
			// c.connectOrClosed()
		}()
		return nil, err
	}

	reader := conn.idle
	if reader != nil {
		// Own the virtual read lock by clearing the idle state.
		conn.idle = nil
		// The receive channel is not used, as we're next in line.
		req.free()
	} else {
		// The virtual read lock is processing the queue.
		c.readQueue <- req.receive
	}

	c.connSem <- conn // release write lock

	if reader == nil {
		// await handover of virtual read lock
		reader = <-req.receive
		req.free()
		if reader == nil {
			// queue abandonment
			return nil, errors.New("errConnLost")
		}
	}

	if !deadline.IsZero() {
		conn.SetReadDeadline(deadline)
	}

	return reader, nil
}

func (c *client2) commandInteger(req *request) (int64, error) {
	r, err := c.submit(req)
	if err != nil {
		return 0, err
	}
	integer, err := decodeInteger(r)
	c.pass(r, err)
	return integer, err
}

// Pass over the virtual read lock to the following command in line.
// If there are no routines waiting for response, then go in idle mode.
func (c *client2) pass(r *bufio.Reader, err error) {
	switch err {
	case nil, errNull:
		break
	default:
		if _, ok := err.(ServerError); !ok {
			c.dropConn()
			return
		}
	}

	// The high-traffic scenario has the optimal flow.
	select {
	case next := <-c.readQueue:
		next <- r // pass read lock
		return

	default:
		break
	}

	select {
	case next := <-c.readQueue:
		next <- r // pass read lock

	// Write is locked to make the idle decision atomic,
	// as readQueue is fed while holding the write lock.
	case conn := <-c.connSem:
		select {
		case next := <-c.readQueue:
			// lost race recovery
			next <- r // pass read lock

		default:
			// set read lock to idle
			conn.idle = r
		}
		c.connSem <- conn // unlock write

	case <-c.readInterrupt:
		// halt accepted
		break // read lock discard
	}
}

func (c *client2) dropConn() {
	for {
		select {
		case <-c.readInterrupt:
			return // accept halt

		// A write (lock owner) blocks on a full queue,
		// so include discard here to prevent deadlock.
		case next := <-c.readQueue:
			// signal connection loss
			next <- (*bufio.Reader)(nil)

		case conn := <-c.connSem:
			// write locked
			if conn.offline != nil {
				if conn.offline == ErrClosed {
					// confirm by accept
					<-c.readInterrupt
				}
				c.connSem <- conn // restore
			} else {
				// write remains locked
				go func() {
					conn.Close()
					c.cancelQueue()
					// c.connectOrClosed()
				}()
			}

			return
		}
	}
}

func (c *client2) haltReceive(writeLock *conn) {
	if writeLock.offline != nil || writeLock.idle != nil {
		// read routine not running
		return
	}
	// Read routine needs the write lock to idle.

	readHandover := make(chan *bufio.Reader)
	select {
	case c.readInterrupt <- struct{}{}:
		// The read routine accepted the halt,
		// while awaiting the write lock.
		break

	case c.readQueue <- readHandover:
		select {
		case c.readInterrupt <- struct{}{}:
			// The read routine accepted the halt,
			// while awaiting the write lock.
			break

		case <-readHandover:
			// All reads are done. We have the read lock.
			break
		}
	}
}
