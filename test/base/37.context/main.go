package main

import (
	"context"
	"fmt"
)

// 参考：<https://segmentfault.com/a/1190000020549820>
//
// 如果ctx中带有通用的上下文信息，需要写个函数生成一个新的ctx，同时把原来ctx的kv复制出来。否则直接使用context.Background()就好了。

var ctxKeyBar = CtxKeyBar{}
var ctxKeyFoo = CtxKeyFoo{}

type CtxKeyBar struct {
}
type CtxKeyFoo struct {
}

func CopyCtx(ctx context.Context) context.Context {
	newCtx := context.Background()
	newCtx = context.WithValue(newCtx, ctxKeyBar, ctx.Value(ctxKeyBar))
	newCtx = context.WithValue(newCtx, ctxKeyFoo, ctx.Value(ctxKeyFoo))
	return newCtx
}

func main() {
	copiedCtx := CopyCtx(context.Background())
	fmt.Println(copiedCtx)
}
