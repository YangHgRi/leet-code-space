/*
 * @lc app=leetcode.cn id=70 lang=java
 *
 * [70] 爬楼梯
 */

// @lc code=start
class Solution {
    public int climbStairs(int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        if (n == 2) {
            return 2;
        }

        int last = 2;
        int current = 3;
        for (int i = 0; i < n - 3; i++) {
            int temp = current;
            current += last;
            last = temp;
        }

        return current;
    }
}
// @lc code=end
