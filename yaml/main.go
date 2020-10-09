package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: 
  - 3
  - 4
`

// B 结构体，字段注释写不写都行
type B struct {
	RenamedC int   `yaml:"c"`
	D        []int `yaml:",flow"`
}

// T 结构体
// Note: struct 字段必须是全局生效的，以便可以正确 unmarshal 该结构体并讲 yaml 数据填充进来
type T struct {
	A string
	B
}

func main() {
	// 声明结构体，单行 yaml 格式无法显示 key
	t := T{}

	// Unmarshal 具有打散、解包等含义。用于将参数2进行解码后生成的值，传给参数1。
	// 用白话说就是将多行 yaml 格式合并为1行 yaml 格式。
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 单行 yaml 格式:\n%v\n\n", t)
	// Marshel 具有整理、排列、打包、编组等含义。用于给定的值序列化为 YAML 文档格式。
	// Marshel() 参数可以接受的数据格式为 `(...) yaml:"[<key>][,<flag1>[,<flag2>]]" (...)`
	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 多行 yaml 格式:\n%s\n\n", string(d))

	// 通过 make 声明结构体，单行 yaml 格式可以显示 key
	m := make(map[interface{}]interface{})

	// Unmarshal 的作用与上面一样
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 单行 yaml 格式:\n%v\n\n", m)

	// Marshal 的作用与上面一样
	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 多行 yaml 格式:\n%s\n\n", string(d))
}
