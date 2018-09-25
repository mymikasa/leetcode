class Solution:
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        a = [1 for i in range(n + 5)]
        a[1] = 1
        a[2] = 2

        for i in range(3, n + 1):
            a[i] = a[i - 1] + a[i - 2]

        return a[n]




