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

// とりあえず今はsleepして実行するだけ
func setTimeout(ms int, callback func()) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	callback()
}

func main() {
	runtime.run(func() {
		println("開始")
		setTimeout(100, func() {
			println("OK")
		})
		println("終わり")
	})
}
