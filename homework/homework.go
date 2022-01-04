package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	//value := []interface{}{0, 1, 2, 3, 4, 5}
	//newValue, err := Add(value, 4, 4)
	//// 输出 0，1，2，3，4，5，6
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//for i, i2 := range newValue {
	//	print(i, i2)
	//	fmt.Printf("下标 %d, 值: %d, \n", i, value)
	//}

	h := Hello {
		endpoint: "http://httpbin.org/get",
	}

	_ = reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	num := v.NumField()

	fmt.Println(num)

	//for i, d := range f {
	//	println(i)
	//	println(i2)
	//}
	//for i := 0; i < num; i++ {
	//	f := v.Field(i)
	//
	//}


	//msg, _ := h.SayHello("golang")
	//print(msg)
}

func SetFuncField(val interface{}) {
	v := reflect.ValueOf(val)
	ele := v.Elem()
	t := ele.Type()

	numField := t.NumField()
	for i := 0; i < numField; i++ {
		field := t.Field(i)
		fieldValue := ele.Field(i)
		if fieldValue.CanSet() {
			//fmt.Println("aaa" + f.String())
			fn := func(args [] reflect.Value) (results []reflect.Value) {
				name := args[0].Interface().(string)

				client := http.Client{}

				resp, err := client.Get("http://localhost:8080/" + name)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				return []reflect.Value{reflect.ValueOf(string(data)), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			fieldValue.Set(reflect.MakeFunc(field.Type, fn))
		}
	}
}
func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {
	var res []interface{}

	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}

	res = append(res, val)
	return res, nil
}

func Delete(values []interface{}, index int) []interface{} {
	if index < 0 || index > len(values) {
		return values
	}
	var res []interface{}
	for i := 0; i < len(values); i++ {
		if index == i {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}

type HelloService interface {
	SayHello(name string) (string, error)
}

type Hello struct {
	endpoint string
}

func (h Hello) SayHello(name string) (string, error) {

	client := http.Client{}
	resp, _ := client.Get(h.endpoint)
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}