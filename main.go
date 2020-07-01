package main

import (
	"time"
)

// TODO: そのうちフィールドを用意する
type Runtime struct{}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (rt *Runtime) run(program func()) {
	program()
}

var runtime = NewRuntime()

// 並行処理をするようにはなったけど・・・
func setTimeout(ms int, callback func()) {
	go func() {
		time.Sleep(time.Duration(ms) * time.Millisecond)
		callback()
	}()
}

func main() {
	runtime.run(func() {
		println("開始")
		setTimeout(100, func() {
			println("OK")
		})
		// 待つ秒数が少ないこっちが先に実行されたい
		setTimeout(20, func() {
			println("まずはこっち")
		})
		setTimeout(1000, func() {
			println("終わり")
		})
	})
}
