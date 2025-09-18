/*
 * @lc app=leetcode.cn id=49 lang=java
 *
 * [49] 字母异位词分组
 */

// @lc code=start

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        // 如果输入为空或长度为0，直接返回一个空的列表
        if (strs == null || strs.length == 0) {
            return new ArrayList<>();
        }
        // 使用HashMap来存储分组结果
        // key: 排序后的字符串（作为字母异位词的唯一标识）
        // value: 包含所有属于该字母异位词组的原始字符串的列表
        Map<String, List<String>> anagramGroups = new HashMap<>();
        // 遍历输入的每一个字符串
        for (String str : strs) {
            // 将字符串转换为字符数组
            char[] charArray = str.toCharArray();
            // 对字符数组进行排序，例如 "eat" -> "aet", "tea" -> "aet"
            Arrays.sort(charArray);
            // 将排序后的字符数组转换回字符串，作为哈希表的键
            String sortedStr = new String(charArray);
            // 检查哈希表中是否已经存在该键
            // 如果不存在，为该键创建一个新的列表
            if (!anagramGroups.containsKey(sortedStr)) {
                anagramGroups.put(sortedStr, new ArrayList<>());
            }
            // 将原始字符串str添加到该键对应的列表中
            anagramGroups.get(sortedStr).add(str);
        }
        // 最后，将哈希表中所有的值（即所有的字母异位词列表）收集起来，返回结果
        return new ArrayList<>(anagramGroups.values());
    }
}
// @lc code=end
