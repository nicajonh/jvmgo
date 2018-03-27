package main

import "fmt"
import Cmd "jvmgo/cmd"

func main() {
	cmd := Cmd.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		Cmd.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
}
