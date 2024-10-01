package main

import "fmt"
import "time"

var Arr = [8]string{"|", "/", "-", "\\", "|", "/", "-", "\\"}

func ResetScreen() {
	fmt.Print("\033[0m")
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintInGreen(s string) {
	var Green = "\033[32m"
	fmt.Print(Green)
	fmt.Print(fmt.Sprintf("Loading %s", s))
}

type Loader struct {
	Progress int
}

func (l *Loader) Load() {
	if l.Progress < len(Arr) {
		l.Progress++
	}
}

func (l *Loader) Display() {
	ClearScreen()
	ResetScreen()
	PrintInGreen(Arr[l.Progress%len(Arr)])
}

func (l *Loader) IsComplete() bool {
	return l.Progress == len(Arr)
}
func main() {
	l := Loader{
		Progress: 0,
	}
	ch := make(chan bool)
	go func() {
		for !l.IsComplete() {
			l.Load()
			l.Display()
			time.Sleep(time.Second)
		}
		ch <- true
	}()
	<-ch
}
