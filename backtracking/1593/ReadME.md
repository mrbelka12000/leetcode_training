# 1593. Split a String Into the Max Number of Unique Substrings

## SOLVED
### https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/description/
Given a strings,return the maximumnumber of unique substrings that the given string can be split into.



You can split strings into any list ofnon-empty substrings, where the concatenation of the substrings forms the original string.However, you must split the substrings such that all of them are unique.



A substring is a contiguous sequence of characters within a string.





### Example 1:





Input: s = &quot;ababccc&quot;


Output: 5



Explanation: One way to split maximally is [a, b, ab, c, cc]. Splitting like [a, b, a, b, c, cc] is not valid as you have a and b multiple times.





### Example 2:





Input: s = &quot;aba&quot;


Output: 2



Explanation: One way to split maximally is [a, ba].





### Example 3:





Input: s = &quot;aa&quot;


Output: 1



Explanation: It is impossible to split the string any further.







Constraints:





	

	1 <= s.length<= 16

	

	

	s containsonly lower case English letters.

	



