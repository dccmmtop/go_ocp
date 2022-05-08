package web_arg

import (
	"fmt"
	"ocp/arg_parse"
)

type WebParse struct {}

func init(){
	// 注册转换器
	fmt.Println("注册 WebParse")
	arg_parse.Driver = &WebParse{}

}
func (c WebParse) ParseArg(args []string) *arg_parse.Args{
	fmt.Println("我是网页的参数转换器")
	return &arg_parse.Args{}
}
