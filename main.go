package main

import (
	v2 "generics-poc/charts/v2"
	v3 "generics-poc/charts/v3"
	"generics-poc/cmd"
)

func main() {

	chart2 := v2.Chart{
		Metadata: nil,
		Version:  "1.1",
	}

	chart3 := v3.Chart{
		Metadata: nil,
		Version:  2.1,
	}

	cmd.ProcessChart(chart2, true)
	cmd.ProcessChart(chart3, true)
}
