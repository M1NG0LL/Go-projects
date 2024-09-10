package main

import (
	"fmt"
)

type tvstate interface {
	state()
}

type on struct {}

func (o *on) state() {
	fmt.Println("TV is ON")
}

type off struct {}

func (o *off) state() {
	fmt.Println("TV is OFF")
}

type statecontext struct {
	tv tvstate
}

func getcontext() *statecontext {
	return &statecontext{
		tv: &off{},
	}
}

func (sc *statecontext) setstate(state tvstate) {
	sc.tv = state
}

func (sc *statecontext) getstate() {
	sc.tv.state()
}

// func main() {
// 	tvcontext := getcontext()

// 	tvcontext.getstate()

// 	tvcontext.setstate(&on{})

// 	tvcontext.getstate()
// }