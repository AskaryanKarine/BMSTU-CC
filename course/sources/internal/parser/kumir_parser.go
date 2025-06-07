// Code generated from KumirParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // KumirParser
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

type KumirParser struct {
	*antlr.BaseParser
}

var KumirParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kumirparserParserInit() {
	staticData := &KumirParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'\\u0430\\u043B\\u0433'", "'\\u043D\\u0430\\u0447'", "'\\u043A\\u043E\\u043D'",
		"'\\u0434\\u0430\\u043D\\u043E'", "'\\u043D\\u0430\\u0434\\u043E'",
		"'\\u0443\\u0442\\u0432'", "'\\u043D\\u0446'", "", "'\\u043A\\u0446'",
		"'\\u0435\\u0441\\u043B\\u0438'", "'\\u0442\\u043E'", "'\\u0438\\u043D\\u0430\\u0447\\u0435'",
		"'\\u0432\\u0441\\u0435'", "'\\u0432\\u044B\\u0431\\u043E\\u0440'",
		"'\\u043F\\u0440\\u0438'", "'\\u0432\\u0432\\u043E\\u0434'", "'\\u0432\\u044B\\u0432\\u043E\\u0434'",
		"':='", "'\\u0432\\u044B\\u0445\\u043E\\u0434'", "'\\u043F\\u0430\\u0443\\u0437\\u0430'",
		"'\\u0441\\u0442\\u043E\\u043F'", "'\\u0434\\u043B\\u044F'", "'\\u043F\\u043E\\u043A\\u0430'",
		"'\\u0440\\u0430\\u0437'", "'\\u043E\\u0442'", "'\\u0434\\u043E'", "'\\u0448\\u0430\\u0433'",
		"'\\u043D\\u0441'", "'\\u043D\\u0435'", "'\\u0438'", "'\\u0438\\u043B\\u0438'",
		"'\\u0440\\u0435\\u0437'", "'\\u0430\\u0440\\u0433'", "", "'\\u0437\\u043D\\u0430\\u0447'",
		"'\\u0446\\u0435\\u043B'", "'\\u0432\\u0435\\u0449'", "'\\u043B\\u043E\\u0433'",
		"'\\u0441\\u0438\\u043C'", "'\\u043B\\u0438\\u0442'", "'\\u0442\\u0430\\u0431'",
		"", "", "", "", "", "'\\u0434\\u0430'", "'\\u043D\\u0435\\u0442'", "'**'",
		"", "", "", "'+'", "'-'", "'*'", "'/'", "'='", "'<'", "'>'", "'('",
		"')'", "'['", "']'", "'{'", "'}'", "','", "':'", "';'", "'@@'", "'@'",
		"'div'", "'mod'",
	}
	staticData.SymbolicNames = []string{
		"", "ALG_HEADER", "ALG_BEGIN", "ALG_END", "PRE_CONDITION", "POST_CONDITION",
		"ASSERTION", "LOOP", "ENDLOOP_COND", "ENDLOOP", "IF", "THEN", "ELSE",
		"FI", "SWITCH", "CASE", "INPUT", "OUTPUT", "ASSIGN", "EXIT", "PAUSE",
		"STOP", "FOR", "WHILE", "TIMES", "FROM", "TO", "STEP", "NEWLINE_CONST",
		"NOT", "AND", "OR", "OUT_PARAM", "IN_PARAM", "INOUT_PARAM", "RETURN_VALUE",
		"INTEGER_TYPE", "REAL_TYPE", "BOOLEAN_TYPE", "CHAR_TYPE", "STRING_TYPE",
		"TABLE_SUFFIX", "INTEGER_ARRAY_TYPE", "REAL_ARRAY_TYPE", "CHAR_ARRAY_TYPE",
		"STRING_ARRAY_TYPE", "BOOLEAN_ARRAY_TYPE", "TRUE", "FALSE", "POWER",
		"GE", "LE", "NE", "PLUS", "MINUS", "MUL", "DIV", "EQ", "LT", "GT", "LPAREN",
		"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "COMMA", "COLON",
		"SEMICOLON", "ATAT", "AT", "DIV_OP", "MOD_OP", "CHAR_LITERAL", "STRING",
		"REAL", "INTEGER", "ID", "LINE_COMMENT", "DOC_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"qualifiedIdentifier", "literal", "expressionList", "arrayLiteral",
		"primaryExpression", "argumentList", "indexList", "postfixExpression",
		"unaryExpression", "powerExpression", "multiplicativeExpression", "additiveExpression",
		"relationalExpression", "equalityExpression", "logicalAndExpression",
		"logicalOrExpression", "expression", "typeSpecifier", "basicType", "arrayType",
		"arrayBounds", "variableDeclarationItem", "variableList", "variableDeclaration",
		"globalDeclaration", "globalAssignment", "parameterDeclaration", "parameterList",
		"algorithmNameTokens", "algorithmName", "algorithmHeader", "preCondition",
		"postCondition", "algorithmBody", "statementSequence", "lvalue", "assignmentStatement",
		"ioArgument", "ioArgumentList", "ioStatement", "ifStatement", "caseBlock",
		"switchStatement", "endLoopCondition", "loopSpecifier", "loopStatement",
		"exitStatement", "pauseStatement", "stopStatement", "assertionStatement",
		"procedureCallStatement", "statement", "algorithmDefinition", "programItem",
		"program",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 80, 540, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 5, 2, 118, 8, 2, 10, 2, 12, 2, 121, 9, 2, 1, 3, 1, 3, 3, 3, 125, 8,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 137,
		8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 142, 8, 5, 10, 5, 12, 5, 145, 9, 5, 1, 6,
		1, 6, 1, 6, 3, 6, 150, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 156, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 165, 8, 7, 1, 7, 5, 7, 168,
		8, 7, 10, 7, 12, 7, 171, 9, 7, 1, 8, 1, 8, 1, 8, 3, 8, 176, 8, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 181, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10, 186, 8, 10, 10,
		10, 12, 10, 189, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 194, 8, 11, 10, 11,
		12, 11, 197, 9, 11, 1, 12, 1, 12, 1, 12, 5, 12, 202, 8, 12, 10, 12, 12,
		12, 205, 9, 12, 1, 13, 1, 13, 1, 13, 5, 13, 210, 8, 13, 10, 13, 12, 13,
		213, 9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 218, 8, 14, 10, 14, 12, 14, 221,
		9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 226, 8, 15, 10, 15, 12, 15, 229, 9,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 236, 8, 17, 3, 17, 238, 8,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 5, 21, 253, 8, 21, 10, 21, 12, 21, 256, 9, 21, 1,
		21, 1, 21, 3, 21, 260, 8, 21, 1, 21, 1, 21, 3, 21, 264, 8, 21, 1, 22, 1,
		22, 1, 22, 5, 22, 269, 8, 22, 10, 22, 12, 22, 272, 9, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 280, 8, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 3, 25, 287, 8, 25, 1, 25, 3, 25, 290, 8, 25, 1, 26, 3, 26, 293,
		8, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 301, 8, 27, 10,
		27, 12, 27, 304, 9, 27, 1, 28, 4, 28, 307, 8, 28, 11, 28, 12, 28, 308,
		1, 29, 4, 29, 312, 8, 29, 11, 29, 12, 29, 313, 1, 30, 1, 30, 3, 30, 318,
		8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 323, 8, 30, 1, 30, 3, 30, 326, 8, 30,
		1, 30, 3, 30, 329, 8, 30, 1, 31, 1, 31, 1, 31, 3, 31, 334, 8, 31, 1, 32,
		1, 32, 1, 32, 3, 32, 339, 8, 32, 1, 33, 1, 33, 1, 34, 5, 34, 344, 8, 34,
		10, 34, 12, 34, 347, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 354,
		8, 35, 1, 35, 3, 35, 357, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3,
		36, 364, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 371, 8, 37, 3,
		37, 373, 8, 37, 1, 37, 3, 37, 376, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 381,
		8, 38, 10, 38, 12, 38, 384, 9, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 3, 40, 395, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 4, 42, 406, 8, 42, 11, 42, 12, 42, 407,
		1, 42, 1, 42, 3, 42, 412, 8, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 427, 8, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 434, 8, 44, 1, 45, 1, 45, 3,
		45, 438, 8, 45, 1, 45, 1, 45, 1, 45, 3, 45, 443, 8, 45, 1, 46, 1, 46, 1,
		47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 3, 50,
		457, 8, 50, 1, 50, 3, 50, 460, 8, 50, 1, 51, 1, 51, 3, 51, 464, 8, 51,
		1, 51, 1, 51, 3, 51, 468, 8, 51, 1, 51, 1, 51, 3, 51, 472, 8, 51, 1, 51,
		1, 51, 3, 51, 476, 8, 51, 1, 51, 1, 51, 3, 51, 480, 8, 51, 1, 51, 1, 51,
		3, 51, 484, 8, 51, 1, 51, 1, 51, 3, 51, 488, 8, 51, 1, 51, 1, 51, 3, 51,
		492, 8, 51, 1, 51, 1, 51, 3, 51, 496, 8, 51, 1, 51, 1, 51, 3, 51, 500,
		8, 51, 1, 51, 3, 51, 503, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 509,
		8, 52, 10, 52, 12, 52, 512, 9, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 518,
		8, 52, 1, 52, 3, 52, 521, 8, 52, 1, 53, 1, 53, 3, 53, 525, 8, 53, 1, 54,
		5, 54, 528, 8, 54, 10, 54, 12, 54, 531, 9, 54, 1, 54, 4, 54, 534, 8, 54,
		11, 54, 12, 54, 535, 1, 54, 1, 54, 1, 54, 0, 0, 55, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 0, 11, 3, 0, 28,
		28, 47, 48, 73, 76, 2, 0, 29, 29, 53, 54, 2, 0, 55, 56, 71, 72, 1, 0, 53,
		54, 2, 0, 50, 51, 58, 59, 2, 0, 52, 52, 57, 57, 1, 0, 36, 40, 1, 0, 42,
		46, 1, 0, 32, 34, 4, 1, 2, 2, 4, 5, 60, 60, 68, 68, 1, 0, 16, 17, 570,
		0, 110, 1, 0, 0, 0, 2, 112, 1, 0, 0, 0, 4, 114, 1, 0, 0, 0, 6, 122, 1,
		0, 0, 0, 8, 136, 1, 0, 0, 0, 10, 138, 1, 0, 0, 0, 12, 155, 1, 0, 0, 0,
		14, 157, 1, 0, 0, 0, 16, 175, 1, 0, 0, 0, 18, 177, 1, 0, 0, 0, 20, 182,
		1, 0, 0, 0, 22, 190, 1, 0, 0, 0, 24, 198, 1, 0, 0, 0, 26, 206, 1, 0, 0,
		0, 28, 214, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 230, 1, 0, 0, 0, 34, 237,
		1, 0, 0, 0, 36, 239, 1, 0, 0, 0, 38, 241, 1, 0, 0, 0, 40, 243, 1, 0, 0,
		0, 42, 247, 1, 0, 0, 0, 44, 265, 1, 0, 0, 0, 46, 273, 1, 0, 0, 0, 48, 276,
		1, 0, 0, 0, 50, 281, 1, 0, 0, 0, 52, 292, 1, 0, 0, 0, 54, 297, 1, 0, 0,
		0, 56, 306, 1, 0, 0, 0, 58, 311, 1, 0, 0, 0, 60, 315, 1, 0, 0, 0, 62, 330,
		1, 0, 0, 0, 64, 335, 1, 0, 0, 0, 66, 340, 1, 0, 0, 0, 68, 345, 1, 0, 0,
		0, 70, 356, 1, 0, 0, 0, 72, 363, 1, 0, 0, 0, 74, 375, 1, 0, 0, 0, 76, 377,
		1, 0, 0, 0, 78, 385, 1, 0, 0, 0, 80, 388, 1, 0, 0, 0, 82, 398, 1, 0, 0,
		0, 84, 403, 1, 0, 0, 0, 86, 415, 1, 0, 0, 0, 88, 433, 1, 0, 0, 0, 90, 435,
		1, 0, 0, 0, 92, 444, 1, 0, 0, 0, 94, 446, 1, 0, 0, 0, 96, 448, 1, 0, 0,
		0, 98, 450, 1, 0, 0, 0, 100, 453, 1, 0, 0, 0, 102, 502, 1, 0, 0, 0, 104,
		504, 1, 0, 0, 0, 106, 524, 1, 0, 0, 0, 108, 529, 1, 0, 0, 0, 110, 111,
		5, 77, 0, 0, 111, 1, 1, 0, 0, 0, 112, 113, 7, 0, 0, 0, 113, 3, 1, 0, 0,
		0, 114, 119, 3, 32, 16, 0, 115, 116, 5, 66, 0, 0, 116, 118, 3, 32, 16,
		0, 117, 115, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 5, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 124, 5,
		64, 0, 0, 123, 125, 3, 4, 2, 0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 65, 0, 0, 127, 7, 1, 0, 0, 0, 128,
		137, 3, 2, 1, 0, 129, 137, 3, 0, 0, 0, 130, 137, 5, 35, 0, 0, 131, 132,
		5, 60, 0, 0, 132, 133, 3, 32, 16, 0, 133, 134, 5, 61, 0, 0, 134, 137, 1,
		0, 0, 0, 135, 137, 3, 6, 3, 0, 136, 128, 1, 0, 0, 0, 136, 129, 1, 0, 0,
		0, 136, 130, 1, 0, 0, 0, 136, 131, 1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137,
		9, 1, 0, 0, 0, 138, 143, 3, 32, 16, 0, 139, 140, 5, 66, 0, 0, 140, 142,
		3, 32, 16, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1,
		0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 11, 1, 0, 0, 0, 145, 143, 1, 0, 0,
		0, 146, 149, 3, 32, 16, 0, 147, 148, 5, 67, 0, 0, 148, 150, 3, 32, 16,
		0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 156, 1, 0, 0, 0, 151,
		152, 3, 32, 16, 0, 152, 153, 5, 66, 0, 0, 153, 154, 3, 32, 16, 0, 154,
		156, 1, 0, 0, 0, 155, 146, 1, 0, 0, 0, 155, 151, 1, 0, 0, 0, 156, 13, 1,
		0, 0, 0, 157, 169, 3, 8, 4, 0, 158, 159, 5, 62, 0, 0, 159, 160, 3, 12,
		6, 0, 160, 161, 5, 63, 0, 0, 161, 168, 1, 0, 0, 0, 162, 164, 5, 60, 0,
		0, 163, 165, 3, 10, 5, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 168, 5, 61, 0, 0, 167, 158, 1, 0, 0, 0, 167, 162,
		1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 15, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 173, 7, 1, 0, 0,
		173, 176, 3, 16, 8, 0, 174, 176, 3, 14, 7, 0, 175, 172, 1, 0, 0, 0, 175,
		174, 1, 0, 0, 0, 176, 17, 1, 0, 0, 0, 177, 180, 3, 16, 8, 0, 178, 179,
		5, 49, 0, 0, 179, 181, 3, 18, 9, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1,
		0, 0, 0, 181, 19, 1, 0, 0, 0, 182, 187, 3, 18, 9, 0, 183, 184, 7, 2, 0,
		0, 184, 186, 3, 18, 9, 0, 185, 183, 1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187,
		185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 21, 1, 0, 0, 0, 189, 187, 1,
		0, 0, 0, 190, 195, 3, 20, 10, 0, 191, 192, 7, 3, 0, 0, 192, 194, 3, 20,
		10, 0, 193, 191, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0,
		195, 196, 1, 0, 0, 0, 196, 23, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 203,
		3, 22, 11, 0, 199, 200, 7, 4, 0, 0, 200, 202, 3, 22, 11, 0, 201, 199, 1,
		0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0,
		0, 204, 25, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 211, 3, 24, 12, 0, 207,
		208, 7, 5, 0, 0, 208, 210, 3, 24, 12, 0, 209, 207, 1, 0, 0, 0, 210, 213,
		1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 27, 1, 0,
		0, 0, 213, 211, 1, 0, 0, 0, 214, 219, 3, 26, 13, 0, 215, 216, 5, 30, 0,
		0, 216, 218, 3, 26, 13, 0, 217, 215, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0,
		219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 29, 1, 0, 0, 0, 221, 219,
		1, 0, 0, 0, 222, 227, 3, 28, 14, 0, 223, 224, 5, 31, 0, 0, 224, 226, 3,
		28, 14, 0, 225, 223, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227, 225, 1, 0,
		0, 0, 227, 228, 1, 0, 0, 0, 228, 31, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0,
		230, 231, 3, 30, 15, 0, 231, 33, 1, 0, 0, 0, 232, 238, 3, 38, 19, 0, 233,
		235, 3, 36, 18, 0, 234, 236, 5, 41, 0, 0, 235, 234, 1, 0, 0, 0, 235, 236,
		1, 0, 0, 0, 236, 238, 1, 0, 0, 0, 237, 232, 1, 0, 0, 0, 237, 233, 1, 0,
		0, 0, 238, 35, 1, 0, 0, 0, 239, 240, 7, 6, 0, 0, 240, 37, 1, 0, 0, 0, 241,
		242, 7, 7, 0, 0, 242, 39, 1, 0, 0, 0, 243, 244, 3, 32, 16, 0, 244, 245,
		5, 67, 0, 0, 245, 246, 3, 32, 16, 0, 246, 41, 1, 0, 0, 0, 247, 259, 5,
		77, 0, 0, 248, 249, 5, 62, 0, 0, 249, 254, 3, 40, 20, 0, 250, 251, 5, 66,
		0, 0, 251, 253, 3, 40, 20, 0, 252, 250, 1, 0, 0, 0, 253, 256, 1, 0, 0,
		0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256,
		254, 1, 0, 0, 0, 257, 258, 5, 63, 0, 0, 258, 260, 1, 0, 0, 0, 259, 248,
		1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 262, 5, 57,
		0, 0, 262, 264, 3, 32, 16, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0,
		0, 264, 43, 1, 0, 0, 0, 265, 270, 3, 42, 21, 0, 266, 267, 5, 66, 0, 0,
		267, 269, 3, 42, 21, 0, 268, 266, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270,
		268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 45, 1, 0, 0, 0, 272, 270, 1,
		0, 0, 0, 273, 274, 3, 34, 17, 0, 274, 275, 3, 44, 22, 0, 275, 47, 1, 0,
		0, 0, 276, 277, 3, 34, 17, 0, 277, 279, 3, 44, 22, 0, 278, 280, 5, 68,
		0, 0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 49, 1, 0, 0, 0,
		281, 282, 3, 0, 0, 0, 282, 286, 5, 18, 0, 0, 283, 287, 3, 2, 1, 0, 284,
		287, 3, 16, 8, 0, 285, 287, 3, 6, 3, 0, 286, 283, 1, 0, 0, 0, 286, 284,
		1, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288, 290, 5, 68,
		0, 0, 289, 288, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 51, 1, 0, 0, 0,
		291, 293, 7, 8, 0, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 295, 3, 34, 17, 0, 295, 296, 3, 44, 22, 0, 296, 53,
		1, 0, 0, 0, 297, 302, 3, 52, 26, 0, 298, 299, 5, 66, 0, 0, 299, 301, 3,
		52, 26, 0, 300, 298, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0,
		0, 0, 302, 303, 1, 0, 0, 0, 303, 55, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0,
		305, 307, 8, 9, 0, 0, 306, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308,
		306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 57, 1, 0, 0, 0, 310, 312, 5,
		77, 0, 0, 311, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 311, 1, 0, 0,
		0, 313, 314, 1, 0, 0, 0, 314, 59, 1, 0, 0, 0, 315, 317, 5, 1, 0, 0, 316,
		318, 3, 34, 17, 0, 317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 319,
		1, 0, 0, 0, 319, 325, 3, 56, 28, 0, 320, 322, 5, 60, 0, 0, 321, 323, 3,
		54, 27, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 1, 0,
		0, 0, 324, 326, 5, 61, 0, 0, 325, 320, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0,
		326, 328, 1, 0, 0, 0, 327, 329, 5, 68, 0, 0, 328, 327, 1, 0, 0, 0, 328,
		329, 1, 0, 0, 0, 329, 61, 1, 0, 0, 0, 330, 331, 5, 4, 0, 0, 331, 333, 3,
		32, 16, 0, 332, 334, 5, 68, 0, 0, 333, 332, 1, 0, 0, 0, 333, 334, 1, 0,
		0, 0, 334, 63, 1, 0, 0, 0, 335, 336, 5, 5, 0, 0, 336, 338, 3, 32, 16, 0,
		337, 339, 5, 68, 0, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339,
		65, 1, 0, 0, 0, 340, 341, 3, 68, 34, 0, 341, 67, 1, 0, 0, 0, 342, 344,
		3, 102, 51, 0, 343, 342, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1,
		0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 69, 1, 0, 0, 0, 347, 345, 1, 0, 0,
		0, 348, 353, 3, 0, 0, 0, 349, 350, 5, 62, 0, 0, 350, 351, 3, 12, 6, 0,
		351, 352, 5, 63, 0, 0, 352, 354, 1, 0, 0, 0, 353, 349, 1, 0, 0, 0, 353,
		354, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 357, 5, 35, 0, 0, 356, 348,
		1, 0, 0, 0, 356, 355, 1, 0, 0, 0, 357, 71, 1, 0, 0, 0, 358, 359, 3, 70,
		35, 0, 359, 360, 5, 18, 0, 0, 360, 361, 3, 32, 16, 0, 361, 364, 1, 0, 0,
		0, 362, 364, 3, 32, 16, 0, 363, 358, 1, 0, 0, 0, 363, 362, 1, 0, 0, 0,
		364, 73, 1, 0, 0, 0, 365, 372, 3, 32, 16, 0, 366, 367, 5, 67, 0, 0, 367,
		370, 3, 32, 16, 0, 368, 369, 5, 67, 0, 0, 369, 371, 3, 32, 16, 0, 370,
		368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 373, 1, 0, 0, 0, 372, 366,
		1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 376, 1, 0, 0, 0, 374, 376, 5, 28,
		0, 0, 375, 365, 1, 0, 0, 0, 375, 374, 1, 0, 0, 0, 376, 75, 1, 0, 0, 0,
		377, 382, 3, 74, 37, 0, 378, 379, 5, 66, 0, 0, 379, 381, 3, 74, 37, 0,
		380, 378, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 77, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 386, 7,
		10, 0, 0, 386, 387, 3, 76, 38, 0, 387, 79, 1, 0, 0, 0, 388, 389, 5, 10,
		0, 0, 389, 390, 3, 32, 16, 0, 390, 391, 5, 11, 0, 0, 391, 394, 3, 68, 34,
		0, 392, 393, 5, 12, 0, 0, 393, 395, 3, 68, 34, 0, 394, 392, 1, 0, 0, 0,
		394, 395, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 5, 13, 0, 0, 397,
		81, 1, 0, 0, 0, 398, 399, 5, 15, 0, 0, 399, 400, 3, 32, 16, 0, 400, 401,
		5, 67, 0, 0, 401, 402, 3, 68, 34, 0, 402, 83, 1, 0, 0, 0, 403, 405, 5,
		14, 0, 0, 404, 406, 3, 82, 41, 0, 405, 404, 1, 0, 0, 0, 406, 407, 1, 0,
		0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 411, 1, 0, 0, 0,
		409, 410, 5, 12, 0, 0, 410, 412, 3, 68, 34, 0, 411, 409, 1, 0, 0, 0, 411,
		412, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 414, 5, 13, 0, 0, 414, 85,
		1, 0, 0, 0, 415, 416, 5, 8, 0, 0, 416, 417, 3, 32, 16, 0, 417, 87, 1, 0,
		0, 0, 418, 419, 5, 22, 0, 0, 419, 420, 5, 77, 0, 0, 420, 421, 5, 25, 0,
		0, 421, 422, 3, 32, 16, 0, 422, 423, 5, 26, 0, 0, 423, 426, 3, 32, 16,
		0, 424, 425, 5, 27, 0, 0, 425, 427, 3, 32, 16, 0, 426, 424, 1, 0, 0, 0,
		426, 427, 1, 0, 0, 0, 427, 434, 1, 0, 0, 0, 428, 429, 5, 23, 0, 0, 429,
		434, 3, 32, 16, 0, 430, 431, 3, 32, 16, 0, 431, 432, 5, 24, 0, 0, 432,
		434, 1, 0, 0, 0, 433, 418, 1, 0, 0, 0, 433, 428, 1, 0, 0, 0, 433, 430,
		1, 0, 0, 0, 434, 89, 1, 0, 0, 0, 435, 437, 5, 7, 0, 0, 436, 438, 3, 88,
		44, 0, 437, 436, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0,
		439, 442, 3, 68, 34, 0, 440, 443, 5, 9, 0, 0, 441, 443, 3, 86, 43, 0, 442,
		440, 1, 0, 0, 0, 442, 441, 1, 0, 0, 0, 443, 91, 1, 0, 0, 0, 444, 445, 5,
		19, 0, 0, 445, 93, 1, 0, 0, 0, 446, 447, 5, 20, 0, 0, 447, 95, 1, 0, 0,
		0, 448, 449, 5, 21, 0, 0, 449, 97, 1, 0, 0, 0, 450, 451, 5, 6, 0, 0, 451,
		452, 3, 32, 16, 0, 452, 99, 1, 0, 0, 0, 453, 459, 3, 0, 0, 0, 454, 456,
		5, 60, 0, 0, 455, 457, 3, 10, 5, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1,
		0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 460, 5, 61, 0, 0, 459, 454, 1, 0, 0,
		0, 459, 460, 1, 0, 0, 0, 460, 101, 1, 0, 0, 0, 461, 463, 3, 46, 23, 0,
		462, 464, 5, 68, 0, 0, 463, 462, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464,
		503, 1, 0, 0, 0, 465, 467, 3, 72, 36, 0, 466, 468, 5, 68, 0, 0, 467, 466,
		1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 503, 1, 0, 0, 0, 469, 471, 3, 78,
		39, 0, 470, 472, 5, 68, 0, 0, 471, 470, 1, 0, 0, 0, 471, 472, 1, 0, 0,
		0, 472, 503, 1, 0, 0, 0, 473, 475, 3, 80, 40, 0, 474, 476, 5, 68, 0, 0,
		475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 503, 1, 0, 0, 0, 477,
		479, 3, 84, 42, 0, 478, 480, 5, 68, 0, 0, 479, 478, 1, 0, 0, 0, 479, 480,
		1, 0, 0, 0, 480, 503, 1, 0, 0, 0, 481, 483, 3, 90, 45, 0, 482, 484, 5,
		68, 0, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 503, 1, 0, 0,
		0, 485, 487, 3, 92, 46, 0, 486, 488, 5, 68, 0, 0, 487, 486, 1, 0, 0, 0,
		487, 488, 1, 0, 0, 0, 488, 503, 1, 0, 0, 0, 489, 491, 3, 94, 47, 0, 490,
		492, 5, 68, 0, 0, 491, 490, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 503,
		1, 0, 0, 0, 493, 495, 3, 96, 48, 0, 494, 496, 5, 68, 0, 0, 495, 494, 1,
		0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 503, 1, 0, 0, 0, 497, 499, 3, 98, 49,
		0, 498, 500, 5, 68, 0, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500,
		503, 1, 0, 0, 0, 501, 503, 5, 68, 0, 0, 502, 461, 1, 0, 0, 0, 502, 465,
		1, 0, 0, 0, 502, 469, 1, 0, 0, 0, 502, 473, 1, 0, 0, 0, 502, 477, 1, 0,
		0, 0, 502, 481, 1, 0, 0, 0, 502, 485, 1, 0, 0, 0, 502, 489, 1, 0, 0, 0,
		502, 493, 1, 0, 0, 0, 502, 497, 1, 0, 0, 0, 502, 501, 1, 0, 0, 0, 503,
		103, 1, 0, 0, 0, 504, 510, 3, 60, 30, 0, 505, 509, 3, 62, 31, 0, 506, 509,
		3, 64, 32, 0, 507, 509, 3, 46, 23, 0, 508, 505, 1, 0, 0, 0, 508, 506, 1,
		0, 0, 0, 508, 507, 1, 0, 0, 0, 509, 512, 1, 0, 0, 0, 510, 508, 1, 0, 0,
		0, 510, 511, 1, 0, 0, 0, 511, 513, 1, 0, 0, 0, 512, 510, 1, 0, 0, 0, 513,
		514, 5, 2, 0, 0, 514, 515, 3, 66, 33, 0, 515, 517, 5, 3, 0, 0, 516, 518,
		3, 58, 29, 0, 517, 516, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 520, 1,
		0, 0, 0, 519, 521, 5, 68, 0, 0, 520, 519, 1, 0, 0, 0, 520, 521, 1, 0, 0,
		0, 521, 105, 1, 0, 0, 0, 522, 525, 3, 48, 24, 0, 523, 525, 3, 50, 25, 0,
		524, 522, 1, 0, 0, 0, 524, 523, 1, 0, 0, 0, 525, 107, 1, 0, 0, 0, 526,
		528, 3, 106, 53, 0, 527, 526, 1, 0, 0, 0, 528, 531, 1, 0, 0, 0, 529, 527,
		1, 0, 0, 0, 529, 530, 1, 0, 0, 0, 530, 533, 1, 0, 0, 0, 531, 529, 1, 0,
		0, 0, 532, 534, 3, 104, 52, 0, 533, 532, 1, 0, 0, 0, 534, 535, 1, 0, 0,
		0, 535, 533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537,
		538, 5, 0, 0, 1, 538, 109, 1, 0, 0, 0, 71, 119, 124, 136, 143, 149, 155,
		164, 167, 169, 175, 180, 187, 195, 203, 211, 219, 227, 235, 237, 254, 259,
		263, 270, 279, 286, 289, 292, 302, 308, 313, 317, 322, 325, 328, 333, 338,
		345, 353, 356, 363, 370, 372, 375, 382, 394, 407, 411, 426, 433, 437, 442,
		456, 459, 463, 467, 471, 475, 479, 483, 487, 491, 495, 499, 502, 508, 510,
		517, 520, 524, 529, 535,
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

// KumirParserInit initializes any static state used to implement KumirParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKumirParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KumirParserInit() {
	staticData := &KumirParserParserStaticData
	staticData.once.Do(kumirparserParserInit)
}

// NewKumirParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKumirParser(input antlr.TokenStream) *KumirParser {
	KumirParserInit()
	this := new(KumirParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KumirParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "KumirParser.g4"

	return this
}

// KumirParser tokens.
const (
	KumirParserEOF                = antlr.TokenEOF
	KumirParserALG_HEADER         = 1
	KumirParserALG_BEGIN          = 2
	KumirParserALG_END            = 3
	KumirParserPRE_CONDITION      = 4
	KumirParserPOST_CONDITION     = 5
	KumirParserASSERTION          = 6
	KumirParserLOOP               = 7
	KumirParserENDLOOP_COND       = 8
	KumirParserENDLOOP            = 9
	KumirParserIF                 = 10
	KumirParserTHEN               = 11
	KumirParserELSE               = 12
	KumirParserFI                 = 13
	KumirParserSWITCH             = 14
	KumirParserCASE               = 15
	KumirParserINPUT              = 16
	KumirParserOUTPUT             = 17
	KumirParserASSIGN             = 18
	KumirParserEXIT               = 19
	KumirParserPAUSE              = 20
	KumirParserSTOP               = 21
	KumirParserFOR                = 22
	KumirParserWHILE              = 23
	KumirParserTIMES              = 24
	KumirParserFROM               = 25
	KumirParserTO                 = 26
	KumirParserSTEP               = 27
	KumirParserNEWLINE_CONST      = 28
	KumirParserNOT                = 29
	KumirParserAND                = 30
	KumirParserOR                 = 31
	KumirParserOUT_PARAM          = 32
	KumirParserIN_PARAM           = 33
	KumirParserINOUT_PARAM        = 34
	KumirParserRETURN_VALUE       = 35
	KumirParserINTEGER_TYPE       = 36
	KumirParserREAL_TYPE          = 37
	KumirParserBOOLEAN_TYPE       = 38
	KumirParserCHAR_TYPE          = 39
	KumirParserSTRING_TYPE        = 40
	KumirParserTABLE_SUFFIX       = 41
	KumirParserINTEGER_ARRAY_TYPE = 42
	KumirParserREAL_ARRAY_TYPE    = 43
	KumirParserCHAR_ARRAY_TYPE    = 44
	KumirParserSTRING_ARRAY_TYPE  = 45
	KumirParserBOOLEAN_ARRAY_TYPE = 46
	KumirParserTRUE               = 47
	KumirParserFALSE              = 48
	KumirParserPOWER              = 49
	KumirParserGE                 = 50
	KumirParserLE                 = 51
	KumirParserNE                 = 52
	KumirParserPLUS               = 53
	KumirParserMINUS              = 54
	KumirParserMUL                = 55
	KumirParserDIV                = 56
	KumirParserEQ                 = 57
	KumirParserLT                 = 58
	KumirParserGT                 = 59
	KumirParserLPAREN             = 60
	KumirParserRPAREN             = 61
	KumirParserLBRACK             = 62
	KumirParserRBRACK             = 63
	KumirParserLBRACE             = 64
	KumirParserRBRACE             = 65
	KumirParserCOMMA              = 66
	KumirParserCOLON              = 67
	KumirParserSEMICOLON          = 68
	KumirParserATAT               = 69
	KumirParserAT                 = 70
	KumirParserDIV_OP             = 71
	KumirParserMOD_OP             = 72
	KumirParserCHAR_LITERAL       = 73
	KumirParserSTRING             = 74
	KumirParserREAL               = 75
	KumirParserINTEGER            = 76
	KumirParserID                 = 77
	KumirParserLINE_COMMENT       = 78
	KumirParserDOC_COMMENT        = 79
	KumirParserWS                 = 80
)

