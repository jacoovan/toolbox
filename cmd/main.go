package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jacoovan/toolbox/interface/web"
)

const (
	addr = ":80"
)

func main() {
	ctx := context.Background()

	srv := web.NewGin(addr)
	if err := srv.Run(ctx); err != nil {
		exit(2, fmt.Errorf(`srv.Run(err):%v`, err))
	}
}

func exit(errno int, err error) {
	fmt.Println(err)
	os.Exit(errno)
}
