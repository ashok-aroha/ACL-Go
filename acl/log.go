package acl

import (
	"fmt"
	"reflect"
)

type Log struct{}

func (l *Log) PrintLog(obj interface{}) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Struct {
		fmt.Println("Invalid type")
		return
	}

	statusCodeField := val.FieldByName("StatusCode")
	messageField := val.FieldByName("Message")

	if !statusCodeField.IsValid() || !messageField.IsValid() {
		fmt.Println("Invalid type")
		return
	}

	statusCode := statusCodeField.Int()
	message := messageField.String()

	fmt.Println("Status Code:", statusCode)
	fmt.Println("Message:", message)
}
