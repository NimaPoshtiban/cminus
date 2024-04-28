# cminus

This is an interpreter for a minimal dynamically typed programming language. It currently supports the following data types:

Integers (no floats)
Boolean (true, false)
Strings (enclosed with double quotes)
All C defined Operators except ++ & --

## Built-in Functions :

The following built-in functions are currently available to use:

- len() -> returns the length of the string or the array.
- push() -> adds the element to the end of the array.
- join() -> join two arrays into a single array.
- delete() -> returns the array without the deleted item.
- print() -> prints the data to the console.
- exit() -> exits the program by causing it to panic.
- exec() -> execute system command
- info() -> returns system info
- flush() -> clears the console
- str() -> converts types to string
- int() -> converts string to integer

## Operators:
- "=" -> assignment operator
- "%" -> modulo operator
- "-" -> minus operator
- "/" -> division operator
- "*" -> multiplaction operator
- "+" -> sum operator
- ">>" -> right shift operator
- "<<" -> left shift operator
- "!" -> not operator
- "and" -> logical and operator (&&)
- "or" -> logical or operator (||)
- All LT GT Operators (<= => > < == != )

## Built-in Macro Functions:
- quote(expression) -> returns the unevaluated expression
- unquote(expression) -> evaluate the expression inside quote

### Syntax

- Defining a variable:
  ```
  let foo = 20;
  ```
- Defining a function:
  ```
  let sum = func(x,y) { x + y };
  let sum = func(x,y) {
     return x + y
  };
  ```
- Using recursion:
  ```
  let factorial = func(n) { if (n == 0) { 1 } else { n * factorial(n - 1) } };
  ```
- Using while loop:
  ```
  let j = 10
  while(j > 0){
      print(j)
      let j = j - 1
  }
  ```
  
- Using Arrays:
  ```
  let numbers = [1,3,7,9];
  ```
- Using Hashes:
  ```
  let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
  ```
- Defining a macro:
  ```
  let unless = macro(condition, consequence, alternative){ quote(if (!(unquote(condition))) { unquote(consequence); }else { unquote(alternative); }); };
  ```
#### Interpret a file

```
cminus  --interpret filename
```
