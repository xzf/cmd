package main

import (
    "fmt"
    "github.com/xzf/cmd"
)

func main() {
    cg := cmd.NewCmdGroup()
    cg.Add("a", func() {
        fmt.Println("aaa")
    })
    cg.Add("b", func() {
        fmt.Println("bbb")
    })
    type CReq struct {
        ConfigFile string
    }
    cg.Add("c", func(req CReq) {
        fmt.Println("ccc", req.ConfigFile)
    })
    cg.Run()
}
