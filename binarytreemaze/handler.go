package function

import (
	"github.com/fogleman/gg"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func drawLine(dc *gg.Context, x1, y1, x2, y2 int) {
	dc.SetRGBA(0, 0.4, 0.7, 1)
	dc.SetLineWidth(8.0)
	dc.DrawLine(float64(x1), float64(y1), float64(x2), float64(y2))
	dc.Stroke()
}

func drawSegmentNorth(dc *gg.Context, w, h, distance int) {
	x1 := w
	x2 := w + distance
	y1 := h + distance
	y2 := h + distance
	drawLine(dc, x1, y1, x2, y2)
}
func drawSegmentEast(dc *gg.Context, w, h, distance int) {
	x1 := w + distance
	x2 := w + distance
	y1 := h
	y2 := h + distance
	drawLine(dc, x1, y1, x2, y2)
}

func putLines(dc *gg.Context, W, H, distance int) {
	for w := distance; w < W; w += distance {
		for h := distance; h < H; h += distance {
			randInt := rand.Int() % 2
			if randInt == 0 {
				drawSegmentNorth(dc, w, h, distance)
			} else {
				drawSegmentEast(dc, w, h, distance)
			}
		}
	}
}

func buildWalls(dc *gg.Context, W, H, distance int) {
	drawLine(dc, distance, distance, W+distance, distance)
	drawLine(dc, distance, distance, distance, H+distance)
	drawLine(dc, distance, H+distance, W+distance, H+distance)
	drawLine(dc, W+distance, distance, W+distance, H+distance)
}

func Handle(req []byte) string {
	rand.Seed(time.Now().UnixNano())
	const W = 1024
	const H = 1024
	const cells = 16
	distance := W / cells

	dc := gg.NewContext(W+(distance*2), H+(distance*2))
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	buildWalls(dc, W, H, distance)
	putLines(dc, W, H, distance)

	pngImg := dc.Image()
	png.Encode(os.Stdout, pngImg)
	return ""
}
