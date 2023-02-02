package main

import (
	"fifteen_puzzle/wrapper"
	"path"
	"runtime"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	resources := wrapper.Resources{}

	const gameWidth = 256
	const gameHeight = 256

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(gameWidth, gameHeight, "15-Puzzle!", option, 60)

	t, err := wrapper.FileToTexture(fullname("15.png"), &resources)
	if err != nil {
		panic("Couldn't load 15.png")
	}

	w := 64
	grid := [6][6]int{}
	sprite := [20]*wrapper.Sprite{}

	for i := 0; i < 20; i++ {
		sprite[i] = wrapper.NewSprite()
	}

	n := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			n++
			sprite[n].SetTexture(t)
			sprite[n].SetTextureRect(i*w, j*w, w, w)
			grid[i+1][j+1] = n
		}
	}

	for wnd.IsOpen() {
		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Mouse_ButtonPressed() {
				if wnd.Mouse_ButtonIs(window.SfMouseLeft) {
					vec := graphics.SfMouse_getPositionRenderWindow(wnd.Get_Window())
					x := vec.GetX()/w + 1
					y := vec.GetY()/w + 1

					dx := 0
					dy := 0

					if grid[x+1][y] == 16 {
						dx = 1
						dy = 0
					}
					if grid[x][y+1] == 16 {
						dx = 0
						dy = 1
					}
					if grid[x][y-1] == 16 {
						dx = 0
						dy = -1
					}
					if grid[x-1][y] == 16 {
						dx = -1
						dy = 0
					}

					n = grid[x][y]
					grid[x][y] = 16
					grid[x+dx][y+dy] = n

					sprite[16].Move(float32(-dx*w), float32(-dy*w))
					speed := 3

					for i := 0; i < w; i += speed {
						sprite[n].Move(float32(speed*dx), float32(speed*dy))
					}
				}
			}
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				n = grid[i+1][j+1]
				sprite[n].SetPosition(float32(i*w), float32(j*w))
				sprite[n].Draw(wnd.Get_Window())
			}
		}

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
