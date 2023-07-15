package main

import "github.com/xzf/cmd"

func main() {
    binTree := cmd.CMD{
        Name: "A",
        SubCMD: []cmd.CMD{
            {
                Name: "a",
                Logic: func(i interface{}) {

                },
            },
        },
    }
}
