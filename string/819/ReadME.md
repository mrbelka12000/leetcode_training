# 819. Most Common Word

## SOLVED
### https://leetcode.com/problems/most-common-word/description/
Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned. It is guaranteed there is at least one word that is not banned, and that the answer is unique.



The words in paragraph are case-insensitive and the answer should be returned in lowercase.





### Example 1:





Input: paragraph = &quot;Bob hit a ball, the hit BALL flew far after it was hit.&quot;, banned = [&quot;hit&quot;]


Output: &quot;ball&quot;



Explanation: 

&quot;hit&quot; occurs 3 times, but it is a banned word.

&quot;ball&quot; occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 

Note that words in the paragraph are not case sensitive,

that punctuation is ignored (even if adjacent to words, such as &quot;ball,&quot;), 

and that &quot;hit&quot; isnt the answer even though it occurs more because it is banned.





### Example 2:





Input: paragraph = &quot;a.&quot;, banned = []


Output: &quot;a&quot;







Constraints:





	1 <= paragraph.length <= 1000

	paragraph consists of English letters, space  , or one of the symbols: &quot;!?,;.&quot;.

	0 <= banned.length <= 100

	1 <= banned[i].length <= 10

	banned[i] consists of only lowercase English letters.



