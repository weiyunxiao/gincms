package vardump

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/dump"
)

// 打印变量
func P(v any) {
	dump.P(v)
}

// 打印并退出
func Dd(v any) {
	dump.P(v)
	os.Exit(8008)
}

// 变量信息非常长时，可以将调试信息输出到log文件中，方便查看
func Fprint(v any) {
	f, err := os.OpenFile("log/vardump.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("调试信息到文件中出错:%s\r\n", err)
		return
	}
	defer f.Close()
	dump.Config(dump.WithoutColor())
	dump.Fprint(f, v)
}
