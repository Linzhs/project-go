package leetcode

// Symbol       Value
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000

// special:
// 	IV 4
// 	IX 9
//	XL 40
//	XC 90
//	CD 400
//	CM 900
func romanToInt(s string) int {

	var result int

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			if i < len(s)-1 {
				if string(s[i+1]) == "V" {
					result += 4
					i++
					continue
				} else if string(s[i+1]) == "X" {
					result += 9
					i++
					continue
				}
			}
			result += 1
		case "V":
			result += 5
		case "X":
			if i < len(s)-1 {
				if string(s[i+1]) == "L" {
					result += 40
					i++
					continue
				} else if string(s[i+1]) == "C" {
					result += 90
					i++
					continue
				}
			}
			result += 10
		case "L":
			result += 50
		case "C":
			if i < len(s)-1 {
				if string(s[i+1]) == "D" {
					result += 400
					i++
					continue
				} else if string(s[i+1]) == "M" {
					result += 900
					i++
					continue
				}
			}
			result += 100
		case "D":
			result += 500
		case "M":
			result += 1000
		default:
			return 0
		}
	}

	return result
}
