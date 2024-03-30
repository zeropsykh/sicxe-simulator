package main

import (
    "fmt"
    "os"

    "sicxesimulator/sicxelexer"
)

func main() {
    fileName := "example.asm"
    fileContent, err := os.ReadFile(fileName)
    if err != nil {
        fmt.Printf("Could not open file %s\n", fileName)
        os.Exit(1)
    }  
    
    lexer := sicxelexer.NewLexer(string(fileContent))
    for token := lexer.NextToken(); token.Type != sicxelexer.TokenEOF; token = lexer.NextToken() {
        lexer.Tokens = append(lexer.Tokens, token) 
    }
    lexer.PrintTokens()
}

