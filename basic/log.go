package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	_, err := os.Open("nothing.go")
	if err != nil {
		log.Fatalln("exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!") // ここでexit

	fmt.Println("ok!") // 表示されない
}
