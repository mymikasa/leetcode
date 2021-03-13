class Solution:
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        for i in matrix:
            if len(i) == 0:
                break
            if i[0] < target and i[-1] > target:
                temp = set(i)
                if target in temp:
                    return True
                return False

