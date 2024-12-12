# 1935. Maximum Number of Words You Can Type

## SOLVED
### https://leetcode.com/problems/maximum-number-of-words-you-can-type/description/
There is a malfunctioning keyboard where some letter keys do not work. All other keys on the keyboard work properly.



Given a string text of words separated by a single space (no leading or trailing spaces) and a string brokenLetters of all distinct letter keys that are broken, return the number of words in text you can fully type using this keyboard.





### Example 1:





Input: text = &quot;hello world&quot;, brokenLetters = &quot;ad&quot;


Output: 1



Explanation: We cannot type &quot;world&quot; because the d key is broken.





### Example 2:





Input: text = &quot;leet code&quot;, brokenLetters = &quot;lt&quot;


Output: 1



Explanation: We cannot type &quot;leet&quot; because the l and t keys are broken.





### Example 3:





Input: text = &quot;leet code&quot;, brokenLetters = &quot;e&quot;


Output: 0



Explanation: We cannot type either word because the e key is broken.







Constraints:





	1 <= text.length <= 104

	0 <= brokenLetters.length <= 26

	text consists of words separated by a single space without any leading or trailing spaces.

	Each word only consists of lowercase English letters.

	brokenLetters consists of distinct lowercase English letters.



