package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// User 用户信息
type User struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

//Nginx nginx  配置
type Nginx struct {
	Port    int    `yaml:"port"`
	LogPath string `yaml:"logPath"`
	Path    string `yaml:"path"`
}

//YamlInfo yaml文件信息
type YamlInfo struct {
	Name  string `yaml:"name"`
	User  []User `yaml:"user"`
	Nginx Nginx  `yaml:"nginx"`
}

func main() {
	// 创建并解析模板文件
	t := template.Must(template.ParseFiles("template/yaml.tmpl"))

	// 读取 yaml 文件内容，并将内容放入 config 中后，通过 Unmarshal 处理内容，再放入 yamlInfo 中。
	config, errRead := ioutil.ReadFile("./info.yaml")
	if errRead != nil {
		fmt.Print(errRead)
	}

	var yamlInfo YamlInfo
	errUnmarshal := yaml.Unmarshal(config, &yamlInfo)
	if errUnmarshal != nil {
		log.Fatalf("error: %v", errUnmarshal)
	}

	// 将处理好的所有 yaml 文件中的内容，应用于模板，并输出。
	t.Execute(os.Stdout, yamlInfo)
}
