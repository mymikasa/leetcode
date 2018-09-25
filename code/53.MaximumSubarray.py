class Solution:
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        maxx = nums[0]
        current_sum = nums[0]
        for i in range(1, len(nums)):
            if current_sum <= 0:
                current_sum = nums[i]
            else:
                current_sum += a[i]

            if current_sum > maxx:
                maxx = current_sum
        return maxx



