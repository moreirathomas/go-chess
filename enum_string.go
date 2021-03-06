// Code generated by "stringer -type=Color,Type"; DO NOT EDIT.

package chess

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[White-8]
	_ = x[Black-16]
}

const _Color_name = "WhiteBlack"

var _Color_index = [...]uint8{0, 5, 10}

func (i Color) String() string {
	switch i {
	case 8:
		i = 0
	case 16:
		i = 1
	default:
		i = -1
	}
	if i < 0 || i >= Color(len(_Color_index)-1) {
		return "Color(" + strconv.FormatInt(int64(i+8), 10) + ")"
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[None-0]
	_ = x[King-1]
	_ = x[Pawn-2]
	_ = x[Knight-3]
	_ = x[Bishop-4]
	_ = x[Rook-5]
	_ = x[Queen-6]
}

const _Type_name = "NoneKingPawnKnightBishopRookQueen"

var _Type_index = [...]uint8{0, 4, 8, 12, 18, 24, 28, 33}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
