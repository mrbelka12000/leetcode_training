from typing import List

class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        ans = 0
        for i in nums:
            if i % 3 != 0:
                ans += 1

        return ans


def main():
    nums = [1, 2, 3]
    print(Solution().minimumOperations(nums))

if __name__ == '__main__':
    main()