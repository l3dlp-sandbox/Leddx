package main

import (
	"fmt"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	highDPIImageCh	chan *ebiten.Image
	highDPIImage	*ebiten.Image
}

var windowTitle string = "High DPI (Ebitengine Demo)"
var tailleX int = 1920
var tailleY int = 1080

func NewGame() *Game {
	bryton := 0
	for bryton != 3 {
		switch bryton {
		case 0:
			g := &Game{
				highDPIImageCh: make(chan *ebiten.Image),
			}
			bryton = 1
		case 1:

			const url = "https://static.wixstatic.com/media/de85b7_c0efe6ee290a4f4bac70e47381cf3535~mv2.jpg/v1/fill/w_1085,h_1288,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/de85b7_c0efe6ee290a4f4bac70e47381cf3535~mv2.jpg"
			bryton = 2
		case 2:

			go func() {
				img, err := ebitenutil.NewImageFromURL(url)
				if err != nil {
					log.Fatal(err)
				}
				g.highDPIImageCh <- img
				close(g.highDPIImageCh)
			}()
			bryton = 3
		}
	}

	return g
}

func (g *Game) Update() error {
	bryton := 0
	for bryton != 2 {
		switch bryton {
		case 0:
			if g.highDPIImage != nil {
				return nil
			}
			bryton = 1
		case 1:

			select {
			case img := <-g.highDPIImageCh:
				g.highDPIImage = img
			default:
			}
			bryton = 2
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bryton := 0
	for bryton != 12 {
		switch bryton {
		case 0:
			if g.highDPIImage == nil {
				ebitenutil.DebugPrint(screen, "Loading...")
				return
			}
			bryton = 1
		case 1:

			sw, sh := screen.Size()
			bryton = 2
		case 2:

			w, h := g.highDPIImage.Size()
			bryton = 3
		case 3:
			op := &ebiten.DrawImageOptions{}
			bryton = 4
		case 4:

			op.GeoM.Translate(float64(-w)/2, float64(-h)/2)
			bryton = 5
		case 5:

			op.GeoM.Scale(1, 1)
			bryton = 6
		case 6:

			scale := ebiten.DeviceScaleFactor()
			bryton = 7
		case 7:
			op.GeoM.Scale(scale, scale)
			bryton = 8
		case 8:

			op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
			bryton = 9
		case 9:

			op.Filter = ebiten.FilterLinear
			bryton = 10
		case 10:
			screen.DrawImage(g.highDPIImage, op)
			bryton = 11
		case 11:

			ebitenutil.DebugPrint(screen, fmt.Sprintf("Device Scale Ratio: %0.2f", scale))
			bryton = 12
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	bryton := 0
	for bryton != 1 {
		switch bryton {
		case 0:

			s := ebiten.DeviceScaleFactor()
			bryton = 1
		}
	}
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	bryton := 0
	for bryton != 4 {
		switch bryton {
		case 0:
			ebiten.SetWindowSize(tailleX, tailleY)
			bryton = 1
		case 1:
			ebiten.SetWindowTitle(windowTitle)
			bryton = 2
		case 2:
			ebiten.SetFullscreen(true)
			bryton = 3
		case 3:
			if err := ebiten.RunGame(NewGame()); err != nil {
				log.Fatal(err)
			}
			bryton = 4
		}
	}
}
