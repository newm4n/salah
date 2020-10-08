package salah

import "testing"

func TestPrintStackTrace(t *testing.T) {
	Func1()
}

func Func1() {
	Func2()
}

func Func2() {
	Func3()
}

func Func3() {
	PrintStackTrace()
}