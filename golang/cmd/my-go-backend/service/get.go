package service

import (
	"database/sql"
	"fmt"
	"strings"

	"my-go-backend/db"
)

type GetParams struct {
	ServicePath string
	Table       string
	Id          int
	Columns     []string
}

type GetResult struct {
	Item *sql.Row
}

func Get(params GetParams) (*GetResult, error) {
	fmt.Println("Running Query")
	columnList := strings.Join(params.Columns, ", ")
	selectString := fmt.Sprintf("SELECT %s FROM %s WHERE id = %d", columnList, params.Table, params.Id)
	fmt.Println(selectString)
	row := db.GetDB().QueryRow(selectString)

	res := GetResult{
		Item: row,
	}
	return &res, nil
}
