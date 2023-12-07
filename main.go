package main

import (
	"gamecode/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Game Code Go")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1200, 600))

	mainMenuScene := scenes.NewMainMenuScene(myWindow)

	mainMenuScene.Show()
	myWindow.ShowAndRun()
}