package render

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Table(header []string, rows [][]string) {
	t := table.New().Width(getWidth(80)).StyleFunc(rowStyle)

	t.Headers(header...)
	t.Rows(addBreakToColumns(rows)...)

	fmt.Print(t)
}

func rowStyle(row, col int) lipgloss.Style {
	switch {
	case row == 0:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE11A"))
	case row%2 == 0:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#A4C3E4"))
	default:
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#AF9BEA"))
	}
}

func addBreakToColumns(rows [][]string) [][]string {
	width := getWidth(80) / 3

	colStyle := lipgloss.NewStyle().MaxWidth(width)

	for i, row := range rows {
		for index, col := range row {
			row[index] = colStyle.Render(breakLines(col, width))
		}

		rows[i] = row
	}

	return rows
}

func breakLines(input string, count int) string {
	var out = []rune{}

	for i, char := range input {
		out = append(out, char)

		if i%count == 0 && i != 0 {
			out = append(out, '\n')
		}
	}

	return string(out)
}
