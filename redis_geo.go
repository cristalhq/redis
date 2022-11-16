package redis

import (
	"context"
	"fmt"
)

// Geo represents a Redis GeoHash structure.
type Geo struct {
	name string
	c    *Client
}

// NewGeo instantiates a new Geo structure client for Redis.
func NewGeo(name string, client *Client) Geo {
	return Geo{name: name, c: client}
}

type (
	// GeoLocation is used with GeoAdd to add geospatial location.
	GeoLocation         = interface{} // todo redis.GeoLocation
	GeoSearchQuery      = interface{}
	GeoSearchStoreQuery = interface{}
	// GeoPos is used with GeoPos to add geospatial position.
	GeoPos = interface{} // todo redis.GeoPos

	// GeoRadiusQuery is used with GeoRadius to query geospatial index.
	GeoRadiusQuery = interface{} // todo redis.GeoRadiusQuery
)

// GeoUnit ...
type GeoUnit string

// GeoUnit consts.
const (
	MetersUnit     GeoUnit = "m"
	KilometersUnit GeoUnit = "km"
	MilesUnit      GeoUnit = "mi"
	FeetsUnit      GeoUnit = "ft"
)

// Add add one or more geospatial items in the geospatial index represented using a sorted set.
// TODO: XX, NX, CH https://redis.io/commands/geoadd
func (g Geo) Add(ctx context.Context, geoLocation ...*GeoLocation) (int64, error) {
	panic("redis: Geo.Add not implemented")
}

// Dist returns the distance between two members in the geospatial index in units.
// See: https://redis.io/commands/geodist
func (g Geo) Dist(ctx context.Context, member1, member2 string, unit GeoUnit) (float64, error) {
	if unit != MetersUnit && unit != KilometersUnit && unit != MilesUnit && unit != FeetsUnit {
		return 0, fmt.Errorf("unknown GeoUnit unit: %s", unit)
	}
	req := newRequest("*5\r\n$7\r\nGEODIST\r\n$")
	req.addString4(g.name, member1, member2, string(unit))
	return g.c.cmdFloat(ctx, req)
}

// Hash returns Geohash strings representing the position of one or more elements.
// See: https://redis.io/commands/geohash
func (g Geo) Hash(ctx context.Context, members ...string) ([]string, error) {
	req := newRequestSize(2+len(members), "\r\n$7\r\nGEOHASH\r\n$")
	req.addStringAndStrings(g.name, members)
	return g.c.cmdStrings(ctx, req)
}

// Position returns the positions (longitude,latitude) of all the specified members.
// See: https://redis.io/commands/geopos
func (g Geo) Position(ctx context.Context, members ...string) ([][2]float64, error) {
	req := newRequestSize(2+len(members), "\r\n$6\r\nGEOPOS\r\n$")
	req.addStringAndStrings(g.name, members)
	// TODO(oleg): decode array
	_, err := g.c.cmdStrings(ctx, req)
	return nil, err
}

// Radius query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point.
func (g Geo) Radius(ctx context.Context, longitude, latitude float64, query GeoRadiusQuery) ([]GeoLocation, error) {
	panic("redis: Geo.Radius not implemented")
}

// RadiusByMember query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member.
func (g Geo) RadiusByMember(ctx context.Context, member string, query GeoRadiusQuery) ([]GeoLocation, error) {
	panic("redis: Geo.RadiusByMember not implemented")
}

// Search ...
// TODO: https://redis.io/commands/geosearch
func (g Geo) Search(ctx context.Context) error {
	return nil
}

// SearchStore ...
// TODO: https://redis.io/commands/geosearchstore
func (g Geo) SearchStore(ctx context.Context) error {
	return nil
}

/*
GEOPOS
GEORADIUS
GEORADIUSBYMEMBER
GEORADIUSBYMEMBER_RO
GEORADIUS_RO
*/
