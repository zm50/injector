package main

import (
	"fmt"
	"github.com/zm50/injector"
	"reflect"
)

type TwoString struct {
	s1 *string
	s2 *string
}

func main() {
	// 通过类型注入变量，注入一个string类型的变量
	var injectString string = "醉墨居士"
	err := injector.Inject(injectString, "")
	if err != nil {
		panic(err)
	}
	// 通过类型装配变量，通过string类型自动装配变量
	var autowiseString string
	err = injector.Autowise(&autowiseString, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("类型注入和装配的演示结果")
	fmt.Println("注入的变量：", injectString, "装配的变量：", autowiseString)

	// 通过名称注入变量
	var injectName string = "醉墨"
	var injectString2 string = "居士"
	err = injector.Inject(injectString2, injectName)
	if err != nil {
		panic(err)
	}
	// 通过名称装配变量
	var autowiseString2 string
	err = injector.Autowise(&autowiseString2, "醉墨")
	if err != nil {
		panic(err)
	}

	fmt.Println("名称注入和装配的演示结果")
	fmt.Println("注入的变量：", injectString2, "装配的变量：", autowiseString2)

	// 通过类型注入结构体指针
	injectStruct := &TwoString{}
	injectStruct.s1 = new(string)
	injectStruct.s2 = new(string)
	*injectStruct.s1 = "醉墨"
	*injectStruct.s2 = "居士"
	err = injector.Inject(injectStruct, "")
	if err != nil {
        panic(err)
    }
	// 通过类型装配结构体指针
	autowiseStruct := &TwoString{}
	err = injector.Autowise(&autowiseStruct, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("结构体指针注入和装配的演示结果")
	fmt.Println("注入的变量：", injectStruct, *(injectStruct.s1), *(injectStruct.s2))
	fmt.Println("装配的变量：", autowiseStruct, *(autowiseStruct.s1), *(autowiseStruct.s2))
	fmt.Println("是否相等：", injectStruct == autowiseStruct, injectStruct.s1 == autowiseStruct.s1, injectStruct.s2 == autowiseStruct.s2)

	// 自定义依赖注入和装配的能力，演示自定义依赖注入和装配的能力实现深拷贝，大家可以也根据自己的需求自定义依赖注入和装配的能力
	injectStruct2 := &TwoString{
		s1: new(string), s2: new(string),
	}
	*(injectStruct2.s1) = "醉墨"
	*(injectStruct2.s2) = "居士"
	provider := func() reflect.Value {
		twoString := TwoString{}
		twoString.s1 = new(string)
		twoString.s2 = new(string)
		*twoString.s1 = *injectStruct.s1
		*twoString.s2 = *injectStruct.s2
		return reflect.ValueOf(twoString)
	}
	err = injector.DeepInject(provider, "")
	if err != nil {
		panic(err)
	}
	autowiseStruct2 := &TwoString{}
	err = injector.Autowise(autowiseStruct2, "")
	if err != nil {
		panic(err)
	}

	fmt.Println("自定义规则实现结构体深拷贝注入和装配的演示结果")
	fmt.Println("注入的变量：", injectStruct2, *(injectStruct2.s1), *(injectStruct2.s2))
	fmt.Println("装配的变量：", autowiseStruct2, *(autowiseStruct2.s1), *(autowiseStruct2.s2))
	fmt.Println("是否相等：", injectStruct2 == autowiseStruct2, injectStruct2.s1 == autowiseStruct2.s1, injectStruct2.s2 == autowiseStruct2.s2)
}
