package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string) string
}

type ReadWrite interface {
	Reader
	Writer
}

// FileReader 是 ReadWriter 接口的实现
type FileReadWriter struct {
	// 文件读取器的具体实现
 }
 
 // Read 实现了 ReadWriter 接口的 Read 方法
 func (fr FileReadWriter) Read() string {
	// 实现文件读取逻辑
	return "Data from file"
 }
 
 // Write 实现了 ReadWriter 接口的 Write 方法
 func (cw FileReadWriter) Write(data string) {
	// 实现控制台写入逻辑
	fmt.Println("Writing data to console:", data)
 }

func main() {
	frw := FileReadWriter{}
	frw.Read()
	frw.Write(11)
}
