package main

import (
	"time"
	"math/rand"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/mitsuji/kawakudari_pixel/std15"
)

var frame int32 = 0
var x int32 = 15
var running bool = true

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Kawakudari Pixel",
		Bounds: pixel.R(0, 0, 512, 384),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	std15 := std15.New(512, 384, 32, 24)
	for !win.Closed() {
		if !running {continue}
		if frame % 5 == 0 {
		  if win.Pressed(pixelgl.KeyLeft) {
		    x--
		  }
		  if win.Pressed(pixelgl.KeyRight) {
		    x++
		  }
        	  std15.Locate(x,5)
        	  std15.Putc('0')
        	  std15.Locate(rand.Int31n(32),23)
        	  std15.Putc('*')
        	  std15.Scroll()
		  if std15.Scr(x,5) != 0 {
		    running = false
		  }
		}
		win.Clear(colornames.Black)
		imd := imdraw.New(nil)
		std15.PAppletDraw(imd)
		imd.Draw(win)
		win.Update()
		frame++
	}
}

func main() {
        rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}

