package cmd

import (
    "fmt"
    "os"
    "reflect"
)

//Add add command
//name: command name
//logicFunc: only support type func() func(struct)
func (cmd *cmdGroup) Add(name string, logicFunc interface{}) {
    err := cmd.checkInput(name, logicFunc)
    if err != nil {
        panic(err)
    }
    cmd.logicMap[name] = logicFunc
}

//Run call in main package
func (cmd *cmdGroup) Run() {
    argLen := len(os.Args)
    if argLen == 1 {
        cmd.printHelp("")
        return
    }
    subCommName := os.Args[1]
    if subCommName == "--help" {
        cmd.printSubCommand()
        return
    }
    logic, ok := cmd.logicMap[subCommName]
    if ok == false {
        cmd.printSubCommand()
        fmt.Println("command", "["+subCommName+"]", "not fund")
        return
    }
    help := false
    for _, item := range os.Args {
        if item == "--help" {
            help = true
            break
        }
    }
    if help {
        cmd.printHelp(subCommName)
        return
    }

    logicCal := reflect.ValueOf(logic)
    logicType := reflect.TypeOf(logic)
    if logicType.NumIn() == 0 {
        logicCal.Call(nil)
        return
    }
    argInfo, err := parseArgs(os.Args[1:])
    if err != nil {
        fmt.Println("[kc9q6p24df] parseArgs failed:", err)
        return
    }
    valPtr, err := cmd.argsToParaObjValue(subCommName, argInfo)
    if err != nil {
        fmt.Println("[k2bz3ybv7v]", err)
        return
    }
    logicCal.Call([]reflect.Value{*valPtr})
}
