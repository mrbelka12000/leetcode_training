# 506. Relative Ranks

## SOLVED
### https://leetcode.com/problems/relative-ranks/description/
You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.



The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on. The placement of each athlete determines their rank:





	The 1st place athletes rank is &quot;Gold Medal&quot;.

	The 2nd place athletes rank is &quot;Silver Medal&quot;.

	The 3rd place athletes rank is &quot;Bronze Medal&quot;.

	For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athletes rank is &quot;x&quot;).





Return an array answer of size n where answer[i] is the rank of the ith athlete.





### Example 1:





Input: score = [5,4,3,2,1]


Output: [&quot;Gold Medal&quot;,&quot;Silver Medal&quot;,&quot;Bronze Medal&quot;,&quot;4&quot;,&quot;5&quot;]



Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].



### Example 2:





Input: score = [10,3,8,9,4]


Output: [&quot;Gold Medal&quot;,&quot;5&quot;,&quot;Bronze Medal&quot;,&quot;Silver Medal&quot;,&quot;4&quot;]



Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].









Constraints:





	n == score.length

	1 <= n <= 104

	0 <= score[i] <= 106

	All the values in score are unique.



