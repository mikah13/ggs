package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func getTable(stocks []ChartResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Ticker", "Last Price", "Change", "Change %", "Previous Close", "Currency"})

	for _, stock := range stocks {
		row := getRow(stock)
		t.AppendRow(row)
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name: "Change",
			Transformer: text.Transformer(func(val interface{}) string {
				return text.Colors{text.FgRed}.Sprint(val)
			}),
		},
	})

	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}

func getRow(stock ChartResponse) table.Row {
	data := stock.Chart.Result[0].Meta
	diff := data.RegularMarketPrice - data.PreviousClose
	ticker := data.Symbol
	lastPrice := data.RegularMarketPrice
	change := appendPlus(diff)
	changePercent := appendPlus(diff / data.PreviousClose * 100)
	currency := data.Currency
	previousClose := data.PreviousClose

	return table.Row{ticker, lastPrice, change, changePercent, previousClose, currency}
}

func appendPlus(num float64) string {
	if num >= 0 {
		return fmt.Sprintf("+%.2f", num)
	}
	return fmt.Sprintf("%.2f", num)
}
