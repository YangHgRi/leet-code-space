/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30204
 *
 * [21] 合并两个有序链表
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 基本情况：如果任一链表为nil，则返回另一个链表
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 递归步骤：比较当前节点的值
	if list1.Val <= list2.Val {
		// 如果list1的值更小或相等，它成为当前合并的头节点
		// 递归合并list1的剩余部分 (list1.Next) 与 list2
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		// 如果list2的值更小，它成为当前合并的头节点
		// 递归合并 list1 与 list2的剩余部分 (list2.Next)
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
