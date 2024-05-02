# 1647. Minimum Deletions to Make Character Frequencies Unique

## SOLVED
### https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/description/
A string s is called good if there are no two different characters in s that have the same frequency.



Given a string s, return the minimum number of characters you need to delete to make s good.



The frequency of a character in a string is the number of times it appears in the string. For example, in the string &quot;aab&quot;, the frequency of a is 2, while the frequency of b is 1.





### Example 1:





Input: s = &quot;aab&quot;


Output: 0



Explanation: s is already good.





### Example 2:





Input: s = &quot;aaabbbcc&quot;


Output: 2



Explanation: You can delete two bs resulting in the good string &quot;aaabcc&quot;.

Another way it to delete one b and one c resulting in the good string &quot;aaabbc&quot;.



### Example 3:





Input: s = &quot;ceabaacb&quot;


Output: 2



Explanation: You can delete both cs resulting in the good string &quot;eabaab&quot;.

Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).







Constraints:





	1 <= s.length <= 105

	scontains only lowercase English letters.



