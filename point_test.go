package geom

import (
	"testing"
)

func TestVertexAngle(t *testing.T) {
	A := Coord{0, 0}
	B := Coord{1, 0}
	C := Coord{1, 1}
	D := Coord{1, -1}
	r1 := VertexAngle(A, B, C)
	r2 := VertexAngle(A, B, D)
	if r1 != -r2 {
		t.Error("r1 != -r2", r1, r2)
	}

	p1 := Coord{1, 2}
	p2 := Coord{0, 3}
	p3 := Coord{0, 0}
	p4 := Coord{1, 1}

	rp := VertexAngle(p1, p2, p3)
	rn := VertexAngle(p4, p2, p3)
	t.Log(rp, rn)
}

func TestVectorAngle(t *testing.T) {
	v1 := Coord{0, -1}
	v2 := Coord{-1, 0}
	v3 := Coord{1, -1}
	v4 := Coord{-1, 0}

	a12 := VectorAngle(v1, v2)
	a34 := VectorAngle(v3, v4)

	t.Log(a12, a34)
}
