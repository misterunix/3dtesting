package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var width, height int
var colors []color.RGBA
var px, py, pz float64

type Vector struct {
	X, Y, Z float64
}

type Pixel struct {
	X, Y float64
}

var p [4]Vector
var s [4]Pixel
var eye Vector

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	//ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	p[0] = Vector{-50, -50, 0}
	p[1] = Vector{50, -50, 0}
	p[2] = Vector{50, 50, 0}
	p[3] = Vector{-50, 50, 0}

	for i := 0; i < len(p); i++ {
		ebitenutil.DrawCircle(screen, p[i].X, p[i].Y, 5.0, colors[0])
	}

	x1 := s[0].X
	y1 := s[0].Y
	x2 := s[1].X
	y2 := s[1].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	x1 = s[0].X
	y1 = s[0].Y
	x2 = s[2].X
	y2 = s[2].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	x1 = s[0].X
	y1 = s[0].Y
	x2 = s[3].X
	y2 = s[3].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	x1 = s[1].X
	y1 = s[1].Y
	x2 = s[2].X
	y2 = s[2].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	x1 = s[1].X
	y1 = s[1].Y
	x2 = s[3].X
	y2 = s[3].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

	x1 = s[2].X
	y1 = s[2].Y
	x2 = s[3].X
	y2 = s[3].Y
	ebitenutil.DrawLine(screen, x1, y1, x2, y2, colors[0])

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func setup() {

	p[0] = Vector{X: 30, Y: 1, Z: 1}
	p[1] = Vector{X: 1, Y: 30, Z: 1}
	p[2] = Vector{X: 1, Y: 1, Z: 30}
	p[3] = Vector{X: -30, Y: -30, Z: -30}

	for i := 0; i < 4; i++ {
		s[i].X = (p[i].X/p[i].Z)*float64(width) + float64(width)/2
		s[i].Y = (p[i].Y/p[i].Z)*float64(height) + float64(height)/2
	}
}

func main() {

	setup()

	//c := mat.Dot(projection.At(),zero)

	c := color.RGBA{R: uint8(0xff), G: uint8(0xff), B: uint8(0xff), A: 0xff}
	colors = append(colors, c)

	width = 640
	height = 480

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("3d test")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
