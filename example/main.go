// Copyright 2017 <614422099@qq.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

// dingdayu/golangtools/config package example.
//
// https://github.com/golangunit/config
package main

import (
	"flag"
	"fmt"
	"github.com/golangunit/config"
)

/**
实例
*/
type Instance struct {
	Name string
	Path string
	Cmd  []string
	User string
}

// 配置
type Config struct {
	Port     int
	Instance []Instance
}

func main() {

	file := flag.String("c", "conf.json", "config file path")
	flag.Parse()

	var conf Config
	err := config.New(*file, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}
