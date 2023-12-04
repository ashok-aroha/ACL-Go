package acl

import "fmt"

type Log struct{}

type Obj struct {
	StatusCode int
	Message    string
}

func (l Log) PrintLog(obj Obj) {
	fmt.Println("Status Code:", obj.StatusCode)
	fmt.Println("Message:", obj.Message)
}

// Example

// func main() {
//      log := Log{}
//      obj := Obj{StatusCode: 200, Message: "OK"}
//      log.PrintLog(obj)
// # }
