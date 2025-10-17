class Solution:
    def reverseDegree(self, s: str) -> int:
        bytess = s.encode("utf8")
        ans = 0
        for i in range(len(bytess)):
            num = bytess[i] - 97
            ans += (26 - num) * (i+1)

        return ans


def main():
    print(Solution().reverseDegree('abc'))



if __name__ == '__main__':
    main()

