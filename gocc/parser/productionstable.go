// Code generated by gocc; DO NOT EDIT.

package parser

import (
    "github.com/uliseslugo/ilusa_pp/ast"
    "github.com/uliseslugo/ilusa_pp/types")

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : P	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `P : "Program" id ":" P1 P2 P3 B	<< ast.NewProgram(X[1],X[6]) >>`,
		Id:         "P",
		NTType:     1,
		Index:      1,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProgram(X[1],X[6])
		},
	},
	ProdTabEntry{
		String: `P1 : empty	<<  >>`,
		Id:         "P1",
		NTType:     2,
		Index:      2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `P1 : CL	<<  >>`,
		Id:         "P1",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `P2 : empty	<<  >>`,
		Id:         "P2",
		NTType:     3,
		Index:      4,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `P2 : DV	<< ast.GlobalVarDec(X[0]) >>`,
		Id:         "P2",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GlobalVarDec(X[0])
		},
	},
	ProdTabEntry{
		String: `P3 : empty	<<  >>`,
		Id:         "P3",
		NTType:     4,
		Index:      6,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `P3 : F	<<  >>`,
		Id:         "P3",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `B : "main" "(" ")" "{" B1 "}"	<< X[4], nil >>`,
		Id:         "B",
		NTType:     5,
		Index:      8,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[4], nil
		},
	},
	ProdTabEntry{
		String: `B1 : empty	<<  >>`,
		Id:         "B1",
		NTType:     6,
		Index:      9,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `B1 : EST B1	<< ast.NewStatements(X[0],X[1]) >>`,
		Id:         "B1",
		NTType:     6,
		Index:      10,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStatements(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `CL : "Class" id ":" INH "{" "atributes" "{" ATR_CL "}" MET_CL "}"	<< ast.NewClass(X[1]) >>`,
		Id:         "CL",
		NTType:     7,
		Index:      11,
		NumSymbols: 11,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewClass(X[1])
		},
	},
	ProdTabEntry{
		String: `INH : empty	<<  >>`,
		Id:         "INH",
		NTType:     8,
		Index:      12,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `INH : "inherits" id	<<  >>`,
		Id:         "INH",
		NTType:     8,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MET_CL : empty	<<  >>`,
		Id:         "MET_CL",
		NTType:     9,
		Index:      14,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `MET_CL : "methods" "{" F MET_CL1 "}"	<<  >>`,
		Id:         "MET_CL",
		NTType:     9,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `MET_CL1 : empty	<<  >>`,
		Id:         "MET_CL1",
		NTType:     10,
		Index:      16,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `MET_CL1 : F MET_CL1	<<  >>`,
		Id:         "MET_CL1",
		NTType:     10,
		Index:      17,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ATR_CL : "public" ":" ATR_SIMP "private" ":" ATR_SIMP "protected" ":" ATR_SIMP	<<  >>`,
		Id:         "ATR_CL",
		NTType:     11,
		Index:      18,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ATR_SIMP : TIP_SIMP id DV2 ";" ATR_SIMP1	<<  >>`,
		Id:         "ATR_SIMP",
		NTType:     12,
		Index:      19,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ATR_SIMP1 : empty	<<  >>`,
		Id:         "ATR_SIMP1",
		NTType:     13,
		Index:      20,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ATR_SIMP1 : ATR_SIMP ATR_SIMP1	<<  >>`,
		Id:         "ATR_SIMP1",
		NTType:     13,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DV : "variables" "{" DV1 "}"	<< X[2], nil >>`,
		Id:         "DV",
		NTType:     14,
		Index:      22,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `DV1 : empty	<<  >>`,
		Id:         "DV1",
		NTType:     15,
		Index:      23,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `DV1 : DV2 ";" DV1	<< ast.NewBlockVariables(X[0],X[2]) >>`,
		Id:         "DV1",
		NTType:     15,
		Index:      24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBlockVariables(X[0],X[2])
		},
	},
	ProdTabEntry{
		String: `DV2 : TIP_COMP id DV4	<<  >>`,
		Id:         "DV2",
		NTType:     16,
		Index:      25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DV2 : DV6 DV5	<< ast.NewTypeVariables(X[0],X[1]) >>`,
		Id:         "DV2",
		NTType:     16,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTypeVariables(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `DV3 : "[" cte_i "]"	<< ast.NewIntConst(X[1]) >>`,
		Id:         "DV3",
		NTType:     17,
		Index:      27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntConst(X[1])
		},
	},
	ProdTabEntry{
		String: `DV4 : empty	<<  >>`,
		Id:         "DV4",
		NTType:     18,
		Index:      28,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `DV4 : "," id DV4	<<  >>`,
		Id:         "DV4",
		NTType:     18,
		Index:      29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DV5 : empty	<<  >>`,
		Id:         "DV5",
		NTType:     19,
		Index:      30,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `DV5 : "," id DV5	<< ast.NewVariable(nil,X[1],0,0, X[2]) >>`,
		Id:         "DV5",
		NTType:     19,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(nil,X[1],0,0, X[2])
		},
	},
	ProdTabEntry{
		String: `DV5 : "," id DV3 DV5	<< ast.NewVariable(nil,X[1],X[2],0, X[3]) >>`,
		Id:         "DV5",
		NTType:     19,
		Index:      32,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(nil,X[1],X[2],0, X[3])
		},
	},
	ProdTabEntry{
		String: `DV5 : "," id DV3 DV3 DV5	<< ast.NewVariable(nil,X[1],X[2],X[3], X[4]) >>`,
		Id:         "DV5",
		NTType:     19,
		Index:      33,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(nil,X[1],X[2],X[3], X[4])
		},
	},
	ProdTabEntry{
		String: `DV6 : TIP_SIMP id	<< ast.NewVariable(X[0],X[1],0,0, nil) >>`,
		Id:         "DV6",
		NTType:     20,
		Index:      34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(X[0],X[1],0,0, nil)
		},
	},
	ProdTabEntry{
		String: `DV6 : TIP_SIMP id DV3	<< ast.NewVariable(X[0],X[1],X[2],0, nil) >>`,
		Id:         "DV6",
		NTType:     20,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(X[0],X[1],X[2],0, nil)
		},
	},
	ProdTabEntry{
		String: `DV6 : TIP_SIMP id DV3 DV3	<< ast.NewVariable(X[0],X[1],X[2],X[3], nil) >>`,
		Id:         "DV6",
		NTType:     20,
		Index:      36,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVariable(X[0],X[1],X[2],X[3], nil)
		},
	},
	ProdTabEntry{
		String: `TIP_SIMP : "int"	<< types.Integer, nil >>`,
		Id:         "TIP_SIMP",
		NTType:     21,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.Integer, nil
		},
	},
	ProdTabEntry{
		String: `TIP_SIMP : "char"	<< types.Char, nil >>`,
		Id:         "TIP_SIMP",
		NTType:     21,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.Char, nil
		},
	},
	ProdTabEntry{
		String: `TIP_SIMP : "float"	<< types.Float, nil >>`,
		Id:         "TIP_SIMP",
		NTType:     21,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return types.Float, nil
		},
	},
	ProdTabEntry{
		String: `TIP_COMP : id	<<  >>`,
		Id:         "TIP_COMP",
		NTType:     22,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `F : "function" F1 id "(" F2 ")" "{" DV "body" "{" EST B1 "}" "}" F4	<< ast.NewFunction(X[2], X[7]) >>`,
		Id:         "F",
		NTType:     23,
		Index:      41,
		NumSymbols: 15,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[2], X[7])
		},
	},
	ProdTabEntry{
		String: `F1 : TIP_SIMP	<<  >>`,
		Id:         "F1",
		NTType:     24,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `F1 : "void"	<<  >>`,
		Id:         "F1",
		NTType:     24,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `F2 : empty	<<  >>`,
		Id:         "F2",
		NTType:     25,
		Index:      44,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `F2 : TIP_SIMP id F3	<< ast.NewFunction(X[1], nil) >>`,
		Id:         "F2",
		NTType:     25,
		Index:      45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `F3 : empty	<<  >>`,
		Id:         "F3",
		NTType:     26,
		Index:      46,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `F3 : "," TIP_SIMP id F3	<< ast.NewFunction(X[2], nil) >>`,
		Id:         "F3",
		NTType:     26,
		Index:      47,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFunction(X[2], nil)
		},
	},
	ProdTabEntry{
		String: `F4 : empty	<<  >>`,
		Id:         "F4",
		NTType:     27,
		Index:      48,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `F4 : F	<<  >>`,
		Id:         "F4",
		NTType:     27,
		Index:      49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EST : AS	<<  >>`,
		Id:         "EST",
		NTType:     28,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EST : DEC	<< ast.FinishIf(X[0]) >>`,
		Id:         "EST",
		NTType:     28,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FinishIf(X[0])
		},
	},
	ProdTabEntry{
		String: `EST : ESC	<< ast.FinishOutput(X[0]) >>`,
		Id:         "EST",
		NTType:     28,
		Index:      52,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FinishOutput(X[0])
		},
	},
	ProdTabEntry{
		String: `EST : LEC	<< ast.FinishInput(X[0]) >>`,
		Id:         "EST",
		NTType:     28,
		Index:      53,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FinishInput(X[0])
		},
	},
	ProdTabEntry{
		String: `EST : LLAM	<<  >>`,
		Id:         "EST",
		NTType:     28,
		Index:      54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EST : RET	<<  >>`,
		Id:         "EST",
		NTType:     28,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `EST : REP	<<  >>`,
		Id:         "EST",
		NTType:     28,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AS : id "=" H_EXP ";"	<< ast.NewAssignation(X[0],X[2]) >>`,
		Id:         "AS",
		NTType:     29,
		Index:      57,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAssignation(X[0],X[2])
		},
	},
	ProdTabEntry{
		String: `LLAM : id "(" LLAM1 ")"	<<  >>`,
		Id:         "LLAM",
		NTType:     30,
		Index:      58,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LLAM1 : empty	<<  >>`,
		Id:         "LLAM1",
		NTType:     31,
		Index:      59,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `LLAM1 : H_EXP LLAM2	<<  >>`,
		Id:         "LLAM1",
		NTType:     31,
		Index:      60,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LLAM2 : empty	<<  >>`,
		Id:         "LLAM2",
		NTType:     32,
		Index:      61,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `LLAM2 : "," H_EXP LLAM2	<<  >>`,
		Id:         "LLAM2",
		NTType:     32,
		Index:      62,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `RET : "return" "(" H_EXP ")" ";"	<< ast.Return(X[2]) >>`,
		Id:         "RET",
		NTType:     33,
		Index:      63,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Return(X[2])
		},
	},
	ProdTabEntry{
		String: `LEC : "input" "(" VAR LEC1 ")" ";"	<< ast.NewInput(X[2],X[3]) >>`,
		Id:         "LEC",
		NTType:     34,
		Index:      64,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInput(X[2],X[3])
		},
	},
	ProdTabEntry{
		String: `LEC1 : empty	<<  >>`,
		Id:         "LEC1",
		NTType:     35,
		Index:      65,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `LEC1 : "," VAR LEC1	<< ast.NewInput(X[1],X[2]) >>`,
		Id:         "LEC1",
		NTType:     35,
		Index:      66,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInput(X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `ESC : "output" "(" ESC1 ESC2 ")" ";"	<< ast.NewOutput(X[2], X[3]) >>`,
		Id:         "ESC",
		NTType:     36,
		Index:      67,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOutput(X[2], X[3])
		},
	},
	ProdTabEntry{
		String: `ESC1 : H_EXP	<<  >>`,
		Id:         "ESC1",
		NTType:     37,
		Index:      68,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ESC1 : cte_string	<<  >>`,
		Id:         "ESC1",
		NTType:     37,
		Index:      69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ESC2 : empty	<<  >>`,
		Id:         "ESC2",
		NTType:     38,
		Index:      70,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `ESC2 : "," ESC1 ESC2	<< ast.NewOutput(X[1], X[2]) >>`,
		Id:         "ESC2",
		NTType:     38,
		Index:      71,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOutput(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `DEC : "if" "(" H_EXP ")" "{" EST B1 "}" DEC1	<< ast.NewIf(X[2],X[5],X[6],X[8]) >>`,
		Id:         "DEC",
		NTType:     39,
		Index:      72,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIf(X[2],X[5],X[6],X[8])
		},
	},
	ProdTabEntry{
		String: `DEC1 : empty	<<  >>`,
		Id:         "DEC1",
		NTType:     40,
		Index:      73,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `DEC1 : "else" "{" EST B1 "}"	<< ast.NewElse(X[2],X[3]) >>`,
		Id:         "DEC1",
		NTType:     40,
		Index:      74,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewElse(X[2],X[3])
		},
	},
	ProdTabEntry{
		String: `REP : "while" "(" H_EXP REP1	<< ast.NewWhile(X[2],X[3]) >>`,
		Id:         "REP",
		NTType:     41,
		Index:      75,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewWhile(X[2],X[3])
		},
	},
	ProdTabEntry{
		String: `REP : "for" "(" AS ";" H_EXP ";" H_EXP REP1	<<  >>`,
		Id:         "REP",
		NTType:     41,
		Index:      76,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REP1 : ")" "{" EST B1 "}"	<< ast.LoopStatements(X[2],X[3]) >>`,
		Id:         "REP1",
		NTType:     42,
		Index:      77,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LoopStatements(X[2],X[3])
		},
	},
	ProdTabEntry{
		String: `VAR : id VAR1 VAR1	<< ast.GetIdDimConst(X[0], X[1], X[2]) >>`,
		Id:         "VAR",
		NTType:     43,
		Index:      78,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GetIdDimConst(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `VAR1 : empty	<<  >>`,
		Id:         "VAR1",
		NTType:     44,
		Index:      79,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `VAR1 : "[" H_EXP "]"	<< X[1], nil >>`,
		Id:         "VAR1",
		NTType:     44,
		Index:      80,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `H_EXP : S_EXP H_EXP1	<< ast.NewExpression(X[0],X[1]) >>`,
		Id:         "H_EXP",
		NTType:     45,
		Index:      81,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `H_EXP1 : empty	<<  >>`,
		Id:         "H_EXP1",
		NTType:     46,
		Index:      82,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `H_EXP1 : LOG S_EXP	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "H_EXP1",
		NTType:     46,
		Index:      83,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `S_EXP : EXP S_EXP1	<< ast.NewExpression(X[0],X[1]) >>`,
		Id:         "S_EXP",
		NTType:     47,
		Index:      84,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `S_EXP1 : empty	<<  >>`,
		Id:         "S_EXP1",
		NTType:     48,
		Index:      85,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `S_EXP1 : REL EXP	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "S_EXP1",
		NTType:     48,
		Index:      86,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `EXP : TERM EXP1	<< ast.NewExpression(X[0],X[1]) >>`,
		Id:         "EXP",
		NTType:     49,
		Index:      87,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `EXP1 : empty	<<  >>`,
		Id:         "EXP1",
		NTType:     50,
		Index:      88,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `EXP1 : "+" EXP	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "EXP1",
		NTType:     50,
		Index:      89,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `EXP1 : "-" EXP	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "EXP1",
		NTType:     50,
		Index:      90,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `TERM : FAC TERM1	<< ast.NewExpression(X[0],X[1]) >>`,
		Id:         "TERM",
		NTType:     51,
		Index:      91,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExpression(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `TERM1 : empty	<<  >>`,
		Id:         "TERM1",
		NTType:     52,
		Index:      92,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `TERM1 : "*" TERM	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "TERM1",
		NTType:     52,
		Index:      93,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `TERM1 : "/" TERM	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "TERM1",
		NTType:     52,
		Index:      94,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `FAC : "(" H_EXP ")"	<<  >>`,
		Id:         "FAC",
		NTType:     53,
		Index:      95,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FAC : VAR_CTE	<<  >>`,
		Id:         "FAC",
		NTType:     53,
		Index:      96,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `FAC : "+" VAR_CTE	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "FAC",
		NTType:     53,
		Index:      97,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `FAC : "-" VAR_CTE	<< ast.NewOperation(X[0],X[1]) >>`,
		Id:         "FAC",
		NTType:     53,
		Index:      98,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOperation(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `VAR_CTE : id	<< ast.NewIdConst(X[0]) >>`,
		Id:         "VAR_CTE",
		NTType:     54,
		Index:      99,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIdConst(X[0])
		},
	},
	ProdTabEntry{
		String: `VAR_CTE : cte_i	<< ast.NewIntConst(X[0]) >>`,
		Id:         "VAR_CTE",
		NTType:     54,
		Index:      100,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntConst(X[0])
		},
	},
	ProdTabEntry{
		String: `VAR_CTE : cte_float	<< ast.NewFloatConst(X[0]) >>`,
		Id:         "VAR_CTE",
		NTType:     54,
		Index:      101,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFloatConst(X[0])
		},
	},
	ProdTabEntry{
		String: `VAR_CTE : cte_char	<< ast.NewCharConst(X[0]) >>`,
		Id:         "VAR_CTE",
		NTType:     54,
		Index:      102,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCharConst(X[0])
		},
	},
	ProdTabEntry{
		String: `LOG : "&&"	<<  >>`,
		Id:         "LOG",
		NTType:     55,
		Index:      103,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LOG : "||"	<<  >>`,
		Id:         "LOG",
		NTType:     55,
		Index:      104,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `LOG : "!="	<<  >>`,
		Id:         "LOG",
		NTType:     55,
		Index:      105,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REL : "<"	<<  >>`,
		Id:         "REL",
		NTType:     56,
		Index:      106,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REL : ">"	<<  >>`,
		Id:         "REL",
		NTType:     56,
		Index:      107,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REL : ">="	<<  >>`,
		Id:         "REL",
		NTType:     56,
		Index:      108,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REL : "<="	<<  >>`,
		Id:         "REL",
		NTType:     56,
		Index:      109,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `REL : "=="	<<  >>`,
		Id:         "REL",
		NTType:     56,
		Index:      110,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}
