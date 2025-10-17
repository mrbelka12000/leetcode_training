from typing import  List

class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        total = 0
        for num in nums:
            total += num


        return total % k


def main():

    print(Solution().minOperations([3,9,7], 5))



if __name__ == '__main__':
    main()

