class Solution:
    def balancedStringSplit(self, s: str) -> int:
        l = r = result = 0

        for ch in s:
            if ch == 'L':
                l += 1
            else:
                r += 1

            if l ==r:
                l = r = 0
                result+= 1

        return result