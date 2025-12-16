package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Point struct {
	X float64
	Y float64
}

// cubic Bezier curve calculation
func bezier(p0, p1, p2, p3 Point, t float64) Point {
	u := 1 - t
	tt := t * t
	uu := u * u
	uuu := uu * u
	ttt := tt * t

	return Point{
		X: uuu*p0.X + 3*uu*t*p1.X + 3*u*tt*p2.X + ttt*p3.X,
		Y: uuu*p0.Y + 3*uu*t*p1.Y + 3*u*tt*p2.Y + ttt*p3.Y,
	}
}

// MoveMouseHumanLike simulates natural mouse movement
func MoveMouseHumanLike(page *rod.Page, start, end Point) {
	rand.Seed(time.Now().UnixNano())

	ctrl1 := Point{
		X: start.X + rand.Float64()*120 - 60,
		Y: start.Y + rand.Float64()*120 - 60,
	}
	ctrl2 := Point{
		X: end.X + rand.Float64()*120 - 60,
		Y: end.Y + rand.Float64()*120 - 60,
	}

	steps := rand.Intn(40) + 40

	for i := 0; i <= steps; i++ {
	t := float64(i) / float64(steps)
	pos := bezier(start, ctrl1, ctrl2, end, t)

	page.Mouse.MoveTo(proto.Point{
		X: pos.X,
		Y: pos.Y,
	})

	delay := time.Duration(
		math.Abs(rand.NormFloat64())*12+6,
	) * time.Millisecond

	time.Sleep(delay)
}



}
