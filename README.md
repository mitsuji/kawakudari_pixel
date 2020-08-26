# kawakudari-pixel

This project implements part of the [std15.h](https://github.com/IchigoJam/c4ij/blob/master/src/std15.h) API (from [c4ij](https://github.com/IchigoJam/c4ij)) with [pixel](https://github.com/faiface/pixel), and [Kawakudari Game](https://ichigojam.github.io/print/en/KAWAKUDARI.html) on top of it.

It will allow programming for [IchigoJam](https://ichigojam.net/index-en.html)-like targets using a Go programming language.
```
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
                  std15.Putc(65)
                  std15.Locate(rand.Int31n(32),23)
                  std15.Putc(66)
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
```

## Prerequisite

* This project using programming language Go, so you need [Go](https://golang.org/doc/install) build tool properly installd to run example code.


## How to use

To just run example
```
$ go run main.go
```

To build executeble and run example
```
$ go build
$ ./kawakudari_pixel
```