// KumirParser rules.
const (
	KumirParserRULE_qualifiedIdentifier      = 0
	KumirParserRULE_literal                  = 1
	KumirParserRULE_expressionList           = 2
	KumirParserRULE_arrayLiteral             = 3
	KumirParserRULE_primaryExpression        = 4
	KumirParserRULE_argumentList             = 5
	KumirParserRULE_indexList                = 6
	KumirParserRULE_postfixExpression        = 7
	KumirParserRULE_unaryExpression          = 8
	KumirParserRULE_powerExpression          = 9
	KumirParserRULE_multiplicativeExpression = 10
	KumirParserRULE_additiveExpression       = 11
	KumirParserRULE_relationalExpression     = 12
	KumirParserRULE_equalityExpression       = 13
	KumirParserRULE_logicalAndExpression     = 14
	KumirParserRULE_logicalOrExpression      = 15
	KumirParserRULE_expression               = 16
	KumirParserRULE_typeSpecifier            = 17
	KumirParserRULE_basicType                = 18
	KumirParserRULE_arrayType                = 19
	KumirParserRULE_arrayBounds              = 20
	KumirParserRULE_variableDeclarationItem  = 21
	KumirParserRULE_variableList             = 22
	KumirParserRULE_variableDeclaration      = 23
	KumirParserRULE_globalDeclaration        = 24
	KumirParserRULE_globalAssignment         = 25
	KumirParserRULE_parameterDeclaration     = 26
	KumirParserRULE_parameterList            = 27
	KumirParserRULE_algorithmNameTokens      = 28
	KumirParserRULE_algorithmName            = 29
	KumirParserRULE_algorithmHeader          = 30
	KumirParserRULE_preCondition             = 31
	KumirParserRULE_postCondition            = 32
	KumirParserRULE_algorithmBody            = 33
	KumirParserRULE_statementSequence        = 34
	KumirParserRULE_lvalue                   = 35
	KumirParserRULE_assignmentStatement      = 36
	KumirParserRULE_ioArgument               = 37
	KumirParserRULE_ioArgumentList           = 38
	KumirParserRULE_ioStatement              = 39
	KumirParserRULE_ifStatement              = 40
	KumirParserRULE_caseBlock                = 41
	KumirParserRULE_switchStatement          = 42
	KumirParserRULE_endLoopCondition         = 43
	KumirParserRULE_loopSpecifier            = 44
	KumirParserRULE_loopStatement            = 45
	KumirParserRULE_exitStatement            = 46
	KumirParserRULE_pauseStatement           = 47
	KumirParserRULE_stopStatement            = 48
	KumirParserRULE_assertionStatement       = 49
	KumirParserRULE_procedureCallStatement   = 50
	KumirParserRULE_statement                = 51
	KumirParserRULE_algorithmDefinition      = 52
	KumirParserRULE_programItem              = 53
	KumirParserRULE_program                  = 54
)

// IQualifiedIdentifierContext is an interface to support dynamic dispatch.
type IQualifiedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsQualifiedIdentifierContext differentiates from other interfaces.
	IsQualifiedIdentifierContext()
}

type QualifiedIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentifierContext() *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_qualifiedIdentifier
	return p
}

func InitEmptyQualifiedIdentifierContext(p *QualifiedIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_qualifiedIdentifier
}

func (*QualifiedIdentifierContext) IsQualifiedIdentifierContext() {}

func NewQualifiedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_qualifiedIdentifier

	return p
}

func (s *QualifiedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(KumirParserID, 0)
}

func (s *QualifiedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterQualifiedIdentifier(s)
	}
}

func (s *QualifiedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitQualifiedIdentifier(s)
	}
}

func (s *QualifiedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitQualifiedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) QualifiedIdentifier() (localctx IQualifiedIdentifierContext) {
	localctx = NewQualifiedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KumirParserRULE_qualifiedIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(KumirParserID)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	REAL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	CHAR_LITERAL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	NEWLINE_CONST() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(KumirParserINTEGER, 0)
}

func (s *LiteralContext) REAL() antlr.TerminalNode {
	return s.GetToken(KumirParserREAL, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(KumirParserSTRING, 0)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(KumirParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(KumirParserTRUE, 0)
}

func (s *LiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(KumirParserFALSE, 0)
}

func (s *LiteralContext) NEWLINE_CONST() antlr.TerminalNode {
	return s.GetToken(KumirParserNEWLINE_CONST, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KumirParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-28)) & ^0x3f) == 0 && ((int64(1)<<(_la-28))&527765582905345) != 0) {
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

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KumirParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Expression()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserCOMMA {
		{
			p.SetState(115)
			p.Match(KumirParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Expression()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	ExpressionList() IExpressionListContext

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(KumirParserLBRACE, 0)
}

func (s *ArrayLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(KumirParserRBRACE, 0)
}

func (s *ArrayLiteralContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KumirParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(KumirParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-28)) & ^0x3f) == 0 && ((int64(1)<<(_la-28))&1090788651434115) != 0 {
		{
			p.SetState(123)
			p.ExpressionList()
		}

	}
	{
		p.SetState(126)
		p.Match(KumirParserRBRACE)
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	QualifiedIdentifier() IQualifiedIdentifierContext
	RETURN_VALUE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryExpressionContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *PrimaryExpressionContext) RETURN_VALUE() antlr.TerminalNode {
	return s.GetToken(KumirParserRETURN_VALUE, 0)
}

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserLPAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserRPAREN, 0)
}

