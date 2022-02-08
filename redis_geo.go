package redis

import (
	"context"
	"errors"
)

// Geo represents a Redis GeoHash structure.
type Geo struct {
	name string
	c    *Client
}

// NewGeo instantiates a new Geo structure client for Redis.
func NewGeo(name string, client *Client) *Geo {
	return &Geo{name: name, c: client}
}

type (
	// GeoLocation is used with GeoAdd to add geospatial location.
	GeoLocation = interface{} // todo redis.GeoLocation

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
func (g *Geo) Add(ctx context.Context, geoLocation ...*GeoLocation) (int64, error) {
	panic("") // todo return g.client.GeoAdd(ctx, g.name, geoLocation...).Result()
}

// Dist returns the distance between two members of a geospatial index.
func (g *Geo) Dist(ctx context.Context, member1, member2 string, unit GeoUnit) (float64, error) {
	if unit != MetersUnit && unit != KilometersUnit && unit != MilesUnit && unit != FeetsUnit {
		return 0, errors.New("unknown GeoUnit unit")
	}
	panic("") // todo return g.client.GeoDist(ctx, g.name, member1, member2, string(unit)).Result()
}

// Hash returns members of a geospatial index as standard geohash strings.
func (g *Geo) Hash(ctx context.Context, members ...string) ([]string, error) {
	panic("") // todo return g.client.GeoHash(ctx, g.name, members...).Result()
}

// Position returns longitude and latitude of members of a geospatial index.
func (g *Geo) Position(ctx context.Context, members ...string) ([]*GeoPos, error) {
	panic("") // todo return g.client.GeoPos(ctx, g.name, members...).Result()
}

// Radius query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point.
func (g *Geo) Radius(ctx context.Context, longitude, latitude float64, query GeoRadiusQuery) ([]GeoLocation, error) {
	panic("")
	// if unit != MetersUnit && unit != KilometersUnit && unit != MilesUnit && unit != FeetsUnit {
	// 	return 0, errors.New("unknown GeoUnit unit")
	// }
}

// RadiusByMember query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member.
func (g *Geo) RadiusByMember(ctx context.Context, member string, query GeoRadiusQuery) ([]GeoLocation, error) {
	panic("")
	// if unit != MetersUnit && unit != KilometersUnit && unit != MilesUnit && unit != FeetsUnit {
	// 	return 0, errors.New("unknown GeoUnit unit")
	// }
}

// Search ...
// TODO: https://redis.io/commands/geosearch
func (g *Geo) Search(ctx context.Context) error {
	return nil
}

// SearchStore ...
// TODO: https://redis.io/commands/geosearchstore
func (g *Geo) SearchStore(ctx context.Context) error {
	return nil
}

/*
GEOADD
GEODIST
GEOHASH
GEOPOS
GEORADIUS
GEORADIUSBYMEMBER
GEORADIUSBYMEMBER_RO
GEORADIUS_RO
GEOSEARCH
GEOSEARCHSTORE
*/
