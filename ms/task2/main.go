package main

import (
	"image/color"
	"math/rand"
    "math"
    "fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
    var n int = 5000
    points := randomPoints(n)

    p := plot.New()
    p.Add(plotter.NewGrid())

    s, err := plotter.NewScatter(points)
    if err != nil {
        panic(err)
    }

    s.Color = color.RGBA{R: 0, G: 0, B: 0, A: 255}
    s.Shape = draw.CircleGlyph{}
    
    var count int = 0
    sinPoints := make(plotter.XYs, n)
    for i := range sinPoints {
        if points[i].Y <= math.Sin(points[i].X) {
            sinPoints[i].X = points[i].X
            sinPoints[i].Y = points[i].Y
            count++
        } else {
            sinPoints[i].X = 0 
            sinPoints[i].Y = 0
        }
    }

    sin, err := plotter.NewScatter(sinPoints)
    if err != nil {
        panic(err)
    }
    sin.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
    sin.Shape = draw.CircleGlyph{}
    

    p.Add(s, sin)
	if err = p.Save(8*vg.Inch, 8*vg.Inch, "sin.png"); err != nil {
		panic(err)
	}

    fmt.Print(4*(count / 5000))
}

func randomPoints(n int) plotter.XYs {
    pts := make(plotter.XYs, n)
    for i := range pts {
        pts[i].X = -1 + rand.Float64() * 2
        pts[i].Y = -1 + rand.Float64() * 2
    }
    return pts
}
