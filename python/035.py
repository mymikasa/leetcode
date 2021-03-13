class Solution:
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        begin = 0
        end = len(nums) - 1
        mid = 0
        while(begin <= end):
            mid = int((begin + end) / 2)
            if nums[mid] > target:
                end = mid - 1
            elif nums[mid] < target:
                begin = mid + 1
            else:
                return mid

        if nums[mid] > target:
            return mid
        else:
            return mid + 1

A = Solution()
print(A.searchInsert([1,3], -2))

