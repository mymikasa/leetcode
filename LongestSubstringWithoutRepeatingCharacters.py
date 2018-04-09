class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        """
        wa了无数遍,踩了无数个坑
        大致思路就是从字符串左边开始
        发现一个未出现的字母就将其作为字典的键加入到字典中去,它本身所在字符串的下标位置记为值
        用一个变量cnt来记录此刻已经有多少字母没有出现过
        如果发现相同字母出现了记录一下结果ans,取值为cnt 与 此字母与字典中存放的字母下标之差的最大值
        然后更新字典以及cnt
        """
        A = {}
        ans = 0
        cnt = 0
        temp = 0  #temp是用来记录  在查找过程中无重复字符串的起点(随着i的改变而改变)
        for i in range(len(s)):
            if A.get(s[i]) is None:
                A[s[i]] = i
                cnt +=1
            else:
                if A[s[i]] < temp:
                    ans = max(cnt,i - temp,ans)
                    cnt = i - temp
                    temp = max(A[s[i]],temp)
                    A[s[i]] = i
                else:
                    #temp = A[s[i]]
                    ans = max(ans,cnt)
                    ans = max(ans,i-A[s[i]])
                    cnt = i - A[s[i]]
                    temp = A[s[i]]
                    A[s[i]] = i
        return max(cnt,ans)
