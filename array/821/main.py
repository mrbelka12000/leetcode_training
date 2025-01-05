class Solution:
    def shortestToChar(self, s: str, c: str) -> List[int]:
        ans = [0] * len(s)
        occur = []

        for idx, ch in enumerate(s):
            if ch == c:
                occur.append(idx)

        for idx, ch in enumerate(s):
            min_dist = 100000
            for oc in occur:
                val = self.abs(idx - oc)
                if val > min_dist:
                    break
                min_dist = min(min_dist, val)
            ans[idx] = min_dist

        return ans

    def abs(self, a):
        if a < 0:
            return a * -1
        return a