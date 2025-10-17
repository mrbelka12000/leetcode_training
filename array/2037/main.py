from typing import List

class Solution:
    def minMovesToSeat(self, seats: List[int], students: List[int]) -> int:
        ans = 0
        seats.sort()
        students.sort()

        for v in seats:
            ans += abs(v - students[0])
            students.pop(0)

        return ans

def main():

    print(Solution().minMovesToSeat(seats=[3,1,5], students=[2,7,4]))



if __name__ == '__main__':
    main()

