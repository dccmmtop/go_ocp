package commandline_arg

import (
	"fmt"
	"ocp/arg_parse"
)

type CommandlineParse struct {}

func init(){
	// 注册转换器
	fmt.Println("注册 CommandlineParse")
	arg_parse.Driver = &CommandlineParse{}

}
func (c CommandlineParse) ParseArg(args []string) *arg_parse.Args{
	fmt.Println("我是命令行的参数转换器")
	return &arg_parse.Args{}
}
