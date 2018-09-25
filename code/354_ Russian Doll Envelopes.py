class Solution:
    def maxEnvelopes(self, envelopes):
        """
        :type envelopes: List[List[int]]
        :rtype: int
        """
        envelopes.sort(key = lambda x:x[1])
        res = 1
        start = envelopes[0]
        for i in envelopes:
            if i[0] > start[0] and i[1] > start[1]:
                res += 1
                start = i
        return res


Solution().maxEnvelopes([[5,4],[6,4],[6,7],[2,3]])