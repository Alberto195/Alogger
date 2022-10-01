package main

import (
	"flag"
	"fmt"
	"hello/ALogger/alog"
	"log"
	"os"
	"sync"
	"time"
)

var (
	threads     *int
	logCount    *int
	aLogEnabled *bool
)

func init() {
	threads = flag.Int("threads_count", 3, "number of threads to run")
	logCount = flag.Int("log_count", 50, "number of logs to write in each thread")
	aLogEnabled = flag.Bool("async_logger_enabled", true, "enable async logger")
}

func main() {
	flag.Parse()
	wg := sync.WaitGroup{}
	if *aLogEnabled {
		aLogger := alog.NewAsyncLogger("[INFO]", os.Stdout)
		for i := 0; i < *threads; i++ {
			wg.Add(1)
			i := i
			go func() {
				for j := 0; j < *logCount; j++ {
					aLogger.Info(fmt.Sprintf("Hi! it is %v in %v", j, i))
				}
				wg.Done()
			}()
		}
	} else {
		logger := log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
		for i := 0; i < *threads; i++ {
			wg.Add(1)
			i := i
			go func() {
				for j := 0; j < *logCount; j++ {
					_ = logger.Output(2, fmt.Sprintf("Hi! it is %v in %v", j, i))
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
}
