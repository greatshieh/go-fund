package utils

import (
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

type Writer struct {
	WorkBook     *excelize.File
	StreamWriter *excelize.StreamWriter
}

func structVal2interface(elemVal reflect.Value) (val []interface{}) {
	for i := 0; i < elemVal.NumField(); i++ {
		if elemVal.Field(i).Type().Kind().String() == "struct" {
			val = append(val, structVal2interface(elemVal.Field(i))...)
		} else {
			if elemVal.Field(i).IsValid() {
				v := elemVal.Field(i).Interface()
				val = append(val, v)
				// if v != "" {
				// 	val = append(val, elemVal.Field(i).Interface())
				// } else {
				// 	val = append(val, "--")
				// }
			} else {
				val = append(val, "")
			}
		}
	}
	return val
}

func conver2interface(content interface{}) (val []interface{}) {
	elemVal := reflect.ValueOf(content).Elem()
	val = structVal2interface(elemVal)
	// for i := 0; i < elemVal.NumField(); i++ {
	// 	if elemVal.Field(i).Type().Kind().String() == "struct" {
	// 		for j := 0; j < elemVal.Field(i).NumField(); j++ {
	// 			if elemVal.Field(i).Field(j).IsValid() {
	// 				v := elemVal.Field(i).Field(j).Interface().(string)
	// 				if v != "" {
	// 					val = append(val, elemVal.Field(i).Field(j).Interface())
	// 				} else {
	// 					val = append(val, "--")
	// 				}
	// 			} else {
	// 				val = append(val, "--")
	// 			}
	// 		}
	// 	} else {
	// 		if elemVal.Field(i).IsValid() {
	// 			v := elemVal.Field(i).Interface().(string)
	// 			if v != "" {
	// 				val = append(val, elemVal.Field(i).Interface())
	// 			} else {
	// 				val = append(val, "--")
	// 			}
	// 		} else {
	// 			val = append(val, "--")
	// 		}
	// 	}
	// }
	return val
}

func createHeader(model interface{}) []interface{} {
	typeVal := reflect.TypeOf(model)

	var headers []interface{}

	for i := 0; i < typeVal.NumField(); i++ {
		field := typeVal.Field(i)
		if field.Type.Kind().String() == "struct" {
			for j := 0; j < field.Type.NumField(); j++ {
				headers = append(headers, field.Type.Field(j).Tag.Get("excel"))
			}
		} else {
			headers = append(headers, field.Tag.Get("excel"))
		}
	}
	return headers
}

// New 创建新的工作簿
func (w *Writer) New(name string) {
	// 生成xlsx文件
	w.WorkBook = excelize.NewFile()
	w.WorkBook.Path = fmt.Sprintf("%s.xlsx", name)
}

// NewStreamWriter 生成新的名为 name 的工作表, 并创建新的流式写入器
func (w *Writer) NewStreamWriter(name string) {
	w.WorkBook.NewSheet(name)
	streamSheet, err := w.WorkBook.NewStreamWriter(name)
	if err != nil {
		panic(err)
	}
	w.StreamWriter = streamSheet
}

func (w *Writer) WriteHeader(model interface{}) int {
	headers := createHeader(model)
	w.StreamWriter.SetRow(fmt.Sprintf("A%d", 1), headers)
	return len(headers)
	// w.StreamWriter.Flush()
}

func (w *Writer) WriteRow(model interface{}, index int) {
	cell_pre, _ := excelize.CoordinatesToCellName(1, index)

	val := conver2interface(model)
	if err := w.StreamWriter.SetRow(cell_pre, val); err != nil {
		panic(err)
	}
}

func (w *Writer) WriteRows(models interface{}, index int, name string) {
	rowVal := reflect.ValueOf(models)
	for i := 0; i < rowVal.Len(); i++ {
		model := rowVal.Index(i).Interface()
		val := conver2interface(model)
		w.StreamWriter.SetRow(fmt.Sprintf("A%d", index+i), val)
	}
	w.StreamWriter.Flush()
}

func (w *Writer) Save() {
	if err := w.WorkBook.Save(); err != nil {
		panic(err)
	}

	if err := w.WorkBook.Close(); err != nil {
		panic(err)
	}
}
