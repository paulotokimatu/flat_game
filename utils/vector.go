package utils

type Vec2 struct {
	X float32
	Y float32
}

func (vec2 *Vec2) SetX(x float32) {
	vec2.X = x
}

func (vec2 *Vec2) SetY(y float32) {
	vec2.Y = y
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (vec *Vec3) SetX(x float32) {
	vec.X = x
}

func (vec *Vec3) SetY(y float32) {
	vec.Y = y
}

func (vec *Vec3) SetZ(z float32) {
	vec.Z = z
}
