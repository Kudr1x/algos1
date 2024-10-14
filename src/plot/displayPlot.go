package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	"image"
	"image/png"
	"os"
)

func DisplayPlot(p *plot.Plot, name string) {
	img := image.NewRGBA(image.Rect(0, 0, 1200, 1200))
	c := vgimg.NewWith(vgimg.UseImage(img))
	p.Draw(draw.New(c))

	err := os.MkdirAll("./image/", os.ModePerm)

	f, err := os.Create("./image/" + name + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, c.Image()); err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
}
