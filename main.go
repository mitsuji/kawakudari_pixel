package main

import (
  "time"
  "math/rand"
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  ij "github.com/mitsuji/kawakudari_pixel/ichigojam"
)

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

