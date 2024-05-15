package animation

import (
	"math"
)

// isSamePoint checks if two points are the same.
// It returns true if the X and Y coordinates of
// the two points are equal, otherwise false.
func isSamePoint(point1, point2 Point) bool {
	return point1.X == point2.X && point1.Y == point2.Y
}

// calculateWidthAndHeight calculates the width and height
// of the based image which is used for the animation.
func calculateWidthAndHeight(start, end Point) (int, int) {
	width := int(start.X + end.X)
	height := int(start.Y + end.Y)
	return width, height
}

// getVectorOfTwoPoints calculates the vector between two points.
func getVectorOfTwoPoints(point1, point2 Point) Point {
	return Point{
		X: point2.X - point1.X,
		Y: point2.Y - point1.Y,
	}
}

// getDistanceOfTwoPoints calculates the distance between two points.
func getDistanceOfTwoPoints(point1, point2 Point) float32 {
	x := point2.X - point1.X
	y := point2.Y - point1.Y
	return float32(math.Sqrt(float64(x*x + y*y)))
}

func getVectorStep(start, end Point, speed float32) Point {
	vectorOfTwoPoints := getVectorOfTwoPoints(start, end)
	distance := getDistanceOfTwoPoints(start, end)

	lineStep := speed / distance

	return Point{
		X: vectorOfTwoPoints.X * lineStep,
		Y: vectorOfTwoPoints.Y * lineStep,
	}
}
