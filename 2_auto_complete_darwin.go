//go:build darwin
// +build darwin

package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "reflect"
    "strings"
)

func (cmd *cmdGroup) _setUpAutoComplete() {
    args := os.Args
    binName := args[0]
    oldBinName := binName
    binName = strings.ToLower(binName)
    homeDir, err := os.UserHomeDir()
    if err != nil {
        panic("[sl5iyabpoh] " + err.Error())
    }
    fileName := fmt.Sprintf("%s-completion.sh", binName)
    fileName = filepath.Join(homeDir, fileName)
    fileContent := fmt.Sprintf(`#!/bin/bash

function _%s() {
        COMPREPLY=( $(%s %s $COMP_CWORD ${COMP_WORDS[@]})  )
}

complete -o default -F _%s %s
`, oldBinName, oldBinName, commandAutoComplete, oldBinName, oldBinName)

    err = os.WriteFile(fileName, []byte(fileContent), 0777)
    if err != nil {
        panic("[wq77mdo7fp] " + err.Error())
    }
    exec.Command("source", fileName).Run()
}

func (cmd *cmdGroup) _autoComplete() {
    if len(os.Args) <= 3 {
        return
    }
    binName := os.Args[0]
    args := os.Args[3:]
    if args[0] != binName {
        return
    }
    if len(args) == 1 || len(args) == 2 {
        var prefix string
        if len(args) == 2 {
            prefix = args[1]
            logic, ok := cmd.logicMap[prefix]
            if ok {
                cmd.walkLogicParaGoodField(logic, func(index int, field reflect.StructField) {
                    fmt.Println("-" + field.Name)
                })
            }
        }
        for _, subName := range cmd.subNameSlice() {
            if prefix == "" {
                fmt.Println(subName)
                continue
            }
            if subName == prefix {
                continue
            }
            if strings.HasPrefix(subName, prefix) ||
                strings.HasPrefix(strings.ToLower(subName), strings.ToLower(prefix)) {
                fmt.Println(subName)
            }
        }
        return
    }
    sub := args[1]
    prefix := args[len(args)-1]
    isConfig := strings.HasPrefix(prefix, "-")
    prefix = strings.TrimLeft(prefix, "-")
    prefix = strings.ToLower(prefix)
    logic, ok := cmd.logicMap[sub]
    if ok == false {
        return
    }
    havePara := map[string]struct{}{}
    for _, one := range args {
        if strings.HasPrefix(one, "-") {
            havePara[strings.ToLower(strings.TrimLeft(one, "-"))] = struct{}{}
        }
    }
    if isConfig {
        cmd.walkLogicParaGoodField(logic, func(index int, field reflect.StructField) {
            fieldNameLow := strings.ToLower(field.Name)
            _, ok := havePara[fieldNameLow]
            if ok {
                return
            }
            if strings.HasPrefix(fieldNameLow, prefix) {
                fmt.Println("-" + field.Name)
            }
        })
        return
    }
    cmd.walkLogicParaGoodField(logic, func(index int, field reflect.StructField) {
        fieldNameLow := strings.ToLower(field.Name)
        _, ok := havePara[fieldNameLow]
        if ok {
            return
        }
        fmt.Println("-" + field.Name)
    })
}
