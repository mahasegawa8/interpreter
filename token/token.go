package token 

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT = "INT"     // 123456

    // Operators
    ASSIGN = "="
    PLUS ="+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"

    MINUS       = "-"
    BANG        = "!"
    SLASH       = "/"
    ASTERISK    = "*"
    LT          = "<"
    GT          = ">"
    ILLEGAL     = "ILLEGAL"
    EQ          = "=="
    NOT_EQ      = "!="
)

var keywords = map[string]TokenType{
    "fn":   FUNCTION,
    "let":  LET,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}

