# 4. Median of Two Sorted Arrays

class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        m = len(nums1)
        n = len(nums2)

        # 令nums1数组始终为较短的那一个
        if m > n:
            nums1, nums2, m, n = nums2, nums1, n, m

        begin, end, half = 0, m, int((m + n + 1) / 2)
        while begin <= end:
            i = int((begin + end) / 2)
            j = half - i
            # 此时i的值过小，增大i的值
            if i < m and nums2[j - 1] > nums1[i]:
                begin = i + 1
            # i的值过达， 减小i
            elif i > 0 and nums2[j] < nums1[i - 1]:
                end = i - 1
            # i的值合法，查找中位数
            else:
                # 此时共有以下三种情况
                # 1. 当i位于本组最左边， 或者j位于本组最左边
                # 2. i位于本组最右边， 或者j位于本组最右边
                # 3. i位于本组中间， j位于本组中间
                if i == 0:
                    max_of_left = nums2[j - 1]
                elif j == 0:
                    max_of_left = nums1[i - 1]
                else:
                    max_of_left = max(nums1[i - 1], nums2[j - 1])

                if (m + n) % 2 == 1:
                    return max_of_left

                if i == m:
                    min_of_right = nums2[j]
                elif j == n:
                    min_of_right = nums1[i]
                else:
                    min_of_right = min(nums1[i], nums2[j])

                return (max_of_left + min_of_right) / 2.0

A = Solution()
a = [1,3]
b = [2]

print(A.findMedianSortedArrays(a,b))