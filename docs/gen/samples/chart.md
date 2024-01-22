# Chart

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/chart"
)
```
## Functions


### chartBar

```go
func chartBar() {
	newChart := &chart.BarChartOpts{
		Output:   "bar.html",
		Title:    "Bar chart in Go",
		Subtitle: "This is fun to use!",
		XAxis:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"},
		Items: []chart.BarChartItem{
			{Name: "Category A", Values: []float64{100, 200, 300, 400, 500, 600}},
			{Name: "Category B", Values: []float64{200, 300, 400, 500, 600, 700}},
			{Name: "Category C", Values: []float64{300, 400, 500, 600, 700, 800}},
		},
	}

	chart.CreateBarChart(newChart)
}
```

### chartLine

```go
func chartLine() {
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
