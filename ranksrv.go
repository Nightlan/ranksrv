package main

import (
	"fmt"
	"os"
	"os/signal"
	"ranksrv/rank"
	"syscall"
)

const(
	httpPort = "10031"
	cacheDir = `d:/cache`
)

func main() {
	if err := rank.InitRankMgr(cacheDir, httpPort); err != nil {
		fmt.Println("Init rank manager failed", err)
		return
	}
	// 等待退出信号
	fmt.Println("server is running...")
	Wait()
	rank.CloseRankMgr()
	fmt.Println("server is stop.")
}

func Wait() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
