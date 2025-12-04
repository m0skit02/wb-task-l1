package main

import "fmt"

type OldPrinter struct{}

func (OldPrinter) Print(text string) {
	fmt.Println("Old printer output:", text)
}

type Logger interface {
	Log(message string)
}

type PrinterAdapter struct {
	oldPrinter OldPrinter
}

func (a PrinterAdapter) Log(message string) {
	a.oldPrinter.Print(message)
}

func main() {
	old := OldPrinter{}

	var logger Logger = PrinterAdapter{oldPrinter: old}

	logger.Log("Привет! Это лог через адаптер.")
}
