package charts

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type BarConfiguration struct {
	*components.BaseConfiguration
	Series *series.BarSeries0 `json:"series,omitempty"`
	XAxis  *components.XAxis  `json:"xAxis,omitempty,reserved"`
	YAxis  *components.YAxis  `json:"yAxis,omitempty,reserved"`
}

// Bar represents a bar chart.
type Bar struct {
	*BarConfiguration
}

func (bar *Bar) GetChartName() string {
	return types.ChartBar
}

func (bar *Bar) GetChart() interface{} {
	return bar
}

func (bar *Bar) GetContainer() *core.Container {
	if bar.Container != nil {
		return bar.Container
	}

	bar.Container = core.NewContainer(bar)
	return bar.Container

}

func (bar *Bar) GetPage() *core.Page {
	if bar.Page != nil {
		return bar.Page
	}

	bar.Page = core.NewPage(bar.GetContainer())
	return bar.Page
}

func NewBar() *Bar {
	bar := &Bar{}

	bar.BarConfiguration = &BarConfiguration{
		BaseConfiguration: components.BaseConfiguration{}.New(),
		XAxis:             components.XAxis{}.New(),
		YAxis:             components.YAxis{}.New(),
		Series:            &series.BarSeries0{},
	}

	bar.Container = core.NewContainer(bar)
	bar.Page = core.NewPage(bar.Container)

	return bar
}

func (bar *Bar) AddSeries(series ...*series.BarSeries) {
	for _, s := range series {
		c := append(*bar.Series, s)
		bar.Series = &c
	}
}
