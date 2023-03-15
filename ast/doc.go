/* AST Package Documentation
The ast package defines different types of nodes that interoperate with the lexer and parser in order to build an Abstract Syntax Tree (AST) for the input program.

Types:

Node interface
The Node interface represents a node in the AST. It has two methods:

TokenLiteral() string returns the literal value of the token associated with this node.
String() string returns the string representation of the Node that can be used for debugging and testing purposes.
Statement interface
The Statement interface is implemented by all AST nodes that represent statements in the input program. It extends the Node interface and also adds one method:

statementNode()
Expression interface
The Expression interface is implemented by all AST nodes that represent expressions in the input program. It extends the Node interface and also adds one method:

expressionNode()
Program struct
The Program struct represents the top-level node of the AST. It contains zero or more Statements.
It has two methods:

TokenLiteral() string returns the literal value of the token associated with the first statement in the Statements array.
String() string returns a string that concatenates the strings of all its children using a buffer.
LetStatement struct
The LetStatement struct implements the Statement interface. It represents a let statement in the input program.
It has four properties:

Token token.Token: A LET token from the lexer.
Name *Identifier: The name of the variable that we are assigning to.
Value Expression: The expression that provides the value for the new variable.
It has four methods:
statementNode() is used to implement the Statement interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the LetStatement.
Identifier struct
The Identifier struct implements the Expression interface. It represents an identifier in the input program.
It has two properties:

Token token.Token: An IDENT token from the lexer.
Value string: The actual value of the identifier.
It has three methods:
expressionNode() is used to implement the Expression interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the Identifier.
ReturnStatement struct
The ReturnStatement struct implements the Statement interface. It represents a return statement in the input program.
It has two properties:

Token token.Token: A RETURN token from the lexer.
ReturnValue Expression: The expression that represents the return value.
It has four methods:
statementNode() is used to implement the Statement interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the ReturnStatement.
ExpressionStatement struct
The ExpressionStatement struct implements the Statement interface. It represents an expression statement in the input program.
It has two properties:

Token token.Token: The first token of the expression.
Expression Expression: The expression itself.
It has four methods:
statementNode() is used to implement the Statement interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the ExpressionStatement.
IntegerLiteral struct
The IntegerLiteral struct implements the Expression interface. It represents an integer literal in the input program.
It has two properties:

Token token.Token: the INT token from the lexer.
Value int64: The actual value of the integer.
It has three methods:

expressionNode() is used to implement the Expression interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the IntegerLiteral.
PrefixExpression struct
The PrefixExpression struct implements the Expression interface. It represents prefix expressions in the input program such as !, - , ...
It has three properties:

Token token.Token: The corresponding operator token.
Operator string: the operator itself.
Right Expression: the right hand side expression.
It has four methods:
expressionNode() is used to implement the Expression interface.
TokenLiteral() string returns the literal value of the token.
String() string returns the string representation of the PrefixExpression.
InfixExpression struct
The InfixExpression struct implements the Expression interface. It represents infix expressions in the input program like 1 + 2, a - b ,....
It has four properties:

Token token.Token: Token that represents the operator (+, -, ...)
Left Expression: The left expression of the
*/
package ast