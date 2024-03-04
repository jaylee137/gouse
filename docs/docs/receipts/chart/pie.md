# Pie

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/chart"
)
```
## Functions


### SampleChartPie

```go
func SampleChartPie() {
	newChart := &chart.PieChartOpts{
		Output:    "pie.html",
		Title:     "Pie chart in Go",
		Subtitle:  "This is fun to use!",
		Radius:    200,
		Format:    "{b}: {c} ({d}%)",
		ShowLabel: true,
		Items: []chart.PieChartItem{
			{Name: "Category A", Values: 335},
			{Name: "Category B", Values: 310},
			{Name: "Category C", Values: 234},
			{Name: "Category D", Values: 135},
			{Name: "Category E", Values: 1548},
			{Name: "Category F", Values: 1548},
		},
	}

	chart.CreatePieChart(newChart)
}
```
