package zigzag_conversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if numRows == 0 {
		return ""
	}

	columns := make([]*strings.Builder, numRows)
	goToDown := true
	rowIndex := 0
	for i, c := range s {
		if i < numRows {
			column := &strings.Builder{}
			columns[i] = column
		}

		columns[rowIndex].WriteRune(c)

		if (goToDown && rowIndex == numRows-1) || ((!goToDown) && rowIndex == 0) {
			goToDown = !goToDown
		}

		if goToDown {
			rowIndex++
		} else {
			rowIndex--
		}
	}

	var answer strings.Builder
	for _, column := range columns {
		if column != nil {
			answer.WriteString(column.String())
		}
	}

	return answer.String()
}
