# Animation

My implementation of drawing animation for ebitengine. Currently **only** support drawing line.

Example:
```go
package main

import (
	"image/color"
	"log"

	"github.com/weGAME9/animation/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	colorWhite = color.White
	colorBlue  = color.RGBA{0, 0, 255, 255}
)

type Game struct {
	line   animation.Animation
}

func (g *Game) Update() error {
	g.line.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colorWhite)

	screen.DrawImage(g.line.Image(), &ebiten.DrawImageOptions{})
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Drawing Animation.")

	game := &Game{
		line: animation.NewLine(
			animation.Point{X: 10, Y: 10},
			animation.Point{X: 300, Y: 300},
			10,
			2,
			colorBlue,
			true,
		),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
```
