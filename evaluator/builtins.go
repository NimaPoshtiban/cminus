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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
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
			for _, arg := range args {
				fmt.Print(arg.Inspect() + " ")
			}
			fmt.Println()
			return nil
		},
	},
	"push": &object.Builtin{
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
	"delete": &object.Builtin{
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
	"join": &object.Builtin{
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
}
