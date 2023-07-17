package cmd

func NewGroup() *cmdGroup {
    return &cmdGroup{
        logicMap: map[string]interface{}{},
    }
}