func (s *PrimaryExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KumirParserRULE_primaryExpression)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserNEWLINE_CONST, KumirParserTRUE, KumirParserFALSE, KumirParserCHAR_LITERAL, KumirParserSTRING, KumirParserREAL, KumirParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Literal()
		}

	case KumirParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.QualifiedIdentifier()
		}

	case KumirParserRETURN_VALUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(KumirParserRETURN_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KumirParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.Match(KumirParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Expression()
		}
		{
			p.SetState(133)
			p.Match(KumirParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KumirParserLBRACE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.ArrayLiteral()
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

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KumirParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Expression()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserCOMMA {
		{
			p.SetState(139)
			p.Match(KumirParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Expression()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IIndexListContext is an interface to support dynamic dispatch.
type IIndexListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsIndexListContext differentiates from other interfaces.
	IsIndexListContext()
}

type IndexListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexListContext() *IndexListContext {
	var p = new(IndexListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_indexList
	return p
}

func InitEmptyIndexListContext(p *IndexListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_indexList
}

func (*IndexListContext) IsIndexListContext() {}

func NewIndexListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexListContext {
	var p = new(IndexListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_indexList

	return p
}

func (s *IndexListContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexListContext) COLON() antlr.TerminalNode {
	return s.GetToken(KumirParserCOLON, 0)
}

func (s *IndexListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, 0)
}

func (s *IndexListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterIndexList(s)
	}
}

func (s *IndexListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitIndexList(s)
	}
}

func (s *IndexListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitIndexList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) IndexList() (localctx IIndexListContext) {
	localctx = NewIndexListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KumirParserRULE_indexList)
	var _la int

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Expression()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == KumirParserCOLON {
			{
				p.SetState(147)
				p.Match(KumirParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(148)
				p.Expression()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Expression()
		}
		{
			p.SetState(152)
			p.Match(KumirParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
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

// IPostfixExpressionContext is an interface to support dynamic dispatch.
type IPostfixExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllIndexList() []IIndexListContext
	IndexList(i int) IIndexListContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllArgumentList() []IArgumentListContext
	ArgumentList(i int) IArgumentListContext

	// IsPostfixExpressionContext differentiates from other interfaces.
	IsPostfixExpressionContext()
}

type PostfixExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixExpressionContext() *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_postfixExpression
	return p
}

func InitEmptyPostfixExpressionContext(p *PostfixExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_postfixExpression
}

func (*PostfixExpressionContext) IsPostfixExpressionContext() {}

func NewPostfixExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixExpressionContext {
	var p = new(PostfixExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_postfixExpression

	return p
}

func (s *PostfixExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PostfixExpressionContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(KumirParserLBRACK)
}

func (s *PostfixExpressionContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserLBRACK, i)
}

func (s *PostfixExpressionContext) AllIndexList() []IIndexListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexListContext); ok {
			len++
		}
	}

	tst := make([]IIndexListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexListContext); ok {
			tst[i] = t.(IIndexListContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) IndexList(i int) IIndexListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexListContext)
}

func (s *PostfixExpressionContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(KumirParserRBRACK)
}

func (s *PostfixExpressionContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserRBRACK, i)
}

func (s *PostfixExpressionContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(KumirParserLPAREN)
}

func (s *PostfixExpressionContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserLPAREN, i)
}

func (s *PostfixExpressionContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(KumirParserRPAREN)
}

func (s *PostfixExpressionContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserRPAREN, i)
}

func (s *PostfixExpressionContext) AllArgumentList() []IArgumentListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentListContext); ok {
			len++
		}
	}

	tst := make([]IArgumentListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentListContext); ok {
			tst[i] = t.(IArgumentListContext)
			i++
		}
	}

	return tst
}

func (s *PostfixExpressionContext) ArgumentList(i int) IArgumentListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *PostfixExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPostfixExpression(s)
	}
}

func (s *PostfixExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPostfixExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PostfixExpression() (localctx IPostfixExpressionContext) {
	localctx = NewPostfixExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KumirParserRULE_postfixExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.PrimaryExpression()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case KumirParserLBRACK:
				{
					p.SetState(158)
					p.Match(KumirParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(159)
					p.IndexList()
				}
				{
					p.SetState(160)
					p.Match(KumirParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case KumirParserLPAREN:
				{
					p.SetState(162)
					p.Match(KumirParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(164)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64((_la-28)) & ^0x3f) == 0 && ((int64(1)<<(_la-28))&1090788651434115) != 0 {
					{
						p.SetState(163)
						p.ArgumentList()
					}

				}
				{
					p.SetState(166)
					p.Match(KumirParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
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

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	PostfixExpression() IPostfixExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(KumirParserPLUS, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(KumirParserMINUS, 0)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(KumirParserNOT, 0)
}

func (s *UnaryExpressionContext) PostfixExpression() IPostfixExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KumirParserRULE_unaryExpression)
	var _la int

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserNOT, KumirParserPLUS, KumirParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&27021598301093888) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(173)
			p.UnaryExpression()
		}

	case KumirParserNEWLINE_CONST, KumirParserRETURN_VALUE, KumirParserTRUE, KumirParserFALSE, KumirParserLPAREN, KumirParserLBRACE, KumirParserCHAR_LITERAL, KumirParserSTRING, KumirParserREAL, KumirParserINTEGER, KumirParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.PostfixExpression()
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

// IPowerExpressionContext is an interface to support dynamic dispatch.
type IPowerExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	POWER() antlr.TerminalNode
	PowerExpression() IPowerExpressionContext

	// IsPowerExpressionContext differentiates from other interfaces.
	IsPowerExpressionContext()
}

type PowerExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerExpressionContext() *PowerExpressionContext {
	var p = new(PowerExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_powerExpression
	return p
}

func InitEmptyPowerExpressionContext(p *PowerExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_powerExpression
}

func (*PowerExpressionContext) IsPowerExpressionContext() {}

func NewPowerExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_powerExpression

	return p
}

func (s *PowerExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *PowerExpressionContext) POWER() antlr.TerminalNode {
	return s.GetToken(KumirParserPOWER, 0)
}

func (s *PowerExpressionContext) PowerExpression() IPowerExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExpressionContext)
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPowerExpression(s)
	}
}

func (s *PowerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPowerExpression(s)
	}
}

func (s *PowerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPowerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PowerExpression() (localctx IPowerExpressionContext) {
	localctx = NewPowerExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KumirParserRULE_powerExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.UnaryExpression()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserPOWER {
		{
			p.SetState(178)
			p.Match(KumirParserPOWER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.PowerExpression()
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

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllPowerExpression() []IPowerExpressionContext
	PowerExpression(i int) IPowerExpressionContext
	AllMUL() []antlr.TerminalNode
	MUL(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode
	AllDIV_OP() []antlr.TerminalNode
	DIV_OP(i int) antlr.TerminalNode
	AllMOD_OP() []antlr.TerminalNode
	MOD_OP(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeExpressionContext) AllPowerExpression() []IPowerExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPowerExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPowerExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPowerExpressionContext); ok {
			tst[i] = t.(IPowerExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) PowerExpression(i int) IPowerExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(KumirParserMUL)
}

func (s *MultiplicativeExpressionContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserMUL, i)
}

func (s *MultiplicativeExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(KumirParserDIV)
}

func (s *MultiplicativeExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserDIV, i)
}

func (s *MultiplicativeExpressionContext) AllDIV_OP() []antlr.TerminalNode {
	return s.GetTokens(KumirParserDIV_OP)
}

func (s *MultiplicativeExpressionContext) DIV_OP(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserDIV_OP, i)
}

func (s *MultiplicativeExpressionContext) AllMOD_OP() []antlr.TerminalNode {
	return s.GetTokens(KumirParserMOD_OP)
}

func (s *MultiplicativeExpressionContext) MOD_OP(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserMOD_OP, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KumirParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.PowerExpression()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&196611) != 0 {
		{
			p.SetState(183)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MultiplicativeExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-55)) & ^0x3f) == 0 && ((int64(1)<<(_la-55))&196611) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MultiplicativeExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(184)
			p.PowerExpression()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(KumirParserPLUS)
}

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserPLUS, i)
}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(KumirParserMINUS)
}

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserMINUS, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KumirParserRULE_additiveExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.MultiplicativeExpression()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AdditiveExpressionContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == KumirParserPLUS || _la == KumirParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AdditiveExpressionContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(192)
				p.MultiplicativeExpression()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
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

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *RelationalExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(KumirParserLT)
}

func (s *RelationalExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserLT, i)
}

func (s *RelationalExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(KumirParserGT)
}

func (s *RelationalExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserGT, i)
}

func (s *RelationalExpressionContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(KumirParserLE)
}

func (s *RelationalExpressionContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserLE, i)
}

func (s *RelationalExpressionContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(KumirParserGE)
}

func (s *RelationalExpressionContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserGE, i)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KumirParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.AdditiveExpression()
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&868068828175663104) != 0 {
		{
			p.SetState(199)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*RelationalExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&868068828175663104) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*RelationalExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(200)
			p.AdditiveExpression()
		}

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllRelationalExpression() []IRelationalExpressionContext
	RelationalExpression(i int) IRelationalExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNE() []antlr.TerminalNode
	NE(i int) antlr.TerminalNode

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_equalityExpression
	return p
}

func InitEmptyEqualityExpressionContext(p *EqualityExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_equalityExpression
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExpressionContext) AllRelationalExpression() []IRelationalExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExpressionContext); ok {
			tst[i] = t.(IRelationalExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) RelationalExpression(i int) IRelationalExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(KumirParserEQ)
}

func (s *EqualityExpressionContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserEQ, i)
}

func (s *EqualityExpressionContext) AllNE() []antlr.TerminalNode {
	return s.GetTokens(KumirParserNE)
}

func (s *EqualityExpressionContext) NE(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserNE, i)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KumirParserRULE_equalityExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.RelationalExpression()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserNE || _la == KumirParserEQ {
		{
			p.SetState(207)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*EqualityExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == KumirParserNE || _la == KumirParserEQ) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*EqualityExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(208)
			p.RelationalExpression()
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEqualityExpression() []IEqualityExpressionContext
	EqualityExpression(i int) IEqualityExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_logicalAndExpression
	return p
}

func InitEmptyLogicalAndExpressionContext(p *LogicalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_logicalAndExpression
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) AllEqualityExpression() []IEqualityExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			len++
		}
	}

	tst := make([]IEqualityExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityExpressionContext); ok {
			tst[i] = t.(IEqualityExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) EqualityExpression(i int) IEqualityExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *LogicalAndExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(KumirParserAND)
}

func (s *LogicalAndExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserAND, i)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLogicalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KumirParserRULE_logicalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.EqualityExpression()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserAND {
		{
			p.SetState(215)
			p.Match(KumirParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.EqualityExpression()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpression() []ILogicalAndExpressionContext
	LogicalAndExpression(i int) ILogicalAndExpressionContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_logicalOrExpression
	return p
}

func InitEmptyLogicalOrExpressionContext(p *LogicalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_logicalOrExpression
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) AllLogicalAndExpression() []ILogicalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExpressionContext); ok {
			tst[i] = t.(ILogicalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) LogicalAndExpression(i int) ILogicalAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(KumirParserOR)
}

func (s *LogicalOrExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserOR, i)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLogicalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KumirParserRULE_logicalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.LogicalAndExpression()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserOR {
		{
			p.SetState(223)
			p.Match(KumirParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.LogicalAndExpression()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpression() ILogicalOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KumirParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.LogicalOrExpression()
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

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArrayType() IArrayTypeContext
	BasicType() IBasicTypeContext
	TABLE_SUFFIX() antlr.TerminalNode

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_typeSpecifier
	return p
}

func InitEmptyTypeSpecifierContext(p *TypeSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_typeSpecifier
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *TypeSpecifierContext) BasicType() IBasicTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasicTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasicTypeContext)
}

func (s *TypeSpecifierContext) TABLE_SUFFIX() antlr.TerminalNode {
	return s.GetToken(KumirParserTABLE_SUFFIX, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitTypeSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KumirParserRULE_typeSpecifier)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserINTEGER_ARRAY_TYPE, KumirParserREAL_ARRAY_TYPE, KumirParserCHAR_ARRAY_TYPE, KumirParserSTRING_ARRAY_TYPE, KumirParserBOOLEAN_ARRAY_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.ArrayType()
		}

	case KumirParserINTEGER_TYPE, KumirParserREAL_TYPE, KumirParserBOOLEAN_TYPE, KumirParserCHAR_TYPE, KumirParserSTRING_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.BasicType()
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(234)
				p.Match(KumirParserTABLE_SUFFIX)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

// IBasicTypeContext is an interface to support dynamic dispatch.
type IBasicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_TYPE() antlr.TerminalNode
	REAL_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	CHAR_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode

	// IsBasicTypeContext differentiates from other interfaces.
	IsBasicTypeContext()
}

type BasicTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicTypeContext() *BasicTypeContext {
	var p = new(BasicTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_basicType
	return p
}

func InitEmptyBasicTypeContext(p *BasicTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_basicType
}

func (*BasicTypeContext) IsBasicTypeContext() {}

func NewBasicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicTypeContext {
	var p = new(BasicTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_basicType

	return p
}

func (s *BasicTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicTypeContext) INTEGER_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserINTEGER_TYPE, 0)
}

func (s *BasicTypeContext) REAL_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserREAL_TYPE, 0)
}

func (s *BasicTypeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserBOOLEAN_TYPE, 0)
}

func (s *BasicTypeContext) CHAR_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserCHAR_TYPE, 0)
}

func (s *BasicTypeContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserSTRING_TYPE, 0)
}

func (s *BasicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterBasicType(s)
	}
}

func (s *BasicTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitBasicType(s)
	}
}

func (s *BasicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitBasicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) BasicType() (localctx IBasicTypeContext) {
	localctx = NewBasicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KumirParserRULE_basicType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2130303778816) != 0) {
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

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_ARRAY_TYPE() antlr.TerminalNode
	REAL_ARRAY_TYPE() antlr.TerminalNode
	BOOLEAN_ARRAY_TYPE() antlr.TerminalNode
	CHAR_ARRAY_TYPE() antlr.TerminalNode
	STRING_ARRAY_TYPE() antlr.TerminalNode

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) INTEGER_ARRAY_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserINTEGER_ARRAY_TYPE, 0)
}

func (s *ArrayTypeContext) REAL_ARRAY_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserREAL_ARRAY_TYPE, 0)
}

func (s *ArrayTypeContext) BOOLEAN_ARRAY_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserBOOLEAN_ARRAY_TYPE, 0)
}

func (s *ArrayTypeContext) CHAR_ARRAY_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserCHAR_ARRAY_TYPE, 0)
}

func (s *ArrayTypeContext) STRING_ARRAY_TYPE() antlr.TerminalNode {
	return s.GetToken(KumirParserSTRING_ARRAY_TYPE, 0)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KumirParserRULE_arrayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&136339441844224) != 0) {
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

// IArrayBoundsContext is an interface to support dynamic dispatch.
type IArrayBoundsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsArrayBoundsContext differentiates from other interfaces.
	IsArrayBoundsContext()
}

type ArrayBoundsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayBoundsContext() *ArrayBoundsContext {
	var p = new(ArrayBoundsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayBounds
	return p
}

func InitEmptyArrayBoundsContext(p *ArrayBoundsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_arrayBounds
}

func (*ArrayBoundsContext) IsArrayBoundsContext() {}

func NewArrayBoundsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayBoundsContext {
	var p = new(ArrayBoundsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_arrayBounds

	return p
}

func (s *ArrayBoundsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayBoundsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayBoundsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayBoundsContext) COLON() antlr.TerminalNode {
	return s.GetToken(KumirParserCOLON, 0)
}

func (s *ArrayBoundsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayBoundsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayBoundsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterArrayBounds(s)
	}
}

func (s *ArrayBoundsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitArrayBounds(s)
	}
}

func (s *ArrayBoundsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitArrayBounds(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ArrayBounds() (localctx IArrayBoundsContext) {
	localctx = NewArrayBoundsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KumirParserRULE_arrayBounds)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Expression()
	}
	{
		p.SetState(244)
		p.Match(KumirParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Expression()
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

// IVariableDeclarationItemContext is an interface to support dynamic dispatch.
type IVariableDeclarationItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllArrayBounds() []IArrayBoundsContext
	ArrayBounds(i int) IArrayBoundsContext
	RBRACK() antlr.TerminalNode
	EQ() antlr.TerminalNode
	Expression() IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVariableDeclarationItemContext differentiates from other interfaces.
	IsVariableDeclarationItemContext()
}

type VariableDeclarationItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationItemContext() *VariableDeclarationItemContext {
	var p = new(VariableDeclarationItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableDeclarationItem
	return p
}

func InitEmptyVariableDeclarationItemContext(p *VariableDeclarationItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableDeclarationItem
}

func (*VariableDeclarationItemContext) IsVariableDeclarationItemContext() {}

func NewVariableDeclarationItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationItemContext {
	var p = new(VariableDeclarationItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_variableDeclarationItem

	return p
}

func (s *VariableDeclarationItemContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationItemContext) ID() antlr.TerminalNode {
	return s.GetToken(KumirParserID, 0)
}

func (s *VariableDeclarationItemContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(KumirParserLBRACK, 0)
}

func (s *VariableDeclarationItemContext) AllArrayBounds() []IArrayBoundsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayBoundsContext); ok {
			len++
		}
	}

	tst := make([]IArrayBoundsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayBoundsContext); ok {
			tst[i] = t.(IArrayBoundsContext)
			i++
		}
	}

	return tst
}

func (s *VariableDeclarationItemContext) ArrayBounds(i int) IArrayBoundsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayBoundsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayBoundsContext)
}

func (s *VariableDeclarationItemContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(KumirParserRBRACK, 0)
}

func (s *VariableDeclarationItemContext) EQ() antlr.TerminalNode {
	return s.GetToken(KumirParserEQ, 0)
}

func (s *VariableDeclarationItemContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationItemContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *VariableDeclarationItemContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *VariableDeclarationItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterVariableDeclarationItem(s)
	}
}

func (s *VariableDeclarationItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitVariableDeclarationItem(s)
	}
}

func (s *VariableDeclarationItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitVariableDeclarationItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) VariableDeclarationItem() (localctx IVariableDeclarationItemContext) {
	localctx = NewVariableDeclarationItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KumirParserRULE_variableDeclarationItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(KumirParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserLBRACK {
		{
			p.SetState(248)
			p.Match(KumirParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.ArrayBounds()
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == KumirParserCOMMA {
			{
				p.SetState(250)
				p.Match(KumirParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(251)
				p.ArrayBounds()
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(257)
			p.Match(KumirParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserEQ {
		{
			p.SetState(261)
			p.Match(KumirParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Expression()
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

// IVariableListContext is an interface to support dynamic dispatch.
type IVariableListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDeclarationItem() []IVariableDeclarationItemContext
	VariableDeclarationItem(i int) IVariableDeclarationItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVariableListContext differentiates from other interfaces.
	IsVariableListContext()
}

type VariableListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableListContext() *VariableListContext {
	var p = new(VariableListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableList
	return p
}

func InitEmptyVariableListContext(p *VariableListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableList
}

func (*VariableListContext) IsVariableListContext() {}

func NewVariableListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableListContext {
	var p = new(VariableListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_variableList

	return p
}

func (s *VariableListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableListContext) AllVariableDeclarationItem() []IVariableDeclarationItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationItemContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationItemContext); ok {
			tst[i] = t.(IVariableDeclarationItemContext)
			i++
		}
	}

	return tst
}

func (s *VariableListContext) VariableDeclarationItem(i int) IVariableDeclarationItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationItemContext)
}

func (s *VariableListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *VariableListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *VariableListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterVariableList(s)
	}
}

func (s *VariableListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitVariableList(s)
	}
}

func (s *VariableListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitVariableList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) VariableList() (localctx IVariableListContext) {
	localctx = NewVariableListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, KumirParserRULE_variableList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.VariableDeclarationItem()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(266)
				p.Match(KumirParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(267)
				p.VariableDeclarationItem()
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
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

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSpecifier() ITypeSpecifierContext
	VariableList() IVariableListContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *VariableDeclarationContext) VariableList() IVariableListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableListContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, KumirParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.TypeSpecifier()
	}
	{
		p.SetState(274)
		p.VariableList()
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

// IGlobalDeclarationContext is an interface to support dynamic dispatch.
type IGlobalDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSpecifier() ITypeSpecifierContext
	VariableList() IVariableListContext
	SEMICOLON() antlr.TerminalNode

	// IsGlobalDeclarationContext differentiates from other interfaces.
	IsGlobalDeclarationContext()
}

type GlobalDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalDeclarationContext() *GlobalDeclarationContext {
	var p = new(GlobalDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_globalDeclaration
	return p
}

func InitEmptyGlobalDeclarationContext(p *GlobalDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_globalDeclaration
}

func (*GlobalDeclarationContext) IsGlobalDeclarationContext() {}

func NewGlobalDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalDeclarationContext {
	var p = new(GlobalDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_globalDeclaration

	return p
}

func (s *GlobalDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *GlobalDeclarationContext) VariableList() IVariableListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableListContext)
}

func (s *GlobalDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *GlobalDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterGlobalDeclaration(s)
	}
}

func (s *GlobalDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitGlobalDeclaration(s)
	}
}

func (s *GlobalDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitGlobalDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) GlobalDeclaration() (localctx IGlobalDeclarationContext) {
	localctx = NewGlobalDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, KumirParserRULE_globalDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.TypeSpecifier()
	}
	{
		p.SetState(277)
		p.VariableList()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(278)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IGlobalAssignmentContext is an interface to support dynamic dispatch.
type IGlobalAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QualifiedIdentifier() IQualifiedIdentifierContext
	ASSIGN() antlr.TerminalNode
	Literal() ILiteralContext
	UnaryExpression() IUnaryExpressionContext
	ArrayLiteral() IArrayLiteralContext
	SEMICOLON() antlr.TerminalNode

	// IsGlobalAssignmentContext differentiates from other interfaces.
	IsGlobalAssignmentContext()
}

type GlobalAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalAssignmentContext() *GlobalAssignmentContext {
	var p = new(GlobalAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_globalAssignment
	return p
}

func InitEmptyGlobalAssignmentContext(p *GlobalAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_globalAssignment
}

func (*GlobalAssignmentContext) IsGlobalAssignmentContext() {}

func NewGlobalAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalAssignmentContext {
	var p = new(GlobalAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_globalAssignment

	return p
}

func (s *GlobalAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalAssignmentContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *GlobalAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KumirParserASSIGN, 0)
}

func (s *GlobalAssignmentContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *GlobalAssignmentContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *GlobalAssignmentContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *GlobalAssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *GlobalAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterGlobalAssignment(s)
	}
}

func (s *GlobalAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitGlobalAssignment(s)
	}
}

func (s *GlobalAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitGlobalAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) GlobalAssignment() (localctx IGlobalAssignmentContext) {
	localctx = NewGlobalAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, KumirParserRULE_globalAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.QualifiedIdentifier()
	}
	{
		p.SetState(282)
		p.Match(KumirParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(283)
			p.Literal()
		}

	case 2:
		{
			p.SetState(284)
			p.UnaryExpression()
		}

	case 3:
		{
			p.SetState(285)
			p.ArrayLiteral()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(288)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IParameterDeclarationContext is an interface to support dynamic dispatch.
type IParameterDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeSpecifier() ITypeSpecifierContext
	VariableList() IVariableListContext
	IN_PARAM() antlr.TerminalNode
	OUT_PARAM() antlr.TerminalNode
	INOUT_PARAM() antlr.TerminalNode

	// IsParameterDeclarationContext differentiates from other interfaces.
	IsParameterDeclarationContext()
}

type ParameterDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclarationContext() *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_parameterDeclaration
	return p
}

func InitEmptyParameterDeclarationContext(p *ParameterDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_parameterDeclaration
}

func (*ParameterDeclarationContext) IsParameterDeclarationContext() {}

func NewParameterDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclarationContext {
	var p = new(ParameterDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_parameterDeclaration

	return p
}

func (s *ParameterDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ParameterDeclarationContext) VariableList() IVariableListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableListContext)
}

func (s *ParameterDeclarationContext) IN_PARAM() antlr.TerminalNode {
	return s.GetToken(KumirParserIN_PARAM, 0)
}

func (s *ParameterDeclarationContext) OUT_PARAM() antlr.TerminalNode {
	return s.GetToken(KumirParserOUT_PARAM, 0)
}

func (s *ParameterDeclarationContext) INOUT_PARAM() antlr.TerminalNode {
	return s.GetToken(KumirParserINOUT_PARAM, 0)
}

func (s *ParameterDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterParameterDeclaration(s)
	}
}

func (s *ParameterDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitParameterDeclaration(s)
	}
}

