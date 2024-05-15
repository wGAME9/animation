package animation

import "github.com/hajimehoshi/ebiten/v2"

type Animation interface {
	Update()
	Image() *ebiten.Image
}
