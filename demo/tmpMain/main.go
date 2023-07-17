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
    comm.Run()
}
