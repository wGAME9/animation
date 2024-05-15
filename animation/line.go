package animation

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func NewLine(
	start, end Point,
	speed,
	strokeWidth float32,
	clr color.Color,
	antialias bool,
) Animation {
	vectorStep := getVectorStep(start, end, speed)

	return &line{
		start:       start,
		end:         end,
		current:     start,
		speed:       speed,
		basedImage:  ebiten.NewImage(calculateWidthAndHeight(start, end)),
		clr:         clr,
		strokeWidth: strokeWidth,
		antialias:   antialias,
		vectorStep:  vectorStep,
	}
}

type line struct {
	start   Point
	end     Point
	current Point

	speed float32

	basedImage  *ebiten.Image
	strokeWidth float32
	clr         color.Color
	antialias   bool

	vectorStep Point

	isDone bool
}

func (l *line) Update() {
	// stop the animation if it's done
	if l.isDone {
		return
	}

	current := l.current
	next := l.getNextPoint()

	vector.StrokeLine(
		l.basedImage,
		float32(current.X), float32(current.Y),
		float32(next.X), float32(next.Y),
		float32(l.strokeWidth),
		l.clr,
		l.antialias,
	)

	if isSamePoint(current, next) {
		l.isDone = true
		return
	}

	l.current.X = next.X
	l.current.Y = next.Y
}

func (l *line) Image() *ebiten.Image {
	return l.basedImage
}

func (l *line) getNextPoint() Point {
	vectorStep := l.vectorStep

	next := Point{
		X: l.current.X + vectorStep.X,
		Y: l.current.Y + vectorStep.Y,
	}

	distanceOfStartEnd := getDistanceOfTwoPoints(l.start, l.end)
	distanceOfStartNext := getDistanceOfTwoPoints(l.start, next)

	if distanceOfStartNext >= distanceOfStartEnd {
		next.X = l.end.X
		next.Y = l.end.Y
	}

	return next
}
