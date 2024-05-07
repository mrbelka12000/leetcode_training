# 1980. Find Unique Binary String

## SOLVED
### https://leetcode.com/problems/find-unique-binary-string/description/
Given an array of strings nums containing n unique binary strings each of length n, return a binary string of length n that does not appear in nums. If there are multiple answers, you may return any of them.





### Example 1:





Input: nums = [&quot;01&quot;,&quot;10&quot;]


Output: &quot;11&quot;



Explanation: &quot;11&quot; does not appear in nums. &quot;00&quot; would also be correct.





### Example 2:





Input: nums = [&quot;00&quot;,&quot;01&quot;]


Output: &quot;11&quot;



Explanation: &quot;11&quot; does not appear in nums. &quot;10&quot; would also be correct.





### Example 3:





Input: nums = [&quot;111&quot;,&quot;011&quot;,&quot;001&quot;]


Output: &quot;101&quot;



Explanation: &quot;101&quot; does not appear in nums. &quot;000&quot;, &quot;010&quot;, &quot;100&quot;, and &quot;110&quot; would also be correct.







Constraints:





	n == nums.length

	1 <= n <= 16

	nums[i].length == n

	nums[i] is either 0 or 1.

	All the strings of nums are unique.



