package main

import (
	"fmt"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	highDPIImageCh chan *ebiten.Image
	highDPIImage   *ebiten.Image
}

var windowTitle string = "High DPI (Ebitengine Demo)"
var tailleX int = 1920
var tailleY int = 1080

func NewGame() *Game {
	g := &Game{
		highDPIImageCh: make(chan *ebiten.Image),
	}

	// const url = "https://upload.wikimedia.org/wikipedia/commons/1/1f/As08-16-2593.jpg"
	const url = "https://static.wixstatic.com/media/de85b7_c0efe6ee290a4f4bac70e47381cf3535~mv2.jpg/v1/fill/w_1085,h_1288,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/de85b7_c0efe6ee290a4f4bac70e47381cf3535~mv2.jpg"

	// Load the image asynchronously.
	go func() {
		img, err := ebitenutil.NewImageFromURL(url)
		if err != nil {
			log.Fatal(err)
		}
		g.highDPIImageCh <- img
		close(g.highDPIImageCh)
	}()

	return g
}

func (g *Game) Update() error {
	if g.highDPIImage != nil {
		return nil
	}

	// Use select and 'default' clause for non-blocking receiving.
	select {
	case img := <-g.highDPIImageCh:
		g.highDPIImage = img
	default:
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.highDPIImage == nil {
		ebitenutil.DebugPrint(screen, "Loading...")
		return
	}

	sw, sh := screen.Size()

	w, h := g.highDPIImage.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the images's center to the upper left corner.
	op.GeoM.Translate(float64(-w)/2, float64(-h)/2)

	// The image is just too big. Adjust the scale.
	op.GeoM.Scale(1, 1)

	// Scale the image by the device ratio so that the rendering result can be same
	// on various (different-DPI) environments.
	scale := ebiten.DeviceScaleFactor()
	op.GeoM.Scale(scale, scale)

	// Move the image's center to the screen's center.
	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)

	op.Filter = ebiten.FilterLinear
	screen.DrawImage(g.highDPIImage, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Device Scale Ratio: %0.2f", scale))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// The unit of outsideWidth/Height is device-independent pixels.
	// By multiplying them by the device scale factor, we can get a hi-DPI screen size.
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	ebiten.SetWindowSize(tailleX, tailleY)
	ebiten.SetWindowTitle(windowTitle)
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
