# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findSecondMinimumValue(self, root: Optional[TreeNode]) -> int:
        a = {}

        def f(x):
            if x:
                a[x.val] = True
            if x.left:
                f(x.left)
            if x.right:
                f(x.right)

        f(root)

        keys = list(a.keys())
        keys.sort()
        # print(keys)
        if len(keys) <= 1:
            return -1

        return keys[1]