package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
    var Ts float64 = 22     // Температура окружающей среды
    var dt float64 = 0.1    // Интервал времени
    var eps float64 = 0.3   // Погрешность


    var cup = coffeeCup {
        T: 83,
        k: -0.09,
        M: 0.5,
    }

    // Остывание чашки кофе в огромной комнате
    var onlyCup []float64 = cup.coolingProcess(Ts, dt, eps)
   
    onlyCupData := make(plotter.XYs, len(onlyCup))
    for i, cup := range onlyCup {
        onlyCupData[i] = plotter.XY{ X: float64(i), Y: cup }
    }
    drawLine([]plotter.XYs{onlyCupData}, "onlyCup.png")

    // Остывание чашки кофе в коробке
    var cupArr, boxArr []float64 = cup.coolingInABox(Ts, dt, eps)

    cupData := make(plotter.XYs, len(cupArr))
    boxData := make(plotter.XYs, len(boxArr))
    for i := range cupArr {
        cupData[i] = plotter.XY{ X: float64(i), Y: cupArr[i] }
        boxData[i] = plotter.XY{ X: float64(i), Y: boxArr[i] }
    }

    drawLine([]plotter.XYs{cupData, boxData}, "cupAndBox.png")
    
    // Разбавление чашки кофе водой
    var T_desired float64 = 40 // Желаемая температура кофе

    var water = liquid {
        M: 0.025,
        T: 85,
    }
   
    var cupAndWater []float64 = cup.diluteProcess(T_desired, Ts, dt, eps, water)
    cupAndWaterData := make(plotter.XYs, len(cupAndWater))
    for i, cup := range cupAndWater {
        cupAndWaterData[i] = plotter.XY{ X: float64(i), Y: cup }
    }
    drawLine([]plotter.XYs{cupAndWaterData}, "cupAndWater.png")

    // Разбавление чашки кофе молоком
    var milk = liquid {
        M: 0.025,
        T: 4,
    }
    var cupAndMilk []float64 = cup.diluteProcess(T_desired, Ts, dt, eps, milk)
    cupAndMilkData := make(plotter.XYs, len(cupAndMilk))
    for i, cup := range cupAndMilk {
        cupAndMilkData[i] = plotter.XY{ X: float64(i), Y: cup }
    }
    drawLine([]plotter.XYs{cupAndMilkData}, "cupAndMilk.png")
}

func drawLine(data[] plotter.XYs, name string) error {
    plot := plot.New()
    plot.Title.Text = name
    plot.X.Label.Text = "Количество шагов"
    plot.Y.Label.Text = "Температура"
    plot.X.Tick.Length = 25

    for _, d := range data {
        line, err := plotter.NewLine(d)
        if err != nil {
            return err;
        }
        
        plot.Add(line)
    }

    if err := plot.Save(8*vg.Inch, 8*vg.Inch, name); err != nil {
        return err
    }

    return nil
}

type liquid struct {
    M float64   // Масса жидкости
    T float64   // Температура жидкости
}

type coffeeCup struct {
    T float64   // Начальная температура кофе
    k float64   // Коэффициент скорости остывания кофе
    M float64   // Масса кофе в кружке
}

func (cc *coffeeCup) coolingFormula(Tc float64, Ts float64, dt float64) float64 {
    return Tc + cc.k * (Tc - Ts) * dt
}

func (cc *coffeeCup) coolingProcess(T float64, dt float64, eps float64) []float64 {
    var result = make([]float64, 0)
    result = append(result, cc.T)

    last := result[len(result) - 1]
    for ; last - T > eps; {
        result = append(result, cc.coolingFormula(last, T, dt))
        last = result[len(result) - 1]
    }

    return result
}

func (cc *coffeeCup) coolingInABox(T float64, dt float64, eps float64) ([]float64, []float64) {
    var cupData = make([]float64, 0)
    var boxData = make([]float64, 0)

    cupData = append(cupData, cc.T)
    boxData = append(boxData, T)

    var lastCup = cupData[len(cupData) - 1]
    var lastBox = boxData[len(boxData) - 1]
    for lastCup - lastBox > eps {
        cupData = append(cupData, cc.coolingFormula(lastCup, lastBox, dt))
        boxData = append(boxData, lastBox - (cupData[len(cupData) - 1] - cupData[len(cupData) - 2]))
        
        lastCup = cupData[len(cupData) - 1]
        lastBox = boxData[len(boxData) - 1]
    }

    return cupData, boxData
}

func (cc *coffeeCup) diluteProcess(T_desired float64, T float64, dt float64, eps float64, liq liquid) []float64 {
    var coffeM = cc.M
    var result = make([]float64, 0)
    result = append(result, cc.T)

    last := result[len(result) - 1]
    for i := 1; last - T_desired > eps; i++ {
        curCofTemp := cc.coolingFormula(last, T, dt)
        if i % 20 != 0 {
            result = append(result, curCofTemp)
        } else {
            result = append(result, ((cc.M * curCofTemp) + (liq.M * liq.T)) / (liq.M + cc.M))
            cc.M += liq.M
        }
        last = result[len(result) - 1]
    }
    
    cc.M = coffeM
    return result
}
