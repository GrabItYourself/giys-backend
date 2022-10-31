package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/pkg/errors"
)

func runService(serviceName string) (io.WriteCloser, error) {
	cmd := exec.Command("go", "run", serviceName+"/cmd/main.go")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get stdin pipe: "+serviceName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get stdout pipe: "+serviceName)
	}

	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "Failed to start service: "+serviceName)
	}

	reader := bufio.NewReader(stdout)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "Failed to start service: "+serviceName)
	}

	for err == nil {
		fmt.Print("[", serviceName, "]:", line)
		line, err = reader.ReadString('\n')
	}

	return stdin, nil
}

func cleanUpServices(stdinList chan io.WriteCloser) {
	fmt.Println("Shutting down...")
	for stdin := range stdinList {
		if err := stdin.Close(); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Done")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	serviceNames := flag.String("services", "", "List of services to start seperated by comma")

	flag.Parse()

	wg := &sync.WaitGroup{}
	services := strings.Split(*serviceNames, ",")
	stdinList := make(chan io.WriteCloser, len(services))
	defer cleanUpServices(stdinList)

	for _, serviceName := range services {
		wg.Add(1)
		go func(serviceName string) {
			defer wg.Done()
			stdin, err := runService(serviceName)
			if err != nil {
				panic(err)
			}
			stdinList <- stdin
		}(serviceName)
	}
	wg.Wait()
	close(stdinList)

	<-ctx.Done()
}
