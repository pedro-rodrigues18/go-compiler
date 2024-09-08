// Code generated from MyGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MyGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MyGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mygrammarlexerLexerInit() {
	staticData := &MyGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'['", "']'", "','", "'func'", "'||'", "'&&'", "'!'", "", "", "",
		"", "", "'main'", "'int'", "'float'", "'char'", "'boolean'", "'void'",
		"'return'", "'if'", "'else'", "'while'", "'scanf'", "'println'", "'('",
		"')'", "'{'", "'}'", "';'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "CONSTANTE", "NUM_INT", "NUM_DEC", "TEXTO",
		"COMP", "MAIN", "INT", "FLOAT", "CHAR", "BOOLEAN", "VOID", "RETURN",
		"IF", "ELSE", "WHILE", "SCANF", "PRINTLN", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "SEMI", "ASSIGN", "PLUS", "MINUS", "MUL", "DIV", "MOD", "ID",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "CONSTANTE",
		"NUM_INT", "NUM_DEC", "TEXTO", "COMP", "MAIN", "INT", "FLOAT", "CHAR",
		"BOOLEAN", "VOID", "RETURN", "IF", "ELSE", "WHILE", "SCANF", "PRINTLN",
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "SEMI", "ASSIGN", "PLUS", "MINUS",
		"MUL", "DIV", "MOD", "ID", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 238, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 97, 8,
		7, 1, 8, 4, 8, 100, 8, 8, 11, 8, 12, 8, 101, 1, 9, 4, 9, 105, 8, 9, 11,
		9, 12, 9, 106, 1, 9, 1, 9, 4, 9, 111, 8, 9, 11, 9, 12, 9, 112, 1, 10, 1,
		10, 5, 10, 117, 8, 10, 10, 10, 12, 10, 120, 9, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 133, 8,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 227,
		8, 35, 10, 35, 12, 35, 230, 9, 35, 1, 36, 4, 36, 233, 8, 36, 11, 36, 12,
		36, 234, 1, 36, 1, 36, 1, 118, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 1, 0, 5, 1, 0, 48, 57, 2, 0, 60, 60, 62,
		62, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 248, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0,
		0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0,
		0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1,
		0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51,
		1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0,
		59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0,
		0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 81, 1, 0,
		0, 0, 9, 86, 1, 0, 0, 0, 11, 89, 1, 0, 0, 0, 13, 92, 1, 0, 0, 0, 15, 96,
		1, 0, 0, 0, 17, 99, 1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 114, 1, 0, 0,
		0, 23, 132, 1, 0, 0, 0, 25, 134, 1, 0, 0, 0, 27, 139, 1, 0, 0, 0, 29, 143,
		1, 0, 0, 0, 31, 149, 1, 0, 0, 0, 33, 154, 1, 0, 0, 0, 35, 162, 1, 0, 0,
		0, 37, 167, 1, 0, 0, 0, 39, 174, 1, 0, 0, 0, 41, 177, 1, 0, 0, 0, 43, 182,
		1, 0, 0, 0, 45, 188, 1, 0, 0, 0, 47, 194, 1, 0, 0, 0, 49, 202, 1, 0, 0,
		0, 51, 204, 1, 0, 0, 0, 53, 206, 1, 0, 0, 0, 55, 208, 1, 0, 0, 0, 57, 210,
		1, 0, 0, 0, 59, 212, 1, 0, 0, 0, 61, 214, 1, 0, 0, 0, 63, 216, 1, 0, 0,
		0, 65, 218, 1, 0, 0, 0, 67, 220, 1, 0, 0, 0, 69, 222, 1, 0, 0, 0, 71, 224,
		1, 0, 0, 0, 73, 232, 1, 0, 0, 0, 75, 76, 5, 91, 0, 0, 76, 2, 1, 0, 0, 0,
		77, 78, 5, 93, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 44, 0, 0, 80, 6, 1,
		0, 0, 0, 81, 82, 5, 102, 0, 0, 82, 83, 5, 117, 0, 0, 83, 84, 5, 110, 0,
		0, 84, 85, 5, 99, 0, 0, 85, 8, 1, 0, 0, 0, 86, 87, 5, 124, 0, 0, 87, 88,
		5, 124, 0, 0, 88, 10, 1, 0, 0, 0, 89, 90, 5, 38, 0, 0, 90, 91, 5, 38, 0,
		0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 33, 0, 0, 93, 14, 1, 0, 0, 0, 94, 97,
		3, 17, 8, 0, 95, 97, 3, 19, 9, 0, 96, 94, 1, 0, 0, 0, 96, 95, 1, 0, 0,
		0, 97, 16, 1, 0, 0, 0, 98, 100, 7, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 18, 1, 0, 0,
		0, 103, 105, 7, 0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110,
		5, 46, 0, 0, 109, 111, 7, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0,
		0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 20, 1, 0, 0, 0,
		114, 118, 5, 34, 0, 0, 115, 117, 9, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117,
		120, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 121,
		1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122, 5, 34, 0, 0, 122, 22, 1, 0,
		0, 0, 123, 133, 7, 1, 0, 0, 124, 125, 5, 61, 0, 0, 125, 133, 5, 61, 0,
		0, 126, 127, 5, 60, 0, 0, 127, 133, 5, 61, 0, 0, 128, 129, 5, 62, 0, 0,
		129, 133, 5, 61, 0, 0, 130, 131, 5, 33, 0, 0, 131, 133, 5, 61, 0, 0, 132,
		123, 1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 132, 126, 1, 0, 0, 0, 132, 128,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 24, 1, 0, 0, 0, 134, 135, 5, 109,
		0, 0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 105, 0, 0, 137, 138, 5, 110,
		0, 0, 138, 26, 1, 0, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5, 110, 0,
		0, 141, 142, 5, 116, 0, 0, 142, 28, 1, 0, 0, 0, 143, 144, 5, 102, 0, 0,
		144, 145, 5, 108, 0, 0, 145, 146, 5, 111, 0, 0, 146, 147, 5, 97, 0, 0,
		147, 148, 5, 116, 0, 0, 148, 30, 1, 0, 0, 0, 149, 150, 5, 99, 0, 0, 150,
		151, 5, 104, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153, 5, 114, 0, 0, 153,
		32, 1, 0, 0, 0, 154, 155, 5, 98, 0, 0, 155, 156, 5, 111, 0, 0, 156, 157,
		5, 111, 0, 0, 157, 158, 5, 108, 0, 0, 158, 159, 5, 101, 0, 0, 159, 160,
		5, 97, 0, 0, 160, 161, 5, 110, 0, 0, 161, 34, 1, 0, 0, 0, 162, 163, 5,
		118, 0, 0, 163, 164, 5, 111, 0, 0, 164, 165, 5, 105, 0, 0, 165, 166, 5,
		100, 0, 0, 166, 36, 1, 0, 0, 0, 167, 168, 5, 114, 0, 0, 168, 169, 5, 101,
		0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 117, 0, 0, 171, 172, 5, 114,
		0, 0, 172, 173, 5, 110, 0, 0, 173, 38, 1, 0, 0, 0, 174, 175, 5, 105, 0,
		0, 175, 176, 5, 102, 0, 0, 176, 40, 1, 0, 0, 0, 177, 178, 5, 101, 0, 0,
		178, 179, 5, 108, 0, 0, 179, 180, 5, 115, 0, 0, 180, 181, 5, 101, 0, 0,
		181, 42, 1, 0, 0, 0, 182, 183, 5, 119, 0, 0, 183, 184, 5, 104, 0, 0, 184,
		185, 5, 105, 0, 0, 185, 186, 5, 108, 0, 0, 186, 187, 5, 101, 0, 0, 187,
		44, 1, 0, 0, 0, 188, 189, 5, 115, 0, 0, 189, 190, 5, 99, 0, 0, 190, 191,
		5, 97, 0, 0, 191, 192, 5, 110, 0, 0, 192, 193, 5, 102, 0, 0, 193, 46, 1,
		0, 0, 0, 194, 195, 5, 112, 0, 0, 195, 196, 5, 114, 0, 0, 196, 197, 5, 105,
		0, 0, 197, 198, 5, 110, 0, 0, 198, 199, 5, 116, 0, 0, 199, 200, 5, 108,
		0, 0, 200, 201, 5, 110, 0, 0, 201, 48, 1, 0, 0, 0, 202, 203, 5, 40, 0,
		0, 203, 50, 1, 0, 0, 0, 204, 205, 5, 41, 0, 0, 205, 52, 1, 0, 0, 0, 206,
		207, 5, 123, 0, 0, 207, 54, 1, 0, 0, 0, 208, 209, 5, 125, 0, 0, 209, 56,
		1, 0, 0, 0, 210, 211, 5, 59, 0, 0, 211, 58, 1, 0, 0, 0, 212, 213, 5, 61,
		0, 0, 213, 60, 1, 0, 0, 0, 214, 215, 5, 43, 0, 0, 215, 62, 1, 0, 0, 0,
		216, 217, 5, 45, 0, 0, 217, 64, 1, 0, 0, 0, 218, 219, 5, 42, 0, 0, 219,
		66, 1, 0, 0, 0, 220, 221, 5, 47, 0, 0, 221, 68, 1, 0, 0, 0, 222, 223, 5,
		37, 0, 0, 223, 70, 1, 0, 0, 0, 224, 228, 7, 2, 0, 0, 225, 227, 7, 3, 0,
		0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 72, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 233, 7,
		4, 0, 0, 232, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 232, 1, 0, 0,
		0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 6, 36, 0, 0, 237,
		74, 1, 0, 0, 0, 9, 0, 96, 101, 106, 112, 118, 132, 228, 234, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MyGrammarLexerInit initializes any static state used to implement MyGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMyGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MyGrammarLexerInit() {
	staticData := &MyGrammarLexerLexerStaticData
	staticData.once.Do(mygrammarlexerLexerInit)
}

// NewMyGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMyGrammarLexer(input antlr.CharStream) *MyGrammarLexer {
	MyGrammarLexerInit()
	l := new(MyGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MyGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MyGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MyGrammarLexer tokens.
const (
	MyGrammarLexerT__0      = 1
	MyGrammarLexerT__1      = 2
	MyGrammarLexerT__2      = 3
	MyGrammarLexerT__3      = 4
	MyGrammarLexerT__4      = 5
	MyGrammarLexerT__5      = 6
	MyGrammarLexerT__6      = 7
	MyGrammarLexerCONSTANTE = 8
	MyGrammarLexerNUM_INT   = 9
	MyGrammarLexerNUM_DEC   = 10
	MyGrammarLexerTEXTO     = 11
	MyGrammarLexerCOMP      = 12
	MyGrammarLexerMAIN      = 13
	MyGrammarLexerINT       = 14
	MyGrammarLexerFLOAT     = 15
	MyGrammarLexerCHAR      = 16
	MyGrammarLexerBOOLEAN   = 17
	MyGrammarLexerVOID      = 18
	MyGrammarLexerRETURN    = 19
	MyGrammarLexerIF        = 20
	MyGrammarLexerELSE      = 21
	MyGrammarLexerWHILE     = 22
	MyGrammarLexerSCANF     = 23
	MyGrammarLexerPRINTLN   = 24
	MyGrammarLexerLPAREN    = 25
	MyGrammarLexerRPAREN    = 26
	MyGrammarLexerLBRACE    = 27
	MyGrammarLexerRBRACE    = 28
	MyGrammarLexerSEMI      = 29
	MyGrammarLexerASSIGN    = 30
	MyGrammarLexerPLUS      = 31
	MyGrammarLexerMINUS     = 32
	MyGrammarLexerMUL       = 33
	MyGrammarLexerDIV       = 34
	MyGrammarLexerMOD       = 35
	MyGrammarLexerID        = 36
	MyGrammarLexerWS        = 37
)