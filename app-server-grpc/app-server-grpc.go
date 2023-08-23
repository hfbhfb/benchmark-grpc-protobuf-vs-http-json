package main

import (
	"fmt"
	"os"
	"time"

	grpcprotobuf "github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf"
	httpjson "github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

var (
	envServerGrpc = ""
	envServerHttp = ""

	envClientGrpc = ""
	envClientHttp = ""

	envServerUrl = ""
)

func checkMoreFlag() int {
	i := 0
	if envServerGrpc != "" {
		i++
	}
	if envServerHttp != "" {
		i++
	}
	if envClientGrpc != "" {
		i++
	}
	if envClientHttp != "" {
		i++
	}
	return i
}

func main() {
	envServerGrpc = os.Getenv("ENV-SERVER-GRPC")
	envServerHttp = os.Getenv("ENV-SERVER-HTTP")

	envClientGrpc = os.Getenv("ENV-CLIENT-GRPC")
	envClientHttp = os.Getenv("ENV-CLIENT-HTTP")

	// envServerUrl := os.Getenv("ENV-SERVER-URL")
	// fmt.Println(envServerGrpc)
	// fmt.Println(envServerHttp)
	// fmt.Println(envClientGrpc)
	// fmt.Println(envClientHttp)

	if envServerGrpc == "" && envServerHttp == "" && envClientGrpc == "" && envClientHttp == "" {
		fmt.Println("需要选择一种角色: error")
		fmt.Println("环境变量 ENV-SERVER-GRPC 有值: 对应A")
		fmt.Println("环境变量 ENV-SERVER-HTTP 有值: 对应B")
		fmt.Println("环境变量 ENV-CLIENT-GRPC 有值: 对应C")
		fmt.Println("环境变量 ENV-CLIENT-HTTP 有值: 对应D")
		time.Sleep(time.Second * 10)
		return
	}
	if checkMoreFlag() >= 2 {
		fmt.Println("只能选择一种角色...error,,,  存在多个环境变量")
		fmt.Println("环境变量 ENV-SERVER-GRPC 有值: 对应A")
		fmt.Println("环境变量 ENV-SERVER-HTTP 有值: 对应B")
		fmt.Println("环境变量 ENV-CLIENT-GRPC 有值: 对应C")
		fmt.Println("环境变量 ENV-CLIENT-HTTP 有值: 对应D")
		time.Sleep(time.Second * 10)
		return
	}

	if envServerGrpc != "" {
		grpcprotobuf.Start()
	}
	if envServerHttp != "" {
		httpjson.Start()
	}

	// envServerUrl

	// // 根据环境变量的值执行不同的逻辑
	// if environment == "production" {
	//     fmt.Println("Running in production environment.")
	// } else if environment == "development" {
	//     fmt.Println("Running in development environment.")
	// } else {
	//     fmt.Println("Running in an unknown environment.")
	// }
	// fmt.Println("111")
}
