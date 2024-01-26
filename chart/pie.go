package chart

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type PieChartItem struct {
	Name   string
	Values float64
}

type PieChartOpts struct {
	Output    string
	Title     string
	Subtitle  string
	Radius    float64
	Format    string
	ShowLabel bool
	Items     []PieChartItem
}

func generatePieItems(ele []PieChartItem) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < len(ele); i++ {
		items = append(items, opts.PieData{
			Name:  ele[i].Name,
			Value: ele[i].Values,
		})
	}
	return items
}

func CreatePieChart(options *PieChartOpts) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    options.Title,
				Subtitle: options.Subtitle,
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems(options.Items)).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: options.Radius,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      options.ShowLabel,
					Formatter: options.Format,
				},
			),
		)
	f, _ := os.Create(options.Output)
	_ = pie.Render(f)
}
