package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

type pos struct {
	x, y int
}

type ball struct {
	pos    // Composition : now we can use ball.x instead of ball.pos.x
	radius int
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.x += int(ball.xv)
	ball.y += int(ball.yv)

	if int(ball.y)-ball.radius < 0 || int(ball.y)+ball.radius > winHeight {
		ball.yv = -ball.yv
	}

	if int(ball.x)+ball.radius < 0 || int(ball.x)+ball.radius > winWidth {
		ball.x = 400
		ball.y = 300
	}

	if ball.x-ball.radius/2 < leftPaddle.x+leftPaddle.w/2 &&
		ball.x+ball.radius/2 > leftPaddle.x-leftPaddle.w/2 &&
		ball.y+ball.radius >= leftPaddle.y-leftPaddle.h/2 &&
		ball.y-ball.radius <= leftPaddle.y+leftPaddle.h/2 {
		ball.xv = -ball.xv
	} else if ball.x+ball.radius/2 > rightPaddle.x-rightPaddle.w/2 &&
		ball.x-ball.radius/2 < rightPaddle.x+rightPaddle.w/2 &&
		ball.y+ball.radius >= rightPaddle.y-rightPaddle.h/2 &&
		ball.y-ball.radius <= rightPaddle.y+rightPaddle.h/2 {
		ball.xv = -ball.xv
	}
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

type paddle struct {
	pos
	w     int
	h     int
	color color
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= 7
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += 7
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	for i := 1; i < 8; i++ {
		if ball.y > paddle.y {
			paddle.y++
		} else if ball.y < paddle.y {
			paddle.y--
		}
	}
	// paddle.y = ball.y
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Executes this command just before program exits
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Executes this command just before program exits
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	// 4 bytes for each pixel, RGB Alpha
	pixels := make([]byte, winWidth*winHeight*4)

	player1 := paddle{pos{100, 100}, 20, 100, color{255, 255, 255}}
	player2 := paddle{pos{700, 100}, 20, 100, color{255, 255, 255}}
	ball := ball{pos{400, 300}, 12, 4, 4, color{255, 255, 255}}
	keyState := sdl.GetKeyboardState()

	// Check for events
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
				break
			}
		}
		clear(pixels)

		player1.update(keyState)
		player2.aiUpdate(&ball)
		ball.update(&player1, &player2)

		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(16)
	}

}
