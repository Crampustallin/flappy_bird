package main

import (
	"fmt"

	"github.com/Crampustallin/flappy_bird/internal/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mode int

const (
	MODE_TITLE = iota
	MODE_GAME
)

type Game struct {
	player *models.Player
	mode   Mode
}

func NewGame() ebiten.Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.player = models.NewPlayer()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	switch g.mode {
	case MODE_TITLE:
		if g.isKeyJustPressed() {
			g.mode = MODE_GAME
		}
	case MODE_GAME:
		if g.isKeyJustPressed() {
			g.player.Jump()
			g.player.Success()
		}
		g.player.Fall()
		g.handleMovement()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.mode != MODE_TITLE {
		g.player.Draw(screen)
	}
	x, y := g.player.GetPos()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f \nScore: %v\nPosX: %.2f \nPosY: %.2f", ebiten.ActualFPS(), g.player.GetScore(), x, y))
}

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.mode = MODE_TITLE
		return true
	}
	return false
}

func (g *Game) handleMovement() {
	move := float32(10.0)
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		move = move * 25.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.player.Move(move)
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		g.player.Move(-move)
	}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
