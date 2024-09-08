// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
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
		"programa", "listaFuncoes", "decFuncao", "tipoRetorno", "tipo", "tipoBase",
		"dimensao", "parametros", "listaParametros", "principal", "bloco", "listaVariaveis",
		"listaId", "comandos", "comando", "leitura", "termoLeitura", "novoTermoLeitura",
		"dimensao2", "escrita", "termoEscrita", "novoTermoEscrita", "selecao",
		"senao", "enquanto", "atribuicao", "complemento", "funcao", "argumentos",
		"novoArgumento", "retorno", "expressao", "exprOu", "exprOu2", "exprE",
		"exprE2", "exprRelacional", "exprRelacional2", "exprAditiva", "exprAditiva2",
		"opAditivo", "exprMultiplicativa", "exprMultiplicativa2", "opMultiplicativo",
		"fator", "termo", "sinal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 354, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 103, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 126, 8, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 133, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 141, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 160,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 166, 8, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 172, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 3, 14, 181, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 198,
		8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 206, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 219, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 226, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3,
		23, 238, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 253, 8, 26, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 265, 8, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 272, 8, 29, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 3, 33, 288, 8, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 3, 35, 298, 8, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		3, 37, 306, 8, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 3, 39, 316, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 3, 42, 328, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 342, 8, 44,
		1, 45, 1, 45, 1, 45, 3, 45, 347, 8, 45, 1, 46, 1, 46, 1, 46, 3, 46, 352,
		8, 46, 1, 46, 0, 0, 47, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 3, 1,
		0, 14, 17, 1, 0, 31, 32, 1, 0, 33, 35, 340, 0, 94, 1, 0, 0, 0, 2, 102,
		1, 0, 0, 0, 4, 104, 1, 0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 115, 1, 0, 0, 0,
		10, 118, 1, 0, 0, 0, 12, 125, 1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 140,
		1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 147, 1, 0, 0, 0, 22, 159, 1, 0, 0,
		0, 24, 165, 1, 0, 0, 0, 26, 171, 1, 0, 0, 0, 28, 180, 1, 0, 0, 0, 30, 182,
		1, 0, 0, 0, 32, 189, 1, 0, 0, 0, 34, 197, 1, 0, 0, 0, 36, 205, 1, 0, 0,
		0, 38, 207, 1, 0, 0, 0, 40, 218, 1, 0, 0, 0, 42, 225, 1, 0, 0, 0, 44, 227,
		1, 0, 0, 0, 46, 237, 1, 0, 0, 0, 48, 239, 1, 0, 0, 0, 50, 245, 1, 0, 0,
		0, 52, 252, 1, 0, 0, 0, 54, 254, 1, 0, 0, 0, 56, 264, 1, 0, 0, 0, 58, 271,
		1, 0, 0, 0, 60, 273, 1, 0, 0, 0, 62, 277, 1, 0, 0, 0, 64, 279, 1, 0, 0,
		0, 66, 287, 1, 0, 0, 0, 68, 289, 1, 0, 0, 0, 70, 297, 1, 0, 0, 0, 72, 299,
		1, 0, 0, 0, 74, 305, 1, 0, 0, 0, 76, 307, 1, 0, 0, 0, 78, 315, 1, 0, 0,
		0, 80, 317, 1, 0, 0, 0, 82, 319, 1, 0, 0, 0, 84, 327, 1, 0, 0, 0, 86, 329,
		1, 0, 0, 0, 88, 341, 1, 0, 0, 0, 90, 346, 1, 0, 0, 0, 92, 351, 1, 0, 0,
		0, 94, 95, 3, 2, 1, 0, 95, 96, 3, 18, 9, 0, 96, 97, 5, 0, 0, 1, 97, 1,
		1, 0, 0, 0, 98, 99, 3, 4, 2, 0, 99, 100, 3, 2, 1, 0, 100, 103, 1, 0, 0,
		0, 101, 103, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103,
		3, 1, 0, 0, 0, 104, 105, 3, 6, 3, 0, 105, 106, 5, 36, 0, 0, 106, 107, 5,
		25, 0, 0, 107, 108, 3, 14, 7, 0, 108, 109, 5, 26, 0, 0, 109, 110, 3, 20,
		10, 0, 110, 5, 1, 0, 0, 0, 111, 114, 3, 8, 4, 0, 112, 114, 5, 18, 0, 0,
		113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 7, 1, 0, 0, 0, 115, 116,
		3, 10, 5, 0, 116, 117, 3, 12, 6, 0, 117, 9, 1, 0, 0, 0, 118, 119, 7, 0,
		0, 0, 119, 11, 1, 0, 0, 0, 120, 121, 5, 1, 0, 0, 121, 122, 5, 9, 0, 0,
		122, 123, 5, 2, 0, 0, 123, 126, 3, 12, 6, 0, 124, 126, 1, 0, 0, 0, 125,
		120, 1, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126, 13, 1, 0, 0, 0, 127, 128, 3,
		8, 4, 0, 128, 129, 5, 36, 0, 0, 129, 130, 3, 16, 8, 0, 130, 133, 1, 0,
		0, 0, 131, 133, 1, 0, 0, 0, 132, 127, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0,
		133, 15, 1, 0, 0, 0, 134, 135, 5, 3, 0, 0, 135, 136, 3, 8, 4, 0, 136, 137,
		5, 36, 0, 0, 137, 138, 3, 16, 8, 0, 138, 141, 1, 0, 0, 0, 139, 141, 1,
		0, 0, 0, 140, 134, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141, 17, 1, 0, 0,
		0, 142, 143, 5, 13, 0, 0, 143, 144, 5, 25, 0, 0, 144, 145, 5, 26, 0, 0,
		145, 146, 3, 20, 10, 0, 146, 19, 1, 0, 0, 0, 147, 148, 5, 27, 0, 0, 148,
		149, 3, 22, 11, 0, 149, 150, 3, 26, 13, 0, 150, 151, 5, 28, 0, 0, 151,
		21, 1, 0, 0, 0, 152, 153, 3, 8, 4, 0, 153, 154, 5, 36, 0, 0, 154, 155,
		3, 24, 12, 0, 155, 156, 5, 29, 0, 0, 156, 157, 3, 22, 11, 0, 157, 160,
		1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 152, 1, 0, 0, 0, 159, 158, 1, 0,
		0, 0, 160, 23, 1, 0, 0, 0, 161, 162, 5, 3, 0, 0, 162, 163, 5, 36, 0, 0,
		163, 166, 3, 24, 12, 0, 164, 166, 1, 0, 0, 0, 165, 161, 1, 0, 0, 0, 165,
		164, 1, 0, 0, 0, 166, 25, 1, 0, 0, 0, 167, 168, 3, 28, 14, 0, 168, 169,
		3, 26, 13, 0, 169, 172, 1, 0, 0, 0, 170, 172, 1, 0, 0, 0, 171, 167, 1,
		0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 27, 1, 0, 0, 0, 173, 181, 3, 30, 15,
		0, 174, 181, 3, 38, 19, 0, 175, 181, 3, 50, 25, 0, 176, 181, 3, 54, 27,
		0, 177, 181, 3, 44, 22, 0, 178, 181, 3, 48, 24, 0, 179, 181, 3, 60, 30,
		0, 180, 173, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180, 175, 1, 0, 0, 0, 180,
		176, 1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 179,
		1, 0, 0, 0, 181, 29, 1, 0, 0, 0, 182, 183, 5, 23, 0, 0, 183, 184, 5, 25,
		0, 0, 184, 185, 3, 32, 16, 0, 185, 186, 3, 34, 17, 0, 186, 187, 5, 26,
		0, 0, 187, 188, 5, 29, 0, 0, 188, 31, 1, 0, 0, 0, 189, 190, 5, 36, 0, 0,
		190, 191, 3, 36, 18, 0, 191, 33, 1, 0, 0, 0, 192, 193, 5, 3, 0, 0, 193,
		194, 3, 32, 16, 0, 194, 195, 3, 34, 17, 0, 195, 198, 1, 0, 0, 0, 196, 198,
		1, 0, 0, 0, 197, 192, 1, 0, 0, 0, 197, 196, 1, 0, 0, 0, 198, 35, 1, 0,
		0, 0, 199, 200, 5, 1, 0, 0, 200, 201, 3, 76, 38, 0, 201, 202, 5, 2, 0,
		0, 202, 203, 3, 36, 18, 0, 203, 206, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0,
		205, 199, 1, 0, 0, 0, 205, 204, 1, 0, 0, 0, 206, 37, 1, 0, 0, 0, 207, 208,
		5, 24, 0, 0, 208, 209, 5, 25, 0, 0, 209, 210, 3, 40, 20, 0, 210, 211, 3,
		42, 21, 0, 211, 212, 5, 26, 0, 0, 212, 213, 5, 29, 0, 0, 213, 39, 1, 0,
		0, 0, 214, 215, 5, 36, 0, 0, 215, 219, 3, 36, 18, 0, 216, 219, 5, 8, 0,
		0, 217, 219, 5, 11, 0, 0, 218, 214, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218,
		217, 1, 0, 0, 0, 219, 41, 1, 0, 0, 0, 220, 221, 5, 3, 0, 0, 221, 222, 3,
		40, 20, 0, 222, 223, 3, 42, 21, 0, 223, 226, 1, 0, 0, 0, 224, 226, 1, 0,
		0, 0, 225, 220, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 43, 1, 0, 0, 0,
		227, 228, 5, 20, 0, 0, 228, 229, 5, 25, 0, 0, 229, 230, 3, 62, 31, 0, 230,
		231, 5, 26, 0, 0, 231, 232, 3, 20, 10, 0, 232, 233, 3, 46, 23, 0, 233,
		45, 1, 0, 0, 0, 234, 235, 5, 21, 0, 0, 235, 238, 3, 20, 10, 0, 236, 238,
		1, 0, 0, 0, 237, 234, 1, 0, 0, 0, 237, 236, 1, 0, 0, 0, 238, 47, 1, 0,
		0, 0, 239, 240, 5, 22, 0, 0, 240, 241, 5, 25, 0, 0, 241, 242, 3, 62, 31,
		0, 242, 243, 5, 26, 0, 0, 243, 244, 3, 20, 10, 0, 244, 49, 1, 0, 0, 0,
		245, 246, 5, 36, 0, 0, 246, 247, 5, 30, 0, 0, 247, 248, 3, 52, 26, 0, 248,
		249, 5, 29, 0, 0, 249, 51, 1, 0, 0, 0, 250, 253, 3, 62, 31, 0, 251, 253,
		3, 54, 27, 0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0, 0, 0, 253, 53, 1, 0,
		0, 0, 254, 255, 5, 4, 0, 0, 255, 256, 5, 36, 0, 0, 256, 257, 5, 25, 0,
		0, 257, 258, 3, 56, 28, 0, 258, 259, 5, 26, 0, 0, 259, 55, 1, 0, 0, 0,
		260, 261, 3, 62, 31, 0, 261, 262, 3, 58, 29, 0, 262, 265, 1, 0, 0, 0, 263,
		265, 1, 0, 0, 0, 264, 260, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265, 57, 1,
		0, 0, 0, 266, 267, 5, 3, 0, 0, 267, 268, 3, 62, 31, 0, 268, 269, 3, 58,
		29, 0, 269, 272, 1, 0, 0, 0, 270, 272, 1, 0, 0, 0, 271, 266, 1, 0, 0, 0,
		271, 270, 1, 0, 0, 0, 272, 59, 1, 0, 0, 0, 273, 274, 5, 19, 0, 0, 274,
		275, 3, 62, 31, 0, 275, 276, 5, 29, 0, 0, 276, 61, 1, 0, 0, 0, 277, 278,
		3, 64, 32, 0, 278, 63, 1, 0, 0, 0, 279, 280, 3, 68, 34, 0, 280, 281, 3,
		66, 33, 0, 281, 65, 1, 0, 0, 0, 282, 283, 5, 5, 0, 0, 283, 284, 3, 68,
		34, 0, 284, 285, 3, 66, 33, 0, 285, 288, 1, 0, 0, 0, 286, 288, 1, 0, 0,
		0, 287, 282, 1, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 67, 1, 0, 0, 0, 289,
		290, 3, 72, 36, 0, 290, 291, 3, 70, 35, 0, 291, 69, 1, 0, 0, 0, 292, 293,
		5, 6, 0, 0, 293, 294, 3, 72, 36, 0, 294, 295, 3, 70, 35, 0, 295, 298, 1,
		0, 0, 0, 296, 298, 1, 0, 0, 0, 297, 292, 1, 0, 0, 0, 297, 296, 1, 0, 0,
		0, 298, 71, 1, 0, 0, 0, 299, 300, 3, 76, 38, 0, 300, 301, 3, 74, 37, 0,
		301, 73, 1, 0, 0, 0, 302, 303, 5, 12, 0, 0, 303, 306, 3, 76, 38, 0, 304,
		306, 1, 0, 0, 0, 305, 302, 1, 0, 0, 0, 305, 304, 1, 0, 0, 0, 306, 75, 1,
		0, 0, 0, 307, 308, 3, 82, 41, 0, 308, 309, 3, 78, 39, 0, 309, 77, 1, 0,
		0, 0, 310, 311, 3, 80, 40, 0, 311, 312, 3, 82, 41, 0, 312, 313, 3, 78,
		39, 0, 313, 316, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0,
		315, 314, 1, 0, 0, 0, 316, 79, 1, 0, 0, 0, 317, 318, 7, 1, 0, 0, 318, 81,
		1, 0, 0, 0, 319, 320, 3, 88, 44, 0, 320, 321, 3, 84, 42, 0, 321, 83, 1,
		0, 0, 0, 322, 323, 3, 86, 43, 0, 323, 324, 3, 88, 44, 0, 324, 325, 3, 84,
		42, 0, 325, 328, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327, 322, 1, 0, 0, 0,
		327, 326, 1, 0, 0, 0, 328, 85, 1, 0, 0, 0, 329, 330, 7, 2, 0, 0, 330, 87,
		1, 0, 0, 0, 331, 332, 3, 92, 46, 0, 332, 333, 3, 90, 45, 0, 333, 342, 1,
		0, 0, 0, 334, 342, 5, 11, 0, 0, 335, 336, 5, 7, 0, 0, 336, 342, 3, 88,
		44, 0, 337, 338, 5, 25, 0, 0, 338, 339, 3, 62, 31, 0, 339, 340, 5, 26,
		0, 0, 340, 342, 1, 0, 0, 0, 341, 331, 1, 0, 0, 0, 341, 334, 1, 0, 0, 0,
		341, 335, 1, 0, 0, 0, 341, 337, 1, 0, 0, 0, 342, 89, 1, 0, 0, 0, 343, 344,
		5, 36, 0, 0, 344, 347, 3, 36, 18, 0, 345, 347, 5, 8, 0, 0, 346, 343, 1,
		0, 0, 0, 346, 345, 1, 0, 0, 0, 347, 91, 1, 0, 0, 0, 348, 352, 5, 31, 0,
		0, 349, 352, 5, 32, 0, 0, 350, 352, 1, 0, 0, 0, 351, 348, 1, 0, 0, 0, 351,
		349, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 93, 1, 0, 0, 0, 25, 102, 113,
		125, 132, 140, 159, 165, 171, 180, 197, 205, 218, 225, 237, 252, 264, 271,
		287, 297, 305, 315, 327, 341, 346, 351,
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

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF       = antlr.TokenEOF
	GrammarParserT__0      = 1
	GrammarParserT__1      = 2
	GrammarParserT__2      = 3
	GrammarParserT__3      = 4
	GrammarParserT__4      = 5
	GrammarParserT__5      = 6
	GrammarParserT__6      = 7
	GrammarParserCONSTANTE = 8
	GrammarParserNUM_INT   = 9
	GrammarParserNUM_DEC   = 10
	GrammarParserTEXTO     = 11
	GrammarParserCOMP      = 12
	GrammarParserMAIN      = 13
	GrammarParserINT       = 14
	GrammarParserFLOAT     = 15
	GrammarParserCHAR      = 16
	GrammarParserBOOLEAN   = 17
	GrammarParserVOID      = 18
	GrammarParserRETURN    = 19
	GrammarParserIF        = 20
	GrammarParserELSE      = 21
	GrammarParserWHILE     = 22
	GrammarParserSCANF     = 23
	GrammarParserPRINTLN   = 24
	GrammarParserLPAREN    = 25
	GrammarParserRPAREN    = 26
	GrammarParserLBRACE    = 27
	GrammarParserRBRACE    = 28
	GrammarParserSEMI      = 29
	GrammarParserASSIGN    = 30
	GrammarParserPLUS      = 31
	GrammarParserMINUS     = 32
	GrammarParserMUL       = 33
	GrammarParserDIV       = 34
	GrammarParserMOD       = 35
	GrammarParserID        = 36
	GrammarParserWS        = 37
)

