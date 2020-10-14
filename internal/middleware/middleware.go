// @Author: 2014BDuck
// @Date: 2020/10/14

package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"runtime/debug"
)

func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if e := recover(); e != nil {
			recoveryLog := "recovery log: method: %s, message: %v, stack: %s"
			log.Printf(recoveryLog, info.FullMethod, e, string(debug.Stack()[:]))
		}
	}()
	return handler(ctx, req)
}
