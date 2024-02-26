# Line

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/chart"
)
```
## Functions


### SampleChartLine

```go
func SampleChartLine() {
	newChart := &chart.LineChartOpts{
		Output:   "line.html",
		Title:    "Line chart in Go",
		Subtitle: "This is fun to use!",
		XAxis:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		Items: []chart.LineChartItem{
			{Name: "Category A", Values: []float64{70, 200, 10, 300, 310, 900}},
			{Name: "Category B", Values: []float64{680, 290, 356, 434, 900, 100}},
		},
	}

	chart.CreateLineChart(newChart)
}
```
