'''
直接模拟会超时，所以根据变换顺序找每行字母的规律
'''

class Solution:

    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1 or not s:
            return s
        res = ''
        cnt = 0
        while cnt<len(s):
            res += s[cnt]
            cnt += 2 * numRows - 2
        for i in range(1,numRows-1):
            cnt = i
            while cnt < len(s):
                #print(s[cnt])
                res += s[cnt]
                sss = (numRows - i - 1)*2 -1
                cnt = cnt + sss + 1
                if not cnt < len(s):
                    break
                #print(s[cnt])
                res += s[cnt]
                cnt += 2 * numRows - 2 - sss - 1
        cnt = numRows - 1
        while cnt<len(s):
            res += s[cnt]
            cnt += 2 * numRows - 2
        return (res)
A = Solution()
ss = "PAYPALISHIRING"
n = 3
print(A.convert(ss,n))

