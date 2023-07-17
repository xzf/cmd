package cmd

func NewCmdGroup() *cmdGroup {
    result := &cmdGroup{
        logicMap: map[string]interface{}{},
    }
    result.logicMap[commandSetUpAutoComplete] = result._setUpAutoComplete
    result.logicMap[commandAutoComplete] = result._autoComplete
    return result
}
