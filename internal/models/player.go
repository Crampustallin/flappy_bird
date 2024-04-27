package models

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GRAVITY = 5.432
	JUMP    = 50
)

type Player struct {
	posX, poxY, vPoxY float32
	width, height     float32
	score             int
}

func NewPlayer() *Player {
	return &Player{
		posX:  60,
		poxY:  190,
		score: 0,
	}
}

func (p *Player) Success() {
	p.score++
}

func (p *Player) GetScore() int {
	return p.score
}

func (p *Player) GetPos() (float32, float32) {
	return p.posX, p.poxY
}

func (p *Player) Jump() {
	p.vPoxY = -JUMP
}

func (p *Player) Fall() {
	p.poxY += p.vPoxY
	if p.poxY >= 325 {
		p.vPoxY = 0
	} else {
		p.vPoxY += GRAVITY
	}
	if p.vPoxY > JUMP {
		p.vPoxY = JUMP
	}
}

func (p *Player) Move(dir float32) {
	p.posX += dir
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, p.posX, p.poxY, 22.0, color.RGBA{R: 255, A: 255}, true)
}
