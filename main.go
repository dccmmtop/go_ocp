package main

import (
	"ocp/arg_parse"
	//安装命令行转换器
	//_ "ocp/commandline_arg"

	// 或者安装网页参数转换器
	_ "ocp/web_arg"
	"os"
)

func main(){
	arg_parse.Driver.ParseArg(os.Args)
}
