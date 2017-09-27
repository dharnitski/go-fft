package main

import (
	"bufio"
	"math"
	"os"

	"github.com/mjibson/go-dsp/spectral"
	"github.com/wcharczuk/go-chart"
)

func main() {

	input := make([]float64, 65536)
	for i := range input {
		input[i] = math.Sin(float64(i)) + 2*math.Sin(float64(i)*3) + 3*math.Sin(float64(i)*4)
	}
	p, freqs := spectral.Pwelch(input, 2, &spectral.PwelchOptions{
		NFFT: 65536,
	})

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: freqs,
				YValues: p,
			},
		},
	}

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	err = graph.Render(chart.PNG, w)
	if err != nil {
		panic(err)
	}

	w.Flush()
}
