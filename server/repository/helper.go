package repository

import (
	"fmt"
	"strings"
)

func generateInsertQuery(table_name string, column_names []string, values [][]interface{}) (string, []interface{}) {
	colNames := "(" + strings.Join(column_names, ", ") + ")"

	valStrings := []string{}
	args := []interface{}{}
	argPos := 1
	for _, row := range values {
		placeholders := []string{}
		for _, v := range row {
			args = append(args, v)
			placeholders = append(placeholders, fmt.Sprintf("$%d", argPos))
			argPos++
		}
		valStrings = append(valStrings, "("+strings.Join(placeholders, ", ")+")")
	}

	return fmt.Sprintf("INSERT INTO %s %s VALUES %s RETURNING *", table_name, colNames, strings.Join(valStrings, ", ")), args

}
