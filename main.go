package main

import (
	"fmt"
	"time"
)

type ITimeManager interface {
	Now() time.Time
}

type TimeManager struct{}

func (t *TimeManager) Now() time.Time {
	return time.Now()
}

func main() {
	fmt.Println(UnixTimeSample(&TimeManager{}))
}

func UnixTimeSample(timeManager ITimeManager) string {
	unixNow := timeManager.Now().Unix()
	if unixNow%2 != 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func PrintSample() {
	var a, b MyInterface
	a = &ImplA{}
	b = &ImplB{}
	PrintAny(a) // AをDI!!
	PrintAny(b) // BをDI!!
}

func PrintAny(any MyInterface) {
	msg, err := any.Method([]string{})
	if err != nil {
		return
	}
	fmt.Println("any message is ", msg)
}

type MyInterface interface {
	Method(args []string) (string, error)
}

type ImplA struct{}

func (a *ImplA) Method(args []string) (string, error) {
	return fmt.Sprintf("I am A"), nil
}

type ImplB struct{}

func (a *ImplB) Method(args []string) (string, error) {
	return fmt.Sprintf("I am B"), nil
}
