package models

import (
	/*"fmt"*/
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Asteroid struct {
	posX   float32
	posY   float32
	status bool
	image  *canvas.Image
}

var posInitialX float32 = 0

func NewAsteroid(posX float32, posY float32, image *canvas.Image) *Asteroid {
	posInitialX = posX
	return &Asteroid{
		posX:   posX,
		posY:   posY,
		status: true,
		image:  image,
	}

}

func (a *Asteroid) InitialPositionAsteroid() {
	a.posX = posInitialX
	a.image.Move(fyne.NewPos(a.posX, a.posY))
}

func (a *Asteroid) MoveAsteroid(steps float32) {

	var incX float32 = steps
	a.posX -= incX
	a.image.Move(fyne.NewPos(a.posX, a.posY))
	time.Sleep(30 * time.Millisecond)

}

func (a *Asteroid) StopAsteroid() {
	a.status = false
}

func (a *Asteroid) StarAsteroid() {
	a.status = true
}

func (a *Asteroid) GetStatus() bool {
	return a.status
}

func (a *Asteroid) GetPositionInitialX() float32 {
	return posInitialX
}


func (a *Asteroid) GetPositionShip() fyne.Position {
	return fyne.NewPos(a.posX, a.posY)
}