func (s *ParameterDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitParameterDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ParameterDeclaration() (localctx IParameterDeclarationContext) {
	localctx = NewParameterDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, KumirParserRULE_parameterDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0 {
		{
			p.SetState(291)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30064771072) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(294)
		p.TypeSpecifier()
	}
	{
		p.SetState(295)
		p.VariableList()
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

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameterDeclaration() []IParameterDeclarationContext
	ParameterDeclaration(i int) IParameterDeclarationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameterDeclaration() []IParameterDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclarationContext); ok {
			tst[i] = t.(IParameterDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) ParameterDeclaration(i int) IParameterDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclarationContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, KumirParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.ParameterDeclaration()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserCOMMA {
		{
			p.SetState(298)
			p.Match(KumirParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.ParameterDeclaration()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAlgorithmNameTokensContext is an interface to support dynamic dispatch.
type IAlgorithmNameTokensContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllALG_BEGIN() []antlr.TerminalNode
	ALG_BEGIN(i int) antlr.TerminalNode
	AllPRE_CONDITION() []antlr.TerminalNode
	PRE_CONDITION(i int) antlr.TerminalNode
	AllPOST_CONDITION() []antlr.TerminalNode
	POST_CONDITION(i int) antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	AllEOF() []antlr.TerminalNode
	EOF(i int) antlr.TerminalNode

	// IsAlgorithmNameTokensContext differentiates from other interfaces.
	IsAlgorithmNameTokensContext()
}

type AlgorithmNameTokensContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmNameTokensContext() *AlgorithmNameTokensContext {
	var p = new(AlgorithmNameTokensContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmNameTokens
	return p
}

func InitEmptyAlgorithmNameTokensContext(p *AlgorithmNameTokensContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmNameTokens
}

func (*AlgorithmNameTokensContext) IsAlgorithmNameTokensContext() {}

func NewAlgorithmNameTokensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmNameTokensContext {
	var p = new(AlgorithmNameTokensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_algorithmNameTokens

	return p
}

func (s *AlgorithmNameTokensContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmNameTokensContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(KumirParserLPAREN)
}

func (s *AlgorithmNameTokensContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserLPAREN, i)
}

func (s *AlgorithmNameTokensContext) AllALG_BEGIN() []antlr.TerminalNode {
	return s.GetTokens(KumirParserALG_BEGIN)
}

func (s *AlgorithmNameTokensContext) ALG_BEGIN(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserALG_BEGIN, i)
}

func (s *AlgorithmNameTokensContext) AllPRE_CONDITION() []antlr.TerminalNode {
	return s.GetTokens(KumirParserPRE_CONDITION)
}

func (s *AlgorithmNameTokensContext) PRE_CONDITION(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserPRE_CONDITION, i)
}

func (s *AlgorithmNameTokensContext) AllPOST_CONDITION() []antlr.TerminalNode {
	return s.GetTokens(KumirParserPOST_CONDITION)
}

func (s *AlgorithmNameTokensContext) POST_CONDITION(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserPOST_CONDITION, i)
}

func (s *AlgorithmNameTokensContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(KumirParserSEMICOLON)
}

func (s *AlgorithmNameTokensContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, i)
}

func (s *AlgorithmNameTokensContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(KumirParserEOF)
}

func (s *AlgorithmNameTokensContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserEOF, i)
}

func (s *AlgorithmNameTokensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmNameTokensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmNameTokensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAlgorithmNameTokens(s)
	}
}

func (s *AlgorithmNameTokensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAlgorithmNameTokens(s)
	}
}

func (s *AlgorithmNameTokensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAlgorithmNameTokens(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AlgorithmNameTokens() (localctx IAlgorithmNameTokensContext) {
	localctx = NewAlgorithmNameTokensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, KumirParserRULE_algorithmNameTokens)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(305)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || ((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&2305843009213694057) != 0) || _la == KumirParserSEMICOLON {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
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

// IAlgorithmNameContext is an interface to support dynamic dispatch.
type IAlgorithmNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsAlgorithmNameContext differentiates from other interfaces.
	IsAlgorithmNameContext()
}

type AlgorithmNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmNameContext() *AlgorithmNameContext {
	var p = new(AlgorithmNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmName
	return p
}

func InitEmptyAlgorithmNameContext(p *AlgorithmNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmName
}

func (*AlgorithmNameContext) IsAlgorithmNameContext() {}

func NewAlgorithmNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmNameContext {
	var p = new(AlgorithmNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_algorithmName

	return p
}

func (s *AlgorithmNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(KumirParserID)
}

func (s *AlgorithmNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserID, i)
}

func (s *AlgorithmNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAlgorithmName(s)
	}
}

func (s *AlgorithmNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAlgorithmName(s)
	}
}

func (s *AlgorithmNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAlgorithmName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AlgorithmName() (localctx IAlgorithmNameContext) {
	localctx = NewAlgorithmNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, KumirParserRULE_algorithmName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KumirParserID {
		{
			p.SetState(310)
			p.Match(KumirParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAlgorithmHeaderContext is an interface to support dynamic dispatch.
type IAlgorithmHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALG_HEADER() antlr.TerminalNode
	AlgorithmNameTokens() IAlgorithmNameTokensContext
	TypeSpecifier() ITypeSpecifierContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	ParameterList() IParameterListContext

	// IsAlgorithmHeaderContext differentiates from other interfaces.
	IsAlgorithmHeaderContext()
}

type AlgorithmHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmHeaderContext() *AlgorithmHeaderContext {
	var p = new(AlgorithmHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmHeader
	return p
}

func InitEmptyAlgorithmHeaderContext(p *AlgorithmHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmHeader
}

func (*AlgorithmHeaderContext) IsAlgorithmHeaderContext() {}

func NewAlgorithmHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmHeaderContext {
	var p = new(AlgorithmHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_algorithmHeader

	return p
}

func (s *AlgorithmHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmHeaderContext) ALG_HEADER() antlr.TerminalNode {
	return s.GetToken(KumirParserALG_HEADER, 0)
}

func (s *AlgorithmHeaderContext) AlgorithmNameTokens() IAlgorithmNameTokensContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmNameTokensContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmNameTokensContext)
}

func (s *AlgorithmHeaderContext) TypeSpecifier() ITypeSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *AlgorithmHeaderContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserLPAREN, 0)
}

func (s *AlgorithmHeaderContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserRPAREN, 0)
}

func (s *AlgorithmHeaderContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *AlgorithmHeaderContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *AlgorithmHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAlgorithmHeader(s)
	}
}

func (s *AlgorithmHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAlgorithmHeader(s)
	}
}

