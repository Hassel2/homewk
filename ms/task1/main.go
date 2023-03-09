package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Задание 1
	result := [8][6]float64{}
	h := [8]int{10, 100, 500, 1000, 2000, 3000, 4000, 5000}

	buf := [6]float64{0, 0, 0, 0, 0, 0}
	for i := 1; i <= 5000; i++ {
		buf[rand.Intn(6)]++

		if pos := Contains(h, i); pos != -1 {
			for j := 0; j < 6; j++ {
				result[pos][j] = buf[j] / float64(i)
			}

		}
	}

	if err := MakePlots(result); err != nil {
		panic(err)
	}

	// Задание 2 вариант 1
	result1 := [8]float64{}

	var count float64 = 0
	for i := 1; i <= 5000; i++ {
		if (rand.Intn(6) % 2 == 0 && rand.Intn(6) % 2 == 0) {
			count++;
		}

		if pos := Contains(h, i); pos != -1 {
			result1[pos] = count / float64(i)
		}
	}

	if err := MakePlot(result1); err != nil {
		panic(err)
	}
}

func Contains(a [8]int, x int) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func MakePlot(arr [8]float64) error {
	p := plot.New()

	p.Title.Text = "Кубик задание 2 график"
	p.X.Label.Text = "Попытка"
	p.Y.Label.Text = "Вероятность"

	x := [8]float64{10, 100, 500, 1000, 2000, 3000, 4000, 5000}

	points := make(plotter.XYs, 8)

	for i := 0; i < 8; i++ {
		points[i].X = x[i]
		points[i].Y = arr[i]
	}

	err := plotutil.AddLinePoints(p,
		"Два кубика показывают четное одновременно", points)
	
	if err != nil {
		return err
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "second.png"); err != nil {
		return err
	}

	return nil
}

func MakePlots(arr [8][6]float64) error {
	p := plot.New()

	p.Title.Text = "Кубик задание 1 график"
	p.X.Label.Text = "Попытка"
	p.Y.Label.Text = "Вероятность"

	x := [8]float64{10, 100, 500, 1000, 2000, 3000, 4000, 5000}

	var points [6]plotter.XYs

	for i := 0; i < 6; i++ {
		points[i] = make(plotter.XYs, 8)
	}
	
	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			points[j][i].X = x[i]
			points[j][i].Y = arr[i][j]
		}
	}

	err := plotutil.AddLinePoints(p,
		"Один", points[0],
		"Два", points[1],
		"Три", points[2],
		"Четыре", points[3],
		"Пять", points[4],
		"Шесть", points[5])

	if err != nil {
		return err
	}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "first.png"); err != nil {
		return err
	}

	return nil
}
