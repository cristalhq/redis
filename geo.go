package redis

import (
	"errors"

	"github.com/go-redis/redis"
)

// Geo represents a Redis GeoHash structure.
type Geo struct {
	name   string
	client *redis.Client
}

// NewGeo instantiates a new Geo structure client for Redis.
func NewGeo(name string, client *redis.Client) *Geo {
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

type GeoUnit string

const (
	MetersUnit     GeoUnit = "m"
	KiloMetersUnit GeoUnit = "km"
	MilesUnit      GeoUnit = "mi"
	FeetsUnit      GeoUnit = "ft"
)

func IsValidGeoUnit(unit GeoUnit) bool {
	switch unit {
	case MetersUnit, KiloMetersUnit, MilesUnit, FeetsUnit:
		return true
	default:
		return false
	}
}

// Add add one or more geospatial items in the geospatial index represented using a sorted set.
func (g *Geo) Add(geoLocation ...*GeoLocation) (int64, error) {
	return g.client.GeoAdd(g.name, geoLocation...).Result()
}

// Dist returns the distance between two members of a geospatial index.
func (g *Geo) Dist(member1, member2 string, unit GeoUnit) (float64, error) {
	if IsValidGeoUnit(unit) {
		return g.client.GeoDist(g.name, member1, member2, string(unit)).Result()
	}
	return 0, errors.New("unknown unit value")
}

// Hash returns members of a geospatial index as standard geohash strings.
func (g *Geo) Hash(members ...string) ([]string, error) {
	return g.client.GeoHash(g.name, members...).Result()
}

// Position returns longitude and latitude of members of a geospatial index.
func (g *Geo) Position(members ...string) ([]*GeoPos, error) {
	return g.client.GeoPos(g.name, members...).Result()
}

// Radius query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point.
func (g *Geo) Radius(longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error) {
	if IsValidGeoUnit(GeoUnit(query.Unit)) {
		return g.client.GeoRadius(g.name, longitude, latitude, query).Result()
	}
	return nil, errors.New("unknown unit value")
}

// RadiusByMember query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member.
func (g *Geo) RadiusByMember(member string) ([]GeoLocation, error) {
	return g.client.GeoRadiusByMember(g.name, member, nil).Result()
}
