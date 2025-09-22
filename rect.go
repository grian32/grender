package grender

type Vec2I struct {
	X int32
	Y int32
}

type RectP struct {
	X int32
	Y int32
	W int32
	H int32
}

func (r RectP) IsContainedIn(o RectP) bool {
	return r.X >= o.X && r.Y >= o.Y && r.X+r.W <= o.X+o.W && r.Y+r.H <= o.Y+o.H
}
