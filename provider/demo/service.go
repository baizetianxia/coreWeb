package demo

import (
	"fmt"

	"github.com/baizetianxia/coreWeb/framework"
)

type DemoService struct {

	//显示实现接口
	Service

	c framework.Container
}

func (ds *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am foo",
	}
}

// 初始化实例的方法
func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}
