package main

import (
    "fmt"
    "github.com/xzf/cmd"
)

func main() {
    cg := cmd.NewCmdGroup()
    cg.Add("aa", func() {
        fmt.Println("aaa")
    })
    cg.Add("ab", func() {
        fmt.Println("ab")
    })
    cg.Add("ac", func() {
        fmt.Println("ac")
    })
    cg.Add("bb", func() {
        fmt.Println("bbb")
    })
    type CReq struct {
        ConfigFile string `cmd:"config file path"`
        LogPath    string `cmd:"log file write path"`
    }
    cg.Add("cc", func(req CReq) {
        fmt.Println("ccc", req.ConfigFile)
    })
    cg.Run()
}
