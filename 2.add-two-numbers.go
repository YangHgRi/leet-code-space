/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30204
 *
 * [2] 两数相加
 */

// @lcpr-template-start
package leetcodespace

// @lcpr-template-end
// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{} // 哑节点
	curr := dummyHead        // 当前处理的节点
	carry := 0               // 进位
	// 循环直到两个链表都为空且没有进位
	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10 // 计算新的进位

		// 创建新节点并添加到结果链表
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next // 移动到下一个节点
	}
	return dummyHead.Next // 返回实际的头节点
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/
