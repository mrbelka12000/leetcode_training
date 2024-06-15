# 1839. Longest Substring Of All Vowels in Order

## SOLVED
### https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/description/
A string is considered beautiful if it satisfies the following conditions:





	Each of the 5 English vowels (a, e, i, o, u) must appear at least once in it.

	The letters must be sorted in alphabetical order (i.e. all as before es, all es before is, etc.).





For example, strings &quot;aeiou&quot; and &quot;aaaaaaeiiiioou&quot; are considered beautiful, but &quot;uaeio&quot;, &quot;aeoiu&quot;, and &quot;aaaeeeooo&quot; are not beautiful.



Given a string word consisting of English vowels, return the length of the longest beautiful substring of word. If no such substring exists, return 0.



A substring is a contiguous sequence of characters in a string.





### Example 1:





Input: word = &quot;aeiaaioaaaaeiiiiouuuooaauuaeiu&quot;


Output: 13



Explanation: The longest beautiful substring in word is &quot;aaaaeiiiiouuu&quot; of length 13.



### Example 2:





Input: word = &quot;aeeeiiiioooauuuaeiou&quot;


Output: 5



Explanation: The longest beautiful substring in word is &quot;aeiou&quot; of length 5.





### Example 3:





Input: word = &quot;a&quot;


Output: 0



Explanation: There is no beautiful substring, so return 0.







Constraints:





	1 <= word.length <= 5 * 105

	word consists of characters a, e, i, o, and u.



