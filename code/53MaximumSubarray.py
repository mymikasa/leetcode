class Solution:
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        maxx = nums[0]
        current_sum = 0
        for i in range(len(nums)):
            if current_sum <= 0:
                current_sum = nums[i]
            else:
                current_sum += nums[i]
            if current_sum > maxx:
                maxx = current_sum
        return maxx