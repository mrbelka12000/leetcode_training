# 424. Longest Repeating Character Replacement

## SOLVED
### https://leetcode.com/problems/longest-repeating-character-replacement/description/
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.



Return the length of the longest substring containing the same letter you can get after performing the above operations.





### Example 1:





Input: s = &quot;ABAB&quot;, k = 2


Output: 4



Explanation: Replace the two As with two Bs or vice versa.





### Example 2:





Input: s = &quot;AABABBA&quot;, k = 1


Output: 4



Explanation: Replace the one A in the middle with B and form &quot;AABBBBA&quot;.

The substring &quot;BBBB&quot; has the longest repeating letters, which is 4.

There may exists other ways to achieve this answer too.





Constraints:





	1 <= s.length <= 105

	s consists of only uppercase English letters.

	0 <= k <= s.length



