# 3297. Count Substrings That Can Be Rearranged to Contain a String I

## SOLVED
### https://leetcode.com/problems/count-substrings-that-can-be-rearranged-to-contain-a-string-i/description/
You are given two strings word1 and word2.



A string x is called valid if x can be rearranged to have word2 as a prefix.



Return the total number of valid substrings of word1.





### Example 1:





Input: word1 = &quot;bcca&quot;, word2 = &quot;abc&quot;




Output: 1





Explanation:



The only valid substring is &quot;bcca&quot; which can be rearranged to &quot;abcc&quot; having &quot;abc&quot; as a prefix.





### Example 2:





Input: word1 = &quot;abcabc&quot;, word2 = &quot;abc&quot;




Output: 10





Explanation:



All the substrings except substrings of size 1 and size 2 are valid.





### Example 3:





Input: word1 = &quot;abcabc&quot;, word2 = &quot;aaabc&quot;




Output: 0







Constraints:





	1 <= word1.length <= 105

	1 <= word2.length <= 104

	word1 and word2 consist only of lowercase English letters.



