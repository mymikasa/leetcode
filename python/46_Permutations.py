class Solution:
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        import itertools
        r = itertools.permutations(nums , len(nums))
        ans = []
        for i in r:
            ans.append(i)
        return ans