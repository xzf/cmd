package cmd

import (
    "errors"
    "fmt"
    "strings"
)

func parseArgs(argSlice []string) (*argsInfo, error) {
    fmt.Println("???", argSlice)
    argLen := len(argSlice)
    if argLen == 0 {
        return nil, errors.New(`[4p0ahdq9u7] len(os.Args) == 0`)
    }
    result := &argsInfo{
        name: argSlice[0],
    }
    if argLen == 1 {
        return result, nil
    }
    index := 1
    var oneKey string
    var oneValueSlice []string
    for ; ; index++ {
        fmt.Println("[cnclo07hx8]", index, "["+oneKey+"]", oneValueSlice)
        if index >= argLen {
            break
        }
        thisStr := argSlice[index]
        if thisStr == "" {
            continue
        }
        isArgKey := strings.HasPrefix(thisStr, "-")
        if isArgKey {
            if oneKey != "" {
                fmt.Println("???", oneValueSlice)
                result.configSlice = append(result.configSlice, kv{
                    k: oneKey,
                    v: strings.Join(oneValueSlice, " "),
                })
            }
            oneKey = strings.TrimLeft(thisStr, "-")
            oneValueSlice = nil
            continue
        }
        oneValueSlice = append(oneValueSlice, thisStr)
    }
    return result, nil
}

type argsInfo struct {
    name        string
    configSlice []kv
}

type kv struct {
    k string
    v string
}
