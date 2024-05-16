package main

import (
	"image/color"
	"log"

	"github.com/wGAME9/animation/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	colorWhite = color.White

	colorRed   = color.RGBA{255, 0, 0, 255}
	colorGreen = color.RGBA{0, 255, 0, 255}
	colorBlue  = color.RGBA{0, 0, 255, 255}
)

type Game struct {
	diagonal   animation.Animation
	vertical   animation.Animation
	horizontal animation.Animation
}

func (g *Game) Update() error {
	g.diagonal.Update()
	g.vertical.Update()
	g.horizontal.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colorWhite)

	screen.DrawImage(g.diagonal.Image(), &ebiten.DrawImageOptions{})
	screen.DrawImage(g.vertical.Image(), &ebiten.DrawImageOptions{})
	screen.DrawImage(g.horizontal.Image(), &ebiten.DrawImageOptions{})
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Drawing Animation!")

	diagonal := animation.NewLine(
		animation.Point{X: 10, Y: 10},
		animation.Point{X: 300, Y: 300},
		10,
		2,
		colorBlue,
		true,
	)

	vertical := animation.NewLine(
		animation.Point{X: 10, Y: 10},
		animation.Point{X: 300, Y: 10},
		10,
		2,
		colorRed,
		true,
	)

	horizontal := animation.NewLine(
		animation.Point{X: 10, Y: 10},
		animation.Point{X: 10, Y: 300},
		10,
		2,
		colorGreen,
		true,
	)

	game := &Game{
		diagonal:   diagonal,
		vertical:   vertical,
		horizontal: horizontal,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
