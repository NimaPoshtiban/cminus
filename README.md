# interpreter

an Interpreter for my minimal dynamically typed programming language

Supported data types :

1. Integers (no floats)
2. Boolean (true,false)
3. Strings (surrounded with ")

  <!-- ,----..                         
 /   /   \                        
|   :     :     ,---,.     ,---,. 
.   |  ;. /   ,'  .' |   ,'  .' | 
.   ; /--`  ,---.'   , ,---.'   , 
;   | ;     |   |    | |   |    | 
|   : |     :   :  .'  :   :  .'  
.   | '___  :   |.'    :   |.'    
'   ; : .'| `---'      `---'      
'   | '/  :                       
|   :    /                        
 \   \ .'                         
  `---`                           
                                 -->

## Built-in functions :

- len() -> returns the length of the string
- print() -> prints the data to console
- exit() -> program will panic

### Syntax

- defining a variable:
  ```
  let foo = 20;
  ```
- defining a function:
  ```
  let sum = fn(x,y) { x + y };
  ```
- using recursion:
  ```
  let factorial = fn(n) { if (n == 0) { 1 } else { n * factorial(n - 1) } };
  ```