// GrammarParser rules.
const (
	GrammarParserRULE_programa            = 0
	GrammarParserRULE_listaFuncoes        = 1
	GrammarParserRULE_decFuncao           = 2
	GrammarParserRULE_tipoRetorno         = 3
	GrammarParserRULE_tipo                = 4
	GrammarParserRULE_tipoBase            = 5
	GrammarParserRULE_dimensao            = 6
	GrammarParserRULE_parametros          = 7
	GrammarParserRULE_listaParametros     = 8
	GrammarParserRULE_principal           = 9
	GrammarParserRULE_bloco               = 10
	GrammarParserRULE_listaVariaveis      = 11
	GrammarParserRULE_listaId             = 12
	GrammarParserRULE_comandos            = 13
	GrammarParserRULE_comando             = 14
	GrammarParserRULE_leitura             = 15
	GrammarParserRULE_termoLeitura        = 16
	GrammarParserRULE_novoTermoLeitura    = 17
	GrammarParserRULE_dimensao2           = 18
	GrammarParserRULE_escrita             = 19
	GrammarParserRULE_termoEscrita        = 20
	GrammarParserRULE_novoTermoEscrita    = 21
	GrammarParserRULE_selecao             = 22
	GrammarParserRULE_senao               = 23
	GrammarParserRULE_enquanto            = 24
	GrammarParserRULE_atribuicao          = 25
	GrammarParserRULE_complemento         = 26
	GrammarParserRULE_funcao              = 27
	GrammarParserRULE_argumentos          = 28
	GrammarParserRULE_novoArgumento       = 29
	GrammarParserRULE_retorno             = 30
	GrammarParserRULE_expressao           = 31
	GrammarParserRULE_exprOu              = 32
	GrammarParserRULE_exprOu2             = 33
	GrammarParserRULE_exprE               = 34
	GrammarParserRULE_exprE2              = 35
	GrammarParserRULE_exprRelacional      = 36
	GrammarParserRULE_exprRelacional2     = 37
	GrammarParserRULE_exprAditiva         = 38
	GrammarParserRULE_exprAditiva2        = 39
	GrammarParserRULE_opAditivo           = 40
	GrammarParserRULE_exprMultiplicativa  = 41
	GrammarParserRULE_exprMultiplicativa2 = 42
	GrammarParserRULE_opMultiplicativo    = 43
	GrammarParserRULE_fator               = 44
	GrammarParserRULE_termo               = 45
	GrammarParserRULE_sinal               = 46
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListaFuncoes() IListaFuncoesContext
	Principal() IPrincipalContext
	EOF() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) ListaFuncoes() IListaFuncoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaFuncoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaFuncoesContext)
}

func (s *ProgramaContext) Principal() IPrincipalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrincipalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrincipalContext)
}

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *GrammarParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.ListaFuncoes()
	}
	{
		p.SetState(95)
		p.Principal()
	}
	{
		p.SetState(96)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaFuncoesContext is an interface to support dynamic dispatch.
type IListaFuncoesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DecFuncao() IDecFuncaoContext
	ListaFuncoes() IListaFuncoesContext

	// IsListaFuncoesContext differentiates from other interfaces.
	IsListaFuncoesContext()
}

type ListaFuncoesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaFuncoesContext() *ListaFuncoesContext {
	var p = new(ListaFuncoesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaFuncoes
	return p
}

func InitEmptyListaFuncoesContext(p *ListaFuncoesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaFuncoes
}

func (*ListaFuncoesContext) IsListaFuncoesContext() {}

func NewListaFuncoesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaFuncoesContext {
	var p = new(ListaFuncoesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_listaFuncoes

	return p
}

func (s *ListaFuncoesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaFuncoesContext) DecFuncao() IDecFuncaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecFuncaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecFuncaoContext)
}

func (s *ListaFuncoesContext) ListaFuncoes() IListaFuncoesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaFuncoesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaFuncoesContext)
}

func (s *ListaFuncoesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaFuncoesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaFuncoesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListaFuncoes(s)
	}
}

func (s *ListaFuncoesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListaFuncoes(s)
	}
}

func (p *GrammarParser) ListaFuncoes() (localctx IListaFuncoesContext) {
	localctx = NewListaFuncoesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_listaFuncoes)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserCHAR, GrammarParserBOOLEAN, GrammarParserVOID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.DecFuncao()
		}
		{
			p.SetState(99)
			p.ListaFuncoes()
		}

	case GrammarParserMAIN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecFuncaoContext is an interface to support dynamic dispatch.
type IDecFuncaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TipoRetorno() ITipoRetornoContext
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Parametros() IParametrosContext
	RPAREN() antlr.TerminalNode
	Bloco() IBlocoContext

	// IsDecFuncaoContext differentiates from other interfaces.
	IsDecFuncaoContext()
}

type DecFuncaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecFuncaoContext() *DecFuncaoContext {
	var p = new(DecFuncaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decFuncao
	return p
}

func InitEmptyDecFuncaoContext(p *DecFuncaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decFuncao
}

func (*DecFuncaoContext) IsDecFuncaoContext() {}

func NewDecFuncaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecFuncaoContext {
	var p = new(DecFuncaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_decFuncao

	return p
}

func (s *DecFuncaoContext) GetParser() antlr.Parser { return s.parser }

func (s *DecFuncaoContext) TipoRetorno() ITipoRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoRetornoContext)
}

func (s *DecFuncaoContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DecFuncaoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *DecFuncaoContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *DecFuncaoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *DecFuncaoContext) Bloco() IBlocoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocoContext)
}

func (s *DecFuncaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecFuncaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecFuncaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDecFuncao(s)
	}
}

func (s *DecFuncaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDecFuncao(s)
	}
}

func (p *GrammarParser) DecFuncao() (localctx IDecFuncaoContext) {
	localctx = NewDecFuncaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_decFuncao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.TipoRetorno()
	}
	{
		p.SetState(105)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Parametros()
	}
	{
		p.SetState(108)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Bloco()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoRetornoContext is an interface to support dynamic dispatch.
type ITipoRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	VOID() antlr.TerminalNode

	// IsTipoRetornoContext differentiates from other interfaces.
	IsTipoRetornoContext()
}

type TipoRetornoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoRetornoContext() *TipoRetornoContext {
	var p = new(TipoRetornoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipoRetorno
	return p
}

func InitEmptyTipoRetornoContext(p *TipoRetornoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipoRetorno
}

func (*TipoRetornoContext) IsTipoRetornoContext() {}

func NewTipoRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoRetornoContext {
	var p = new(TipoRetornoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_tipoRetorno

	return p
}

func (s *TipoRetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoRetornoContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *TipoRetornoContext) VOID() antlr.TerminalNode {
	return s.GetToken(GrammarParserVOID, 0)
}

func (s *TipoRetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoRetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoRetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTipoRetorno(s)
	}
}

func (s *TipoRetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTipoRetorno(s)
	}
}

func (p *GrammarParser) TipoRetorno() (localctx ITipoRetornoContext) {
	localctx = NewTipoRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_tipoRetorno)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserCHAR, GrammarParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Tipo()
		}

	case GrammarParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(GrammarParserVOID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TipoBase() ITipoBaseContext
	Dimensao() IDimensaoContext

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) TipoBase() ITipoBaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoBaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoBaseContext)
}

func (s *TipoContext) Dimensao() IDimensaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensaoContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *GrammarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_tipo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.TipoBase()
	}
	{
		p.SetState(116)
		p.Dimensao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoBaseContext is an interface to support dynamic dispatch.
type ITipoBaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsTipoBaseContext differentiates from other interfaces.
	IsTipoBaseContext()
}

type TipoBaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoBaseContext() *TipoBaseContext {
	var p = new(TipoBaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipoBase
	return p
}

func InitEmptyTipoBaseContext(p *TipoBaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_tipoBase
}

func (*TipoBaseContext) IsTipoBaseContext() {}

func NewTipoBaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoBaseContext {
	var p = new(TipoBaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_tipoBase

	return p
}

func (s *TipoBaseContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoBaseContext) INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINT, 0)
}

func (s *TipoBaseContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GrammarParserFLOAT, 0)
}

func (s *TipoBaseContext) CHAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserCHAR, 0)
}

func (s *TipoBaseContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GrammarParserBOOLEAN, 0)
}

func (s *TipoBaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoBaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoBaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTipoBase(s)
	}
}

func (s *TipoBaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTipoBase(s)
	}
}

func (p *GrammarParser) TipoBase() (localctx ITipoBaseContext) {
	localctx = NewTipoBaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_tipoBase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDimensaoContext is an interface to support dynamic dispatch.
type IDimensaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUM_INT() antlr.TerminalNode
	Dimensao() IDimensaoContext

	// IsDimensaoContext differentiates from other interfaces.
	IsDimensaoContext()
}

type DimensaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensaoContext() *DimensaoContext {
	var p = new(DimensaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_dimensao
	return p
}

func InitEmptyDimensaoContext(p *DimensaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_dimensao
}

func (*DimensaoContext) IsDimensaoContext() {}

func NewDimensaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensaoContext {
	var p = new(DimensaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_dimensao

	return p
}

func (s *DimensaoContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensaoContext) NUM_INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserNUM_INT, 0)
}

func (s *DimensaoContext) Dimensao() IDimensaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensaoContext)
}

func (s *DimensaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDimensao(s)
	}
}

func (s *DimensaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDimensao(s)
	}
}

func (p *GrammarParser) Dimensao() (localctx IDimensaoContext) {
	localctx = NewDimensaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_dimensao)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(GrammarParserNUM_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Dimensao()
		}

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	ID() antlr.TerminalNode
	ListaParametros() IListaParametrosContext

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_parametros
	return p
}

func InitEmptyParametrosContext(p *ParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_parametros
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ParametrosContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ParametrosContext) ListaParametros() IListaParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParametrosContext)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (p *GrammarParser) Parametros() (localctx IParametrosContext) {
	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_parametros)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserCHAR, GrammarParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Tipo()
		}
		{
			p.SetState(128)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.ListaParametros()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaParametrosContext is an interface to support dynamic dispatch.
type IListaParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	ID() antlr.TerminalNode
	ListaParametros() IListaParametrosContext

	// IsListaParametrosContext differentiates from other interfaces.
	IsListaParametrosContext()
}

type ListaParametrosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaParametrosContext() *ListaParametrosContext {
	var p = new(ListaParametrosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaParametros
	return p
}

func InitEmptyListaParametrosContext(p *ListaParametrosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaParametros
}

func (*ListaParametrosContext) IsListaParametrosContext() {}

func NewListaParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaParametrosContext {
	var p = new(ListaParametrosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_listaParametros

	return p
}

func (s *ListaParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaParametrosContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ListaParametrosContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ListaParametrosContext) ListaParametros() IListaParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaParametrosContext)
}

func (s *ListaParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListaParametros(s)
	}
}

func (s *ListaParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListaParametros(s)
	}
}

func (p *GrammarParser) ListaParametros() (localctx IListaParametrosContext) {
	localctx = NewListaParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_listaParametros)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Tipo()
		}
		{
			p.SetState(136)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.ListaParametros()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrincipalContext is an interface to support dynamic dispatch.
type IPrincipalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAIN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Bloco() IBlocoContext

	// IsPrincipalContext differentiates from other interfaces.
	IsPrincipalContext()
}

type PrincipalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrincipalContext() *PrincipalContext {
	var p = new(PrincipalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_principal
	return p
}

func InitEmptyPrincipalContext(p *PrincipalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_principal
}

func (*PrincipalContext) IsPrincipalContext() {}

func NewPrincipalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrincipalContext {
	var p = new(PrincipalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_principal

	return p
}

func (s *PrincipalContext) GetParser() antlr.Parser { return s.parser }

func (s *PrincipalContext) MAIN() antlr.TerminalNode {
	return s.GetToken(GrammarParserMAIN, 0)
}

func (s *PrincipalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *PrincipalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *PrincipalContext) Bloco() IBlocoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocoContext)
}

func (s *PrincipalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrincipalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrincipalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPrincipal(s)
	}
}

func (s *PrincipalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPrincipal(s)
	}
}

func (p *GrammarParser) Principal() (localctx IPrincipalContext) {
	localctx = NewPrincipalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_principal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(GrammarParserMAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Bloco()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlocoContext is an interface to support dynamic dispatch.
type IBlocoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	ListaVariaveis() IListaVariaveisContext
	Comandos() IComandosContext
	RBRACE() antlr.TerminalNode

	// IsBlocoContext differentiates from other interfaces.
	IsBlocoContext()
}

type BlocoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlocoContext() *BlocoContext {
	var p = new(BlocoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_bloco
	return p
}

func InitEmptyBlocoContext(p *BlocoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_bloco
}

func (*BlocoContext) IsBlocoContext() {}

func NewBlocoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlocoContext {
	var p = new(BlocoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_bloco

	return p
}

func (s *BlocoContext) GetParser() antlr.Parser { return s.parser }

func (s *BlocoContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserLBRACE, 0)
}

func (s *BlocoContext) ListaVariaveis() IListaVariaveisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaVariaveisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaVariaveisContext)
}

