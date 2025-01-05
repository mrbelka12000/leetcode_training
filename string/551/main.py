class Solution:
    def checkRecord(self, s: str) -> bool:
        ab = lt = 0

        for ch in s:
            if ch == 'L':
                lt += 1
            if ch =='A':
                ab += 1
                lt = 0
            if ch == 'P':
                lt = 0
            if lt >= 3:
                return False

        if ab > 1:
            return False

        return True