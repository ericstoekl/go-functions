package function

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/png"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"
)

func drawLine(dc *gg.Context, x1, y1, x2, y2 int) {
	dc.SetRGBA(0, 0.4, 0.7, 1)
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
	params := os.Getenv("Http_Query")
	m, err := url.ParseQuery(params)
	if err != nil {
		fmt.Printf("parsequery error: %v\n", err)
		return fmt.Sprintf("parsequery error: %v\n", err)
	}
	cells := 16
	if len(m["V"]) > 0 {
		cells, _ = strconv.Atoi(m["V"][0])
	} else {
		strInput := string(req)
		cells, _ = strconv.Atoi(strInput)
	}
	if cells == 0 {
		cells = 16
	}

	rand.Seed(time.Now().UnixNano())
	const W = 1024
	const H = 1024
	distance := W / cells

	dc := gg.NewContext(W+(distance*2), H+(distance*2))
	dc.SetLineWidth(5.0)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	buildWalls(dc, W, H, distance)
	putLines(dc, W, H, distance)

	pngImg := dc.Image()
	png.Encode(os.Stdout, pngImg)
	return ""
}
