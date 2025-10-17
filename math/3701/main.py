from typing import List

class Solution:
    def alternatingSum(self, nums: List[int]) -> int:
        ans = 0

        for i in range(len(nums)):
            if i % 2 == 0:
                ans += nums[i]
            else:
                ans -= nums[i]

        return ans

def main():
    nums = [1, 2, 3, 4]
    print(Solution().alternatingSum(nums))


if __name__ == '__main__':
    main()