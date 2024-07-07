# 1880. Check if Word Equals Summation of Two Words

## SOLVED
### https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/description/
The letter value of a letter is its position in the alphabet starting from 0 (i.e. a -> 0, b -> 1, c -> 2, etc.).



The numerical value of some string of lowercase English letters s is the concatenation of the letter values of each letter in s, which is then converted into an integer.





	For example, if s = &quot;acb&quot;, we concatenate each letters letter value, resulting in &quot;021&quot;. After converting it, we get 21.





You are given three strings firstWord, secondWord, and targetWord, each consisting of lowercase English letters a through j inclusive.



Return true if the summation of the numerical values of firstWord and secondWord equals the numerical value of targetWord, or false otherwise.





### Example 1:





Input: firstWord = &quot;acb&quot;, secondWord = &quot;cba&quot;, targetWord = &quot;cdb&quot;


Output: true



Explanation:

The numerical value of firstWord is &quot;acb&quot; -> &quot;021&quot; -> 21.

The numerical value of secondWord is &quot;cba&quot; -> &quot;210&quot; -> 210.

The numerical value of targetWord is &quot;cdb&quot; -> &quot;231&quot; -> 231.

We return true because 21 + 210 == 231.





### Example 2:





Input: firstWord = &quot;aaa&quot;, secondWord = &quot;a&quot;, targetWord = &quot;aab&quot;


Output: false



Explanation: 

The numerical value of firstWord is &quot;aaa&quot; -> &quot;000&quot; -> 0.

The numerical value of secondWord is &quot;a&quot; -> &quot;0&quot; -> 0.

The numerical value of targetWord is &quot;aab&quot; -> &quot;001&quot; -> 1.

We return false because 0 + 0 != 1.





### Example 3:





Input: firstWord = &quot;aaa&quot;, secondWord = &quot;a&quot;, targetWord = &quot;aaaa&quot;


Output: true



Explanation: 

The numerical value of firstWord is &quot;aaa&quot; -> &quot;000&quot; -> 0.

The numerical value of secondWord is &quot;a&quot; -> &quot;0&quot; -> 0.

The numerical value of targetWord is &quot;aaaa&quot; -> &quot;0000&quot; -> 0.

We return true because 0 + 0 == 0.







Constraints:





	1 <= firstWord.length, secondWord.length, targetWord.length <= 8

	firstWord, secondWord, and targetWord consist of lowercase English letters from a to j inclusive.



