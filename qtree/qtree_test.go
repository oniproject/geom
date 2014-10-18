package qtree

import (
	"github.com/oniproject/geom"
	"testing"
)

func TestInsertCollect(t *testing.T) {
	cfg := ConfigDefault()
	qt := New(cfg, geom.Rect{geom.Coord{0, 0}, geom.Coord{100, 100}})

	r := geom.Rect{geom.Coord{20, 20}, geom.Coord{40, 40}}
	qt.Insert(r)

	collection := make(map[Item]bool)
	qt.CollectIntersect(r, collection)

	t.Logf("%v", collection)
}
