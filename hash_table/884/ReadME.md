# 884. Uncommon Words from Two Sentences

## SOLVED
### https://leetcode.com/problems/uncommon-words-from-two-sentences/description/
A sentence is a string of single-space separated words where each word consists only of lowercase letters.



A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.



Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.





### Example 1:





Input: s1 = &quot;this apple is sweet&quot;, s2 = &quot;this apple is sour&quot;




Output: [&quot;sweet&quot;,&quot;sour&quot;]





Explanation:



The word &quot;sweet&quot; appears only in s1, while the word &quot;sour&quot; appears only in s2.





### Example 2:





Input: s1 = &quot;apple apple&quot;, s2 = &quot;banana&quot;




Output: [&quot;banana&quot;]







Constraints:





	1 <= s1.length, s2.length <= 200

	s1 and s2 consist of lowercase English letters and spaces.

	s1 and s2 do not have leading or trailing spaces.

	All the words in s1 and s2 are separated by a single space.



