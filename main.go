package main

import (
	"fmt"
)

type TokenType string

const (
	TOKEN_METADATA_START   = "METADATA_START"
	TOKEN_METADATA_END     = "METADATA_END"
	TOKEN_METADATA_KEY     = "METADATA_KEY"
	TOKEN_METADATA_VALUE   = "METADATA_VALUE"
	TOKEN_COMMENT          = "COMMENT"
	TOKEN_ROLE             = "ROLE"
	TOKEN_PERSON           = "PERSON"
	TOKEN_PERSON_ALIAS     = "PERSON_ALIAS"
	TOKEN_BRAND            = "BRAND"
	TOKEN_ENTRY            = "ENTRY"
	TOKEN_COMMA            = "COMMA"
	TOKEN_ASSIGNMENT_VALUE = "ASSIGNMENT_VALUE"
	TOKEN_ASSIGNMENT_START = "ASSIGNMENT_START"
	TOKEN_ASSIGNMENT_END   = "ASSIGNMENT_END"
	TOKEN_EOF              = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

func main() {
	fmt.Println("sakugo")
}
