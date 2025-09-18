/*
 * @lc app=leetcode.cn id=21 lang=java
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 * int val;
 * ListNode next;
 * ListNode() {}
 * ListNode(int val) { this.val = val; }
 * ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        // Base case: if either list is null, return the other list
        if (list1 == null) {
            return list2;
        }
        if (list2 == null) {
            return list1;
        }
        // Recursive step: compare the values of the current nodes
        if (list1.val <= list2.val) {
            // If list1's value is smaller or equal, it becomes the head of the current
            // merge.
            // Recursively merge the rest of list1 (list1.next) with list2.
            list1.next = mergeTwoLists(list1.next, list2);
            return list1;
        } else {
            // If list2's value is smaller, it becomes the head of the current merge.
            // Recursively merge list1 with the rest of list2 (list2.next).
            list2.next = mergeTwoLists(list1, list2.next);
            return list2;
        }
    }
}
// @lc code=end
