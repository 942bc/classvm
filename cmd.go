package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0.0"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	classPath   string
	class       string
	XjreOption  string
	args        []string
}

//解析命令行参数
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.classPath, "cp", "", "classpath")
	flag.StringVar(&cmd.classPath, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
