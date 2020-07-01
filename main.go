package main

import (
	"time"
)

type Runtime struct {
	nextID    int
	callbacks map[int]func()
	recv      chan int
}

func NewRuntime() *Runtime {
	recv := make(chan int)
	return &Runtime{recv: recv, callbacks: map[int]func(){}}
}

func (rt *Runtime) run(program func()) {
	program()
	for {
		funcID := <-rt.recv
		callback := rt.callbacks[funcID]
		delete(rt.callbacks, funcID)
		callback()
		if len(rt.callbacks) == 0 {
			break
		}
	}
}

func (rt *Runtime) register(callback func()) int {
	registeredID := rt.nextID
	rt.callbacks[registeredID] = callback
	rt.nextID++
	return registeredID
}

var runtime = NewRuntime()

func setTimeout(ms int, callback func()) {
	funcID := runtime.register(callback)
	go func() {
		time.Sleep(time.Duration(ms) * time.Millisecond)
		runtime.recv <- funcID
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
