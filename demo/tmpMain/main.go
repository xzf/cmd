package main

import (
    "fmt"
    "github.com/xzf/cmd"
)

func main() {
    comm := cmd.NewGroup()
    comm.AddCommand("a", func() {
        fmt.Println("aaa")
    })
    comm.AddCommand("b", func() {
        fmt.Println("bbb")
    })
    type CReq struct {
        ConfigFile string
    }
    comm.AddCommand("c", func(req CReq) {
        fmt.Println("ccc", req.ConfigFile)
    })
    comm.Run()
}