func (s *BlocoContext) Comandos() IComandosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandosContext)
}

func (s *BlocoContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GrammarParserRBRACE, 0)
}

func (s *BlocoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlocoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlocoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterBloco(s)
	}
}

func (s *BlocoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitBloco(s)
	}
}

func (p *GrammarParser) Bloco() (localctx IBlocoContext) {
	localctx = NewBlocoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_bloco)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(GrammarParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.ListaVariaveis()
	}
	{
		p.SetState(149)
		p.Comandos()
	}
	{
		p.SetState(150)
		p.Match(GrammarParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaVariaveisContext is an interface to support dynamic dispatch.
type IListaVariaveisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tipo() ITipoContext
	ID() antlr.TerminalNode
	ListaId() IListaIdContext
	SEMI() antlr.TerminalNode
	ListaVariaveis() IListaVariaveisContext

	// IsListaVariaveisContext differentiates from other interfaces.
	IsListaVariaveisContext()
}

type ListaVariaveisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaVariaveisContext() *ListaVariaveisContext {
	var p = new(ListaVariaveisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaVariaveis
	return p
}

func InitEmptyListaVariaveisContext(p *ListaVariaveisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaVariaveis
}

func (*ListaVariaveisContext) IsListaVariaveisContext() {}

func NewListaVariaveisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaVariaveisContext {
	var p = new(ListaVariaveisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_listaVariaveis

	return p
}

func (s *ListaVariaveisContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaVariaveisContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ListaVariaveisContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ListaVariaveisContext) ListaId() IListaIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIdContext)
}

func (s *ListaVariaveisContext) SEMI() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMI, 0)
}

func (s *ListaVariaveisContext) ListaVariaveis() IListaVariaveisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaVariaveisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaVariaveisContext)
}

func (s *ListaVariaveisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaVariaveisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaVariaveisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListaVariaveis(s)
	}
}

func (s *ListaVariaveisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListaVariaveis(s)
	}
}

func (p *GrammarParser) ListaVariaveis() (localctx IListaVariaveisContext) {
	localctx = NewListaVariaveisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_listaVariaveis)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserCHAR, GrammarParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Tipo()
		}
		{
			p.SetState(153)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.ListaId()
		}
		{
			p.SetState(155)
			p.Match(GrammarParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.ListaVariaveis()
		}

	case GrammarParserT__3, GrammarParserRETURN, GrammarParserIF, GrammarParserWHILE, GrammarParserSCANF, GrammarParserPRINTLN, GrammarParserRBRACE, GrammarParserID:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListaIdContext is an interface to support dynamic dispatch.
type IListaIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ListaId() IListaIdContext

	// IsListaIdContext differentiates from other interfaces.
	IsListaIdContext()
}

type ListaIdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListaIdContext() *ListaIdContext {
	var p = new(ListaIdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaId
	return p
}

func InitEmptyListaIdContext(p *ListaIdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_listaId
}

func (*ListaIdContext) IsListaIdContext() {}

func NewListaIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaIdContext {
	var p = new(ListaIdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_listaId

	return p
}

func (s *ListaIdContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaIdContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ListaIdContext) ListaId() IListaIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListaIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListaIdContext)
}

func (s *ListaIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterListaId(s)
	}
}

func (s *ListaIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitListaId(s)
	}
}

func (p *GrammarParser) ListaId() (localctx IListaIdContext) {
	localctx = NewListaIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_listaId)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.ListaId()
		}

	case GrammarParserSEMI:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandosContext is an interface to support dynamic dispatch.
type IComandosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Comando() IComandoContext
	Comandos() IComandosContext

	// IsComandosContext differentiates from other interfaces.
	IsComandosContext()
}

type ComandosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandosContext() *ComandosContext {
	var p = new(ComandosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_comandos
	return p
}

func InitEmptyComandosContext(p *ComandosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_comandos
}

func (*ComandosContext) IsComandosContext() {}

func NewComandosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandosContext {
	var p = new(ComandosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_comandos

	return p
}

func (s *ComandosContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandosContext) Comando() IComandoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *ComandosContext) Comandos() IComandosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComandosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComandosContext)
}

func (s *ComandosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComandos(s)
	}
}

func (s *ComandosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComandos(s)
	}
}

func (p *GrammarParser) Comandos() (localctx IComandosContext) {
	localctx = NewComandosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_comandos)
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__3, GrammarParserRETURN, GrammarParserIF, GrammarParserWHILE, GrammarParserSCANF, GrammarParserPRINTLN, GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Comando()
		}
		{
			p.SetState(168)
			p.Comandos()
		}

	case GrammarParserRBRACE:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Leitura() ILeituraContext
	Escrita() IEscritaContext
	Atribuicao() IAtribuicaoContext
	Funcao() IFuncaoContext
	Selecao() ISelecaoContext
	Enquanto() IEnquantoContext
	Retorno() IRetornoContext

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoContext() *ComandoContext {
	var p = new(ComandoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_comando
	return p
}

func InitEmptyComandoContext(p *ComandoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_comando
}

func (*ComandoContext) IsComandoContext() {}

func NewComandoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoContext {
	var p = new(ComandoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_comando

	return p
}

func (s *ComandoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoContext) Leitura() ILeituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeituraContext)
}

func (s *ComandoContext) Escrita() IEscritaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEscritaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEscritaContext)
}

func (s *ComandoContext) Atribuicao() IAtribuicaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtribuicaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtribuicaoContext)
}

func (s *ComandoContext) Funcao() IFuncaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncaoContext)
}

func (s *ComandoContext) Selecao() ISelecaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelecaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelecaoContext)
}

func (s *ComandoContext) Enquanto() IEnquantoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnquantoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnquantoContext)
}

func (s *ComandoContext) Retorno() IRetornoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetornoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetornoContext)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComando(s)
	}
}

func (s *ComandoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComando(s)
	}
}

func (p *GrammarParser) Comando() (localctx IComandoContext) {
	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_comando)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSCANF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Leitura()
		}

	case GrammarParserPRINTLN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Escrita()
		}

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Atribuicao()
		}

	case GrammarParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Funcao()
		}

	case GrammarParserIF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Selecao()
		}

	case GrammarParserWHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(178)
			p.Enquanto()
		}

	case GrammarParserRETURN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(179)
			p.Retorno()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILeituraContext is an interface to support dynamic dispatch.
type ILeituraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCANF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TermoLeitura() ITermoLeituraContext
	NovoTermoLeitura() INovoTermoLeituraContext
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsLeituraContext differentiates from other interfaces.
	IsLeituraContext()
}

type LeituraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeituraContext() *LeituraContext {
	var p = new(LeituraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_leitura
	return p
}

func InitEmptyLeituraContext(p *LeituraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_leitura
}

func (*LeituraContext) IsLeituraContext() {}

func NewLeituraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeituraContext {
	var p = new(LeituraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_leitura

	return p
}

func (s *LeituraContext) GetParser() antlr.Parser { return s.parser }

func (s *LeituraContext) SCANF() antlr.TerminalNode {
	return s.GetToken(GrammarParserSCANF, 0)
}

func (s *LeituraContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *LeituraContext) TermoLeitura() ITermoLeituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoLeituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoLeituraContext)
}

func (s *LeituraContext) NovoTermoLeitura() INovoTermoLeituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoTermoLeituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoTermoLeituraContext)
}

func (s *LeituraContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *LeituraContext) SEMI() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMI, 0)
}

func (s *LeituraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeituraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeituraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterLeitura(s)
	}
}

func (s *LeituraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitLeitura(s)
	}
}

func (p *GrammarParser) Leitura() (localctx ILeituraContext) {
	localctx = NewLeituraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_leitura)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(GrammarParserSCANF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.TermoLeitura()
	}
	{
		p.SetState(185)
		p.NovoTermoLeitura()
	}
	{
		p.SetState(186)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(GrammarParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoLeituraContext is an interface to support dynamic dispatch.
type ITermoLeituraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Dimensao2() IDimensao2Context

	// IsTermoLeituraContext differentiates from other interfaces.
	IsTermoLeituraContext()
}

type TermoLeituraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoLeituraContext() *TermoLeituraContext {
	var p = new(TermoLeituraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termoLeitura
	return p
}

func InitEmptyTermoLeituraContext(p *TermoLeituraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termoLeitura
}

func (*TermoLeituraContext) IsTermoLeituraContext() {}

func NewTermoLeituraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoLeituraContext {
	var p = new(TermoLeituraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_termoLeitura

	return p
}

func (s *TermoLeituraContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoLeituraContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TermoLeituraContext) Dimensao2() IDimensao2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensao2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensao2Context)
}

func (s *TermoLeituraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoLeituraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoLeituraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTermoLeitura(s)
	}
}

func (s *TermoLeituraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTermoLeitura(s)
	}
}

func (p *GrammarParser) TermoLeitura() (localctx ITermoLeituraContext) {
	localctx = NewTermoLeituraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_termoLeitura)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Dimensao2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INovoTermoLeituraContext is an interface to support dynamic dispatch.
type INovoTermoLeituraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TermoLeitura() ITermoLeituraContext
	NovoTermoLeitura() INovoTermoLeituraContext

	// IsNovoTermoLeituraContext differentiates from other interfaces.
	IsNovoTermoLeituraContext()
}

type NovoTermoLeituraContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNovoTermoLeituraContext() *NovoTermoLeituraContext {
	var p = new(NovoTermoLeituraContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoTermoLeitura
	return p
}

func InitEmptyNovoTermoLeituraContext(p *NovoTermoLeituraContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoTermoLeitura
}

func (*NovoTermoLeituraContext) IsNovoTermoLeituraContext() {}

func NewNovoTermoLeituraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NovoTermoLeituraContext {
	var p = new(NovoTermoLeituraContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_novoTermoLeitura

	return p
}

func (s *NovoTermoLeituraContext) GetParser() antlr.Parser { return s.parser }

func (s *NovoTermoLeituraContext) TermoLeitura() ITermoLeituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoLeituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoLeituraContext)
}

