package models

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
)

type Ship struct {
    posX, posY   float32
    status       bool
    image        *canvas.Image
    originalImage *canvas.Image // Agrega el campo para la imagen original
}

func NewShip(posX float32, posY float32, image *canvas.Image) *Ship {
    return &Ship{
        posX:          posX,
        posY:          posY,
        image:         image,
        originalImage: image, // Asigna la imagen original
    }
}

func (s *Ship) InitialPositionShip() {
	s.image.Move(fyne.NewPos(s.posX, s.posY))
}

func (s *Ship) MoveUp(posY float32) {
	s.posY -= posY
	print(s.posY)

	s.image.Move(fyne.NewPos(s.posX, s.posY))
}

func (s *Ship) MoveDow(posY float32) {
	s.posY += posY
	s.image.Move(fyne.NewPos(s.posX, s.posY))
}

func (s *Ship) GetPositionShip() fyne.Position {
	return fyne.NewPos(s.posX, s.posY)
}

func (s *Ship) RestoreOriginalImage() {
    s.image = s.originalImage
    s.image.Move(fyne.NewPos(s.posX, s.posY))
}

func (s *Ship) SetImage(image *canvas.Image) {
    s.image = image
    // Actualiza la posición de la nueva imagen en la pantalla
    s.image.Move(fyne.NewPos(s.posX, s.posY))
}