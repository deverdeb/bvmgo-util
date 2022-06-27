package term

import (
	"fmt"
	"os"
	"runtime"
)

func DisplayInformations() {
	fmt.Printf(White.Sprint("System information:\n"))
	fmt.Printf(" - OS: %s\n", runtime.GOOS)
	fmt.Printf(" - Architecture: %s\n", runtime.GOARCH)
	fmt.Printf(" - Number of logical CPUs usable by the current process: %d\n", runtime.NumCPU())

	fmt.Printf(White.Sprint("Build information:\n"))
	fmt.Printf(" - Go ROOT: %s\n", runtime.GOROOT())
	fmt.Printf(" - Go version: %s\n", runtime.Version())
	fmt.Printf(" - Go compiler: %s\n", runtime.Compiler)

	fmt.Printf(White.Sprint("Path information:\n"))
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf(" - Current path: Error - %v", err)
	} else {
		fmt.Printf(" - Current path: %s\n", path)
	}
	exec, err := os.Executable()
	if err != nil {
		fmt.Printf(" - Current executable: Error - %v", err)
	} else {
		fmt.Printf(" - Current executable: %s\n", exec)
	}

	fmt.Printf(White.Sprint("Command line information:\n"))
	// Command line arguments :
	// Note : voir l'utilisation des "flags" pour l'utilisation des arguments - https://gobyexample.com/command-line-flags
	argsWithProg := os.Args
	argsProg := argsWithProg[0]
	argsWithoutProg := argsWithProg[1:]
	fmt.Printf(" - Command line: %s\n", argsWithProg)
	fmt.Printf("   - Command line executanle: %s\n", argsProg)
	fmt.Printf("   - Command line arguments: %s\n", argsWithoutProg)

	fmt.Printf(White.Sprint("Terminal information:\n"))
	_, exists := os.LookupEnv("NO_COLOR")
	nocolor := exists || os.Getenv("TERM") == "dumb"
	fmt.Printf("   - No color: #{color} %v\n", nocolor)

}
