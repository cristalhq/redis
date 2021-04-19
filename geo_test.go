package redis

import (
	"context"
	"testing"
)

func TestGeoNoAlloc(t *testing.T) {
	t.SkipNow()

	c := NewGeo("testgeo", redisTestClient)
	loc1 := &GeoLocation{
		Name:      "mycity",
		Longitude: 42,
		Latitude:  69,
	}
	loc2 := &GeoLocation{
		Name:      "cityyour",
		Longitude: 69,
		Latitude:  42,
	}

	if _, err := c.Add(context.TODO(), loc1); err != nil {
		t.Fatal(err)
	}
	if _, err := c.Add(context.TODO(), loc2); err != nil {
		t.Fatal(err)
	}

	f := func() {
		if _, err := c.Add(context.TODO(), loc1); err != nil {
			t.Fatal(err)
		}
		if _, err := c.Dist(context.TODO(), "mycity", "cityyour", MetersUnit); err != nil {
			t.Fatal(err)
		}
		if _, err := c.Hash(context.TODO(), "mycity"); err != nil {
			t.Fatal(err)
		}
		if _, err := c.Position(context.TODO(), "mycity", "2"); err != nil {
			t.Fatal(err)
		}
		if _, err := c.Radius(context.TODO(), 1, 2, GeoRadiusQuery{
			Radius: 10000,
			Unit:   "m",
		}); err != nil {
			t.Fatal(err)
		}
		if _, err := c.RadiusByMember(context.TODO(), "mycity", GeoRadiusQuery{
			Radius: 10000,
			Unit:   "m",
		}); err != nil {
			t.Fatal(err)
		}
	}

	perRun := testing.AllocsPerRun(1, f)
	if want := 0; perRun != 0 {
		t.Errorf("want %v memory allocations, did %v", want, perRun)
	}
}
