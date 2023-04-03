package main

func movingAverage(data Data) Data {
    result := make([]float64, len(data.Data) - 2)
    for i := 1; i < len(data.Data) - 1; i++ {
        result[i - 1] = (data.Data[i - 1] + data.Data[i] + data.Data[i + 1]) / 3
    }

    return Data{ Step: data.Step, Legend: data.Legend, x: data.x[:len(result)], Data: result }
}

func movingAverageImproved(data Data) Data {
    result := make([]float64, len(data.Data) - 2)
    result[0] = (data.Data[0] + data.Data[1]) / 2

    for i := 1; i < len(data.Data) - 2; i++ {
        result[i] = (result[i - 1] + data.Data[i] + data.Data[i + 1]) / 3
    }

    return Data{ Step: data.Step, Legend: data.Legend, x: data.x[:len(result)], Data: result }
}
