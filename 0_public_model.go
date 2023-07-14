package cmd

type CMD struct {
    name  string
    logic func(interface{}) //
    sub   map[string]*CMD   //sub cmd
}

//type Args interface {
//}
