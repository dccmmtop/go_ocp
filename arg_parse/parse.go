package arg_parse
// 定义参数转换的接口，其他各种参数转换器都要实现本接口的功能。

var Driver Parse

type Args struct {
	Name string
	Age int
	Location string
}

// 待实现的方法
type Parse interface {
	ParseArg([] string) *Args
}
