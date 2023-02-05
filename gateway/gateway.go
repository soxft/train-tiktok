package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"os/exec"
	"train-tiktok/common/errorx"
	"train-tiktok/gateway/internal/config"
	"train-tiktok/gateway/internal/handler"
	"train-tiktok/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	// check if ffmpeg is installed
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		panic("ffmpeg is not installed")
	}

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		_, ok := status.FromError(err)
		if !ok {
			logx.Errorf("error when handler rpc err: %v", err)
			return http.StatusOK, errorx.ErrorResp{Code: 3004, Msg: err.Error()}
		}

		logx.Errorf("error when handler err: %v", err)
		return http.StatusOK, errorx.FromRpcStatus(errorx.ErrSystemError)
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
