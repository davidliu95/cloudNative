package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		//接收客户端 request，并将 request 中带的 header 写入 response header
		header := request.Header
		for key := range header {
			writer.Header().Set(key, header[key][0])
		}
		//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		writer.Header().Set("VERSION", os.Getenv("VERSION"))
		//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		ip, _, err := net.SplitHostPort(request.RemoteAddr)
		if err != nil {
			return
		}
		if net.ParseIP(ip) != nil {
			fmt.Println("ip ===>>%s\n", ip)
			log.Println(ip)
		}

		fmt.Println("http Status Code ===>>%s\n", http.StatusOK)
		log.Println(http.StatusOK)

		//response响应
		writer.WriteHeader(http.StatusOK)

		writer.Write([]byte("Server Access,Success!"))

	})
	//当访问 localhost/healthz 时，应返回 200
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("200"))
	})
	http.ListenAndServe("localhost:8000", nil)
}
