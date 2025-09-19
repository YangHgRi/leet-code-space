/*
 * @lc app=leetcode.cn id=3484 lang=golang
 * @lcpr version=30204
 *
 * [3484] 设计电子表格
 */

// @lcpr-template-start
package leetcodespace

import (
	"strconv"
	"strings"
	"unicode"
)

// @lcpr-template-end
// @lc code=start

type Spreadsheet struct {
	cellValues map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		cellValues: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cellValues[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.cellValues, cell)
}

func (this *Spreadsheet) GetValue(formula string) int {
	i := strings.Index(formula, "+")
	cell1 := formula[1:i]
	cell2 := formula[i+1:]

	var val1, val2 int
	if unicode.IsLetter(rune(cell1[0])) {
		val1 = this.cellValues[cell1]
	} else {
		val1, _ = strconv.Atoi(cell1)
	}
	if unicode.IsLetter(rune(cell2[0])) {
		val2 = this.cellValues[cell2]
	} else {
		val2, _ = strconv.Atoi(cell2)
	}

	return val1 + val2
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
// @lc code=end

/*
// @lcpr case=start
// 5+7"], ["A1", 10]\nA1+6"], ["B2", 15]\nA1+B2"], ["A1"]\nA1+B2"]]\n
// @lcpr case=end

*/
