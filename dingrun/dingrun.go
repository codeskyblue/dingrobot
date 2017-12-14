package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/codeskyblue/dingrobot"
	"github.com/codeskyblue/kexec"
)

func goFunc(f func() error) chan error {
	errC := make(chan error, 1)
	go func() {
		errC <- f()
	}()
	return errC
}

func currentIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err.Error()
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: dingrun command args...\nExample: dingrun sleep 10")
		os.Exit(1)
	}

	dingRobotToken := os.Getenv("DINGROBOT_TOKEN")
	if dingRobotToken != "" {
		log.Println("DingTalk robot notification enabled")
	}
	cmd := kexec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	startTime := time.Now()
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	select {
	case err := <-goFunc(cmd.Wait):
		if err == nil {
			log.Println("Program success exited")
		} else {
			log.Println(err)
		}
		if dingRobotToken != "" {
			cmdline := strings.Join(args, " ")
			dingrobot.New(dingRobotToken).Markdown("cmd quit: "+cmdline,
				fmt.Sprintf(">%s\n\n", cmdline)+
					fmt.Sprintf("**%v**\n", err)+
					fmt.Sprintf("- 运行时间: %s\n", time.Since(startTime))+
					fmt.Sprintf("- IP: %s\n", currentIP())+
					fmt.Sprintf("- OS: %s\n", runtime.GOOS))
		}
	case <-sigC:
		log.Println("Handle interrupt signal, kill program")
		cmd.Terminate(os.Kill)
	}
}
