package cmd

func NewCmdGroup() *cmdGroup {
    return &cmdGroup{
        logicMap: map[string]interface{}{},
    }
}
