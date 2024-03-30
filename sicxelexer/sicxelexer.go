package sicxelexer

import (
    "fmt"
    "unicode"

    "sicxesimulator/sicxelib"
)

type TokenType int

const (
    // default 
    TokenSymbol TokenType = iota
    TokenInstruction
    TokenDirective
    TokenComment
    TokenNumber

    // newline is kind of a delimiter
    TokenNewLine

    // specified before index  
    TokenComma
    
    // specifies immediate mode
    TokenHash
    
    // specifies format 4 instructions or used in EQU assembler directive 
    TokenPlus

    // specifies indirect addressing
    TokenAtSign

    // specifies BYTE assembler directive C'' and X''
    TokenSingleQuote

    // specifies current location or used in EQU assmebler directive  
    TokenAsterisk

    // specifies literal pool
    TokenEqual

    // specified in EQU assembler directives
    TokenMinus
    TokenDivision
    TokenLeftBracket
    TokenRightBracket

    // EOF
    TokenEOF

    // invalid token
    TokenInvalid
)

type Token struct {
    Type TokenType
    Value string
}

type Lexer struct {
    fileContent string
    pos int
    Tokens []Token
}

func NewLexer(fileContent string) *Lexer {
    return &Lexer{
        fileContent: fileContent,
        pos: 0,
        Tokens: []Token{},
    }
}

func (l Lexer) PrintTokens() {
    fmt.Println("TokenType\tToken")
    fmt.Println("----------------------")
    for _, token := range l.Tokens {
        switch token.Type {
        case TokenSymbol:
            fmt.Printf("SYMBOL\t\t%v\n", token.Value)
        case TokenInstruction:
            fmt.Printf("INSTRUCTION\t%v\n", token.Value)
        case TokenDirective:
            fmt.Printf("DIRECTIVE\t%v\n", token.Value)
        case TokenComment:
            fmt.Printf("COMMENT\t\t%v\n", token.Value)
        case TokenNumber:
            fmt.Printf("NUMBER\t\t%v\n", token.Value)
        case TokenNewLine:
            fmt.Printf("NEW LINE\t\\n\n")
        case TokenComma:
            fmt.Printf("COMMA\t\t%v\n", token.Value)
        case TokenHash:
            fmt.Printf("POUND\t\t%v\n", token.Value)
        case TokenPlus:
            fmt.Printf("PLUS\t%v\n", token.Value)
        case TokenAtSign:
            fmt.Printf("AT SIGN\t%v\n", token.Value)
        case TokenSingleQuote:
            fmt.Printf("SINGLE QUOTE\t%v\n", token.Value)
        case TokenAsterisk:
            fmt.Printf("ASTERISK\t%v\n", token.Value)
        case TokenEqual:
            fmt.Printf("EQUAL\t%v\n", token.Value)
        case TokenMinus:
            fmt.Printf("MINUS\t%v\n", token.Value)
        case TokenDivision:
            fmt.Printf("DIVISION\t%v\n", token.Value)
        case TokenLeftBracket:
            fmt.Printf("LEFT BRACKET\t%v\n", token.Value)
        case TokenRightBracket:
            fmt.Printf("RIGHT BRACKET\t%v\n", token.Value)
        case TokenEOF:
            fmt.Printf("EOF\t\tEnd of Line\n")
        case TokenInvalid: 
            fmt.Printf("INVALID")
        default:
            fmt.Printf("UNKNOWN")
        }
    }
}

func isDirective(symbol string) bool {
    switch symbol {
    case "START": 
        fallthrough
    case "WORD": 
        fallthrough
    case "RESW": 
        fallthrough
    case "BYTE": 
        fallthrough
    case "RESB": 
        fallthrough
    case "END":
        return true
    default:
        return false
    }
}

func (l *Lexer) NextToken() Token {
    if l.pos >= len(l.fileContent) {
        return Token{Type: TokenEOF, Value: ""}
    }  
    l.trimWhitespace()

    var token Token
    currentChar := l.fileContent[l.pos]
    switch {
    case unicode.IsLetter(rune(currentChar)):
        token = l.parseSymbol()
        optab := sicxelib.GenerateOptab()
        if _, found := optab.LookUp(token.Value); found {
            token.Type = TokenInstruction
        } else if isDirective(token.Value) {
            token.Type = TokenDirective 
        }
    case currentChar == '.':
        token = l.parseComment()
    case unicode.IsDigit(rune(currentChar)):
        token = l.parseNumber() 
    case currentChar == '\n':
        l.pos++
        token = Token{Type: TokenNewLine, Value: "\n"}
    case currentChar == ',':
        l.pos++
        token = Token{Type: TokenComma, Value: "."}
    case currentChar == '#':
        l.pos++
        token = Token{Type: TokenComma, Value: "#"}
    case currentChar == '+':
        l.pos++
        token = Token{Type: TokenPlus, Value: "+"}
    case currentChar == '@':
        l.pos++
        token = Token{Type: TokenAtSign, Value: "@"}
    case currentChar == '\'':
        l.pos++
        token = Token{Type: TokenSingleQuote, Value: "'"}
    case currentChar == '*':
        l.pos++
        token = Token{Type: TokenAsterisk, Value: "@"}
    case currentChar == '=':
        l.pos++
        token = Token{Type: TokenEqual, Value: "="}
    case currentChar == '-':
        l.pos++
        token = Token{Type: TokenMinus, Value: "-"}
    case currentChar == '\\':
        l.pos++
        token = Token{Type: TokenDivision, Value: "\\"}
    case currentChar == '(':
        l.pos++
        token = Token{Type: TokenLeftBracket, Value: "("}
    case currentChar == ')':
        l.pos++
        token = Token{Type: TokenRightBracket, Value: ")"}
    default:
        l.pos++
        token = Token{Type: TokenInvalid, Value: ""}
    }

    return token
}

func isAlnum(char rune) bool {
    if unicode.IsLetter(char) || unicode.IsDigit(char) {
        return true
    } 
    return false
}

func (l *Lexer) trimWhitespace() {
    for l.pos < len(l.fileContent) && l.fileContent[l.pos] == ' ' {
        l.pos++
    }
}

func (l *Lexer) parseSymbol() Token {
    start := l.pos 
    for l.pos < len(l.fileContent) && isAlnum(rune(l.fileContent[l.pos])) {
        l.pos++
    }
    return Token{Type: TokenSymbol, Value: l.fileContent[start:l.pos]}
}

func (l *Lexer) parseComment() Token {
    start := l.pos
    for l.pos < len(l.fileContent) && l.fileContent[l.pos] != '\n' {
        l.pos++
    }
    return Token{Type: TokenComment, Value: l.fileContent[start:l.pos]}
}

func (l *Lexer) parseNumber() Token {
    start := l.pos
    for l.pos < len(l.fileContent) && unicode.IsDigit(rune(l.fileContent[l.pos])) {
        l.pos++ 
    }
    return Token{Type: TokenNumber, Value: l.fileContent[start:l.pos]} 
}
