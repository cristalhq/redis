package redis

// Error represents an error returned by Redis.
type Error string

func (err Error) Error() string { return string(err) }
