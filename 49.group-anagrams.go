/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=30204
 *
 * [49] 字母异位词分组
 */

// @lcpr-template-start
package leetcodespace

import (
	"sort"
)

// @lcpr-template-end
// @lc code=start

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{} // 空输入返回空切片
	}

	// 使用map[string][]string存储结果，key是排序后的字符串，value是原始字符串列表
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		// 将字符串转换为字节切片，方便排序
		charArray := []byte(str)
		// 对字节切片排序，例如 "eat" -> "aet"
		sort.Slice(charArray, func(i, j int) bool {
			return charArray[i] < charArray[j]
		})
		// 排序后的字节切片转回字符串作为键
		sortedStr := string(charArray)

		// 将原始字符串添加到对应键的列表中
		// Go的map在首次访问不存在的key时，会返回该值类型的零值。
		// []string的零值是nil，可以直接对其使用append。
		// append会处理nil切片的情况，创建新的底层数组。
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// 收集map中所有的值（即所有的字母异位词列表）
	result := make([][]string, 0, len(anagramGroups)) // 预分配容量
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// ["eat", "tea", "tan", "ate", "nat", "bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
