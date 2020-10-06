package main

import (
	"os"
	"text/template"
)

func main() {
	// 创建一个名为 hello world 的未定义模板。此时模板中无任何内容，模板内容以及要传递的值都没有，仅仅是生成一个模板指针，用来后续调用。
	// 这种行为类似于 open 一个 文件，或者 new 一个数组之类。
	t := template.New("hello world")

	// 定义模板 t ，将参数解析为模板的主体。
	// 如果用一个文件来类比的话，现在就是有一个名为"hello world"的文件，这个文件的内容为"这是一个 {{ . }} 模板"
	t, _ = t.Parse("这是一个 {{ . }} 模板")

	// 这是一个帮助程序，用来对模板 t 进行验证，并不影响后续对模板的输出，可省略
	// t = template.Must(t, err)

	// 上述三个分开的步骤可以合并为一行代码，示例如下
	// 这是一个没有 Must 验证的代码
	// t, _ := template.New("hello world").Parse("这是一个 {{ . }} 模板")
	// 这是一个包含 Must 验证的代码
	// t := template.Must(template.New("hello world").Parse("这是一个 {{ . }} 模板"))

	// 将指定的数据(第二个参数)应用于已解析的模板 t ，
	t.Execute(os.Stdout, "hello world")
}
