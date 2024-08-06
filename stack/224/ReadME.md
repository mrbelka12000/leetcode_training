# 224. Basic Calculator

## SOLVED
### https://leetcode.com/problems/basic-calculator/description/
Given a string s representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.



Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().





### Example 1:





Input: s = &quot;1 + 1&quot;


Output: 2





### Example 2:





Input: s = &quot; 2-1 + 2 &quot;


Output: 3





### Example 3:





Input: s = &quot;(1+(4+5+2)-3)+(6+8)&quot;


Output: 23







Constraints:





	1 <= s.length <= 3 * 105

	s consists of digits, +, -, (, ), and  .

	s represents a valid expression.

	+ is not used as a unary operation (i.e., &quot;+1&quot; and &quot;+(2 + 3)&quot; is invalid).

	- could be used as a unary operation (i.e., &quot;-1&quot; and &quot;-(2 + 3)&quot; is valid).

	There will be no two consecutive operators in the input.

	Every number and running calculation will fit in a signed 32-bit integer.



