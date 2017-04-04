// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // exected output.
}

var golden = []Golden{
	{"day", day_in, day_out},
	{"offset", offset_in, offset_out},
	{"gap", gap_in, gap_out},
	{"num", num_in, num_out},
	{"unum", unum_in, unum_out},
	{"prime", prime_in, prime_out},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const day_in = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`

const day_out = `
const _Day_name = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _Day_key_slice = []interface{}{
	Day(0),
	Day(1),
	Day(2),
	Day(3),
	Day(4),
	Day(5),
	Day(6),
}

var _Day_val_slice = []string{
	_Day_name[0:6],
	_Day_name[6:13],
	_Day_name[13:22],
	_Day_name[22:30],
	_Day_name[30:36],
	_Day_name[36:44],
	_Day_name[44:50],
}

func (i Day) GetEnumSlices() ([]interface{}, []string) {
	return _Day_key_slice, _Day_val_slice
}
`

// Enumeration with an offset.
// Also includes a duplicate.
const offset_in = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offset_out = `
const _Number_name = "OneTwoThree"

var _Number_key_slice = []interface{}{
	Number(1),
	Number(2),
	Number(3),
}

var _Number_val_slice = []string{
	_Number_name[0:3],
	_Number_name[3:6],
	_Number_name[6:11],
}

func (i Number) GetEnumSlices() ([]interface{}, []string) {
	return _Number_key_slice, _Number_val_slice
}
`

// Gaps and an offset.
const gap_in = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

const gap_out = `
const _Gap_name = "TwoThreeFiveSixSevenEightNineEleven"

var _Gap_key_slice = []interface{}{
	Gap(2),
	Gap(3),
	Gap(5),
	Gap(6),
	Gap(7),
	Gap(8),
	Gap(9),
	Gap(11),
}

var _Gap_val_slice = []string{
	_Gap_name[0:3],
	_Gap_name[3:8],
	_Gap_name[8:12],
	_Gap_name[12:15],
	_Gap_name[15:20],
	_Gap_name[20:25],
	_Gap_name[25:29],
	_Gap_name[29:35],
}

func (i Gap) GetEnumSlices() ([]interface{}, []string) {
	return _Gap_key_slice, _Gap_val_slice
}
`

// Signed integers spanning zero.
const num_in = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const num_out = `
const _Num_name = "m_2m_1m0m1m2"

var _Num_key_slice = []interface{}{
	Num(-2),
	Num(-1),
	Num(0),
	Num(1),
	Num(2),
}

var _Num_val_slice = []string{
	_Num_name[0:3],
	_Num_name[3:6],
	_Num_name[6:8],
	_Num_name[8:10],
	_Num_name[10:12],
}

func (i Num) GetEnumSlices() ([]interface{}, []string) {
	return _Num_key_slice, _Num_val_slice
}
`

// Unsigned integers spanning zero.
const unum_in = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

const unum_out = `
const _Unum_name = "m0m1m2m_2m_1"

var _Unum_key_slice = []interface{}{
	Unum(0),
	Unum(1),
	Unum(2),
	Unum(253),
	Unum(254),
}

var _Unum_val_slice = []string{
	_Unum_name[0:2],
	_Unum_name[2:4],
	_Unum_name[4:6],
	_Unum_name[6:9],
	_Unum_name[9:12],
}

func (i Unum) GetEnumSlices() ([]interface{}, []string) {
	return _Unum_key_slice, _Unum_val_slice
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const prime_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_key_slice = []interface{}{
	Prime(2),
	Prime(3),
	Prime(5),
	Prime(7),
	Prime(11),
	Prime(13),
	Prime(17),
	Prime(19),
	Prime(23),
	Prime(29),
	Prime(31),
	Prime(41),
	Prime(43),
}

var _Prime_val_slice = []string{
	_Prime_name[0:2],
	_Prime_name[2:4],
	_Prime_name[4:6],
	_Prime_name[6:8],
	_Prime_name[8:11],
	_Prime_name[11:14],
	_Prime_name[14:17],
	_Prime_name[17:20],
	_Prime_name[20:23],
	_Prime_name[23:26],
	_Prime_name[26:29],
	_Prime_name[29:32],
	_Prime_name[32:35],
}

func (i Prime) GetEnumSlices() ([]interface{}, []string) {
	return _Prime_key_slice, _Prime_val_slice
}
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		var g Generator
		input := "package test\n" + test.input
		file := test.name + ".go"
		g.parsePackage(".", []string{file}, input)
		// Extract the name and type of the constant from the first line.
		tokens := strings.SplitN(test.input, " ", 3)
		if len(tokens) != 3 {
			t.Fatalf("%s: need type declaration on first line", test.name)
		}
		g.generate(tokens[1])
		got := string(g.format())
		if got != test.output {
			t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
		}
	}
}
