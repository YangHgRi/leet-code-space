/*
 * @lc app=leetcode.cn id=704 lang=java
 *
 * [704] 二分查找
 */

// @lc code=start
class Solution {
    public int search(int[] nums, int target) {
        int leftIndex = 0;
        int rightIndex = nums.length - 1;

        while (leftIndex <= rightIndex) {
            int midIndex = (rightIndex - leftIndex) / 2 + leftIndex;

            int num = nums[midIndex];

            if (num == target) {
                return midIndex;
            }

            if (num > target) {
                rightIndex = midIndex - 1;
            } else {
                leftIndex = midIndex + 1;
            }
        }

        return -1;
    }
}
// @lc code=end