func (s *NovoTermoLeituraContext) NovoTermoLeitura() INovoTermoLeituraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoTermoLeituraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoTermoLeituraContext)
}

func (s *NovoTermoLeituraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NovoTermoLeituraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NovoTermoLeituraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNovoTermoLeitura(s)
	}
}

func (s *NovoTermoLeituraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNovoTermoLeitura(s)
	}
}

func (p *GrammarParser) NovoTermoLeitura() (localctx INovoTermoLeituraContext) {
	localctx = NewNovoTermoLeituraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_novoTermoLeitura)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.TermoLeitura()
		}
		{
			p.SetState(194)
			p.NovoTermoLeitura()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDimensao2Context is an interface to support dynamic dispatch.
type IDimensao2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprAditiva() IExprAditivaContext
	Dimensao2() IDimensao2Context

	// IsDimensao2Context differentiates from other interfaces.
	IsDimensao2Context()
}

type Dimensao2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensao2Context() *Dimensao2Context {
	var p = new(Dimensao2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_dimensao2
	return p
}

func InitEmptyDimensao2Context(p *Dimensao2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_dimensao2
}

func (*Dimensao2Context) IsDimensao2Context() {}

func NewDimensao2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dimensao2Context {
	var p = new(Dimensao2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_dimensao2

	return p
}

func (s *Dimensao2Context) GetParser() antlr.Parser { return s.parser }

func (s *Dimensao2Context) ExprAditiva() IExprAditivaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAditivaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAditivaContext)
}

func (s *Dimensao2Context) Dimensao2() IDimensao2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensao2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensao2Context)
}

func (s *Dimensao2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dimensao2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dimensao2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDimensao2(s)
	}
}

func (s *Dimensao2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDimensao2(s)
	}
}

func (p *GrammarParser) Dimensao2() (localctx IDimensao2Context) {
	localctx = NewDimensao2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_dimensao2)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.ExprAditiva()
		}
		{
			p.SetState(201)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Dimensao2()
		}

	case GrammarParserT__1, GrammarParserT__2, GrammarParserT__4, GrammarParserT__5, GrammarParserCOMP, GrammarParserRPAREN, GrammarParserSEMI, GrammarParserPLUS, GrammarParserMINUS, GrammarParserMUL, GrammarParserDIV, GrammarParserMOD:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEscritaContext is an interface to support dynamic dispatch.
type IEscritaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINTLN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	TermoEscrita() ITermoEscritaContext
	NovoTermoEscrita() INovoTermoEscritaContext
	RPAREN() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsEscritaContext differentiates from other interfaces.
	IsEscritaContext()
}

type EscritaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEscritaContext() *EscritaContext {
	var p = new(EscritaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_escrita
	return p
}

func InitEmptyEscritaContext(p *EscritaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_escrita
}

func (*EscritaContext) IsEscritaContext() {}

func NewEscritaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EscritaContext {
	var p = new(EscritaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_escrita

	return p
}

func (s *EscritaContext) GetParser() antlr.Parser { return s.parser }

func (s *EscritaContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(GrammarParserPRINTLN, 0)
}

func (s *EscritaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *EscritaContext) TermoEscrita() ITermoEscritaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoEscritaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoEscritaContext)
}

func (s *EscritaContext) NovoTermoEscrita() INovoTermoEscritaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoTermoEscritaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoTermoEscritaContext)
}

func (s *EscritaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *EscritaContext) SEMI() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMI, 0)
}

func (s *EscritaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EscritaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EscritaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterEscrita(s)
	}
}

func (s *EscritaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitEscrita(s)
	}
}

func (p *GrammarParser) Escrita() (localctx IEscritaContext) {
	localctx = NewEscritaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_escrita)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(GrammarParserPRINTLN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.TermoEscrita()
	}
	{
		p.SetState(210)
		p.NovoTermoEscrita()
	}
	{
		p.SetState(211)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(GrammarParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoEscritaContext is an interface to support dynamic dispatch.
type ITermoEscritaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Dimensao2() IDimensao2Context
	CONSTANTE() antlr.TerminalNode
	TEXTO() antlr.TerminalNode

	// IsTermoEscritaContext differentiates from other interfaces.
	IsTermoEscritaContext()
}

type TermoEscritaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoEscritaContext() *TermoEscritaContext {
	var p = new(TermoEscritaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termoEscrita
	return p
}

func InitEmptyTermoEscritaContext(p *TermoEscritaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termoEscrita
}

func (*TermoEscritaContext) IsTermoEscritaContext() {}

func NewTermoEscritaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoEscritaContext {
	var p = new(TermoEscritaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_termoEscrita

	return p
}

func (s *TermoEscritaContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoEscritaContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TermoEscritaContext) Dimensao2() IDimensao2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensao2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensao2Context)
}

func (s *TermoEscritaContext) CONSTANTE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONSTANTE, 0)
}

func (s *TermoEscritaContext) TEXTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserTEXTO, 0)
}

func (s *TermoEscritaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoEscritaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoEscritaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTermoEscrita(s)
	}
}

func (s *TermoEscritaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTermoEscrita(s)
	}
}

func (p *GrammarParser) TermoEscrita() (localctx ITermoEscritaContext) {
	localctx = NewTermoEscritaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_termoEscrita)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Dimensao2()
		}

	case GrammarParserCONSTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(GrammarParserCONSTANTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserTEXTO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)
			p.Match(GrammarParserTEXTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INovoTermoEscritaContext is an interface to support dynamic dispatch.
type INovoTermoEscritaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TermoEscrita() ITermoEscritaContext
	NovoTermoEscrita() INovoTermoEscritaContext

	// IsNovoTermoEscritaContext differentiates from other interfaces.
	IsNovoTermoEscritaContext()
}

type NovoTermoEscritaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNovoTermoEscritaContext() *NovoTermoEscritaContext {
	var p = new(NovoTermoEscritaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoTermoEscrita
	return p
}

func InitEmptyNovoTermoEscritaContext(p *NovoTermoEscritaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoTermoEscrita
}

func (*NovoTermoEscritaContext) IsNovoTermoEscritaContext() {}

func NewNovoTermoEscritaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NovoTermoEscritaContext {
	var p = new(NovoTermoEscritaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_novoTermoEscrita

	return p
}

func (s *NovoTermoEscritaContext) GetParser() antlr.Parser { return s.parser }

func (s *NovoTermoEscritaContext) TermoEscrita() ITermoEscritaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoEscritaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoEscritaContext)
}

func (s *NovoTermoEscritaContext) NovoTermoEscrita() INovoTermoEscritaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoTermoEscritaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoTermoEscritaContext)
}

func (s *NovoTermoEscritaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NovoTermoEscritaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NovoTermoEscritaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNovoTermoEscrita(s)
	}
}

func (s *NovoTermoEscritaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNovoTermoEscrita(s)
	}
}

func (p *GrammarParser) NovoTermoEscrita() (localctx INovoTermoEscritaContext) {
	localctx = NewNovoTermoEscritaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_novoTermoEscrita)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.TermoEscrita()
		}
		{
			p.SetState(222)
			p.NovoTermoEscrita()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelecaoContext is an interface to support dynamic dispatch.
type ISelecaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expressao() IExpressaoContext
	RPAREN() antlr.TerminalNode
	Bloco() IBlocoContext
	Senao() ISenaoContext

	// IsSelecaoContext differentiates from other interfaces.
	IsSelecaoContext()
}

type SelecaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelecaoContext() *SelecaoContext {
	var p = new(SelecaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_selecao
	return p
}

func InitEmptySelecaoContext(p *SelecaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_selecao
}

func (*SelecaoContext) IsSelecaoContext() {}

func NewSelecaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelecaoContext {
	var p = new(SelecaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_selecao

	return p
}

func (s *SelecaoContext) GetParser() antlr.Parser { return s.parser }

func (s *SelecaoContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *SelecaoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *SelecaoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *SelecaoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *SelecaoContext) Bloco() IBlocoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocoContext)
}

func (s *SelecaoContext) Senao() ISenaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISenaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISenaoContext)
}

func (s *SelecaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelecaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelecaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSelecao(s)
	}
}

func (s *SelecaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSelecao(s)
	}
}

func (p *GrammarParser) Selecao() (localctx ISelecaoContext) {
	localctx = NewSelecaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrammarParserRULE_selecao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(GrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Expressao()
	}
	{
		p.SetState(230)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Bloco()
	}
	{
		p.SetState(232)
		p.Senao()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISenaoContext is an interface to support dynamic dispatch.
type ISenaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Bloco() IBlocoContext

	// IsSenaoContext differentiates from other interfaces.
	IsSenaoContext()
}

type SenaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySenaoContext() *SenaoContext {
	var p = new(SenaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_senao
	return p
}

func InitEmptySenaoContext(p *SenaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_senao
}

func (*SenaoContext) IsSenaoContext() {}

func NewSenaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SenaoContext {
	var p = new(SenaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_senao

	return p
}

func (s *SenaoContext) GetParser() antlr.Parser { return s.parser }

func (s *SenaoContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *SenaoContext) Bloco() IBlocoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocoContext)
}

func (s *SenaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SenaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SenaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSenao(s)
	}
}

func (s *SenaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSenao(s)
	}
}

func (p *GrammarParser) Senao() (localctx ISenaoContext) {
	localctx = NewSenaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrammarParserRULE_senao)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserELSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Bloco()
		}

	case GrammarParserT__3, GrammarParserRETURN, GrammarParserIF, GrammarParserWHILE, GrammarParserSCANF, GrammarParserPRINTLN, GrammarParserRBRACE, GrammarParserID:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnquantoContext is an interface to support dynamic dispatch.
type IEnquantoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expressao() IExpressaoContext
	RPAREN() antlr.TerminalNode
	Bloco() IBlocoContext

	// IsEnquantoContext differentiates from other interfaces.
	IsEnquantoContext()
}

type EnquantoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnquantoContext() *EnquantoContext {
	var p = new(EnquantoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_enquanto
	return p
}

func InitEmptyEnquantoContext(p *EnquantoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_enquanto
}

func (*EnquantoContext) IsEnquantoContext() {}

func NewEnquantoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnquantoContext {
	var p = new(EnquantoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_enquanto

	return p
}

func (s *EnquantoContext) GetParser() antlr.Parser { return s.parser }

func (s *EnquantoContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *EnquantoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *EnquantoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *EnquantoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *EnquantoContext) Bloco() IBlocoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlocoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlocoContext)
}

func (s *EnquantoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnquantoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnquantoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterEnquanto(s)
	}
}

func (s *EnquantoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitEnquanto(s)
	}
}

func (p *GrammarParser) Enquanto() (localctx IEnquantoContext) {
	localctx = NewEnquantoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GrammarParserRULE_enquanto)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Expressao()
	}
	{
		p.SetState(242)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Bloco()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtribuicaoContext is an interface to support dynamic dispatch.
type IAtribuicaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Complemento() IComplementoContext
	SEMI() antlr.TerminalNode

	// IsAtribuicaoContext differentiates from other interfaces.
	IsAtribuicaoContext()
}

type AtribuicaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtribuicaoContext() *AtribuicaoContext {
	var p = new(AtribuicaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_atribuicao
	return p
}

func InitEmptyAtribuicaoContext(p *AtribuicaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_atribuicao
}

func (*AtribuicaoContext) IsAtribuicaoContext() {}

func NewAtribuicaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtribuicaoContext {
	var p = new(AtribuicaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_atribuicao

	return p
}

func (s *AtribuicaoContext) GetParser() antlr.Parser { return s.parser }

func (s *AtribuicaoContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AtribuicaoContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GrammarParserASSIGN, 0)
}

func (s *AtribuicaoContext) Complemento() IComplementoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComplementoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComplementoContext)
}

func (s *AtribuicaoContext) SEMI() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMI, 0)
}

func (s *AtribuicaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtribuicaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtribuicaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAtribuicao(s)
	}
}

func (s *AtribuicaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAtribuicao(s)
	}
}

func (p *GrammarParser) Atribuicao() (localctx IAtribuicaoContext) {
	localctx = NewAtribuicaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GrammarParserRULE_atribuicao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(GrammarParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Complemento()
	}
	{
		p.SetState(248)
		p.Match(GrammarParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComplementoContext is an interface to support dynamic dispatch.
type IComplementoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressao() IExpressaoContext
	Funcao() IFuncaoContext

	// IsComplementoContext differentiates from other interfaces.
	IsComplementoContext()
}

type ComplementoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComplementoContext() *ComplementoContext {
	var p = new(ComplementoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_complemento
	return p
}

func InitEmptyComplementoContext(p *ComplementoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_complemento
}

func (*ComplementoContext) IsComplementoContext() {}

func NewComplementoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComplementoContext {
	var p = new(ComplementoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_complemento

	return p
}

func (s *ComplementoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComplementoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *ComplementoContext) Funcao() IFuncaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncaoContext)
}

func (s *ComplementoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComplementoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComplementoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComplemento(s)
	}
}

func (s *ComplementoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComplemento(s)
	}
}

func (p *GrammarParser) Complemento() (localctx IComplementoContext) {
	localctx = NewComplementoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GrammarParserRULE_complemento)
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__6, GrammarParserCONSTANTE, GrammarParserTEXTO, GrammarParserLPAREN, GrammarParserPLUS, GrammarParserMINUS, GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Expressao()
		}

	case GrammarParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Funcao()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncaoContext is an interface to support dynamic dispatch.
type IFuncaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Argumentos() IArgumentosContext
	RPAREN() antlr.TerminalNode

	// IsFuncaoContext differentiates from other interfaces.
	IsFuncaoContext()
}

type FuncaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncaoContext() *FuncaoContext {
	var p = new(FuncaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcao
	return p
}

func InitEmptyFuncaoContext(p *FuncaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcao
}

func (*FuncaoContext) IsFuncaoContext() {}

func NewFuncaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncaoContext {
	var p = new(FuncaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_funcao

	return p
}

func (s *FuncaoContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncaoContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *FuncaoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FuncaoContext) Argumentos() IArgumentosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentosContext)
}

func (s *FuncaoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FuncaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFuncao(s)
	}
}

func (s *FuncaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFuncao(s)
	}
}

func (p *GrammarParser) Funcao() (localctx IFuncaoContext) {
	localctx = NewFuncaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GrammarParserRULE_funcao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(GrammarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(GrammarParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Argumentos()
	}
	{
		p.SetState(258)
		p.Match(GrammarParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentosContext is an interface to support dynamic dispatch.
type IArgumentosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressao() IExpressaoContext
	NovoArgumento() INovoArgumentoContext

	// IsArgumentosContext differentiates from other interfaces.
	IsArgumentosContext()
}

type ArgumentosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentosContext() *ArgumentosContext {
	var p = new(ArgumentosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_argumentos
	return p
}

func InitEmptyArgumentosContext(p *ArgumentosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_argumentos
}

func (*ArgumentosContext) IsArgumentosContext() {}

func NewArgumentosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentosContext {
	var p = new(ArgumentosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_argumentos

	return p
}

func (s *ArgumentosContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentosContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *ArgumentosContext) NovoArgumento() INovoArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoArgumentoContext)
}

func (s *ArgumentosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterArgumentos(s)
	}
}

func (s *ArgumentosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitArgumentos(s)
	}
}

func (p *GrammarParser) Argumentos() (localctx IArgumentosContext) {
	localctx = NewArgumentosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrammarParserRULE_argumentos)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__6, GrammarParserCONSTANTE, GrammarParserTEXTO, GrammarParserLPAREN, GrammarParserPLUS, GrammarParserMINUS, GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Expressao()
		}
		{
			p.SetState(261)
			p.NovoArgumento()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INovoArgumentoContext is an interface to support dynamic dispatch.
type INovoArgumentoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressao() IExpressaoContext
	NovoArgumento() INovoArgumentoContext

	// IsNovoArgumentoContext differentiates from other interfaces.
	IsNovoArgumentoContext()
}

type NovoArgumentoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNovoArgumentoContext() *NovoArgumentoContext {
	var p = new(NovoArgumentoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoArgumento
	return p
}

func InitEmptyNovoArgumentoContext(p *NovoArgumentoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_novoArgumento
}

func (*NovoArgumentoContext) IsNovoArgumentoContext() {}

func NewNovoArgumentoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NovoArgumentoContext {
	var p = new(NovoArgumentoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_novoArgumento

	return p
}

func (s *NovoArgumentoContext) GetParser() antlr.Parser { return s.parser }

func (s *NovoArgumentoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *NovoArgumentoContext) NovoArgumento() INovoArgumentoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovoArgumentoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovoArgumentoContext)
}

func (s *NovoArgumentoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NovoArgumentoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NovoArgumentoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNovoArgumento(s)
	}
}

func (s *NovoArgumentoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNovoArgumento(s)
	}
}

func (p *GrammarParser) NovoArgumento() (localctx INovoArgumentoContext) {
	localctx = NewNovoArgumentoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrammarParserRULE_novoArgumento)
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Expressao()
		}
		{
			p.SetState(268)
			p.NovoArgumento()
		}

	case GrammarParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRetornoContext is an interface to support dynamic dispatch.
type IRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expressao() IExpressaoContext
	SEMI() antlr.TerminalNode

	// IsRetornoContext differentiates from other interfaces.
	IsRetornoContext()
}

type RetornoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetornoContext() *RetornoContext {
	var p = new(RetornoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_retorno
	return p
}

func InitEmptyRetornoContext(p *RetornoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_retorno
}

func (*RetornoContext) IsRetornoContext() {}

func NewRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornoContext {
	var p = new(RetornoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_retorno

	return p
}

func (s *RetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornoContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *RetornoContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *RetornoContext) SEMI() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMI, 0)
}

func (s *RetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRetorno(s)
	}
}

func (s *RetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRetorno(s)
	}
}

func (p *GrammarParser) Retorno() (localctx IRetornoContext) {
	localctx = NewRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GrammarParserRULE_retorno)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(GrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Expressao()
	}
	{
		p.SetState(275)
		p.Match(GrammarParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressaoContext is an interface to support dynamic dispatch.
type IExpressaoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprOu() IExprOuContext

	// IsExpressaoContext differentiates from other interfaces.
	IsExpressaoContext()
}

type ExpressaoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressaoContext() *ExpressaoContext {
	var p = new(ExpressaoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expressao
	return p
}

func InitEmptyExpressaoContext(p *ExpressaoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expressao
}

func (*ExpressaoContext) IsExpressaoContext() {}

func NewExpressaoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressaoContext {
	var p = new(ExpressaoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expressao

	return p
}

func (s *ExpressaoContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressaoContext) ExprOu() IExprOuContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOuContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprOuContext)
}

func (s *ExpressaoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressaoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressaoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExpressao(s)
	}
}

func (s *ExpressaoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExpressao(s)
	}
}

func (p *GrammarParser) Expressao() (localctx IExpressaoContext) {
	localctx = NewExpressaoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GrammarParserRULE_expressao)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.ExprOu()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprOuContext is an interface to support dynamic dispatch.
type IExprOuContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprE() IExprEContext
	ExprOu2() IExprOu2Context

	// IsExprOuContext differentiates from other interfaces.
	IsExprOuContext()
}

type ExprOuContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOuContext() *ExprOuContext {
	var p = new(ExprOuContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprOu
	return p
}

func InitEmptyExprOuContext(p *ExprOuContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprOu
}

func (*ExprOuContext) IsExprOuContext() {}

func NewExprOuContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOuContext {
	var p = new(ExprOuContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprOu

	return p
}

func (s *ExprOuContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOuContext) ExprE() IExprEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprEContext)
}

func (s *ExprOuContext) ExprOu2() IExprOu2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOu2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprOu2Context)
}

func (s *ExprOuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOuContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprOu(s)
	}
}

func (s *ExprOuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprOu(s)
	}
}

