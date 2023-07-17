package cmd

import (
    "errors"
    "fmt"
    "reflect"
    "sort"
    "strconv"
)

const (
    commandSetUpAutoComplete = "_set_up_auto_complete"
    commandAutoComplete      = "_auto_complete"
)

func (cmd *cmdGroup) checkInput(name string, logic interface{}) error {
    if name == "" {
        return errors.New(`[piwt0yq1ms] name == ""`)
    }
    _, have := cmd.logicMap[name]
    if have {
        return errors.New("[bvrwop753a] duplicate command name [" + name + "]")
    }
    logicType := reflect.TypeOf(logic)
    if logicType.Kind() != reflect.Func {
        return errors.New("[gdu9x456sv] logic expect func, get: [" + logicType.Kind().String() + "]")
    }
    paraNum := logicType.NumIn()
    switch paraNum {
    case 0:
        return nil
    case 1:
        paraType := logicType.In(0)
        if paraType.Kind() != reflect.Struct {
            return errors.New("[ntk5vsw7g4] logic func have 1 parameter, expect struct get:[" + paraType.Kind().String() + "]")
        }
        fieldNum := paraType.NumField()
        if fieldNum == 0 {
            return nil
        }
        for i := 0; i < fieldNum; i++ {
            field := paraType.Field(i)
            if field.Anonymous {
                continue
            }
            fieldKind := field.Type.Kind()
            if cmd.checkParaFieldKind(fieldKind) {
                //good
                continue
            }
            //bad
            return errors.New("[n5jwpigmhs] parameter field [" + field.Name + "] type [" + fieldKind.String() + "] not support")
        }
        return nil
    default:
        return errors.New("[qtdlrqwznw] logic func expect 0 or 1 parameter, get:[" + strconv.Itoa(paraNum) + "]")
    }
}

var _supportParaFieldKindMap = map[reflect.Kind]struct{}{
    reflect.Bool:    {},
    reflect.Int:     {},
    reflect.Int64:   {},
    reflect.Float64: {},
    reflect.String:  {},
}

//checkParaFieldKind
//return false mean not support
func (cmd *cmdGroup) checkParaFieldKind(kind reflect.Kind) bool {
    _, ok := _supportParaFieldKindMap[kind]
    return ok
}

func (cmd *cmdGroup) walkLogicParaGoodField(obj interface{}, cb func(int, reflect.StructField)) {
    if cb == nil {
        return
    }
    logicType := reflect.TypeOf(obj)
    if logicType.Kind() != reflect.Func {
        panic("[ip5vk4xrq7] obj type expect func,get:[" + logicType.Kind().String() + "]")
    }
    if logicType.NumIn() == 0 {
        return
    }
    paraType := logicType.In(0)
    fieldNum := paraType.NumField()
    if fieldNum == 0 {
        return
    }
    for i := 0; i < fieldNum; i++ {
        field := paraType.Field(i)
        if field.Anonymous {
            continue
        }
        _, ok := _supportParaFieldKindMap[field.Type.Kind()]
        if ok == false {
            continue
        }
        cb(i, field)
    }
}

func (cmd *cmdGroup) argsToParaObjValue(name string, argInfo *argsInfo) (*reflect.Value, error) {
    logic, ok := cmd.logicMap[name]
    if ok == false {
        return nil, errors.New("[f0qiv4u3nm] command [" + name + "] not found")
    }
    kvMap := map[string]string{}
    for _, kv := range argInfo.configSlice {
        kvMap[kv.k] = kv.v
    }
    logicType := reflect.TypeOf(logic)
    paraType := logicType.In(0)
    value := reflect.New(paraType)
    if value.Type().Kind() == reflect.Ptr {
        value = value.Elem()
    }
    cmd.walkLogicParaGoodField(logic, func(index int, field reflect.StructField) {
        fieldName := field.Name
        val, ok := kvMap[fieldName]
        if ok == false {
            return
        }
        fieldValue := value.FieldByName(fieldName)
        if fieldValue.CanSet() == false {
            return
        }
        switch field.Type.Kind() {
        case reflect.String:
            fieldValue.SetString(val)
        case reflect.Bool:
            boolValue, err := strconv.ParseBool(val)
            if err != nil {
                fmt.Println("[o2wqdd40rn] reflect set bool failed,name:[" + fieldName + "] val:[" + val + "]")
                return
            }
            fieldValue.SetBool(boolValue)
        case reflect.Int,
            reflect.Int8,
            reflect.Int16,
            reflect.Int32,
            reflect.Int64:
            int64Val, err := strconv.ParseInt(val, 10, 64)
            if err != nil {
                fmt.Println("[83vnn251yb] reflect set [" + field.Type.Kind().String() + "] failed,name:[" + fieldName + "] val:[" + val + "]")
                return
            }
            fieldValue.SetInt(int64Val)
        case reflect.Float32,
            reflect.Float64:
            floatVal, err := strconv.ParseFloat(val, 64)
            if err != nil {
                fmt.Println("[vme00vhqiw] reflect set [" + field.Type.Kind().String() + "] failed,name:[" + fieldName + "] val:[" + val + "]")
                return
            }
            fieldValue.SetFloat(floatVal)
        default:
            //Generally, never reach here
        }
    })
    return &value, nil
}

func (cmd *cmdGroup) printHelp(name string) {
    if name == "" {
        cmd.printSubCommand()
        return
    }
    logic, ok := cmd.logicMap[name]
    if ok {
        fmt.Println("command [" + name + "]:")
        cmd.walkLogicParaGoodField(logic, func(_ int, field reflect.StructField) {
            helpInfo := "-" + field.Name
            cmdTagContent, ok := field.Tag.Lookup("cmd")
            if ok == false || cmdTagContent == "" {
                fmt.Println("\t" + helpInfo)
                return
            }
            fmt.Println("\t" + helpInfo + " " + cmdTagContent)
        })
        return
    }
    fmt.Println("[51x4p0qiwr]", name)
}

func (cmd *cmdGroup) printSubCommand() {
    fmt.Println("help: --help")
    fmt.Println("sub command:")
    for _, name := range cmd.subNameSlice() {
        fmt.Println(name)
    }
}

func (cmd *cmdGroup) subNameSlice() []string {
    var nameSlice []string
    for name := range cmd.logicMap {
        if name == commandAutoComplete {
            continue
        }
        nameSlice = append(nameSlice, name)
    }
    sort.Strings(nameSlice)
    return nameSlice
}