func (s *AlgorithmHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAlgorithmHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AlgorithmHeader() (localctx IAlgorithmHeaderContext) {
	localctx = NewAlgorithmHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, KumirParserRULE_algorithmHeader)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(KumirParserALG_HEADER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(316)
			p.TypeSpecifier()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(319)
		p.AlgorithmNameTokens()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserLPAREN {
		{
			p.SetState(320)
			p.Match(KumirParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&138499810394112) != 0 {
			{
				p.SetState(321)
				p.ParameterList()
			}

		}
		{
			p.SetState(324)
			p.Match(KumirParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(327)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IPreConditionContext is an interface to support dynamic dispatch.
type IPreConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRE_CONDITION() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsPreConditionContext differentiates from other interfaces.
	IsPreConditionContext()
}

type PreConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreConditionContext() *PreConditionContext {
	var p = new(PreConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_preCondition
	return p
}

func InitEmptyPreConditionContext(p *PreConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_preCondition
}

func (*PreConditionContext) IsPreConditionContext() {}

func NewPreConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreConditionContext {
	var p = new(PreConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_preCondition

	return p
}

func (s *PreConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *PreConditionContext) PRE_CONDITION() antlr.TerminalNode {
	return s.GetToken(KumirParserPRE_CONDITION, 0)
}

func (s *PreConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PreConditionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *PreConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPreCondition(s)
	}
}

func (s *PreConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPreCondition(s)
	}
}

func (s *PreConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPreCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PreCondition() (localctx IPreConditionContext) {
	localctx = NewPreConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, KumirParserRULE_preCondition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(KumirParserPRE_CONDITION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Expression()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(332)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IPostConditionContext is an interface to support dynamic dispatch.
type IPostConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POST_CONDITION() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsPostConditionContext differentiates from other interfaces.
	IsPostConditionContext()
}

type PostConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostConditionContext() *PostConditionContext {
	var p = new(PostConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_postCondition
	return p
}

func InitEmptyPostConditionContext(p *PostConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_postCondition
}

func (*PostConditionContext) IsPostConditionContext() {}

func NewPostConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostConditionContext {
	var p = new(PostConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_postCondition

	return p
}

func (s *PostConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostConditionContext) POST_CONDITION() antlr.TerminalNode {
	return s.GetToken(KumirParserPOST_CONDITION, 0)
}

func (s *PostConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PostConditionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *PostConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPostCondition(s)
	}
}

func (s *PostConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPostCondition(s)
	}
}

func (s *PostConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPostCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PostCondition() (localctx IPostConditionContext) {
	localctx = NewPostConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, KumirParserRULE_postCondition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(KumirParserPOST_CONDITION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Expression()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(337)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IAlgorithmBodyContext is an interface to support dynamic dispatch.
type IAlgorithmBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementSequence() IStatementSequenceContext

	// IsAlgorithmBodyContext differentiates from other interfaces.
	IsAlgorithmBodyContext()
}

type AlgorithmBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmBodyContext() *AlgorithmBodyContext {
	var p = new(AlgorithmBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmBody
	return p
}

func InitEmptyAlgorithmBodyContext(p *AlgorithmBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmBody
}

func (*AlgorithmBodyContext) IsAlgorithmBodyContext() {}

func NewAlgorithmBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmBodyContext {
	var p = new(AlgorithmBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_algorithmBody

	return p
}

func (s *AlgorithmBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmBodyContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *AlgorithmBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAlgorithmBody(s)
	}
}

func (s *AlgorithmBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAlgorithmBody(s)
	}
}

func (s *AlgorithmBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAlgorithmBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AlgorithmBody() (localctx IAlgorithmBodyContext) {
	localctx = NewAlgorithmBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, KumirParserRULE_algorithmBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.StatementSequence()
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

// IStatementSequenceContext is an interface to support dynamic dispatch.
type IStatementSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementSequenceContext differentiates from other interfaces.
	IsStatementSequenceContext()
}

type StatementSequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementSequenceContext() *StatementSequenceContext {
	var p = new(StatementSequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_statementSequence
	return p
}

func InitEmptyStatementSequenceContext(p *StatementSequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_statementSequence
}

func (*StatementSequenceContext) IsStatementSequenceContext() {}

func NewStatementSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementSequenceContext {
	var p = new(StatementSequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_statementSequence

	return p
}

func (s *StatementSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementSequenceContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementSequenceContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterStatementSequence(s)
	}
}

func (s *StatementSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitStatementSequence(s)
	}
}

func (s *StatementSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitStatementSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) StatementSequence() (localctx IStatementSequenceContext) {
	localctx = NewStatementSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KumirParserRULE_statementSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1180503819750687936) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&15889) != 0) {
		{
			p.SetState(342)
			p.Statement()
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QualifiedIdentifier() IQualifiedIdentifierContext
	LBRACK() antlr.TerminalNode
	IndexList() IIndexListContext
	RBRACK() antlr.TerminalNode
	RETURN_VALUE() antlr.TerminalNode

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *LvalueContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(KumirParserLBRACK, 0)
}

func (s *LvalueContext) IndexList() IIndexListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexListContext)
}

func (s *LvalueContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(KumirParserRBRACK, 0)
}

func (s *LvalueContext) RETURN_VALUE() antlr.TerminalNode {
	return s.GetToken(KumirParserRETURN_VALUE, 0)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (s *LvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KumirParserRULE_lvalue)
	var _la int

	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.QualifiedIdentifier()
		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == KumirParserLBRACK {
			{
				p.SetState(349)
				p.Match(KumirParserLBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(350)
				p.IndexList()
			}
			{
				p.SetState(351)
				p.Match(KumirParserRBRACK)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case KumirParserRETURN_VALUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Match(KumirParserRETURN_VALUE)
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

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lvalue() ILvalueContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *AssignmentStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KumirParserASSIGN, 0)
}

func (s *AssignmentStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KumirParserRULE_assignmentStatement)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.Lvalue()
		}
		{
			p.SetState(359)
			p.Match(KumirParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
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

// IIoArgumentContext is an interface to support dynamic dispatch.
type IIoArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	NEWLINE_CONST() antlr.TerminalNode

	// IsIoArgumentContext differentiates from other interfaces.
	IsIoArgumentContext()
}

type IoArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIoArgumentContext() *IoArgumentContext {
	var p = new(IoArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioArgument
	return p
}

func InitEmptyIoArgumentContext(p *IoArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioArgument
}

func (*IoArgumentContext) IsIoArgumentContext() {}

func NewIoArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IoArgumentContext {
	var p = new(IoArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_ioArgument

	return p
}

func (s *IoArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *IoArgumentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IoArgumentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IoArgumentContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOLON)
}

func (s *IoArgumentContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOLON, i)
}

func (s *IoArgumentContext) NEWLINE_CONST() antlr.TerminalNode {
	return s.GetToken(KumirParserNEWLINE_CONST, 0)
}

func (s *IoArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IoArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IoArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterIoArgument(s)
	}
}

func (s *IoArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitIoArgument(s)
	}
}

func (s *IoArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitIoArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) IoArgument() (localctx IIoArgumentContext) {
	localctx = NewIoArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KumirParserRULE_ioArgument)
	var _la int

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)
			p.Expression()
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == KumirParserCOLON {
			{
				p.SetState(366)
				p.Match(KumirParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(367)
				p.Expression()
			}
			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == KumirParserCOLON {
				{
					p.SetState(368)
					p.Match(KumirParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(369)
					p.Expression()
				}

			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.Match(KumirParserNEWLINE_CONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IIoArgumentListContext is an interface to support dynamic dispatch.
type IIoArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIoArgument() []IIoArgumentContext
	IoArgument(i int) IIoArgumentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIoArgumentListContext differentiates from other interfaces.
	IsIoArgumentListContext()
}

type IoArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIoArgumentListContext() *IoArgumentListContext {
	var p = new(IoArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioArgumentList
	return p
}

func InitEmptyIoArgumentListContext(p *IoArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioArgumentList
}

func (*IoArgumentListContext) IsIoArgumentListContext() {}

func NewIoArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IoArgumentListContext {
	var p = new(IoArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_ioArgumentList

	return p
}

func (s *IoArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *IoArgumentListContext) AllIoArgument() []IIoArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIoArgumentContext); ok {
			len++
		}
	}

	tst := make([]IIoArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIoArgumentContext); ok {
			tst[i] = t.(IIoArgumentContext)
			i++
		}
	}

	return tst
}

func (s *IoArgumentListContext) IoArgument(i int) IIoArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIoArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIoArgumentContext)
}

func (s *IoArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KumirParserCOMMA)
}

func (s *IoArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KumirParserCOMMA, i)
}

func (s *IoArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IoArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IoArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterIoArgumentList(s)
	}
}

func (s *IoArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitIoArgumentList(s)
	}
}

func (s *IoArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitIoArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) IoArgumentList() (localctx IIoArgumentListContext) {
	localctx = NewIoArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KumirParserRULE_ioArgumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.IoArgument()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == KumirParserCOMMA {
		{
			p.SetState(378)
			p.Match(KumirParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.IoArgument()
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IIoStatementContext is an interface to support dynamic dispatch.
type IIoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IoArgumentList() IIoArgumentListContext
	INPUT() antlr.TerminalNode
	OUTPUT() antlr.TerminalNode

	// IsIoStatementContext differentiates from other interfaces.
	IsIoStatementContext()
}

type IoStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIoStatementContext() *IoStatementContext {
	var p = new(IoStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioStatement
	return p
}

func InitEmptyIoStatementContext(p *IoStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ioStatement
}

func (*IoStatementContext) IsIoStatementContext() {}

func NewIoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IoStatementContext {
	var p = new(IoStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_ioStatement

	return p
}

func (s *IoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IoStatementContext) IoArgumentList() IIoArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIoArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIoArgumentListContext)
}

func (s *IoStatementContext) INPUT() antlr.TerminalNode {
	return s.GetToken(KumirParserINPUT, 0)
}

func (s *IoStatementContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(KumirParserOUTPUT, 0)
}

func (s *IoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterIoStatement(s)
	}
}

func (s *IoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitIoStatement(s)
	}
}

func (s *IoStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitIoStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) IoStatement() (localctx IIoStatementContext) {
	localctx = NewIoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, KumirParserRULE_ioStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KumirParserINPUT || _la == KumirParserOUTPUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(386)
		p.IoArgumentList()
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	THEN() antlr.TerminalNode
	AllStatementSequence() []IStatementSequenceContext
	StatementSequence(i int) IStatementSequenceContext
	FI() antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(KumirParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) THEN() antlr.TerminalNode {
	return s.GetToken(KumirParserTHEN, 0)
}

func (s *IfStatementContext) AllStatementSequence() []IStatementSequenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			len++
		}
	}

	tst := make([]IStatementSequenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementSequenceContext); ok {
			tst[i] = t.(IStatementSequenceContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) StatementSequence(i int) IStatementSequenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *IfStatementContext) FI() antlr.TerminalNode {
	return s.GetToken(KumirParserFI, 0)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(KumirParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, KumirParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(KumirParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.Expression()
	}
	{
		p.SetState(390)
		p.Match(KumirParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.StatementSequence()
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserELSE {
		{
			p.SetState(392)
			p.Match(KumirParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.StatementSequence()
		}

	}
	{
		p.SetState(396)
		p.Match(KumirParserFI)
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

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_caseBlock
	return p
}

func InitEmptyCaseBlockContext(p *CaseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_caseBlock
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(KumirParserCASE, 0)
}

func (s *CaseBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseBlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(KumirParserCOLON, 0)
}

func (s *CaseBlockContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (s *CaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, KumirParserRULE_caseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(KumirParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Expression()
	}
	{
		p.SetState(400)
		p.Match(KumirParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.StatementSequence()
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

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	FI() antlr.TerminalNode
	AllCaseBlock() []ICaseBlockContext
	CaseBlock(i int) ICaseBlockContext
	ELSE() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(KumirParserSWITCH, 0)
}

func (s *SwitchStatementContext) FI() antlr.TerminalNode {
	return s.GetToken(KumirParserFI, 0)
}

func (s *SwitchStatementContext) AllCaseBlock() []ICaseBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseBlockContext); ok {
			len++
		}
	}

	tst := make([]ICaseBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseBlockContext); ok {
			tst[i] = t.(ICaseBlockContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) CaseBlock(i int) ICaseBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SwitchStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(KumirParserELSE, 0)
}

func (s *SwitchStatementContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, KumirParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(KumirParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KumirParserCASE {
		{
			p.SetState(404)
			p.CaseBlock()
		}

		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserELSE {
		{
			p.SetState(409)
			p.Match(KumirParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.StatementSequence()
		}

	}
	{
		p.SetState(413)
		p.Match(KumirParserFI)
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

// IEndLoopConditionContext is an interface to support dynamic dispatch.
type IEndLoopConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENDLOOP_COND() antlr.TerminalNode
	Expression() IExpressionContext

	// IsEndLoopConditionContext differentiates from other interfaces.
	IsEndLoopConditionContext()
}

type EndLoopConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndLoopConditionContext() *EndLoopConditionContext {
	var p = new(EndLoopConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_endLoopCondition
	return p
}

func InitEmptyEndLoopConditionContext(p *EndLoopConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_endLoopCondition
}

func (*EndLoopConditionContext) IsEndLoopConditionContext() {}

func NewEndLoopConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndLoopConditionContext {
	var p = new(EndLoopConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_endLoopCondition

	return p
}

func (s *EndLoopConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *EndLoopConditionContext) ENDLOOP_COND() antlr.TerminalNode {
	return s.GetToken(KumirParserENDLOOP_COND, 0)
}

func (s *EndLoopConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EndLoopConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndLoopConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndLoopConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterEndLoopCondition(s)
	}
}

func (s *EndLoopConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitEndLoopCondition(s)
	}
}

func (s *EndLoopConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitEndLoopCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) EndLoopCondition() (localctx IEndLoopConditionContext) {
	localctx = NewEndLoopConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, KumirParserRULE_endLoopCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Match(KumirParserENDLOOP_COND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(416)
		p.Expression()
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

// ILoopSpecifierContext is an interface to support dynamic dispatch.
type ILoopSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	FROM() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	TO() antlr.TerminalNode
	STEP() antlr.TerminalNode
	WHILE() antlr.TerminalNode
	TIMES() antlr.TerminalNode

	// IsLoopSpecifierContext differentiates from other interfaces.
	IsLoopSpecifierContext()
}

type LoopSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopSpecifierContext() *LoopSpecifierContext {
	var p = new(LoopSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_loopSpecifier
	return p
}

func InitEmptyLoopSpecifierContext(p *LoopSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_loopSpecifier
}

func (*LoopSpecifierContext) IsLoopSpecifierContext() {}

func NewLoopSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopSpecifierContext {
	var p = new(LoopSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_loopSpecifier

	return p
}

func (s *LoopSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopSpecifierContext) FOR() antlr.TerminalNode {
	return s.GetToken(KumirParserFOR, 0)
}

func (s *LoopSpecifierContext) ID() antlr.TerminalNode {
	return s.GetToken(KumirParserID, 0)
}

func (s *LoopSpecifierContext) FROM() antlr.TerminalNode {
	return s.GetToken(KumirParserFROM, 0)
}

func (s *LoopSpecifierContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LoopSpecifierContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopSpecifierContext) TO() antlr.TerminalNode {
	return s.GetToken(KumirParserTO, 0)
}

func (s *LoopSpecifierContext) STEP() antlr.TerminalNode {
	return s.GetToken(KumirParserSTEP, 0)
}

func (s *LoopSpecifierContext) WHILE() antlr.TerminalNode {
	return s.GetToken(KumirParserWHILE, 0)
}

func (s *LoopSpecifierContext) TIMES() antlr.TerminalNode {
	return s.GetToken(KumirParserTIMES, 0)
}

func (s *LoopSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLoopSpecifier(s)
	}
}

func (s *LoopSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLoopSpecifier(s)
	}
}

func (s *LoopSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLoopSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) LoopSpecifier() (localctx ILoopSpecifierContext) {
	localctx = NewLoopSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, KumirParserRULE_loopSpecifier)
	var _la int

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserFOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Match(KumirParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(KumirParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.Match(KumirParserFROM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Expression()
		}
		{
			p.SetState(422)
			p.Match(KumirParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Expression()
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == KumirParserSTEP {
			{
				p.SetState(424)
				p.Match(KumirParserSTEP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(425)
				p.Expression()
			}

		}

	case KumirParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(KumirParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Expression()
		}

	case KumirParserNEWLINE_CONST, KumirParserNOT, KumirParserRETURN_VALUE, KumirParserTRUE, KumirParserFALSE, KumirParserPLUS, KumirParserMINUS, KumirParserLPAREN, KumirParserLBRACE, KumirParserCHAR_LITERAL, KumirParserSTRING, KumirParserREAL, KumirParserINTEGER, KumirParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.Expression()
		}
		{
			p.SetState(431)
			p.Match(KumirParserTIMES)
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

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOP() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext
	ENDLOOP() antlr.TerminalNode
	EndLoopCondition() IEndLoopConditionContext
	LoopSpecifier() ILoopSpecifierContext

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_loopStatement
	return p
}

func InitEmptyLoopStatementContext(p *LoopStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_loopStatement
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(KumirParserLOOP, 0)
}

func (s *LoopStatementContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *LoopStatementContext) ENDLOOP() antlr.TerminalNode {
	return s.GetToken(KumirParserENDLOOP, 0)
}

func (s *LoopStatementContext) EndLoopCondition() IEndLoopConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndLoopConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEndLoopConditionContext)
}

func (s *LoopStatementContext) LoopSpecifier() ILoopSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopSpecifierContext)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, KumirParserRULE_loopStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Match(KumirParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(436)
			p.LoopSpecifier()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(439)
		p.StatementSequence()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserENDLOOP:
		{
			p.SetState(440)
			p.Match(KumirParserENDLOOP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case KumirParserENDLOOP_COND:
		{
			p.SetState(441)
			p.EndLoopCondition()
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

// IExitStatementContext is an interface to support dynamic dispatch.
type IExitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXIT() antlr.TerminalNode

	// IsExitStatementContext differentiates from other interfaces.
	IsExitStatementContext()
}

type ExitStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExitStatementContext() *ExitStatementContext {
	var p = new(ExitStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_exitStatement
	return p
}

func InitEmptyExitStatementContext(p *ExitStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_exitStatement
}

func (*ExitStatementContext) IsExitStatementContext() {}

func NewExitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExitStatementContext {
	var p = new(ExitStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_exitStatement

	return p
}

func (s *ExitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExitStatementContext) EXIT() antlr.TerminalNode {
	return s.GetToken(KumirParserEXIT, 0)
}

func (s *ExitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterExitStatement(s)
	}
}

func (s *ExitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitExitStatement(s)
	}
}

func (s *ExitStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitExitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ExitStatement() (localctx IExitStatementContext) {
	localctx = NewExitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, KumirParserRULE_exitStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(KumirParserEXIT)
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

// IPauseStatementContext is an interface to support dynamic dispatch.
type IPauseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAUSE() antlr.TerminalNode

	// IsPauseStatementContext differentiates from other interfaces.
	IsPauseStatementContext()
}

type PauseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPauseStatementContext() *PauseStatementContext {
	var p = new(PauseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_pauseStatement
	return p
}

func InitEmptyPauseStatementContext(p *PauseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_pauseStatement
}

func (*PauseStatementContext) IsPauseStatementContext() {}

func NewPauseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PauseStatementContext {
	var p = new(PauseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_pauseStatement

	return p
}

func (s *PauseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PauseStatementContext) PAUSE() antlr.TerminalNode {
	return s.GetToken(KumirParserPAUSE, 0)
}

func (s *PauseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PauseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterPauseStatement(s)
	}
}

func (s *PauseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitPauseStatement(s)
	}
}

func (s *PauseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitPauseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) PauseStatement() (localctx IPauseStatementContext) {
	localctx = NewPauseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, KumirParserRULE_pauseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(KumirParserPAUSE)
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

// IStopStatementContext is an interface to support dynamic dispatch.
type IStopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STOP() antlr.TerminalNode

	// IsStopStatementContext differentiates from other interfaces.
	IsStopStatementContext()
}

type StopStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStopStatementContext() *StopStatementContext {
	var p = new(StopStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_stopStatement
	return p
}

func InitEmptyStopStatementContext(p *StopStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_stopStatement
}

func (*StopStatementContext) IsStopStatementContext() {}

func NewStopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StopStatementContext {
	var p = new(StopStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_stopStatement

	return p
}

func (s *StopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StopStatementContext) STOP() antlr.TerminalNode {
	return s.GetToken(KumirParserSTOP, 0)
}

func (s *StopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterStopStatement(s)
	}
}

func (s *StopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitStopStatement(s)
	}
}

func (s *StopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitStopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) StopStatement() (localctx IStopStatementContext) {
	localctx = NewStopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, KumirParserRULE_stopStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(KumirParserSTOP)
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

// IAssertionStatementContext is an interface to support dynamic dispatch.
type IAssertionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSERTION() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssertionStatementContext differentiates from other interfaces.
	IsAssertionStatementContext()
}

type AssertionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertionStatementContext() *AssertionStatementContext {
	var p = new(AssertionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_assertionStatement
	return p
}

func InitEmptyAssertionStatementContext(p *AssertionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_assertionStatement
}

func (*AssertionStatementContext) IsAssertionStatementContext() {}

func NewAssertionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionStatementContext {
	var p = new(AssertionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_assertionStatement

	return p
}

func (s *AssertionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionStatementContext) ASSERTION() antlr.TerminalNode {
	return s.GetToken(KumirParserASSERTION, 0)
}

func (s *AssertionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssertionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAssertionStatement(s)
	}
}

func (s *AssertionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAssertionStatement(s)
	}
}

func (s *AssertionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAssertionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AssertionStatement() (localctx IAssertionStatementContext) {
	localctx = NewAssertionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, KumirParserRULE_assertionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(KumirParserASSERTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Expression()
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

// IProcedureCallStatementContext is an interface to support dynamic dispatch.
type IProcedureCallStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QualifiedIdentifier() IQualifiedIdentifierContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsProcedureCallStatementContext differentiates from other interfaces.
	IsProcedureCallStatementContext()
}

type ProcedureCallStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureCallStatementContext() *ProcedureCallStatementContext {
	var p = new(ProcedureCallStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_procedureCallStatement
	return p
}

func InitEmptyProcedureCallStatementContext(p *ProcedureCallStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_procedureCallStatement
}

func (*ProcedureCallStatementContext) IsProcedureCallStatementContext() {}

func NewProcedureCallStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureCallStatementContext {
	var p = new(ProcedureCallStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_procedureCallStatement

	return p
}

func (s *ProcedureCallStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureCallStatementContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *ProcedureCallStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserLPAREN, 0)
}

func (s *ProcedureCallStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(KumirParserRPAREN, 0)
}

func (s *ProcedureCallStatementContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ProcedureCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterProcedureCallStatement(s)
	}
}

func (s *ProcedureCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitProcedureCallStatement(s)
	}
}

func (s *ProcedureCallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitProcedureCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ProcedureCallStatement() (localctx IProcedureCallStatementContext) {
	localctx = NewProcedureCallStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, KumirParserRULE_procedureCallStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.QualifiedIdentifier()
	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserLPAREN {
		{
			p.SetState(454)
			p.Match(KumirParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-28)) & ^0x3f) == 0 && ((int64(1)<<(_la-28))&1090788651434115) != 0 {
			{
				p.SetState(455)
				p.ArgumentList()
			}

		}
		{
			p.SetState(458)
			p.Match(KumirParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	SEMICOLON() antlr.TerminalNode
	AssignmentStatement() IAssignmentStatementContext
	IoStatement() IIoStatementContext
	IfStatement() IIfStatementContext
	SwitchStatement() ISwitchStatementContext
	LoopStatement() ILoopStatementContext
	ExitStatement() IExitStatementContext
	PauseStatement() IPauseStatementContext
	StopStatement() IStopStatementContext
	AssertionStatement() IAssertionStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) IoStatement() IIoStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIoStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIoStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) LoopStatement() ILoopStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementContext) ExitStatement() IExitStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExitStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExitStatementContext)
}

func (s *StatementContext) PauseStatement() IPauseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPauseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPauseStatementContext)
}

func (s *StatementContext) StopStatement() IStopStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStopStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStopStatementContext)
}

func (s *StatementContext) AssertionStatement() IAssertionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssertionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssertionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, KumirParserRULE_statement)
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserINTEGER_TYPE, KumirParserREAL_TYPE, KumirParserBOOLEAN_TYPE, KumirParserCHAR_TYPE, KumirParserSTRING_TYPE, KumirParserINTEGER_ARRAY_TYPE, KumirParserREAL_ARRAY_TYPE, KumirParserCHAR_ARRAY_TYPE, KumirParserSTRING_ARRAY_TYPE, KumirParserBOOLEAN_ARRAY_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.VariableDeclaration()
		}
		p.SetState(463)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(462)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserNEWLINE_CONST, KumirParserNOT, KumirParserRETURN_VALUE, KumirParserTRUE, KumirParserFALSE, KumirParserPLUS, KumirParserMINUS, KumirParserLPAREN, KumirParserLBRACE, KumirParserCHAR_LITERAL, KumirParserSTRING, KumirParserREAL, KumirParserINTEGER, KumirParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.AssignmentStatement()
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(466)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserINPUT, KumirParserOUTPUT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(469)
			p.IoStatement()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(470)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(473)
			p.IfStatement()
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(474)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserSWITCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(477)
			p.SwitchStatement()
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(478)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserLOOP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(481)
			p.LoopStatement()
		}
		p.SetState(483)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(482)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserEXIT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(485)
			p.ExitStatement()
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(486)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserPAUSE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(489)
			p.PauseStatement()
		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(490)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserSTOP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(493)
			p.StopStatement()
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(494)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserASSERTION:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(497)
			p.AssertionStatement()
		}
		p.SetState(499)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(498)
				p.Match(KumirParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case KumirParserSEMICOLON:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(501)
			p.Match(KumirParserSEMICOLON)
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

// IAlgorithmDefinitionContext is an interface to support dynamic dispatch.
type IAlgorithmDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AlgorithmHeader() IAlgorithmHeaderContext
	ALG_BEGIN() antlr.TerminalNode
	AlgorithmBody() IAlgorithmBodyContext
	ALG_END() antlr.TerminalNode
	AllPreCondition() []IPreConditionContext
	PreCondition(i int) IPreConditionContext
	AllPostCondition() []IPostConditionContext
	PostCondition(i int) IPostConditionContext
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext
	AlgorithmName() IAlgorithmNameContext
	SEMICOLON() antlr.TerminalNode

	// IsAlgorithmDefinitionContext differentiates from other interfaces.
	IsAlgorithmDefinitionContext()
}

type AlgorithmDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlgorithmDefinitionContext() *AlgorithmDefinitionContext {
	var p = new(AlgorithmDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmDefinition
	return p
}

func InitEmptyAlgorithmDefinitionContext(p *AlgorithmDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_algorithmDefinition
}

func (*AlgorithmDefinitionContext) IsAlgorithmDefinitionContext() {}

func NewAlgorithmDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlgorithmDefinitionContext {
	var p = new(AlgorithmDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_algorithmDefinition

	return p
}

func (s *AlgorithmDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *AlgorithmDefinitionContext) AlgorithmHeader() IAlgorithmHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmHeaderContext)
}

func (s *AlgorithmDefinitionContext) ALG_BEGIN() antlr.TerminalNode {
	return s.GetToken(KumirParserALG_BEGIN, 0)
}

func (s *AlgorithmDefinitionContext) AlgorithmBody() IAlgorithmBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmBodyContext)
}

func (s *AlgorithmDefinitionContext) ALG_END() antlr.TerminalNode {
	return s.GetToken(KumirParserALG_END, 0)
}

func (s *AlgorithmDefinitionContext) AllPreCondition() []IPreConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPreConditionContext); ok {
			len++
		}
	}

	tst := make([]IPreConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPreConditionContext); ok {
			tst[i] = t.(IPreConditionContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmDefinitionContext) PreCondition(i int) IPreConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPreConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPreConditionContext)
}

func (s *AlgorithmDefinitionContext) AllPostCondition() []IPostConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPostConditionContext); ok {
			len++
		}
	}

	tst := make([]IPostConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPostConditionContext); ok {
			tst[i] = t.(IPostConditionContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmDefinitionContext) PostCondition(i int) IPostConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostConditionContext)
}

func (s *AlgorithmDefinitionContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *AlgorithmDefinitionContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *AlgorithmDefinitionContext) AlgorithmName() IAlgorithmNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmNameContext)
}

