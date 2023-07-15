package cmd

import (
    "errors"
    "fmt"
    "os"
    "reflect"
    "strconv"
    "strings"
)

func (cmd CMD) checkInput(name string, logic interface{}) error {
    if name == "" {
        return errors.New("9wnkqp4lts")
    }
    for _, item := range cmd.SubCMD {
        if item.Name == name {
            return errors.New("bvrwop753a")
        }
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

//func (cmd CMD) Sub(name string) *CMD {
//    cmd.sub = append(cmd.sub, &CMD{
//        name: name,
//    })
//}

//func (cmd CMD) SetDescribe(desc string) {
//
//}

func (cmd CMD) subMap() map[string]CMD {
    result := map[string]CMD{}
    for _, item := range cmd.SubCMD {
        result[item.Name] = item
    }
    return result
}

func (cmd CMD) run(args []string) {
    argLen := len(args)
    if argLen == 0 {
        cmd.printHelp()
        return
    }
    firstArgStr := args[1]
    isConfig := strings.HasPrefix(firstArgStr, "-")
    if isConfig == false {
        subMap := cmd.subMap()
        subCmd, ok := subMap[firstArgStr]
        if ok {
            fmt.Println("389ah05kbb")
            return
        }
        subCmd.run(args[1:])
        return
    }
    argInfo, err := parseArgs(args)
    if err != nil {
        fmt.Println("kc9q6p24df", err)
        return
    }
    valPtr, err := cmd.argsToParaObjValue(argInfo)
    if err != nil {
        fmt.Println("o5181h5wtc", err)
        return
    }
    logicCal := reflect.ValueOf(cmd.Logic)
    logicCal.Call([]reflect.Value{*valPtr})
}

func (cmd CMD) Run() {
    cmd.checkInput(cmd.Name, cmd.Logic)
    for _, item := range cmd.SubCMD {
        cmd.checkInput(item.Name, item.Logic)
    }
    cmd.run(os.Args)
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
func (cmd CMD) checkParaFieldKind(kind reflect.Kind) bool {
    _, ok := _supportParaFieldKindMap[kind]
    return ok
}

func (cmd CMD) walkLogicParaGoodField(cb func(int, reflect.StructField)) {
    if cb == nil {
        return
    }
    logicType := reflect.TypeOf(cmd.Logic)
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

func (cmd CMD) argsToParaObjValue(argInfo *argsInfo) (*reflect.Value, error) {
    kvMap := map[string]string{}
    for _, kv := range argInfo.configSlice {
        kvMap[kv.k] = kv.v
    }
    logicType := reflect.TypeOf(cmd.Logic)
    paraType := logicType.In(0)
    value := reflect.New(paraType)
    cmd.walkLogicParaGoodField(func(index int, field reflect.StructField) {
        fieldName := field.Name
        val, ok := kvMap[fieldName]
        if ok == false {
            fmt.Println("9hvmsalhpz")
            return
        }
        fieldValue := value.FieldByName(fieldName)
        if fieldValue.CanSet() == false {
            fmt.Println("92a4f9ace6")
            return
        }
        switch field.Type.Kind() {
        case reflect.String:
            fieldValue.SetString(val)
        case reflect.Bool:
            boolValue, err := strconv.ParseBool(val)
            if err != nil {
                fmt.Println("o2wqdd40rn")
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
                fmt.Println("my5zy9yixy")
                return
            }
            fieldValue.SetInt(int64Val)
        case reflect.Float32,
            reflect.Float64:
            floatVal, err := strconv.ParseFloat(val, 64)
            if err != nil {
                fmt.Println("swu45be03i")
                return
            }
            fieldValue.SetFloat(floatVal)
        default:
            //Generally, never reach here
        }
    })
    return &value, nil
}

func (cmd CMD) printHelp() {
    if len(cmd.SubCMD) == 0 {
        cmd.printLogicHelp()
        return
    }
    fmt.Println("sub command:")
    for _, subCmd := range cmd.SubCMD {
        fmt.Println(subCmd.Name)
    }
}

func (cmd CMD) printLogicHelp() {
    cmd.walkLogicParaGoodField(func(_ int, field reflect.StructField) {
        helpInfo := "-" + field.Name
        cmdTagContent, ok := field.Tag.Lookup("cmd")
        if ok == false || cmdTagContent == "" {
            fmt.Println(helpInfo)
            return
        }
        fmt.Println(helpInfo + "   " + cmdTagContent)
    })
}
