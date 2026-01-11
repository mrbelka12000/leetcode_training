# 85. Maximal Rectangle

## SOLVED_WITH_HINT
### https://leetcode.com/problems/maximal-rectangle/description/
Given a rows x colsbinary matrix filled with 0s and 1s, find the largest rectangle containing only 1s and return its area.





### Example 1:





Input: matrix = [[&quot;1&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;,&quot;0&quot;],[&quot;1&quot;,&quot;0&quot;,&quot;1&quot;,&quot;1&quot;,&quot;1&quot;],[&quot;1&quot;,&quot;1&quot;,&quot;1&quot;,&quot;1&quot;,&quot;1&quot;],[&quot;1&quot;,&quot;0&quot;,&quot;0&quot;,&quot;1&quot;,&quot;0&quot;]]


Output: 6



Explanation: The maximal rectangle is shown in the above picture.





### Example 2:





Input: matrix = [[&quot;0&quot;]]


Output: 0





### Example 3:





Input: matrix = [[&quot;1&quot;]]


Output: 1







Constraints:





	rows == matrix.length

	cols == matrix[i].length

	1 <= rows, cols <= 200

	matrix[i][j] is 0 or 1.



