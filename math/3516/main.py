class Solution:
    def findClosest(self, x: int, y: int, z: int) -> int:
        x_dist = abs(x - z)
        y_dist = abs(y - z)

        if x_dist < y_dist:
            return 1
        elif x_dist > y_dist:
            return 2

        return 0


def main():
    print(Solution().findClosest(2, 7, 4))


if __name__ == '__main__':
    main()



