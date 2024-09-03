# 1945. Sum of Digits of String After Convert

## SOLVED
### https://leetcode.com/problems/sum-of-digits-of-string-after-convert/description/
You are given a string s consisting of lowercase English letters, and an integer k.



First, convert s into an integer by replacing each letter with its position in the alphabet (i.e., replace a with 1, b with 2, ..., z with 26). Then, transform the integer by replacing it with the sum of its digits. Repeat the transform operation k times in total.



For example, if s = &quot;zbax&quot; and k = 2, then the resulting integer would be 8 by the following operations:





	Convert: &quot;zbax&quot; ➝ &quot;(26)(2)(1)(24)&quot; ➝ &quot;262124&quot; ➝ 262124

	Transform #1: 262124 ➝ 2 + 6 + 2 + 1 + 2 + 4➝ 17

	Transform #2: 17 ➝ 1 + 7 ➝ 8





Return the resulting integer after performing the operations described above.





### Example 1:





Input: s = &quot;iiii&quot;, k = 1


Output: 36



Explanation: The operations are as follows:

- Convert: &quot;iiii&quot; ➝ &quot;(9)(9)(9)(9)&quot; ➝ &quot;9999&quot; ➝ 9999

- Transform #1: 9999 ➝ 9 + 9 + 9 + 9 ➝ 36

Thus the resulting integer is 36.





### Example 2:





Input: s = &quot;leetcode&quot;, k = 2


Output: 6



Explanation: The operations are as follows:

- Convert: &quot;leetcode&quot; ➝ &quot;(12)(5)(5)(20)(3)(15)(4)(5)&quot; ➝ &quot;12552031545&quot; ➝ 12552031545

- Transform #1: 12552031545 ➝ 1 + 2 + 5 + 5 + 2 + 0 + 3 + 1 + 5 + 4 + 5 ➝ 33

- Transform #2: 33 ➝ 3 + 3 ➝ 6

Thus the resulting integer is 6.





### Example 3:





Input: s = &quot;zbax&quot;, k = 2


Output: 8







Constraints:





	1 <= s.length <= 100

	1 <= k <= 10

	s consists of lowercase English letters.



