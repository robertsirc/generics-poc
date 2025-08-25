package main

import (
	"fmt"
	v2 "generics-poc/charts/v2"
	v3 "generics-poc/charts/v3"
	"generics-poc/cmd"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	fmt.Println("Testing Go Generics for a poc")

	chart2 := v2.Chart{
		Metadata: nil,
		Version:  "1.1",
	}

	chart3 := v3.Chart{
		Metadata: nil,
		Version:  1.1,
	}

	cmd.ProcessChart(chart2, true)
	cmd.ProcessChart(chart3, true)
}
