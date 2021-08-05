package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.RGBA{
	{
		uint8(244),
		uint8(67),
		uint8(54),
		uint8(255),
	},
	{
		uint8(156),
		uint8(39),
		uint8(176),
		uint8(255),
	},
	{
		uint8(63),
		uint8(81),
		uint8(181),
		uint8(255),
	},
	{
		uint8(3),
		uint8(169),
		uint8(244),
		uint8(255),
	},
	{
		uint8(0),
		uint8(150),
		uint8(136),
		uint8(255),
	},
	{
		uint8(139),
		uint8(195),
		uint8(74),
		uint8(255),
	},
	{
		uint8(255),
		uint8(235),
		uint8(59),
		uint8(255),
	},
	{
		uint8(255),
		uint8(152),
		uint8(0),
		uint8(255),
	},
	{
		uint8(121),
		uint8(85),
		uint8(72),
		uint8(255),
	},
	{
		uint8(96),
		uint8(125),
		uint8(139),
		uint8(255),
	},
	{
		uint8(233),
		uint8(30),
		uint8(99),
		uint8(255),
	},
	{
		uint8(103),
		uint8(58),
		uint8(183),
		uint8(255),
	},
	{
		uint8(33),
		uint8(150),
		uint8(243),
		uint8(255),
	},
	{
		uint8(0),
		uint8(188),
		uint8(212),
		uint8(255),
	},
	{
		uint8(76),
		uint8(175),
		uint8(80),
		uint8(255),
	},
	{
		uint8(205),
		uint8(220),
		uint8(57),
		uint8(255),
	},
	{
		uint8(255),
		uint8(193),
		uint8(7),
		uint8(255),
	},
	{
		uint8(255),
		uint8(87),
		uint8(34),
		uint8(255),
	},
	{
		uint8(158),
		uint8(158),
		uint8(158),
		uint8(255),
	},
}

type point struct {
	x, y int
	c    color.RGBA
}

func drawFireworksInTheFog() {
	sizeX := 800
	sizeY := 800

	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	pc := 10
	points := make([]point, pc)
	thr := float64(450)

	// place circles evenly
	// 	timeout := 20000
	// loop:
	for i := 0; i < pc; i++ {
		bx, rx, by, ry := randomQuadrants(i, sizeX, sizeY)
		points[i] = point{x: rnd.Intn(rx) + bx, y: rnd.Intn(ry) + by, c: palette[rnd.Intn(len(palette))]}
	}

	// 	for x := 0; x < sizeX; x += 10 {
	// 		for y := 0; y < sizeY; y += 10 {
	// 			l := true
	// 			for i := 0; i < pc; i++ {
	// 				d := math.Sqrt(math.Pow(float64(x)-float64(points[i].x), 2) + math.Pow(float64(y)-float64(points[i].y), 2))
	// 				if d < thr {
	// 					l = false
	// 					break
	// 				}
	// 			}
	// 			if l {
	// 				if timeout--; timeout == 0 {
	// 					log.Fatalf("gave up")
	// 				}
	// 				goto loop
	// 			}
	// 		}
	// 	}

	img := image.NewRGBA(image.Rect(0, 0, sizeX-1, sizeY-1))

	// paint png
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			distances := []float64{}
			dsum := float64(0)
			for i := 0; i < pc; i++ {
				d := math.Sqrt(math.Pow(float64(x)-float64(points[i].x), 2) + math.Pow(float64(y)-float64(points[i].y), 2))
				distances = append(distances, d+float64(0.1))
				if d <= thr-1 {
					dsum += (float64(thr) - d)
				}
			}

			// aggregate color
			r, g, b := float64(0), float64(0), float64(0)
			for i := range distances {
				if distances[i] <= thr-1 {
					aux := float64(points[i].c.R) * ((float64(thr) - distances[i]) / dsum)
					r += math.Pow((float64(thr)-distances[i]), 2) / float64(thr/3) / float64(thr) * aux

					aux = float64(points[i].c.G) * ((float64(thr) - distances[i]) / dsum)
					g += math.Pow((float64(thr)-distances[i]), 2) / float64(thr/3) / float64(thr) * aux

					aux = float64(points[i].c.B) * ((float64(thr) - distances[i]) / dsum)
					b += math.Pow((float64(thr)-distances[i]), 2) / float64(thr/3) / float64(thr) * aux

					if r > 255 {
						r = 255
					}
					if g > 255 {
						g = 255
					}
					if b > 255 {
						b = 255
					}
				}
			}

			// //calculate distance to palette
			// pd := make([]float64, len(palette))
			// for i := range palette {
			// 	pd[i] = math.Sqrt(math.Pow(float64(palette[i].R)-r, 2) + math.Pow(float64(palette[i].G)-g, 2) + math.Pow(float64(palette[i].B)-b, 2))
			// }

			// //calculate closest distance to palette
			// min := float64(999999)
			// minI := -1
			// for i := range pd {
			// 	if pd[i] < min {
			// 		minI = i
			// 	}
			// }

			// // move color % closer to closest palette item
			// pct := 0.02
			// dr, dg, db := r-float64(palette[minI].R), g-float64(palette[minI].G), b-float64(palette[minI].B)
			// r, g, b = r+dr*pct, g+dg*pct, b+db*pct

			img.SetRGBA(
				x,
				y,
				color.RGBA{
					uint8(r),
					uint8(g),
					uint8(b),
					uint8(255),
				})
		}
	}

	f, err := os.Create("static/sample.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	png.Encode(w, img)
	w.Flush()
}

func randomQuadrants(i, maxx, maxy int) (int, int, int, int) {
	switch math.Mod(float64(i), float64(4)) {
	case 0:
		return 0 + maxx/16, maxx/2 - maxx/10, 0 + maxx/16, maxy/2 - maxx/10
	case 1:
		return 0 + maxx/16, maxx/2 - maxx/10, maxy/2 + maxx/16, maxy/2 - maxx/10
	case 2:
		return maxx/2 + maxx/16, maxx/2 - maxx/10, 0 + maxx/16, maxy/2 - maxx/10
	default:
		return maxx/2 + maxx/16, maxx/2 - maxx/10, maxy/2 + maxx/16, maxy/2 - maxx/10
	}
}
