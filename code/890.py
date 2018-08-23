class Solution:
    def findAndReplacePattern(self, words, pattern):
        """
        :type words: List[str]
        :type pattern: str
        :rtype: List[str]
        """

        """
        lp = len(pattern)
        flag = [-1 for i in range(lp)]
        unsame = list()
        for i in range(lp):
            if flag[i] == -1:
                temp = True
                for j in range(i, lp):
                    if pattern[i] == pattern[j]:
                        flag[j] = i
                        if temp:
                            unsame.append(j)
                            temp = False
        ans = list()
        for i in words:
            strr = str()
            for j in flag:
                strr += i[j]
            if strr == i:
                tt = list()
                for j in unsame:
                    tt.append(i[j])
                if len(tt) == len(set(tt)):
                    ans.append(strr)
        return ans
        """

        ans = list()
        l = len(pattern)

        for word in words:
            flag = True
            for i in range(l):
                for j in range(i+1, l):
                    if (word[i] == word[j]) != (pattern[i] == pattern[j]):
                        flag = False
                        break
                if not flag:
                    break

            if flag:
                ans.append(word)
        return ans


A = Solution()
words = ["abc","deq","mee","aqq","dkd","ccc"]
pattern = "abb"
print(A.findAndReplacePattern(words, pattern))