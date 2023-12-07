package scenes

import (
	"gamecode/models"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var asteroid1 *models.Asteroid
var asteroid2 *models.Asteroid
var asteroid3 *models.Asteroid
var ship *models.Ship

var topUp int = 5
var topDow int = 8

var stepUp int = 0
var stepDow int = 0

type MainMenuScene struct {
	window fyne.Window
}

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window}
}

var content *fyne.Container

// Define un canal para notificar la detención del juego.
var stopGameCh chan struct{}

func (scene *MainMenuScene) Show() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/space.jpg"))
	backgroundImage.Resize(fyne.NewSize(1200, 600))

	imageAsteroid1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/asteroid1.png"))
	imageAsteroid1.Resize(fyne.NewSize(100, 100))

	imageAsteroid2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/asteroid2.png"))
	imageAsteroid2.Resize(fyne.NewSize(100, 100))

	imageAsteroid3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/asteroid2.png"))
	imageAsteroid3.Resize(fyne.NewSize(100, 100))

	imageShip := canvas.NewImageFromURI(storage.NewFileURI("./assets/ship.png"))
	imageShip.Resize(fyne.NewSize(100, 100))

	ship = models.NewShip(100, 250, imageShip)
	ship.InitialPositionShip()

	buttonStart := widget.NewButton("Iniciar juego", scene.StartGame)
	buttonStart.Resize(fyne.NewSize(150, 30))
	buttonStart.Move(fyne.NewPos(300, 10))

	/*buttonStop := widget.NewButton("Detener juego", scene.StopGame)
	buttonStop.Resize(fyne.NewSize(150, 30))
	buttonStop.Move(fyne.NewPos(300, 50))*/

	buttonReset := widget.NewButton("Finalizar juego", scene.ResetGame)
	buttonReset.Resize(fyne.NewSize(150, 30))
	buttonReset.Move(fyne.NewPos(300, 90)) // Ajusta la posición del botón de reinicio

	asteroid1 = models.NewAsteroid(1200, 100, imageAsteroid1)
	asteroid1.InitialPositionAsteroid()

	asteroid2 = models.NewAsteroid(1200, 250, imageAsteroid2)
	asteroid2.InitialPositionAsteroid()

	asteroid3 = models.NewAsteroid(1200, 400, imageAsteroid3)
	asteroid3.InitialPositionAsteroid()

	go scene.moveShip()

	content = container.NewWithoutLayout(
		backgroundImage,
		imageAsteroid1,
		imageAsteroid2,
		imageAsteroid3,
		imageShip,
		buttonStart,
		//buttonStop,
		buttonReset,
	)

	scene.window.SetContent(content)
}

