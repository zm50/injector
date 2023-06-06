package main

import (
	"fmt"
	"github.com/go75/injector"
	"reflect"
)

type TwoString struct {
	s1 *string
	s2 *string
}

func main() {
	// type inject
	err := injector.Put("java")
	if err!=nil {
		panic(err)
	}
	s, err := injector.Get[string]()
	if err!=nil {
		panic(err)
	}
	// name inject
	err = injector.PutByName("name","jack")
	if err!=nil {
		panic(err)
	}
	name, err := injector.GetByName[string]("name")
	if err!=nil {
		panic(err)
	}
	fmt.Println(s, name)

	// shallow copy
	twoStr := TwoString{}
	twoStr.s1 = new(string)
	twoStr.s2 = new(string)
	*twoStr.s1 = "string1"
	*twoStr.s2 = "string2"
	err = injector.Put(twoStr)
	if err!= nil {
        panic(err)
    }
	twoStr2, err := injector.Get[TwoString]()
	if err!=nil {
		panic(err)
	}

	fmt.Println("twoStr==twoStr2: ", twoStr==twoStr2)
	
	// deep copy
	err = injector.DeepPut(func() reflect.Value {
		twoString := TwoString{}
		twoString.s1 = new(string)
		twoString.s2 = new(string)
		*twoString.s1 = "string1"
		*twoString.s2 = "string2"
		return reflect.ValueOf(twoString)
	})
	if err!=nil {
		panic(err)
	}
	ts, err := injector.DeepGet[TwoString]()
	if err!=nil {
		panic(err)
	}
	twoString, err := injector.DeepGet[TwoString]()
	if err!=nil {
		panic(err)
	}
	fmt.Println("ts==twoString: ",ts==twoString)
	fmt.Println(*ts.s1, *ts.s2)
	fmt.Println(*twoString.s1, *twoString.s2)
	*ts.s1, *ts.s2 = "string2", "string1"
	fmt.Println(*ts.s1, *ts.s2)
	fmt.Println(*twoString.s1, *twoString.s2)
}
