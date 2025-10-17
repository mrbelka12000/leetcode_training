class Solution:
    def findPermutationDifference(self, s: str, t: str) -> int:
        ans = 0
        for i in range(len(s)):
            for j in range(len(t)):
                if s[i] == t[j]:
                    ans += abs(j -i)
        return ans


def main():

    print(Solution().findPermutationDifference("abc", "bac"))



if __name__ == '__main__':
    main()

