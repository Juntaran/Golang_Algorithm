/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/30 17:28
  */

package Merkle_Tree

// 判断 n 是否为 2 的幂
func isPowerOfTwo(n uint64) bool {
	// http://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
	return n != 0 && (n&(n-1)) == 0
}

// 返回比 n 大的最小的 2 的幂
func nextPowerOfTwo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	// http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
	n --
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n ++
	return n
}

// 计算 log2(n)
func logBaseTwo(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	ret := uint64(0)
	for n != 0 {
		n >>= 1
		ret ++
	}
	return ret - 1
}

// Returns the ceil'd log2 value of n
// See: http://stackoverflow.com/a/15327567

var log2lookup []uint64 = []uint64{
	0xFFFFFFFF00000000,
	0x00000000FFFF0000,
	0x000000000000FF00,
	0x00000000000000F0,
	0x000000000000000C,
	0x0000000000000002,
}

func ceilLogBaseTwo(x uint64) uint64 {
	y := uint64(1)
	if (x & (x - 1)) == 0 {
		y = 0
	}
	j := uint64(32)
	i := uint64(0)

	for ; i < 6; i++ {
		k := j
		if (x & log2lookup[i]) == 0 {
			k = 0
		}
		y += k
		x >>= k
		j >>= 1
	}

	return y
}