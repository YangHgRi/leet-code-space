/*
 * @lc app=leetcode.cn id=1 lang=java
 *
 * [1] 两数之和
 */

// @lc code=start
import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            int dif = target - num;
            if (m.containsKey(dif)) {
                return new int[] { i, m.get(dif) };
            }
            m.put(num, i);
        }
        return new int[] {};
    }
}
// @lc code=end
