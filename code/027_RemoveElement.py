class Solution:
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        i = 0
        n = len(nums)
        j = 0
        while(j < n):
            if j >= n:
                break
            if nums[j] != val:
                nums[i] = nums[j]
                i += 1
            j += 1
        return i

A = Solution()
A.removeElement([0,1,2,2,3,0,4,2], val = 2)