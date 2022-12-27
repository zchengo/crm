package common

import (
	"crm/global"
	"log"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const (
	letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenExcelFile[T any](sheet string, columns []string, rowValue []T, filename string) (string, error) {
	f := excelize.NewFile()
	index := f.NewSheet(sheet)
	for i := 0; i < len(columns); i++ {
		axis := string(letter[i]) + strconv.Itoa(1)
		f.SetCellValue(sheet, axis, columns[i])
	}

	num := 2
	for r := 0; r < len(rowValue); r++ {
		v := reflect.ValueOf(rowValue[r])
		for i := 0; i < v.NumField(); i++ {
			axis := string(letter[i]) + strconv.Itoa(num)
			f.SetCellValue(sheet, axis, v.Field(i))
		}
		num++
	}
	f.SetActiveSheet(index)
	path := global.Config.File.Path + filename + ".xlsx"

	if err := f.SaveAs(path); err != nil {
		log.Printf("[ERROR] file save error: %s", err)
		return "", err
	}
	return path, nil
}
