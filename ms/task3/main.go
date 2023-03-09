package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
    crosslights := [1000][7]int{}

    points := [13]int{}

    p := plot.New()

    for i := 0; i < 1000; i++ {
        crosslights[i][6] = 0
        for j := 0; j < 6; j++ {
            crosslights[i][j] = rand.Intn(3)
            crosslights[i][6] += crosslights[i][j]
        }
        points[crosslights[i][6]] += 1
    }

    gaussData := make(plotter.XYs, 13)
    for i := 0; i < 13; i++ {
        gaussData[i].Y = float64(points[i]) / 1000
        gaussData[i].X = float64(i)
    }

    gauss, err := plotter.NewLine(gaussData)

    if err != nil {
        panic(err)
    }

    p.Add(gauss)
    
	if err = p.Save(8*vg.Inch, 8*vg.Inch, "gauss.png"); err != nil {
		panic(err)
	}

    for i := 1; i < 13; i++ {
        gaussData[i].Y += gaussData[i-1].Y
    }

    second, err := plotter.NewLine(gaussData)
    if err != nil {
        panic(err)
    }
    s := plot.New()
    s.Add(second)

	if err = s.Save(8*vg.Inch, 8*vg.Inch, "second.png"); err != nil {
		panic(err)
	}
}
