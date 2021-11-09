package main

import (
	"fmt"
	"strconv"
	"time"

	// "io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/data/binding"
	// "fyne.io/fyne/v2/widget"
)

func getTime() string {
	t := time.Now()
	hour, min, sec := t.Clock()
	return ItoaTwoDigits(hour) + ":" + ItoaTwoDigits(min) + ":" + ItoaTwoDigits(sec)
}

// ItoaTwoDigits time.Clock returns one digit on values, so we make sure to convert to two digits
func ItoaTwoDigits(i int) string {
	b := "0" + strconv.Itoa(i)
	return b[len(b)-2:]
}

// func onReady() {
//   go func() {
//     for {
//       getTime()
//       fmt.Println(getTime())
//       time.Sleep(1 * time.Second)
//     }
//   }()
// }

func onReady() {
	for {
		getTime()
		fmt.Println(getTime())
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Creating App
	myApp := app.New()
	// Creating NewWindow
	myWin := myApp.NewWindow("Time App")
	// Resizing the NewWindow
	myWin.Resize(fyne.NewSize(500, 500))

	go onReady()

	myWin.ShowAndRun()
}
