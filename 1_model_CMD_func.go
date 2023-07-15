package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//Add wrong call will be panic
func (cmd *CMD) Add(name string, logic interface{}) {
	if name == "" {
		panic(`[vrfe087nfq] name == ""`)
	}
	err := cmd.checkLogic(logic)
	if err != nil {
		panic(err)
	}
	//logicType := reflect.TypeOf(logic)
	//if logicType.Kind() != reflect.Func {
	//    return errors.New("[gdu9x456sv] logic expect func, get: [" + logicType.Kind().String() + "]")
	//}
	// return nil
}

//func (cmd *CMD) Sub(name string) *CMD {
//    cmd.sub = append(cmd.sub, &CMD{
//        name: name,
//    })
//}

//func (cmd *CMD) SetDescribe(desc string) {
//
//}

func (cmd *CMD) Run() {
	argLen := len(os.Args)
	if argLen == 1 {
		cmd.printHelp()
		return
	}
	for i := 1; i < argLen; i++ {
		thisArg := os.Args[i]

	}
	//todo
	if len(cmd.sub) == 0 {
		cmd.printHelp()
		return
	}
	//cmd.logic()
}

func (cmd *CMD) checkLogic(logic interface{}) error {
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
func (cmd *CMD) checkParaFieldKind(kind reflect.Kind) bool {
	_, ok := _supportParaFieldKindMap[kind]
	return ok
}

func (cmd *CMD) walkLogicParaGoodField(cb func(reflect.StructField)) {
	if cb == nil {
		return
	}
	logicType := reflect.TypeOf(cmd.logic)
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
		cb(field)
	}
}

func (cmd *CMD) argsToParaObjValue() (reflect.Value, error) {
	logicType := reflect.TypeOf(cmd.logic)
	paraType := logicType.In(0)
	value := reflect.New(paraType)
	cmd.walkLogicParaGoodField(func(field reflect.StructField) {
		switch field.Type.Kind() {
		case reflect.Bool:
		case reflect.Int:
		case reflect.Int64:
		case reflect.Float64:
		case reflect.String:
		default:
			//Generally, never reach here
		}
	})

	fieldNum := paraType.NumField()
	//have not field will not be consider to error
	if fieldNum == 0 {
		return value, nil
	}
	for i := 0; i < fieldNum; i++ {
		field := paraType.Field(i)
		if field.Anonymous {
			continue
		}

	}
	return value, nil
}

func (cmd *CMD) printHelp() {
	if len(cmd.sub) == 0 {
		cmd.printLogicHelp()
		return
	}
	fmt.Println("sub command:")
	for _, subCmd := range cmd.sub {
		fmt.Println(subCmd.name)
	}
}

func (cmd *CMD) printLogicHelp() {
	//if len(cmd.sub) != 0 {
	//    cmd.printHelp()
	//    return
	//}
	cmd.walkLogicParaGoodField(func(field reflect.StructField) {
		helpInfo := "-" + field.Name
		cmdTagContent, ok := field.Tag.Lookup("cmd")
		if ok == false || cmdTagContent == "" {
			fmt.Println(helpInfo)
			return
		}
		fmt.Println(helpInfo + "   " + cmdTagContent)
	})
}