func (p *GrammarParser) ExprOu() (localctx IExprOuContext) {
	localctx = NewExprOuContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GrammarParserRULE_exprOu)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.ExprE()
	}
	{
		p.SetState(280)
		p.ExprOu2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprOu2Context is an interface to support dynamic dispatch.
type IExprOu2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprE() IExprEContext
	ExprOu2() IExprOu2Context

	// IsExprOu2Context differentiates from other interfaces.
	IsExprOu2Context()
}

type ExprOu2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOu2Context() *ExprOu2Context {
	var p = new(ExprOu2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprOu2
	return p
}

func InitEmptyExprOu2Context(p *ExprOu2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprOu2
}

func (*ExprOu2Context) IsExprOu2Context() {}

func NewExprOu2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOu2Context {
	var p = new(ExprOu2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprOu2

	return p
}

func (s *ExprOu2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprOu2Context) ExprE() IExprEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprEContext)
}

func (s *ExprOu2Context) ExprOu2() IExprOu2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOu2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprOu2Context)
}

func (s *ExprOu2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOu2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOu2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprOu2(s)
	}
}

func (s *ExprOu2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprOu2(s)
	}
}

func (p *GrammarParser) ExprOu2() (localctx IExprOu2Context) {
	localctx = NewExprOu2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GrammarParserRULE_exprOu2)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.ExprE()
		}
		{
			p.SetState(284)
			p.ExprOu2()
		}

	case GrammarParserT__2, GrammarParserRPAREN, GrammarParserSEMI:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprEContext is an interface to support dynamic dispatch.
type IExprEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprRelacional() IExprRelacionalContext
	ExprE2() IExprE2Context

	// IsExprEContext differentiates from other interfaces.
	IsExprEContext()
}

type ExprEContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprEContext() *ExprEContext {
	var p = new(ExprEContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprE
	return p
}

func InitEmptyExprEContext(p *ExprEContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprE
}

func (*ExprEContext) IsExprEContext() {}

func NewExprEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprEContext {
	var p = new(ExprEContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprE

	return p
}

func (s *ExprEContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprEContext) ExprRelacional() IExprRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprRelacionalContext)
}

func (s *ExprEContext) ExprE2() IExprE2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprE2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprE2Context)
}

func (s *ExprEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprE(s)
	}
}

func (s *ExprEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprE(s)
	}
}

func (p *GrammarParser) ExprE() (localctx IExprEContext) {
	localctx = NewExprEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GrammarParserRULE_exprE)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.ExprRelacional()
	}
	{
		p.SetState(290)
		p.ExprE2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprE2Context is an interface to support dynamic dispatch.
type IExprE2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprRelacional() IExprRelacionalContext
	ExprE2() IExprE2Context

	// IsExprE2Context differentiates from other interfaces.
	IsExprE2Context()
}

type ExprE2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprE2Context() *ExprE2Context {
	var p = new(ExprE2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprE2
	return p
}

func InitEmptyExprE2Context(p *ExprE2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprE2
}

func (*ExprE2Context) IsExprE2Context() {}

func NewExprE2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprE2Context {
	var p = new(ExprE2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprE2

	return p
}

func (s *ExprE2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprE2Context) ExprRelacional() IExprRelacionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprRelacionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprRelacionalContext)
}

func (s *ExprE2Context) ExprE2() IExprE2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprE2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprE2Context)
}

func (s *ExprE2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprE2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprE2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprE2(s)
	}
}

func (s *ExprE2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprE2(s)
	}
}

func (p *GrammarParser) ExprE2() (localctx IExprE2Context) {
	localctx = NewExprE2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GrammarParserRULE_exprE2)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.ExprRelacional()
		}
		{
			p.SetState(294)
			p.ExprE2()
		}

	case GrammarParserT__2, GrammarParserT__4, GrammarParserRPAREN, GrammarParserSEMI:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprRelacionalContext is an interface to support dynamic dispatch.
type IExprRelacionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprAditiva() IExprAditivaContext
	ExprRelacional2() IExprRelacional2Context

	// IsExprRelacionalContext differentiates from other interfaces.
	IsExprRelacionalContext()
}

type ExprRelacionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprRelacionalContext() *ExprRelacionalContext {
	var p = new(ExprRelacionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprRelacional
	return p
}

func InitEmptyExprRelacionalContext(p *ExprRelacionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprRelacional
}

func (*ExprRelacionalContext) IsExprRelacionalContext() {}

func NewExprRelacionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprRelacionalContext {
	var p = new(ExprRelacionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprRelacional

	return p
}

func (s *ExprRelacionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprRelacionalContext) ExprAditiva() IExprAditivaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAditivaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAditivaContext)
}

func (s *ExprRelacionalContext) ExprRelacional2() IExprRelacional2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprRelacional2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprRelacional2Context)
}

func (s *ExprRelacionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRelacionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprRelacionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprRelacional(s)
	}
}

func (s *ExprRelacionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprRelacional(s)
	}
}

func (p *GrammarParser) ExprRelacional() (localctx IExprRelacionalContext) {
	localctx = NewExprRelacionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GrammarParserRULE_exprRelacional)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.ExprAditiva()
	}
	{
		p.SetState(300)
		p.ExprRelacional2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprRelacional2Context is an interface to support dynamic dispatch.
type IExprRelacional2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMP() antlr.TerminalNode
	ExprAditiva() IExprAditivaContext

	// IsExprRelacional2Context differentiates from other interfaces.
	IsExprRelacional2Context()
}

type ExprRelacional2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprRelacional2Context() *ExprRelacional2Context {
	var p = new(ExprRelacional2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprRelacional2
	return p
}

func InitEmptyExprRelacional2Context(p *ExprRelacional2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprRelacional2
}

func (*ExprRelacional2Context) IsExprRelacional2Context() {}

func NewExprRelacional2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprRelacional2Context {
	var p = new(ExprRelacional2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprRelacional2

	return p
}

func (s *ExprRelacional2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprRelacional2Context) COMP() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMP, 0)
}

func (s *ExprRelacional2Context) ExprAditiva() IExprAditivaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAditivaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAditivaContext)
}

func (s *ExprRelacional2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRelacional2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprRelacional2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprRelacional2(s)
	}
}

func (s *ExprRelacional2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprRelacional2(s)
	}
}

func (p *GrammarParser) ExprRelacional2() (localctx IExprRelacional2Context) {
	localctx = NewExprRelacional2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GrammarParserRULE_exprRelacional2)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserCOMP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(GrammarParserCOMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.ExprAditiva()
		}

	case GrammarParserT__2, GrammarParserT__4, GrammarParserT__5, GrammarParserRPAREN, GrammarParserSEMI:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprAditivaContext is an interface to support dynamic dispatch.
type IExprAditivaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExprMultiplicativa() IExprMultiplicativaContext
	ExprAditiva2() IExprAditiva2Context

	// IsExprAditivaContext differentiates from other interfaces.
	IsExprAditivaContext()
}

type ExprAditivaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprAditivaContext() *ExprAditivaContext {
	var p = new(ExprAditivaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprAditiva
	return p
}

func InitEmptyExprAditivaContext(p *ExprAditivaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprAditiva
}

func (*ExprAditivaContext) IsExprAditivaContext() {}

func NewExprAditivaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprAditivaContext {
	var p = new(ExprAditivaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprAditiva

	return p
}

func (s *ExprAditivaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprAditivaContext) ExprMultiplicativa() IExprMultiplicativaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMultiplicativaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMultiplicativaContext)
}

func (s *ExprAditivaContext) ExprAditiva2() IExprAditiva2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAditiva2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAditiva2Context)
}

func (s *ExprAditivaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAditivaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprAditivaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprAditiva(s)
	}
}

func (s *ExprAditivaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprAditiva(s)
	}
}

func (p *GrammarParser) ExprAditiva() (localctx IExprAditivaContext) {
	localctx = NewExprAditivaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GrammarParserRULE_exprAditiva)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.ExprMultiplicativa()
	}
	{
		p.SetState(308)
		p.ExprAditiva2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprAditiva2Context is an interface to support dynamic dispatch.
type IExprAditiva2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpAditivo() IOpAditivoContext
	ExprMultiplicativa() IExprMultiplicativaContext
	ExprAditiva2() IExprAditiva2Context

	// IsExprAditiva2Context differentiates from other interfaces.
	IsExprAditiva2Context()
}

type ExprAditiva2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprAditiva2Context() *ExprAditiva2Context {
	var p = new(ExprAditiva2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprAditiva2
	return p
}

func InitEmptyExprAditiva2Context(p *ExprAditiva2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprAditiva2
}

func (*ExprAditiva2Context) IsExprAditiva2Context() {}

func NewExprAditiva2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprAditiva2Context {
	var p = new(ExprAditiva2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprAditiva2

	return p
}

func (s *ExprAditiva2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprAditiva2Context) OpAditivo() IOpAditivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpAditivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpAditivoContext)
}

func (s *ExprAditiva2Context) ExprMultiplicativa() IExprMultiplicativaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMultiplicativaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMultiplicativaContext)
}

func (s *ExprAditiva2Context) ExprAditiva2() IExprAditiva2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprAditiva2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprAditiva2Context)
}

func (s *ExprAditiva2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAditiva2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprAditiva2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprAditiva2(s)
	}
}

func (s *ExprAditiva2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprAditiva2(s)
	}
}

func (p *GrammarParser) ExprAditiva2() (localctx IExprAditiva2Context) {
	localctx = NewExprAditiva2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GrammarParserRULE_exprAditiva2)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPLUS, GrammarParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.OpAditivo()
		}
		{
			p.SetState(311)
			p.ExprMultiplicativa()
		}
		{
			p.SetState(312)
			p.ExprAditiva2()
		}

	case GrammarParserT__1, GrammarParserT__2, GrammarParserT__4, GrammarParserT__5, GrammarParserCOMP, GrammarParserRPAREN, GrammarParserSEMI:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpAditivoContext is an interface to support dynamic dispatch.
type IOpAditivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsOpAditivoContext differentiates from other interfaces.
	IsOpAditivoContext()
}

type OpAditivoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpAditivoContext() *OpAditivoContext {
	var p = new(OpAditivoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_opAditivo
	return p
}

func InitEmptyOpAditivoContext(p *OpAditivoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_opAditivo
}

func (*OpAditivoContext) IsOpAditivoContext() {}

func NewOpAditivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpAditivoContext {
	var p = new(OpAditivoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_opAditivo

	return p
}

func (s *OpAditivoContext) GetParser() antlr.Parser { return s.parser }

func (s *OpAditivoContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUS, 0)
}

func (s *OpAditivoContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *OpAditivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpAditivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpAditivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterOpAditivo(s)
	}
}

func (s *OpAditivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitOpAditivo(s)
	}
}

func (p *GrammarParser) OpAditivo() (localctx IOpAditivoContext) {
	localctx = NewOpAditivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GrammarParserRULE_opAditivo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserPLUS || _la == GrammarParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprMultiplicativaContext is an interface to support dynamic dispatch.
type IExprMultiplicativaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fator() IFatorContext
	ExprMultiplicativa2() IExprMultiplicativa2Context

	// IsExprMultiplicativaContext differentiates from other interfaces.
	IsExprMultiplicativaContext()
}

type ExprMultiplicativaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprMultiplicativaContext() *ExprMultiplicativaContext {
	var p = new(ExprMultiplicativaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa
	return p
}

func InitEmptyExprMultiplicativaContext(p *ExprMultiplicativaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa
}

func (*ExprMultiplicativaContext) IsExprMultiplicativaContext() {}

func NewExprMultiplicativaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprMultiplicativaContext {
	var p = new(ExprMultiplicativaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa

	return p
}

func (s *ExprMultiplicativaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprMultiplicativaContext) Fator() IFatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorContext)
}

func (s *ExprMultiplicativaContext) ExprMultiplicativa2() IExprMultiplicativa2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMultiplicativa2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMultiplicativa2Context)
}

func (s *ExprMultiplicativaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMultiplicativaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprMultiplicativaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprMultiplicativa(s)
	}
}

func (s *ExprMultiplicativaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprMultiplicativa(s)
	}
}

func (p *GrammarParser) ExprMultiplicativa() (localctx IExprMultiplicativaContext) {
	localctx = NewExprMultiplicativaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GrammarParserRULE_exprMultiplicativa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Fator()
	}
	{
		p.SetState(320)
		p.ExprMultiplicativa2()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprMultiplicativa2Context is an interface to support dynamic dispatch.
type IExprMultiplicativa2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpMultiplicativo() IOpMultiplicativoContext
	Fator() IFatorContext
	ExprMultiplicativa2() IExprMultiplicativa2Context

	// IsExprMultiplicativa2Context differentiates from other interfaces.
	IsExprMultiplicativa2Context()
}

type ExprMultiplicativa2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprMultiplicativa2Context() *ExprMultiplicativa2Context {
	var p = new(ExprMultiplicativa2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa2
	return p
}

func InitEmptyExprMultiplicativa2Context(p *ExprMultiplicativa2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa2
}

func (*ExprMultiplicativa2Context) IsExprMultiplicativa2Context() {}

func NewExprMultiplicativa2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprMultiplicativa2Context {
	var p = new(ExprMultiplicativa2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprMultiplicativa2

	return p
}

func (s *ExprMultiplicativa2Context) GetParser() antlr.Parser { return s.parser }

func (s *ExprMultiplicativa2Context) OpMultiplicativo() IOpMultiplicativoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMultiplicativoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMultiplicativoContext)
}

func (s *ExprMultiplicativa2Context) Fator() IFatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorContext)
}

func (s *ExprMultiplicativa2Context) ExprMultiplicativa2() IExprMultiplicativa2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMultiplicativa2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMultiplicativa2Context)
}

func (s *ExprMultiplicativa2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMultiplicativa2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprMultiplicativa2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprMultiplicativa2(s)
	}
}

func (s *ExprMultiplicativa2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprMultiplicativa2(s)
	}
}

func (p *GrammarParser) ExprMultiplicativa2() (localctx IExprMultiplicativa2Context) {
	localctx = NewExprMultiplicativa2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GrammarParserRULE_exprMultiplicativa2)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserMUL, GrammarParserDIV, GrammarParserMOD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.OpMultiplicativo()
		}
		{
			p.SetState(323)
			p.Fator()
		}
		{
			p.SetState(324)
			p.ExprMultiplicativa2()
		}

	case GrammarParserT__1, GrammarParserT__2, GrammarParserT__4, GrammarParserT__5, GrammarParserCOMP, GrammarParserRPAREN, GrammarParserSEMI, GrammarParserPLUS, GrammarParserMINUS:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpMultiplicativoContext is an interface to support dynamic dispatch.
type IOpMultiplicativoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsOpMultiplicativoContext differentiates from other interfaces.
	IsOpMultiplicativoContext()
}

type OpMultiplicativoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpMultiplicativoContext() *OpMultiplicativoContext {
	var p = new(OpMultiplicativoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_opMultiplicativo
	return p
}

func InitEmptyOpMultiplicativoContext(p *OpMultiplicativoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_opMultiplicativo
}

func (*OpMultiplicativoContext) IsOpMultiplicativoContext() {}

func NewOpMultiplicativoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpMultiplicativoContext {
	var p = new(OpMultiplicativoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_opMultiplicativo

	return p
}

func (s *OpMultiplicativoContext) GetParser() antlr.Parser { return s.parser }

func (s *OpMultiplicativoContext) MUL() antlr.TerminalNode {
	return s.GetToken(GrammarParserMUL, 0)
}

func (s *OpMultiplicativoContext) DIV() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIV, 0)
}

func (s *OpMultiplicativoContext) MOD() antlr.TerminalNode {
	return s.GetToken(GrammarParserMOD, 0)
}

func (s *OpMultiplicativoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpMultiplicativoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpMultiplicativoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterOpMultiplicativo(s)
	}
}

func (s *OpMultiplicativoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitOpMultiplicativo(s)
	}
}

func (p *GrammarParser) OpMultiplicativo() (localctx IOpMultiplicativoContext) {
	localctx = NewOpMultiplicativoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GrammarParserRULE_opMultiplicativo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60129542144) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFatorContext is an interface to support dynamic dispatch.
type IFatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sinal() ISinalContext
	Termo() ITermoContext
	TEXTO() antlr.TerminalNode
	Fator() IFatorContext
	LPAREN() antlr.TerminalNode
	Expressao() IExpressaoContext
	RPAREN() antlr.TerminalNode

	// IsFatorContext differentiates from other interfaces.
	IsFatorContext()
}

type FatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatorContext() *FatorContext {
	var p = new(FatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_fator
	return p
}

func InitEmptyFatorContext(p *FatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_fator
}

func (*FatorContext) IsFatorContext() {}

func NewFatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatorContext {
	var p = new(FatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_fator

	return p
}

func (s *FatorContext) GetParser() antlr.Parser { return s.parser }

func (s *FatorContext) Sinal() ISinalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISinalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISinalContext)
}

func (s *FatorContext) Termo() ITermoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermoContext)
}

func (s *FatorContext) TEXTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserTEXTO, 0)
}

func (s *FatorContext) Fator() IFatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatorContext)
}

func (s *FatorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLPAREN, 0)
}

func (s *FatorContext) Expressao() IExpressaoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressaoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressaoContext)
}

func (s *FatorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRPAREN, 0)
}

func (s *FatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFator(s)
	}
}

func (s *FatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFator(s)
	}
}

func (p *GrammarParser) Fator() (localctx IFatorContext) {
	localctx = NewFatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GrammarParserRULE_fator)
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserCONSTANTE, GrammarParserPLUS, GrammarParserMINUS, GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(331)
			p.Sinal()
		}
		{
			p.SetState(332)
			p.Termo()
		}

	case GrammarParserTEXTO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(GrammarParserTEXTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(335)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Fator()
		}

	case GrammarParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(337)
			p.Match(GrammarParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Expressao()
		}
		{
			p.SetState(339)
			p.Match(GrammarParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermoContext is an interface to support dynamic dispatch.
type ITermoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Dimensao2() IDimensao2Context
	CONSTANTE() antlr.TerminalNode

	// IsTermoContext differentiates from other interfaces.
	IsTermoContext()
}

type TermoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermoContext() *TermoContext {
	var p = new(TermoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termo
	return p
}

func InitEmptyTermoContext(p *TermoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_termo
}

func (*TermoContext) IsTermoContext() {}

func NewTermoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermoContext {
	var p = new(TermoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_termo

	return p
}

func (s *TermoContext) GetParser() antlr.Parser { return s.parser }

func (s *TermoContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TermoContext) Dimensao2() IDimensao2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensao2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensao2Context)
}

func (s *TermoContext) CONSTANTE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONSTANTE, 0)
}

func (s *TermoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterTermo(s)
	}
}

func (s *TermoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitTermo(s)
	}
}

func (p *GrammarParser) Termo() (localctx ITermoContext) {
	localctx = NewTermoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GrammarParserRULE_termo)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Dimensao2()
		}

	case GrammarParserCONSTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Match(GrammarParserCONSTANTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISinalContext is an interface to support dynamic dispatch.
type ISinalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsSinalContext differentiates from other interfaces.
	IsSinalContext()
}

type SinalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinalContext() *SinalContext {
	var p = new(SinalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_sinal
	return p
}

func InitEmptySinalContext(p *SinalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_sinal
}

func (*SinalContext) IsSinalContext() {}

func NewSinalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinalContext {
	var p = new(SinalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_sinal

	return p
}

func (s *SinalContext) GetParser() antlr.Parser { return s.parser }

func (s *SinalContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserPLUS, 0)
}

func (s *SinalContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMINUS, 0)
}

func (s *SinalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSinal(s)
	}
}

func (s *SinalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSinal(s)
	}
}

func (p *GrammarParser) Sinal() (localctx ISinalContext) {
	localctx = NewSinalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GrammarParserRULE_sinal)
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.Match(GrammarParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Match(GrammarParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserCONSTANTE, GrammarParserID:
		p.EnterOuterAlt(localctx, 3)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
