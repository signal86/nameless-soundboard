package main

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  gui "github.com/gen2brain/raylib-go/raygui"
)

func main() {
  
  rl.InitWindow(1280, 720, "nameless soundboard tool")

  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {

    rl.BeginDrawing()

    rl.ClearBackground(rl.GetColor(uint(gui.GetStyle(gui.DEFAULT, gui.BACKGROUND_COLOR))))

    rl.EndDrawing()

  }

  rl.CloseWindow()

}
