/*
 * @lc app=leetcode.cn id=1935 lang=java
 *
 * [1935] 可以输入的最大单词数
 */

// @lc code=start

// 暴力破解法
// class Solution {
//     public int canBeTypedWords(String text, String brokenLetters) {
//         int count = 0;
//         String[] sa = text.split(" ");
//         for (int i = 0; i < sa.length; i++) {
//             int originalLen = sa[i].length();
//             for (int j = 0; j < brokenLetters.length(); j++) {
//                 sa[i] = sa[i].replace(String.valueOf(brokenLetters.charAt(j)), "");
//             }
//             int newLen = sa[i].length();
//             if (originalLen == newLen) {
//                 count++;
//             }
//         }

//         return count;
//     }
// }

// 单循环哈希检测法
// import java.util.HashSet;
// import java.util.Set;

// class Solution {
//     public int canBeTypedWords(String text, String brokenLetters) {
//         Set<Character> bls = new HashSet<>(brokenLetters.length());
//         for (char c : brokenLetters.toCharArray()) {
//             bls.add(c);
//         }

//         int count = 0;
//         boolean flag = true;
//         for (char c : text.toCharArray()) {
//             if (c == ' ') {
//                 if (flag) {
//                     count++;
//                 }
//                 flag = true;
//             } else if (bls.contains(c)) {
//                 flag = false;
//             }
//         }
//         if (flag) {
//             count++;
//         }

//         return count;
//     }
// }

// bitmap索引标记法
class Solution {
    public int canBeTypedWords(String text, String brokenLetters) {
        // boolean[] brokenSet = new boolean[26];
        // for (char c : brokenLetters.toCharArray()) {
        // brokenSet[c - 'a'] = true;
        // }

        // String[] words = text.split(" ");
        // int count = 0;

        // for (String word : words) {
        // boolean canBeTyped = true;
        // for (char c : word.toCharArray()) {
        // if (brokenSet[c - 'a']) {
        // canBeTyped = false;
        // break;
        // }
        // }

        // if (canBeTyped) {
        // count++;
        // }
        // }

        // return count;

        int brokenBitmask = 0;
        for (char c : brokenLetters.toCharArray()) {
            brokenBitmask |= (1 << (c - 'a'));
        }
        String[] words = text.split(" ");
        int count = 0;
        for (String word : words) {
            boolean canBeTyped = true;
            for (char c : word.toCharArray()) {
                if ((brokenBitmask & (1 << (c - 'a'))) != 0) {
                    canBeTyped = false;
                    break;
                }
            }
            if (canBeTyped) {
                count++;
            }
        }
        return count;
    }
}
// @lc code=end
