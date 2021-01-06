package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

// Geo represents a Redis GeoHash structure.
type Geo struct {
	name   string
	client *redisClient
}

// NewGeo instantiates a new Geo structure client for Redis.
func NewGeo(name string, client *redisClient) *Geo {
	return &Geo{name: name, client: client}
}

type (
	// GeoLocation is used with GeoAdd to add geospatial location.
	GeoLocation = redis.GeoLocation

	// GeoPos is used with GeoPos to add geospatial position.
	GeoPos = redis.GeoPos

	// GeoRadiusQuery is used with GeoRadius to query geospatial index.
	GeoRadiusQuery = redis.GeoRadiusQuery
)

// GeoUnit ...
type GeoUnit string

// GeoUnit consts.
const (
	MetersUnit     GeoUnit = "m"
	KiloMetersUnit GeoUnit = "km"
	MilesUnit      GeoUnit = "mi"
	FeetsUnit      GeoUnit = "ft"
)

// IsValidGeoUnit ...
func IsValidGeoUnit(unit GeoUnit) bool {
	switch unit {
	case MetersUnit, KiloMetersUnit, MilesUnit, FeetsUnit:
		return true
	default:
		return false
	}
}

// Add add one or more geospatial items in the geospatial index represented using a sorted set.
// TODO: XX, NX, CH https://redis.io/commands/geoadd
func (g *Geo) Add(ctx context.Context, geoLocation ...*GeoLocation) (int64, error) {
	return g.client.GeoAdd(ctx, g.name, geoLocation...).Result()
}

// Dist returns the distance between two members of a geospatial index.
func (g *Geo) Dist(ctx context.Context, member1, member2 string, unit GeoUnit) (float64, error) {
	if IsValidGeoUnit(unit) {
		return g.client.GeoDist(ctx, g.name, member1, member2, string(unit)).Result()
	}
	return 0, errors.New("unknown unit value")
}

// Hash returns members of a geospatial index as standard geohash strings.
func (g *Geo) Hash(ctx context.Context, members ...string) ([]string, error) {
	return g.client.GeoHash(ctx, g.name, members...).Result()
}

// Position returns longitude and latitude of members of a geospatial index.
func (g *Geo) Position(ctx context.Context, members ...string) ([]*GeoPos, error) {
	return g.client.GeoPos(ctx, g.name, members...).Result()
}

// Radius query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point.
func (g *Geo) Radius(ctx context.Context, longitude, latitude float64, query GeoRadiusQuery) ([]GeoLocation, error) {
	if !IsValidGeoUnit(GeoUnit(query.Unit)) {
		return nil, errors.New("unknown unit value")
	}
	return g.client.GeoRadius(ctx, g.name, longitude, latitude, &query).Result()
}

// RadiusByMember query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member.
func (g *Geo) RadiusByMember(ctx context.Context, member string, query GeoRadiusQuery) ([]GeoLocation, error) {
	if !IsValidGeoUnit(GeoUnit(query.Unit)) {
		return nil, errors.New("unknown unit value")
	}
	return g.client.GeoRadiusByMember(ctx, g.name, member, &query).Result()
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
