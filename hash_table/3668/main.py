from typing import List

class Solution:
    def recoverOrder(self, order: List[int], friends: List[int]) -> List[int]:
        friends_dict = {friend: None for friend in friends}
        ans = []


        for i in range(len(order)):
            found = False
            if order[i] in friends_dict:
                ans.append(order[i])
                friends_dict.pop(order[i])
                found = True
            if found:
                i = 0
        return ans


def main():

    print(Solution().recoverOrder([3,1,2,5,4], [1,3,4]))



if __name__ == '__main__':
    main()
