package main

import (
	"log"
	"os"
	"text/template"
)

// 当template开始执行时，$变量被默认设置成传递个Execute函数的数据参数，也就是“.”光标的开始值
const templ = `{{$comma := sequence "" ", "}}
{{range $}}{{$comma.Next}}{{.}}{{end}}
{{$comma := sequence "" ", "}}
{{$color := cycle "black" "white" "red"}}
{{range $}}{{$comma.Next}}{{.}} in {{$color.Next}}{{end}}
`

var funcMap = template.FuncMap{
	"sequence": sequenceFunc,
	"cycle": cycleFunc,
}

func main() {
	t, err := template.New("").Funcs(funcMap).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, []string{"a", "b", "c", "d", "e", "f"})
	if err != nil {
		log.Fatal(err)
	}
}

// 一个封装有迭代函数的生成器
type generator struct {
	ss	[]string
	i	int
	f	func(ss []string, i int) string
}

func (seq *generator) Next() string {
	s := seq.f(seq.ss, seq.i)
	seq.i++
	return s
}

func sequenceGen(ss []string, i int) string {
	if i >= len(ss) {
		i = len(ss) - 1
	}
	return ss[i]
}

func cycleGen(ss []string, i int) string {
	return ss[i%len(ss)]
}

func sequenceFunc(ss ...string) *generator {
	if len(ss) == 0 {
		log.Fatal("Sequence must have at least one element")
	}
	return &generator{ss, 0, sequenceGen}
}

func cycleFunc(ss ...string) *generator {
	if len(ss) == 0 {
		log.Fatal("Cycle must have at least one element")
	}
	return &generator{ss, 0, cycleGen}
}
