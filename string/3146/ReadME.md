# 3146. Permutation Difference between Two Strings

## SOLVED
### https://leetcode.com/problems/permutation-difference-between-two-strings/description/
You are given two strings s and t such that every character occurs at most once in s and t is a permutation of s.



The permutation difference between s and t is defined as the sum of the absolute difference between the index of the occurrence of each character in s and the index of the occurrence of the same character in t.



Return the permutation difference between s and t.





### Example 1:





Input: s = &quot;abc&quot;, t = &quot;bac&quot;




Output: 2





Explanation:



For s = &quot;abc&quot; and t = &quot;bac&quot;, the permutation difference of s and t is equal to the sum of:





	The absolute difference between the index of the occurrence of &quot;a&quot; in s and the index of the occurrence of &quot;a&quot; in t.

	The absolute difference between the index of the occurrence of &quot;b&quot; in s and the index of the occurrence of &quot;b&quot; in t.

	The absolute difference between the index of the occurrence of &quot;c&quot; in s and the index of the occurrence of &quot;c&quot; in t.





That is, the permutation difference between s and t is equal to |0 - 1| + |1 - 0| + |2 - 2| = 2.





### Example 2:





Input: s = &quot;abcde&quot;, t = &quot;edbac&quot;




Output: 12





Explanation: The permutation difference between s and t is equal to |0 - 3| + |1 - 2| + |2 - 4| + |3 - 1| + |4 - 0| = 12.







Constraints:





	1 <= s.length <= 26

	Each character occurs at most once in s.

	t is a permutation of s.

	s consists only of lowercase English letters.



