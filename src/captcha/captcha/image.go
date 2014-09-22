package captcha

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gographics/imagick/imagick"
)

var (
	// random colors
	colors []string = []string{"#000000", "#b50000", "#373000", "#827482"}
	fonts  []string
)

func child(path string) []string {
	fullPath, _ := filepath.Abs(path)
	listStr := make([]string, 0)
	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".ttf") {
			listStr = append(listStr, path)
		}
		return err
	})
	return listStr
}

func init() {
	// get all fonts file here
	fonts = child("../assets/fonts/")
}

func randFont() string {
	rand.Seed(int64(time.Now().Nanosecond()))
	return fonts[rand.Intn(len(fonts))]
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}

func drawSetfont(mw *imagick.MagickWand, dw *imagick.DrawingWand) {
	dw.SetFont(randFont())
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	pw.SetColor(colors[rand.Intn(len(colors))])
	dw.SetFontWeight(500)
	dw.SetFillColor(pw)
	dw.SetFontSize(33)

}

func drawMetrics(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, text string) {
	rotation := float64(random(-30, 30))
	mw.AnnotateImage(dw, *dx, 35, rotation, text)
	mw.DrawImage(dw)
	// get the font metrics
	fm := mw.QueryFontMetrics(dw, text)
	if fm != nil {
		// Adjust the new x coordinate
		*dx += fm.TextWidth + 2
	}
}

func writeWord(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, text string) {
	drawSetfont(mw, dw)
	drawMetrics(mw, dw, dx, text)
}

func Draw(text string, name string) {
	imagick.Initialize()
	defer imagick.Terminate()
	// Current coordinates of text
	var dx float64 = 20

	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	defer mw.Destroy()
	defer dw.Destroy()

	// Set the size of the image
	mw.SetSize(285, 50)
	mw.ReadImage("../assets/images/bg.gif")

	// Start near the left edge
	dw.SetFontSize(40)
	// Note that we must free up the fontmetric array once we're done with it
	list := strings.Split(text, " ")
	for _, item := range list {
		writeWord(mw, dw, &dx, item)
	}
	mw.DrawImage(dw)
	// Now write the magickwand image
	mw.WriteImage(name)
}
