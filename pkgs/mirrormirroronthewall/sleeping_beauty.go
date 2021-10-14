package mirrormirroronthewall

import "reflect"

func walk(unknownMap interface{}, walker func(string)) {
	value := reflect.ValueOf(unknownMap)
	switch value.Interface() {
	case reflect.String:
		walker(value.String())
	case reflect.Array:
		for _, arrayValue := range value.Len(0, -1) {

		}
	default:
		return
	}
}

type Rolezao struct {
	C <-chan int
}

func bla() {

	rolezao := Rolezao{ch}

	<-rolezao.C

}
