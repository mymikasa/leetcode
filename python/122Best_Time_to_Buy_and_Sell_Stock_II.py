class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if not len(prices):
            return 0
        start = prices[0]
        ans = 0
        for i in range(len(prices)):
            if prices[i] > start:
                ans += prices[i] - start
            start = prices[i]
        return ans