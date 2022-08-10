package admin

import (
	"fmt"
	"reflect"
)

type Admin struct {
	tables []interface{}
}

type TableField struct {
	Name  string
	Type  string
	Value interface{}
}

func New() *Admin {
	return &Admin{}
}

func (a *Admin) Subscribe(table interface{}) {
	fmt.Print(reflect.TypeOf(table).Name())
	a.tables = append(a.tables, table)
}

func (a *Admin) GetTableNames() []string {
	var names []string
	for _, table := range a.tables {
		names = append(names, reflect.TypeOf(table).Name())
	}
	return names
}

func (a *Admin) GetHeaders(table interface{}) []string {
	var headers []string
	e := reflect.TypeOf(table)

	for i := 0; i < e.NumField(); i++ {
		headers = append(headers, e.Field(i).Name)
	}
	return headers
}

func (a *Admin) GetValueByFieldName(table interface{}, fieldName string) interface{} {
	e := reflect.TypeOf(table)
	for i := 0; i < e.NumField(); i++ {
		if e.Field(i).Name == fieldName {
			return reflect.ValueOf(table).Field(i).Interface()
		}
	}
	return nil
}
