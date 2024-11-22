package main

import (
	"log"
	"os"
	"runtime"
	"strings"
	"fmt"

)

const (
    colorRed   = "\033[31m"
    colorReset = "\033[0m"
)

type customLogger struct {
    logger *log.Logger
}

func (c *customLogger) Println(v ...interface{}) {
    _, file, line, ok := runtime.Caller(2)
    if ok {
        shortFile := file[strings.LastIndex(file, "/")+1:]
        c.logger.Println(fmt.Sprintf("%s%s:%d%s", colorRed, shortFile, line, colorReset), fmt.Sprintln(v...))
    } else {
        c.logger.Println(v...)
    }
}

func newCustomLogger() *customLogger {
    return &customLogger{
        logger: log.New(os.Stdout, "", log.LstdFlags),
    }
}