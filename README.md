# interpreter

This is an interpreter for a minimal dynamically-typed programming language. It currently supports the following data types:

Integers (no floats)
Boolean (true, false)
Strings (enclosed with double quotes)



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

## Built-in Functions :

The following built-in functions are currently available to use:

- len()    -> returns the length of the string or the array.
- push()   -> adds the element to the end of the array.
- join()   -> join two arrays into a single array.
- delete() -> returns the array without the deleted item.
- print()  -> prints the data to console.
- exit()   -> exits the program by causing it to panic.

### Syntax

- Defining a variable:
  ```
  let foo = 20
  ```
- Defining a function:
  ```
  let sum = fn(x,y) { x + y }
  ```
- Using recursion:
  ```
  let factorial = fn(n) { if (n == 0) { 1 } else { n * factorial(n - 1) } }
  ```
- Using Arrays:
 ```
  let numbers = [1,3,7,9]
 ```