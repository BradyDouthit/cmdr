package UIHelpers

import (
	"github.com/gizak/termui/v3/widgets"
	Shell "trendify/utils/shell"
)

func BuildBarChart(commands []Shell.Command) *widgets.BarChart {
	var data []float64
	var labels []string

	for _, command := range commands {
		data = append(data, command.Index)
		labels = append(labels, command.Command)
	}

	bc := widgets.NewBarChart()
	bc.Title = "Common Commands"
	bc.Data = data
	bc.Labels = labels
	bc.SetRect(5, 5, 70, 36)
	bc.BarWidth = 10
	bc.BarGap = 5

	return bc
}
