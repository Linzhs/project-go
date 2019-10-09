package leetcode

// 0101
//+0111
//=1100
//
// xor:
// 0101
//^0111
//=0010
//
// and:
// 0101
//&0111
//=0101
//
// solution:
//  1. 先 异或 得到两个数不进位的相加和
//  2. 使用 与 运算得到两个数需要进位的数 并且进位
//  3. 此时两个数的和被拆分为 1、2步两个和
//  4. 重复以上运算 直到进位和为零 和转义为异或的值
func getSum(a int, b int) int {

	xs := a ^ b
	ac := (a & b) << 1
	for ac != 0 {
		xor := xs
		xs = ac ^ xor
		and := ac & xor
		ac = and << 1
	}

	return xs
}
