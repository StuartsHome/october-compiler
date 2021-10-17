# Overview
Following a book on how to build a compiler in Go.

## Intro
- 3 steps
1. Given the inputted source code, it will break the code up into tokens
2. It will parse the tokens to make sure they are in an order that is allowed
3. Third it will emit the C code that our language will translate to

String -> Tokens -> AST -> Bytecode -> Objects
Lexer -> Parser -> Compiler


####
compiler that compiles to assembly and uses an external assembler to convert to machine code

Compiler
Traverse the AST we pass in, find the nodes, evaluate them by turning them into objects
add the objects to the constant pool, and finally emit instructions that reference the
constants in said pool