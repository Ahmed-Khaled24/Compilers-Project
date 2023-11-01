package Scanner

import (
	"fmt"
	"strconv"
	"strings"
)

type Event string
type Action struct {
	Destination State
	Callback    func(*string, *string, *[]Token) *Token
}
type State string
type Condition string
type Transition map[Event]Action
type StateMap map[State]Transition

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isReservedWord(s string) bool {
	return handle_revserved_words(strings.TrimSpace(s)) != UNKOWN
}

func isSpecialSymbols(s string) bool {
	return handle_special_symbols(strings.TrimSpace(s)) != UNKOWN
}

const (
	START      State = "START"
	IN_COMMENT State = "IN_COMMENT"
	IN_NUM     State = "IN_NUM"
	IN_ID      State = "IN_ID"
	IN_ASSIGN  State = "IN_ASSIGN"
	SPECIAL_S  State = "SPECIAL_S"
	DONE       State = "DONE"
	ERR        State = "ERR"
)

type ScannerStruct struct {
	Initial   State
	Current   State
	StateMap  StateMap
	TokenList []Token
	Pointer   int
	Input     string
}

var Fsm = ScannerStruct{
	Initial: START,
	Pointer: 0,
	Input:   "",
	StateMap: StateMap{
		START: Transition{
			"Number": Action{
				Destination: IN_NUM,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"Letter": Action{
				Destination: IN_ID,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"Colon": Action{
				Destination: IN_ASSIGN,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"SpecialSymbol": Action{
				Destination: SPECIAL_S,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"OpenBrace": Action{
				Destination: IN_COMMENT,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"WhiteSpace": Action{
				Destination: START,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					// *StringInput += *CharAddition
					return nil
				},
			},
			"Other": Action{
				Destination: ERR,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
		},
		IN_NUM: Transition{
			"Number": Action{
				Destination: IN_NUM,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"Other": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					newToken := CreateToken(NUMBER, *StringInput)
					*TokenList = append(*TokenList, *newToken)
					*StringInput = ""
					return newToken
				},
			},
		},
		IN_ID: Transition{
			"Number": Action{
				Destination: IN_ID,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"Letter": Action{
				Destination: IN_ID,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					return nil
				},
			},
			"Other": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					var tokenKeyword BASETOKENTYPE
					fmt.Println("Identifier here " + *StringInput)
					if(isReservedWord(*StringInput)){
						tokenKeyword = RESERVED
					}else {
						tokenKeyword = IDENTIFIER
					}
					fmt.Println(tokenKeyword)
					newToken := CreateToken(tokenKeyword, *StringInput)
					*TokenList = append(*TokenList, *newToken)
					*StringInput = ""
					return newToken
				},
			},
		},
		IN_ASSIGN: Transition{
			"EQ": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					newToken := CreateToken(ASSIGNMENT, *StringInput)
					*TokenList = append(*TokenList, *newToken)
					*StringInput = ""
					return newToken
				},
			},
			"Other": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					*StringInput = ""
					return nil
				},
			},
		},
		SPECIAL_S: Transition{
			"Other": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					newToken := CreateToken(SPECIALSYMBOL, *StringInput)
					*TokenList = append(*TokenList, *newToken)
					*StringInput = ""
					return newToken
				},
			},
		},
		ERR: Transition{
			"Other": Action{
				Destination: DONE,
				Callback: func(StringInput *string, CharAddition *string, TokenList *[]Token) *Token {
					*StringInput += *CharAddition
					newToken := CreateToken(ERROR, *StringInput)
					*TokenList = append(*TokenList, *newToken)
					*StringInput = ""
					return newToken
				},
			},
		},
	},
}

func (S *ScannerStruct) Transition(event Event, eventChar string) error {
	action := S.StateMap[S.Current][event]
	// fmt.Println("Action Destination is "+action.Destination)
	// fmt.Println("Current Input is " +S.Input)
	if fmt.Sprint(action) != fmt.Sprint(Action{}) {

		// fmt.Println("Going from "+S.Current+" to "+action.Destination)
		S.Current = action.Destination
		action.Callback(&S.Input, &eventChar, &S.TokenList)
		if(S.Current == DONE){
			S.Current = START
		}
		return nil
	}
	return fmt.Errorf("transition invalid")

}

func (S *ScannerStruct) Scan(inputString string) {
	// S.Input = inputString
	inputString = strings.Join(strings.Split(inputString, ""), " ")
	S.Current = START
	for i := S.Pointer; i < len(inputString); i++ {
		c := inputString[i]
		
		// fmt.Println("Scanning" + string(c))
		if isNumber(string(c)) {
			// fmt.Println("Number" + string(c))
			S.Transition("Number", string(c))
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			// fmt.Println("Letter" + string(c))
			S.Transition("Letter", string(c))
		} else if c == ':' {
			// fmt.Println("Colon" + string(c))
			S.Transition("Colon", string(c))
		} else if c == '=' {
			// fmt.Println("EQ" + string(c))
			err := S.Transition("EQ", string(c))
			if err != nil {
				S.Transition("SpecialSymbol", string(c))
			}
		} else if c == ' ' {
			// fmt.Println("Space" + string(c))
			err := S.Transition("WhiteSpace", string(c))
			if( err != nil){
				S.Transition("Other", string(c))
			}
		} else if isSpecialSymbols(string(c)) {
			// fmt.Println("SpecialSymbol" + string(c))
			S.Transition("SpecialSymbol", string(c))
		} else {
			// fmt.Println("Other" + string(c))
			S.Transition("Other", string(c))
		}

		
		// S.Transition(string(char), string(char))
	}
	if(S.Current != START) {
		// fmt.Println("Curret state is " + S.Current)
		// fmt.Println("There is still some input left " + S.Input)
		S.Input = strings.TrimSpace(S.Input)
		switch S.Current {
			case IN_ID:{
				var tokenKeyword BASETOKENTYPE
				fmt.Println("Identifier here " + S.Input)
				if(isReservedWord(S.Input)){
					tokenKeyword = RESERVED
				}else {
					tokenKeyword = IDENTIFIER
				}
				fmt.Println(tokenKeyword)
				newToken := CreateToken(tokenKeyword, S.Input)
				S.TokenList = append(S.TokenList, *newToken)
			}
			case IN_NUM:{
				newToken := CreateToken(NUMBER, S.Input)
				S.TokenList = append(S.TokenList, *newToken)
			}
			case SPECIAL_S:{
				newToken := CreateToken(SPECIALSYMBOL, S.Input)
				S.TokenList = append(S.TokenList, *newToken)
			}
			case ERR:{
				newToken := CreateToken(ERROR, S.Input)
				S.TokenList = append(S.TokenList, *newToken)
			}
			default:{
			}
		}
	}
}
