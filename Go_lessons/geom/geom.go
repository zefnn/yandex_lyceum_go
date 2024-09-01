package geom

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Scale(factor float64) (float64, float64) {
	r.Width *= factor
	r.Height *= factor
	return r.Width, r.Height
}
