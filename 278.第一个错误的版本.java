/*
 * @lc app=leetcode.cn id=278 lang=java
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/* The isBadVersion API is defined in the parent class VersionControl.
      boolean isBadVersion(int version); */
class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int begin = 0;
        int end = n - 1;
        int mid = -1;

        while (begin <= end) {
            mid = (end - begin) / 2 + begin;
            if (isBadVersion(mid)) {
                end = mid - 1;
            } else {
                begin = mid + 1;
            }
        }

        return begin;
    }
}
// @lc code=end
