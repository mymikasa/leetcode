class Solution:
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        n = len(pattern)
        str_list = str.split(" ")
        if len(str_list) != n:
            return False
        for i in range(n):
            for j in range(i, n):
                if (pattern[i] == pattern[j]) != (str_list[i] == str_list[j]):
                    return False
        return True

print(Solution().wordPattern('aaaa', 'cat cat cat cat'))

# return len(set(zip(s,t))) == len(set(s)) == len(set(t))