func (s *AlgorithmDefinitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(KumirParserSEMICOLON, 0)
}

func (s *AlgorithmDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlgorithmDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlgorithmDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterAlgorithmDefinition(s)
	}
}

func (s *AlgorithmDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitAlgorithmDefinition(s)
	}
}

func (s *AlgorithmDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitAlgorithmDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) AlgorithmDefinition() (localctx IAlgorithmDefinitionContext) {
	localctx = NewAlgorithmDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, KumirParserRULE_algorithmDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.AlgorithmHeader()
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&138469745623088) != 0 {
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case KumirParserPRE_CONDITION:
			{
				p.SetState(505)
				p.PreCondition()
			}

		case KumirParserPOST_CONDITION:
			{
				p.SetState(506)
				p.PostCondition()
			}

		case KumirParserINTEGER_TYPE, KumirParserREAL_TYPE, KumirParserBOOLEAN_TYPE, KumirParserCHAR_TYPE, KumirParserSTRING_TYPE, KumirParserINTEGER_ARRAY_TYPE, KumirParserREAL_ARRAY_TYPE, KumirParserCHAR_ARRAY_TYPE, KumirParserSTRING_ARRAY_TYPE, KumirParserBOOLEAN_ARRAY_TYPE:
			{
				p.SetState(507)
				p.VariableDeclaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(513)
		p.Match(KumirParserALG_BEGIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.AlgorithmBody()
	}
	{
		p.SetState(515)
		p.Match(KumirParserALG_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserID {
		{
			p.SetState(516)
			p.AlgorithmName()
		}

	}
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == KumirParserSEMICOLON {
		{
			p.SetState(519)
			p.Match(KumirParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IProgramItemContext is an interface to support dynamic dispatch.
type IProgramItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GlobalDeclaration() IGlobalDeclarationContext
	GlobalAssignment() IGlobalAssignmentContext

	// IsProgramItemContext differentiates from other interfaces.
	IsProgramItemContext()
}

type ProgramItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramItemContext() *ProgramItemContext {
	var p = new(ProgramItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_programItem
	return p
}

func InitEmptyProgramItemContext(p *ProgramItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_programItem
}

func (*ProgramItemContext) IsProgramItemContext() {}

func NewProgramItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramItemContext {
	var p = new(ProgramItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_programItem

	return p
}

func (s *ProgramItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramItemContext) GlobalDeclaration() IGlobalDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalDeclarationContext)
}

func (s *ProgramItemContext) GlobalAssignment() IGlobalAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalAssignmentContext)
}

func (s *ProgramItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterProgramItem(s)
	}
}

func (s *ProgramItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitProgramItem(s)
	}
}

func (s *ProgramItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitProgramItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) ProgramItem() (localctx IProgramItemContext) {
	localctx = NewProgramItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, KumirParserRULE_programItem)
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case KumirParserINTEGER_TYPE, KumirParserREAL_TYPE, KumirParserBOOLEAN_TYPE, KumirParserCHAR_TYPE, KumirParserSTRING_TYPE, KumirParserINTEGER_ARRAY_TYPE, KumirParserREAL_ARRAY_TYPE, KumirParserCHAR_ARRAY_TYPE, KumirParserSTRING_ARRAY_TYPE, KumirParserBOOLEAN_ARRAY_TYPE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(522)
			p.GlobalDeclaration()
		}

	case KumirParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
			p.GlobalAssignment()
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

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllProgramItem() []IProgramItemContext
	ProgramItem(i int) IProgramItemContext
	AllAlgorithmDefinition() []IAlgorithmDefinitionContext
	AlgorithmDefinition(i int) IAlgorithmDefinitionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = KumirParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = KumirParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(KumirParserEOF, 0)
}

func (s *ProgramContext) AllProgramItem() []IProgramItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramItemContext); ok {
			len++
		}
	}

	tst := make([]IProgramItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramItemContext); ok {
			tst[i] = t.(IProgramItemContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ProgramItem(i int) IProgramItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramItemContext)
}

func (s *ProgramContext) AllAlgorithmDefinition() []IAlgorithmDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAlgorithmDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IAlgorithmDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAlgorithmDefinitionContext); ok {
			tst[i] = t.(IAlgorithmDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) AlgorithmDefinition(i int) IAlgorithmDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlgorithmDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlgorithmDefinitionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KumirParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KumirParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KumirParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, KumirParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-36)) & ^0x3f) == 0 && ((int64(1)<<(_la-36))&2199023257567) != 0 {
		{
			p.SetState(526)
			p.ProgramItem()
		}

		p.SetState(531)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KumirParserALG_HEADER {
		{
			p.SetState(532)
			p.AlgorithmDefinition()
		}

		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(537)
		p.Match(KumirParserEOF)
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
