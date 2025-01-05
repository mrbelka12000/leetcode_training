# 1221. Split a String in Balanced Strings

## SOLVED
### https://leetcode.com/problems/split-a-string-in-balanced-strings/description/
Balanced strings are those that have an equal quantity of L and R characters.



Given a balanced string s, split it into some number of substrings such that:





	Each substring is balanced.





Return the maximum number of balanced strings you can obtain.





### Example 1:





Input: s = &quot;RLRRLLRLRL&quot;


Output: 4



Explanation: s can be split into &quot;RL&quot;, &quot;RRLL&quot;, &quot;RL&quot;, &quot;RL&quot;, each substring contains same number of L and R.





### Example 2:





Input: s = &quot;RLRRRLLRLL&quot;


Output: 2



Explanation: s can be split into &quot;RL&quot;, &quot;RRRLLRLL&quot;, each substring contains same number of L and R.

Note that s cannot be split into &quot;RL&quot;, &quot;RR&quot;, &quot;RL&quot;, &quot;LR&quot;, &quot;LL&quot;, because the 2nd and 5th substrings are not balanced.



### Example 3:





Input: s = &quot;LLLLRRRR&quot;


Output: 1



Explanation: s can be split into &quot;LLLLRRRR&quot;.







Constraints:





	2 <= s.length <= 1000

	s[i] is either L or R.

	s is a balanced string.



