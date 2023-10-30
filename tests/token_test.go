package tests

import (
	"testing"

	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)

func TestCreateToken(t *testing.T) {
	// test case 1:
	token1 := *Scanner.CreateToken(Scanner.SPECIALSYMBOL, "<")
	if token1.TokenType != Scanner.LT.String() {
		t.Errorf("CreateToken Failed to set the correct type of the given token")
	}
	// test case 2:
	token2 := *Scanner.CreateToken(Scanner.SPECIALSYMBOL, "-")
	if token2.TokenType != Scanner.MINUS.String() {
		t.Errorf("CreateToken Failed to recognize MINUS sign")
	}
	// test case 3:
	token3 := *Scanner.CreateToken(Scanner.RESERVED, "REaD")
	if token3.TokenType != Scanner.UNKOWN {
		t.Errorf("CreateToken Failed to recognize unknown token")
	}

	// test case 4:
	token4 := *Scanner.CreateToken(Scanner.IDENTIFIER, "ahmed")
	if token4.TokenType != Scanner.IDENTIFIER.String() {
		t.Errorf("CreateToken Failed to recognize Identifier")
	}

	// test case 5:
	token5 := *Scanner.CreateToken(Scanner.NUMBER, "55")
	if token5.TokenType != Scanner.NUMBER.String() {
		t.Errorf("CreateToken Failed to recognize Number")

	}
	token6 := *Scanner.CreateToken(Scanner.ERROR, "@")
	if token6.TokenType != Scanner.ERROR.String() {
		t.Errorf("CreateToken Failed to recognize Error State of Token")

	}
}
