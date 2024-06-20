// This file contains built-in functions
// Add Built-in functions here
package evaluator

// #include "cfunctions.h"
import "C"

import (
	"cminus/object"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"unsafe"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to len not supported. got %s", args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return &object.Null{}
		},
	},
	"alert":{
		Fn: func (args... object.Object)object.Object  {
			if runtime.GOOS == "windows" {
				C.alert(C.CString(args[0].Inspect()));
				return &object.Null{}
			}
			return &object.String{"Only implemented for windows"}
		},
	},
	"exit": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect() + " ")
			}
			log.Println()
			return nil
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"delete": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `delete` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			indexObj, ok := args[1].(*object.Integer)
			if !ok {
				return newError("second argument to `delete` must be INTEGER, got %s",
					args[1].Type())
			}
			index := int(indexObj.Value)
			if index < 0 || index >= len(arr.Elements) {
				return NULL
			}
			newElements := make([]object.Object, len(arr.Elements)-1)
			copy(newElements[:index], arr.Elements[:index])
			copy(newElements[index:], arr.Elements[index+1:])
			return &object.Array{Elements: newElements}
		},
	},
	"join": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ || args[1].Type() != object.ARRAY_OBJ {
				return newError("arguments to `join` must be ARRAY, got %s and %s",
					args[0].Type(), args[1].Type())
			}
			arr1 := args[0].(*object.Array)
			arr2 := args[1].(*object.Array)
			newElements := make([]object.Object, len(arr1.Elements)+len(arr2.Elements))
			copy(newElements, arr1.Elements)
			copy(newElements[len(arr1.Elements):], arr2.Elements)
			return &object.Array{Elements: newElements}
		},
	},
	"exec": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("wrong number of arguments. got=%d, want=1 or more", len(args))
			}
			for i := 0; i < len(args); i++ {
				if args[i].Type() != object.STRING_OBJ {
					return newError("argument to `cmd` must be STRING, got=%s", args[i].Type())
				}
			}
			str := args[0].(*object.String)
			cmd := exec.Command(str.Value)
			out, _ := cmd.Output()
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}

			return &object.String{Value: string(out)}
		},
	},
	"info": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}
			host, _ := os.Hostname()
			arch := runtime.GOARCH
			cpu := runtime.NumCPU()
			os := runtime.GOOS
			out := fmt.Sprintf("Hostname: %s\nArchitecture: %s\nCpu Cores: %d\nOperating System: %s", host, arch, cpu, os)
			return &object.String{Value: out}
		},
	},
	"flush": {
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: "\033[H\033[2J"}
		},
	},

	"help": {
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: `
				len() -> returns the length of the string or the array.
				push() -> adds the element to the end of the array.
				join() -> joins two arrays into a single array.
				delete() -> returns the array without the deleted item.
				print() -> prints the data to the console.
				exit() -> exits the program by causing it to panic.
				exec() -> execute system command
				info() -> returns system info
				flush() -> clears the console
				str() -> converts types to string
				int() -> converts string to integer
										`}
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: fmt.Sprintf("%v", args[0].Inspect())}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `int` must be an STRING, got=%s", args[0].Type())
			}
			result, _ := strconv.Atoi(args[0].Inspect())
			return &object.Integer{int64(result)}
		},
	},
	"sizeof": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			size := unsafe.Sizeof(args[0])
			return &object.Integer{int64(size)}
		},
	},
}
