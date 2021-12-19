/*
基于 errgroup 实现一个 http server 的启动和关闭,
以及 linux signal 信号的注册和处理，
要保证能够一个退出，全部注销退出。
*/

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

//启动 HTTP server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

// http server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	ctx := context.Background()
	// 定义 withCancel -> cancel() 方法 去取消下游的 Context
	ctx, cancel := context.WithCancel(ctx)
	// 通过 errgroup 管理 goroutine
	group, errCtx := errgroup.WithContext(ctx)
	//http server
	srv := &http.Server{Addr: ":8080"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx) // 关闭 http server
	})

	//实现对中断信号的处理
	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}
