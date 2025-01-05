class Solution:
    def distributeCandies(self, candyType: List[int]) -> int:
        dict = {}
        for val in candyType:
            if val in dict:
                dict[val] += 1
            else:
                dict[val] = 1

        uniq = len(candyType) / 2

        # print(uniq)
        result = 0
        sorted_val = sorted(dict.items(), key=lambda x: x[1])
        while True or uniq == result:
            if len(sorted_val) == 0:
                if uniq < result:
                    return int(uniq)
                return result
            sorted_val.remove(sorted_val[0])
            result += 1
