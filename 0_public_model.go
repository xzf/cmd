package cmd

type CMD struct {
    Name   string
    Logic  interface{} //
    SubCMD []CMD       //sub cmd
}
