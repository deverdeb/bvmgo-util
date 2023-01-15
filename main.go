package main

import (
	"fmt"
	"github.com/deverdeb/bvmgo-util/logs"
	"github.com/deverdeb/bvmgo-util/term"
)

func main() {
	fmt.Println(term.Blue.Sprint("Hello"), term.White.Sprint("the"), term.Red.Sprint("Hello"))

	fmt.Println(term.Black.Sprint("Black"))
	fmt.Println(term.DarkGray.Sprint("DarkGray"))
	fmt.Println(term.Red.Sprint("Red"))
	fmt.Println(term.LightRed.Sprint("LightRed"))
	fmt.Println(term.Green.Sprint("Green"))
	fmt.Println(term.LightGreen.Sprint("LightGreen"))
	fmt.Println(term.Orange.Sprint("Orange"))
	fmt.Println(term.Yellow.Sprint("Yellow"))
	fmt.Println(term.Blue.Sprint("Blue"))
	fmt.Println(term.LightBlue.Sprint("LightBlue"))
	fmt.Println(term.Purple.Sprint("Purple"))
	fmt.Println(term.LightPurple.Sprint("LightPurple"))
	fmt.Println(term.Cyan.Sprint("Cyan"))
	fmt.Println(term.LightCyan.Sprint("LightCyan"))
	fmt.Println(term.LightGray.Sprint("LightGray"))
	fmt.Println(term.White.Sprint("White"))

	for color := 30; color < 50; color++ {
		for style := 0; style < 10; style++ {
			fmt.Printf("\u001B[%d;%dm %d;%d  \u001B[0m\n", style, color, style, color)
		}
	}

	term.DisplayInformations()

	toto := logs.New("main")
	toto.Info("hello", "world", "...")
}
