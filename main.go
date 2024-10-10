package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionMoveLeft input.Action = iota
	ActionMoveRight
    ActionMoveUp
    ActionMoveDown
    outsideWidth = 640
    outsideHeight = 480
)

var img *ebiten.Image

func init() {
        var err error
        img, _, err = ebitenutil.NewImageFromFile("astronaut.png")
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(newExampleGame()); err != nil {
		log.Fatal(err)
	}
}

type exampleGame struct {
	p           *player
	inputSystem input.System
}

func newExampleGame() *exampleGame {
	g := &exampleGame{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyA},
		ActionMoveRight: {input.KeyD},
		ActionMoveUp: {input.KeyS},
		ActionMoveDown: {input.KeyW},
	}
	g.p = &player{
		input: g.inputSystem.NewHandler(0, keymap),
		pos:   image.Point{X: 96, Y: 96},
	}
	return g
}

func (g *exampleGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *exampleGame) Draw(screen *ebiten.Image) {
	g.p.Draw(screen)
}

func (g *exampleGame) Update() error {
	g.inputSystem.Update()
	g.p.Update(outsideWidth, outsideHeight)
	return nil
}

type player struct {
	input *input.Handler
	pos   image.Point
}

func (p *player) Update(outsideWidth, outsideHeight int) {
    if math.Abs(float64(p.pos.X) - float64(outsideWidth)) < 0.4 || math.Abs(float64(p.pos.X)) < 0.4 {
        p.pos.X = outsideWidth / 2
        p.pos.Y = outsideHeight / 2
    }
    if math.Abs(float64(p.pos.Y) - float64(outsideWidth)) < 0.4 || math.Abs(float64(p.pos.Y)) < 0.4 {
        p.pos.X = outsideWidth / 2
        p.pos.Y = outsideHeight / 2
    }
	if p.input.ActionIsPressed(ActionMoveLeft) {
		p.pos.X -= 4
	}
	if p.input.ActionIsPressed(ActionMoveRight) {
		p.pos.X += 4
	}
	if p.input.ActionIsPressed(ActionMoveUp) {
		p.pos.Y += 4
	}
	if p.input.ActionIsPressed(ActionMoveDown) {
		p.pos.Y -= 4
	}
}

func (p *player) Draw(screen *ebiten.Image) {
    // ebitenutil.DebugPrintAt(screen, "player", p.pos.X, p.pos.Y)
    ebitenutil.DrawCircle(screen, float64(p.pos.X), float64(p.pos.Y), 10., color.Opaque)
}