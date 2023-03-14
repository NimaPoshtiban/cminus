// This file contains built-in functions
// Add Built-in functions here
package evaluator

import (
	"fmt"
	"interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			default:
				return newError("argument to len not supported. got %s", args[0].Type())
			}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {

			case *object.String:
				fmt.Println(arg.Value)
				return &object.Null{}
			case *object.Integer:
				str := fmt.Sprintf("%d", arg.Value)
				fmt.Println(str)
				return &object.Null{}
			default:
				return &object.Null{}
			}
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				fmt.Printf("exit with status code %d", arg.Value)
				panic(arg.Value)
			default:
				return newError("invalid argument")
			}
		},
	},
}
