package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"go-hep.org/x/hep/fit"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type Data struct {
    Legend string `json:"legend"`
    FirstX float64 `json:"firstx"`
    Step float64 `json:"step"`
    Data []float64 `json:"data"`
    x []float64
}

func (d *Data) makeX() {
    d.x = make([]float64, len(d.Data))
    for i := d.FirstX; int(float64(i) - d.FirstX) / int(d.Step) < len(d.Data); i += d.Step {
        d.x[int(i - d.FirstX) / int(d.Step)] = i 
    }
}

func (d *Data) DrawLines(output string) {
    data := make(plotter.XYs, len(d.Data))
    if (len(d.x) == 0) {
        d.makeX()
    }

    for i, dat := range d.Data {
        data[i] = plotter.XY{X: d.x[i], Y: dat}
    }

    line, err := plotter.NewLine(data)
    if err != nil {
        panic(err)
    }

    linearRegression := d.linearRegression()
    exponentialRegression := d.exponentialRegression()
    polynomialRegression := d.polynomialRegression()
    predict := d.predict()

    plot := plot.New()
    plot.Add(line, &linearRegression, &exponentialRegression, &polynomialRegression, &predict)
    plot.Legend.Add(d.Legend, line)
    plot.Legend.Add("Linear regression", &linearRegression)
    plot.Legend.Add("Exponential regression", &exponentialRegression)
    plot.Legend.Add("Polynomial regression", &polynomialRegression)
    plot.Legend.Add("Prediction", &predict)


    if err := plot.Save(8*vg.Inch, 8*vg.Inch, "img/" + output + ".png"); err != nil {
        panic(err)
    }
}

func (d *Data) linearRegression(col ...color.RGBA) plotter.Function {
    colour := color.RGBA{G: 255, A: 255}
    if len(col) > 0 {
       colour = col[0] 
    }

    b, a := stat.LinearRegression(d.x, d.Data, nil, false)
    
    function := plotter.NewFunction(func(x float64) float64 { return a*x + b })
    function.LineStyle.Color = colour
    function.Width = vg.Points(3)
    function.Dashes = []vg.Length{vg.Points(4), vg.Points(4)}

    return *function
}

func (d *Data) exponentialRegression(col ...color.RGBA) plotter.Function {
    colour := color.RGBA{R: 255, A: 255}
    if len(col) > 0 {
       colour = col[0] 
    }

    res, err := fit.Curve1D(
        fit.Func1D{
            F: func(x float64, ps []float64) float64 {
                return math.Exp(ps[0]*x + ps[1])
            },
            X: d.x,
            Y: d.Data,
            N: 2,
        },
        nil, &optimize.NelderMead{},
    )
    if err != nil {
        panic(err)
    }

    function := plotter.NewFunction(func(x float64) float64 {
        return math.Exp(res.X[0]*x + res.X[1])
    })
    function.LineStyle.Color = colour
    function.Width = vg.Points(3)
    function.Dashes = []vg.Length{vg.Points(4), vg.Points(4)}

    return *function
}

func (d *Data) polynomialRegression(col ...color.RGBA) plotter.Function {
    colour := color.RGBA{B: 255, A: 255}
    if len(col) > 0 {
       colour = col[0] 
    }

    res, err := fit.Curve1D(
		fit.Func1D{
            F:  func(x float64, ps []float64) float64{
                return ps[0] + ps[1]*x*x
            },
			X:  d.x,
			Y:  d.Data,
			Ps: []float64{1, 1},
		},
		nil, &optimize.NelderMead{},
	)

    if err != nil {
        panic(err)
    }

    function := plotter.NewFunction(func(x float64) float64 {
        return res.X[0] + res.X[1]*x*x
    })
    function.LineStyle.Color = colour
    function.Width = vg.Points(3)
    function.Dashes = []vg.Length{vg.Points(4), vg.Points(4)}

    return *function
}

func (d *Data) predict() plotter.Scatter {
    data := make(plotter.XYs, 5)
    function := d.linearRegression()
    for i := 0; i < 5; i++ {
        data[i].X = d.x[len(d.x) - 1] + 1 + float64(i)
        data[i].Y = function.F(data[i].X)
    }

    res, error := plotter.NewScatter(data)
    if error != nil {
        panic(error)
    }

    res.GlyphStyle = draw.GlyphStyle{
        Radius: vg.Points(5),
        Color: color.RGBA{R: 120, G: 120, B: 120, A: 255},
        Shape: draw.PyramidGlyph{},
    }

    return *res
}

func drawEach(data ...Data) {
    plot := plot.New()

    for i, d := range data {
        lineData := make(plotter.XYs, len(d.Data))
        if (len(d.x) == 0) {
            d.makeX()
        }

        for i, dat := range d.Data {
            lineData[i] = plotter.XY{X: d.x[i], Y: dat}
        }

        line, err := plotter.NewLine(lineData)
        if err != nil {
            panic(err)
        }

        line.Color = color.RGBA{
            R: uint8(255 * rand.Float64()), 
            G: uint8(255 * rand.Float64()), 
            B: uint8(255 * rand.Float64()), 
            A: 255,
        }

        linearRegression := d.linearRegression(
            color.RGBA{
                R: uint8(255 * rand.Float64()), 
                G: uint8(255 * rand.Float64()), 
                B: uint8(255 * rand.Float64()), 
                A: 255,
            },
        )
        exponentialRegression := d.exponentialRegression(
            color.RGBA{
                R: uint8(255 * rand.Float64()), 
                G: uint8(255 * rand.Float64()), 
                B: uint8(255 * rand.Float64()), 
                A: 255,
            },
        )
        polynomialRegression := d.polynomialRegression(
            color.RGBA{
                R: uint8(255 * rand.Float64()), 
                G: uint8(255 * rand.Float64()), 
                B: uint8(255 * rand.Float64()), 
                A: 255,
            },
        )

        plot.Add(line, &linearRegression, &exponentialRegression, &polynomialRegression)
        plot.Legend.Add(fmt.Sprint(i), line)
        plot.Legend.Add(fmt.Sprint(i) + ") Linear regression", &linearRegression)
        plot.Legend.Add(fmt.Sprint(i) + ") Exponential regression", &exponentialRegression)
        plot.Legend.Add(fmt.Sprint(i) + ") Polynomial regression", &polynomialRegression)
    }

    if err := plot.Save(8*vg.Inch, 8*vg.Inch, "img/all.png"); err != nil {
        panic(err)
    }
}
