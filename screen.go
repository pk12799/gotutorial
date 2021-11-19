package main

//
// import (
// 	"fmt"
// 	"image/png"
// 	"os"
// 	"time"

// 	"github.com/kbinani/screenshot"
// )

// func screen(v int) {
// 	n := screenshot.NumActiveDisplays()

// 	for i := 0; i < n; i++ {
// 		bounds := screenshot.GetDisplayBounds(i)
// 		time.Sleep(2 * time.Second)
// 		img, err := screenshot.CaptureRect(bounds)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fileName := fmt.Sprintf("%d_%dx%d%d.png", i, bounds.Dx(), bounds.Dy(), i)
// 		file, _ := os.Create(fileName)

// 		defer file.Close()
// 		png.Encode(file, img)

// 		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
// 	}
// }
// func main() {
// 	for {
// 		time.Sleep(2 * time.Second)
// 		screen(5)

// 	}
// }
