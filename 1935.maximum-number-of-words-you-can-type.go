/*
 * @lc app=leetcode.cn id=1935 lang=golang
 * @lcpr version=30204
 *
 * [1935] 可以输入的最大单词数
 */

// @lcpr-template-start
package leetcodespace

import (
	"strings"
)

// @lcpr-template-end
// @lc code=start

/*
// 暴力破解法
/*
func canBeTypedWords(text string, brokenLetters string) int {
	count := 0
	// strings.Fields 相当于 text.split(" ")，它会自动处理多个空格并去除首尾空格
	words := strings.Fields(text)

	// 为了模拟Java中replace修改字符串后检查长度的逻辑，这里需要对每个单词进行处理
	for _, word := range words {
		originalLen := len(word)
		currentWord := word // Go字符串是不可变的，ReplaceAll会返回新字符串

		// 遍历所有坏字母，将当前单词中包含的坏字母替换掉
		for _, brokenChar := range brokenLetters { // brokenLetters.toCharArray()
			// strings.ReplaceAll 替换所有出现的子串
			currentWord = strings.ReplaceAll(currentWord, string(brokenChar), "")
		}

		newLen := len(currentWord)
		// 如果替换后长度不变，说明单词中不含任何坏字母
		if originalLen == newLen {
			count++
		}
	}
	return count
}
*/

/*
// 单循环哈希检测法
/*
func canBeTypedWords(text string, brokenLetters string) int {
	// 使用 map[rune]bool 作为哈希集合，存储坏字母，rune 代表Unicode字符
	bls := make(map[rune]bool, len(brokenLetters))
	for _, c := range brokenLetters { // 遍历坏字母字符串，将其加入集合
		bls[c] = true // bls.add(c)
	}

	count := 0
	flag := true // 标记当前正在检查的单词是否可打字 (true表示目前为止没遇到坏字母)

	// 遍历整个文本字符串的每一个字符
	for _, c := range text { // text.toCharArray()
		if c == ' ' { // 遇到空格，表示一个单词结束
			if flag { // 如果当前单词直到空格前都是可打字的
				count++ // 计数加一
			}
			flag = true // 重置标记，准备检查下一个单词
		} else if bls[c] { // 如果当前字符是坏字母
			flag = false // 标记为不可打字
		}
	}
	// 循环结束后，处理文本中最后一个单词（如果文本不是以空格结尾）
	if flag {
		count++
	}

	return count
}
*/

// bitmap索引标记法
// Java原注释的Go语言实现，这是最推荐的Go实现方式，高效且符合Go哲学
func canBeTypedWords(text string, brokenLetters string) int {
	// Java boolean[] brokenSet 版本的等效Go代码 (注释掉的部分)
	/*
		var brokenSet [26]bool // 用于标记26个小写字母是否为坏字母
		for _, c := range brokenLetters { // 遍历坏字母字符串
			if c >= 'a' && c <= 'z' {
				brokenSet[c-'a'] = true // 标记对应的字母为坏字母
			}
		}

		words := strings.Fields(text) // 将文本按空格分割成单词切片
		count := 0

		for _, word := range words { // 遍历每个单词
			canBeTyped := true // 标记当前单词是否可打字
			for _, c := range word { // 遍历单词中的每个字符
				if c >= 'a' && c <= 'z' && brokenSet[c-'a'] {
					canBeTyped = false // 发现坏字母，单词不可打
					break              // 提前跳出内层循环
				}
			}

			if canBeTyped { // 如果单词没有坏字母
				count++ // 计数加一
			}
		}
		return count
	*/

	// Java int brokenBitmask 版本的等效Go代码 (活跃部分)
	// 使用一个整数作为位掩码 (bitmap)，每个位代表一个小写字母
	brokenBitmask := 0
	for _, c := range brokenLetters {
		// 将对应字母的位设置为1，表示该字母是坏字母
		if c >= 'a' && c <= 'z' { // 确保是小写字母，防止意外索引
			brokenBitmask |= (1 << (c - 'a'))
		}
	}

	words := strings.Fields(text) // 将文本按空格分割成单词切片
	count := 0
	for _, word := range words {
		canBeTyped := true
		for _, c := range word { // 遍历单词中的每个字符
			if c >= 'a' && c <= 'z' { // 确保是小写字母以进行位运算
				// 通过位与操作检查当前字符是否为坏字母
				if (brokenBitmask & (1 << (c - 'a'))) != 0 {
					canBeTyped = false
					break // 发现坏字母，当前单词不可打，跳出内层循环
				}
			}
		}
		if canBeTyped {
			count++ // 当前单词没有坏字符，计数加一
		}
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// "hello world"\n"ad"\n
// @lcpr case=end

// @lcpr case=start
// "leet code"\n"lt"\n
// @lcpr case=end

// @lcpr case=start
// "leet code"\n"e"\n
// @lcpr case=end

*/
