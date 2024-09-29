package mwpad

import (
	"errors"

	"golang.org/x/text/width"
)

// pad_string	可选。规定供填充使用的字符串。默认是空白。
// pad_type

// STR_PAD_BOTH - 填充字符串的两侧。如果不是偶数，则右侧获得额外的填充。
// STR_PAD_LEFT - 填充字符串的左侧。
// STR_PAD_RIGHT - 填充字符串的右侧。这是默认的。

const (
	STR_PAD_LEFT = iota
	STR_PAD_RIGHT
	STR_PAD_BOTH_LEFT_FIRST
	STR_PAD_BOTH_RIGHT_FIRST
)

type Option struct {
	PadRune      rune
	PadType      int
	NeedTruncate bool
}

var (
	defaultPadRune      rune = ' '
	defaultPadType      int  = STR_PAD_RIGHT
	defaultNeedTruncate bool = false

	errStr = ""
)

func strPad(s string, width int, opt *Option) (string, error) {
	if opt == nil {
		opt = &Option{
			PadRune: ' ',
			PadType: STR_PAD_RIGHT,
		}
	}

	//todo check opt
	if width <= 0 {
		return errStr, errors.New("width must be greater than 0")
	}

	w := WidthByEA(s)
	if w > width {
		return s, nil
	}

	changeRight := true
	if opt.PadType == STR_PAD_LEFT || opt.PadType == STR_PAD_BOTH_LEFT_FIRST {
		changeRight = false
	}
	for WidthByEA(s) < width {
		if changeRight {
			s = s + string(opt.PadRune)
			if opt.PadType == STR_PAD_BOTH_RIGHT_FIRST {
				changeRight = false
			}
		} else {
			s = string(opt.PadRune) + s
			if opt.PadType == STR_PAD_BOTH_LEFT_FIRST {
				changeRight = true
			}
		}
	}

	return s, nil
}

func WidthByEA(s string) (w int) {
	w = 0
	for _, r := range s {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianFullwidth:
			w += 2
		case width.EastAsianWide:
			w += 2
		case width.EastAsianHalfwidth:
			w += 1
		case width.EastAsianNarrow:
			w += 1
		case width.Neutral:
			w += 1
		case width.EastAsianAmbiguous:
			w += 1
		default:
			w += 1
		}
	}

	return w
}
