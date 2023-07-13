package tool

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// Krand 获取随机字符集
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 79}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			// 当此时还在字符数量限制内，每次从随机字符生成器（rand）里面获取一个字符，字符逻辑自己找rand
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
