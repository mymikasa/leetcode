class Solution:
    def search(self, nums, target):
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
            if self.judge(nums, mid):
                break
            else:
                if nums[mid] > nums[-1]:
                    begin = mid + 1
                else:
                    end = mid - 1

        if self.dichotomy(0, mid - 1,nums, target) != -1:
            return self.dichotomy(0, mid - 1,nums, target)

        if self.dichotomy(mid, len(nums) - 1, nums, target) != -1:
            return self.dichotomy(mid, len(nums) - 1, nums, target)
        return -1


    #
    def judge(self, nums, flag):
        return True if nums[flag - 1] > nums[flag] and nums[flag] < nums[flag] + 1 else False

    # dichotomy
    def dichotomy(self, begin, end, nums, target):
        while(begin <= end):
            mid = int((begin + end) / 2)
            if nums[mid] < target:
                begin = mid + 1
            elif nums[mid] > target:
                end = mid - 1
            else:
                return mid
        return -1


A = Solution()
nums = [0,1,2]
target = 2
print(A.search(nums, target))