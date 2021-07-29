package main

import (
	"classvm/classpath"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version %s \n", VERSION)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.classPath)
	fmt.Printf("classpath: %v class:%v  args:%v\n", cp.String(), cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Cloud not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data :%v \n", classData)
}
