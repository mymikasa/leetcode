"""
定义一个二维数组dp[][],其中dp[j][i]表示从在字符串s中是[j]到s[i]是会问字符串


采用动态规划来解决此问题,复杂度为O(n^2)
其中状态转移方程为:
(1) 每一个单个字符都是回文,即dp[i][i] = 1
(2) 如果二者相差1且s[i] == s[j],则dp[j][i] = 1
(3) i,j相差大于1且s[i] == s[j] dp[j+1][i-1] ==1,则dp[j][i] == 1     这句话的意思是如果s[j]到s[i]是回文,则必须s[j] == s[i],并且去掉s[i] 和 s[j] 以后字符串仍然是回文串
"""
class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        ls = len(s)
        #构造一个二维数组dp[][]并给每一个数组都赋值为0
        dp = [[0 for i in range(ls)] for i in range(ls)]
        

        for i in range(ls):
            dp[i][i] = 1

        begin = 0
        max_len = 1

        for i in range(ls):
            for j in range(i):
                if (i - j) < 2:
                    if s[i] == s[j]:
                        dp[j][i] = 1
                    else:
                        dp[j][i] = 0
                else:
                    if s[i] == s[j] and dp[j + 1][i - 1] == 1:
                        dp[j][i] = 1
                    else:
                        dp[j][i] = 0

                if (dp[j][i] == 1 and max_len < (i - j +1)):
                    #print(j,i)
                    max_len = i - j + 1
                    begin = j

        return s[begin:begin + max_len]
