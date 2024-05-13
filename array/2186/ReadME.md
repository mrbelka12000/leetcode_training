# 2186. Minimum Number of Steps to Make Two Strings Anagram II

## SOLVED
### https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/description/
You are given two strings s and t. In one step, you can append any character to either s or t.



Return the minimum number of steps to make s and t anagrams of each other.



An anagram of a string is a string that contains the same characters with a different (or the same) ordering.





### Example 1:





Input: s = &quot;leetcode&quot;, t = &quot;coats&quot;


Output: 7



Explanation: 

- In 2 steps, we can append the letters in &quot;as&quot; onto s = &quot;leetcode&quot;, forming s = &quot;leetcodeas&quot;.

- In 5 steps, we can append the letters in &quot;leede&quot; onto t = &quot;coats&quot;, forming t = &quot;coatsleede&quot;.

&quot;leetcodeas&quot; and &quot;coatsleede&quot; are now anagrams of each other.

We used a total of 2 + 5 = 7 steps.

It can be shown that there is no way to make them anagrams of each other with less than 7 steps.





### Example 2:





Input: s = &quot;night&quot;, t = &quot;thing&quot;


Output: 0



Explanation: The given strings are already anagrams of each other. Thus, we do not need any further steps.







Constraints:





	1 <= s.length, t.length <= 2 * 105

	s and t consist of lowercase English letters.



