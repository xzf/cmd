package cmd

import (
    "fmt"
    "os"
    "reflect"
)

//Add
//logicFunc only support type func() func(in)
func (cmd *cmdGroup) Add(name string, logicFunc interface{}) {
    err := cmd.checkInput(name, logicFunc)
    if err != nil {
        panic(err)
    }
    cmd.logicMap[name] = logicFunc
}

func (cmd *cmdGroup) Run() {
    argLen := len(os.Args)
    if argLen == 1 {
        cmd.printHelp("")
        return
    }
    subCommName := os.Args[1]
    logic, ok := cmd.logicMap[subCommName]
    if ok == false {
        cmd.printSubCommand()
        fmt.Println("command", "["+subCommName+"]", "not fund")
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
