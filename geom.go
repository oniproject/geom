package geom

type Bounded interface {
	Bounds() (bounds Rect)
}

type Transformable interface {
	Translate(offset Coord)
	Rotate(rad float64)
	Scale(xfactor, yfactor float64)
}
