/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=30204
 *
 * [206] 反转链表
 */

// @lcpr-template-start
package leet_code_space

// @lcpr-template-end
// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // prev 初始化为 nil
	curr := head       // curr 初始化为头节点

	// 遍历链表直到 curr 为 nil
	for curr != nil {
		next := curr.Next // 暂存下一个节点
		curr.Next = prev  // 将当前节点的 Next 指针指向前一个节点 (完成反转)
		prev = curr       // prev 向前移动到当前节点
		curr = next       // curr 向前移动到之前暂存的下一个节点
	}

	return prev // 循环结束后，prev 指向原链表的尾节点，也就是新链表的头节点
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