func (scene *MainMenuScene) checkPositionConcurrent() {
	newShipImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/splash.png"))
	newShipImage.Resize(fyne.NewSize(100, 100))

	margin := 0.99
	for {
		select {
		case <-stopGameCh:
			return
		default:
			shipPosX := ship.GetPositionShip().X
			shipPosY := ship.GetPositionShip().Y

			asteroid2PosX := asteroid2.GetPositionShip().X
			asteroid2PosY := asteroid2.GetPositionShip().Y

			if float64(shipPosX) >= float64(asteroid2PosX)-margin &&
				float64(shipPosX) <= float64(asteroid2PosX)+margin &&
				float64(shipPosY) >= float64(asteroid2PosY)-margin &&
				float64(shipPosY) <= float64(asteroid2PosY)+margin {
				scene.StopGame()
				ship.SetImage(newShipImage)
				break
			}

			asteroid1PosX := asteroid1.GetPositionShip().X
			asteroid1PosY := asteroid1.GetPositionShip().Y

			if float64(shipPosX) >= float64(asteroid1PosX)-margin &&
				float64(shipPosX) <= float64(asteroid1PosX)+margin &&
				float64(shipPosY) >= float64(asteroid1PosY)-margin &&
				float64(shipPosY) <= float64(asteroid1PosY)+margin {
				scene.StopGame()
				ship.SetImage(newShipImage)
				break
			}

			asteroid3PosX := asteroid3.GetPositionShip().X
			asteroid3PosY := asteroid3.GetPositionShip().Y

			if float64(shipPosX) >= float64(asteroid3PosX)-margin &&
				float64(shipPosX) <= float64(asteroid3PosX)+margin &&
				float64(shipPosY) >= float64(asteroid3PosY)-margin &&
				float64(shipPosY) <= float64(asteroid3PosY)+margin {
				scene.StopGame()
				ship.SetImage(newShipImage)
				break
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (scene *MainMenuScene) moveShip() {
	scene.window.Canvas().(desktop.Canvas).SetOnKeyDown(func(event *fyne.KeyEvent) {
		keyName := event.Name
		if keyName == "W" { // Arriba
			if stepUp < 6 {
				stepUp++
				stepDow--
				ship.MoveUp(30)
			}
		}
		if keyName == "S" { // Abajo
			if stepDow < 9 {
				stepUp--
				stepDow++
				ship.MoveDow(30)
			}
		}
	})
}

func (scene *MainMenuScene) StartGame() {
	stopGameCh = make(chan struct{})

	go scene.checkPositionConcurrent()
	go scene.delayedMoveAsteroid(1, scene.MoveAsteroid1)
	go scene.delayedMoveAsteroid(2, scene.MoveAsteroid2)
	go scene.delayedMoveAsteroid(3, scene.MoveAsteroid3)
}

func (scene *MainMenuScene) delayedMoveAsteroid(seconds int, moveFunc func()) {
	time.Sleep(time.Duration(seconds) * time.Second)
	moveFunc()
}

var countAsteroid1 int = 0
var endAsteroid1 bool = false

func (scene *MainMenuScene) MoveAsteroid1() {
	endAsteroid1 = false
	asteroid1.StarAsteroid()
	for asteroid1.GetStatus() {
		println("Bucle")
		countAsteroid1++
		asteroid1.MoveAsteroid(20)
		if countAsteroid1 > 60 {
			asteroid1.InitialPositionAsteroid()
			countAsteroid1 = 0
			println("Tope del asteroide")
			numeroAleatorio := rand.Intn(2) + 1
			time.Sleep(time.Duration(numeroAleatorio) * time.Second)
		}
	}
}

var countAsteroid2 int = 0
var endAsteroid2 bool = false

func (scene *MainMenuScene) MoveAsteroid2() {
	endAsteroid2 = false
	asteroid2.StarAsteroid()
	for asteroid2.GetStatus() {
		countAsteroid2++
		asteroid2.MoveAsteroid(20)
		if countAsteroid2 > 60 {
			asteroid2.InitialPositionAsteroid()
			countAsteroid2 = 0
			numeroAleatorio := rand.Intn(2) + 1
			time.Sleep(time.Duration(numeroAleatorio) * time.Second)
		}
	}
}

var countAsteroid3 int = 0
var endAsteroid3 bool = false

func (scene *MainMenuScene) MoveAsteroid3() {
	endAsteroid3 = false
	asteroid3.StarAsteroid()
	for asteroid3.GetStatus() {
		countAsteroid3++
		asteroid3.MoveAsteroid(20)
		if countAsteroid3 > 60 {
			asteroid3.InitialPositionAsteroid()
			countAsteroid3 = 0
			numeroAleatorio := rand.Intn(2) + 1
			time.Sleep(time.Duration(numeroAleatorio) * time.Second)
		}
	}
}

func (scene *MainMenuScene) StopGame() {
	close(stopGameCh)

	go asteroid1.StopAsteroid()
	go asteroid2.StopAsteroid()
	go asteroid3.StopAsteroid()
}

func (scene *MainMenuScene) ResetGame() {
	close(stopGameCh)

	go asteroid1.StopAsteroid()
	go asteroid1.InitialPositionAsteroid()
	go asteroid2.StopAsteroid()
	go asteroid2.InitialPositionAsteroid()
	go asteroid3.StopAsteroid()
	go asteroid3.InitialPositionAsteroid()
}
