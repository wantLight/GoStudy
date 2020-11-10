package main

import (
	"context"
	"fmt"
	"time"
)

/**
context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，然后当作参数传给goroutine使用，
这样就可以使用这个子Context跟踪这个goroutine。重写比较简单，就是把原来的chan stop 换成Context，使用Context跟踪goroutine，以便进行控制，比如结束等。

在goroutine中，使用select调用<-ctx.Done()判断是否要结束，
如果接受到值的话，就可以返回结束goroutine了；如果接收不到，就会继续进行监控。

那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，
它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，第二个返回值就是这个取消函数，它是CancelFunc类型的。
我们调用它就可以发出取消指令，然后我们的监控goroutine就会收到信号，就会返回结束。
*/
func main() {
	// 这个空的Context一般用于整个Context树的根节点;创建一个可取消的子Context，然后当作参数传给goroutine使用，
	ctx, cancel := context.WithCancel(context.Background())
	//Deadline方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。
	//Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。
	//Err方法返回取消的错误原因，因为什么Context被取消。
	//Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。
	ctx = context.WithValue(ctx, "test", "hahaha")
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	// 它是我们调用context.WithCancel(parent)函数生成子Context的时候返回的，第二个返回值就是这个取消函数
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context, name string) {
	for {
		select {
		// 如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了...")
			return
		default:
			fmt.Println(name, "goroutine监控中...")
			fmt.Println(ctx.Value("test"))
			time.Sleep(2 * time.Second)
		}
	}
}
