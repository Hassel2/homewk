package main

import (
	"encoding/json"
	"flag"
	"os"
)

func main() {
    input := flag.String("i", "", "Input json file")
    output := flag.String("o", "timeline.png", "Name of output image")
    flag.Parse()

    byteValue, _ := os.ReadFile(*input)

    var statistic Data

    json.Unmarshal(byteValue, &statistic)

    statistic.DrawLines(*output)

    statisticSmoothed := movingAverage(statistic)
    statisticSmoothed.DrawLines("smoothed_"+*output)

    statisticVerySmoothed := movingAverage(statistic)
    statisticVerySmoothed.DrawLines("very_smoothed_"+*output)

    drawEach(statistic, statisticSmoothed, statisticVerySmoothed)
}

