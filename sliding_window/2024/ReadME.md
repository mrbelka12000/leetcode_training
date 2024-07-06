# 2024. Maximize the Confusion of an Exam

## SOLVED
### https://leetcode.com/problems/maximize-the-confusion-of-an-exam/description/
A teacher is writing a test with n true/false questions, with T denoting true and F denoting false. He wants to confuse the students by maximizing the number of consecutive questions with the same answer (multiple trues or multiple falses in a row).



You are given a string answerKey, where answerKey[i] is the original answer to the ith question. In addition, you are given an integer k, the maximum number of times you may perform the following operation:





	Change the answer key for any question to T or F (i.e., set answerKey[i] to T or F).





Return the maximum number of consecutive Ts or Fs in the answer key after performing the operation at most k times.





### Example 1:





Input: answerKey = &quot;TTFF&quot;, k = 2


Output: 4



Explanation: We can replace both the Fs with Ts to make answerKey = &quot;TTTT&quot;.

There are four consecutive Ts.





### Example 2:





Input: answerKey = &quot;TFFT&quot;, k = 1


Output: 3



Explanation: We can replace the first T with an F to make answerKey = &quot;FFFT&quot;.

Alternatively, we can replace the second T with an F to make answerKey = &quot;TFFF&quot;.

In both cases, there are three consecutive Fs.





### Example 3:





Input: answerKey = &quot;TTFTTFTT&quot;, k = 1


Output: 5



Explanation: We can replace the first F to make answerKey = &quot;TTTTTFTT&quot;

Alternatively, we can replace the second F to make answerKey = &quot;TTFTTTTT&quot;. 

In both cases, there are five consecutive Ts.







Constraints:





	n == answerKey.length

	1 <= n <= 5 * 104

	answerKey[i] is either T or F

	1 <= k <= n



