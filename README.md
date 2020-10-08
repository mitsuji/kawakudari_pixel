# kawakudari-pixel

This project implements part of the [std15.h](https://github.com/IchigoJam/c4ij/blob/master/src/std15.h) API (from [c4ij](https://github.com/IchigoJam/c4ij)) with [pixel](https://github.com/faiface/pixel), and [Kawakudari Game](https://ichigojam.github.io/print/en/KAWAKUDARI.html) on top of it.

It will allow programming for [IchigoJam](https://ichigojam.net/index-en.html)-like targets that display [IchigoJam FONT](https://mitsuji.github.io/ichigojam-font.json/) on screen using a Go programming language.
```
var frame uint32 = 0
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

  std15 := ij.New(512, 384, 32, 24)
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
      std15.Scroll(ij.Up)
      if std15.Scr(x,5) != 0 {
        std15.Locate(0,23)
        std15.Putstr("Game Over...")
        std15.Putnum(int32(frame))
        running = false
      }
    }
    std15.DrawScreen(win)
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
* In Windows enviroment, you might need [install and configure MSYS2](https://github.com/faiface/pixel/wiki/Building-Pixel-on-Windows) to build and run this project.


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


## License
[![Creative Commons License](https://i.creativecommons.org/l/by/4.0/88x31.png)](http://creativecommons.org/licenses/by/4.0/)
[CC BY](https://creativecommons.org/licenses/by/4.0/) [mitsuji.org](https://mitsuji.org)

This work is licensed under a [Creative Commons Attribution 4.0 International License](http://creativecommons.org/licenses/by/4.0/).
