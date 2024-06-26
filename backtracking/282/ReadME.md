# 282. Expression Add Operators

## SOLVED
### https://leetcode.com/problems/expression-add-operators/description/
Given a string num that contains only digits and an integer target, return all possibilities to insert the binary operators +, -, and/or * between the digits of num so that the resultant expression evaluates to the target value.



Note that operands in the returned expressions should not contain leading zeros.





### Example 1:





Input: num = &quot;123&quot;, target = 6


Output: [&quot;1*2*3&quot;,&quot;1+2+3&quot;]



Explanation: Both &quot;1*2*3&quot; and &quot;1+2+3&quot; evaluate to 6.





### Example 2:





Input: num = &quot;232&quot;, target = 8


Output: [&quot;2*3+2&quot;,&quot;2+3*2&quot;]



Explanation: Both &quot;2*3+2&quot; and &quot;2+3*2&quot; evaluate to 8.





### Example 3:





Input: num = &quot;3456237490&quot;, target = 9191


Output: []



Explanation: There are no expressions that can be created from &quot;3456237490&quot; to evaluate to 9191.







Constraints:





	1 <= num.length <= 10

	num consists of only digits.

	-231 <= target <= 231 - 1



