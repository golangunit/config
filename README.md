# Golang Tools Config

<p align="center">
    <a href="https://travis-ci.org/golangunit/config"><img src="https://travis-ci.org/golangunit/config.svg?branch=master" alt="Build Status"></img></a>
    <a href="https://ci.appveyor.com/project/golangunit/config"><img src="https://ci.appveyor.com/api/projects/status/github/golangunit/config?svg=true&branch=master&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending" alt="Windows Build Status"></a>
    <a href="https://blog.dingxiaoyu.com"><img src="https://img.shields.io/badge/author-@dingdayu-blue.svg?style=flat" alt="Author"></a>
    <a href="https://godoc.org/github.com/golangunit/config"><img src="https://godoc.org/github.com/golangunit/config?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/golangunit/config"><img src="https://goreportcard.com/badge/github.com/golangunit/config" alt="Report"></a>
</p>

Golang Tools Config Support：

- Json
- Yaml
- Toml
- Xml

## Install

```
go get github.com/golangunit/config
```

## Example

### json

```
	// 实例
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	// 配置
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.json", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

Run:

```
go run ../example/config/main.go -c ../example/config/conf.json
{8080 [{gitbook /gitbook [ll ls] dingdayu}]}
```

> 请注意路径！

> 参见 [example/config/main.go](example/config/main.go)

### xml


```
	// 实例
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	// 配置
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.xml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参加 [example/config/main.go](example/config/main.go)

### toml


```
	// 实例
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	// 配置
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.toml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参见 [example/config/main.go](example/config/main.go)

### yaml


```
	// 实例
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	// 配置
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.yaml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

> 参见 [example/config/main.go](example/config/main.go)

### Donate

<p align="center">
    <img width="400" src="https://github.com/dingdayu/golangtools/blob/master/donate.png"></img>
</p>