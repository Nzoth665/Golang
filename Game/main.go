package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// В игре реализован ebiten.Интерфейс игры.
type Game struct{ rs float64 }

// Обновление продолжается в соответствии с состоянием игры.
// Обновление вызывается каждый такт (по умолчанию 1/60 [с]).
func (g *Game) Update() error {
	g.rs += 0.05
	ebiten.SetWindowTitle(strconv.Itoa(int(g.rs)))
	return nil
}

// Draw рисует игровой экран.
// Draw вызывается каждый кадр (обычно 1/60[с] для дисплея с частотой 60 Гц).
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, 160, 120, g.rs, color.RGBA{255, 255, 255, 255})
}

// Layout принимает внешний размер (например, размер окна) и возвращает (логический) размер экрана.
// Если вам не нужно настраивать размер экрана в соответствии с внешним размером, просто верните фиксированный размер.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
