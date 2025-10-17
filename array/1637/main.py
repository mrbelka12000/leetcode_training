from typing import List

class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        sorted_points = sorted(points, key=lambda x: x[0])

        max_width = 0

        for i in range(1, len(sorted_points)):
            max_width= max(max_width, sorted_points[i][0]- sorted_points[i-1][0])

        return max_width

def main():

    print(Solution().maxWidthOfVerticalArea([[8,7],[9,9],[7,4],[9,7]]))



if __name__ == '__main__':
    main()

