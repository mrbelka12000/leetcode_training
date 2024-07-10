# 1598. Crawler Log Folder

## SOLVED
### https://leetcode.com/problems/crawler-log-folder/description/
The Leetcode file system keeps a log each time some user performs a change folder operation.



The operations are described below:





	&quot;../&quot; : Move to the parent folder of the current folder. (If you are already in the main folder, remain in the same folder).

	&quot;./&quot; : Remain in the same folder.

	&quot;x/&quot; : Move to the child folder named x (This folder is guaranteed to always exist).





You are given a list of strings logs where logs[i] is the operation performed by the user at the ith step.



The file system starts in the main folder, then the operations in logs are performed.



Return the minimum number of operations needed to go back to the main folder after the change folder operations.





### Example 1:









Input: logs = [&quot;d1/&quot;,&quot;d2/&quot;,&quot;../&quot;,&quot;d21/&quot;,&quot;./&quot;]


Output: 2



Explanation: Use this change folder operation &quot;../&quot; 2 times and go back to the main folder.





### Example 2:









Input: logs = [&quot;d1/&quot;,&quot;d2/&quot;,&quot;./&quot;,&quot;d3/&quot;,&quot;../&quot;,&quot;d31/&quot;]


Output: 3





### Example 3:





Input: logs = [&quot;d1/&quot;,&quot;../&quot;,&quot;../&quot;,&quot;../&quot;]


Output: 0







Constraints:





	1 <= logs.length <= 103

	2 <= logs[i].length <= 10

	logs[i] contains lowercase English letters, digits, ., and /.

	logs[i] follows the format described in the statement.

	Folder names consist of lowercase English letters and digits.



