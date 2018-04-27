class Solution:
    def intt(self, str):
        res = 0
        for i in range(len(str)):
            res = res * 10 + int(str[i])
        return res

    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = str.strip()
        s = ''
        if len(str)==0:
            return 0
        if str[0] == '-' or (ord('0') <= ord(str[0]) and ord(str[0]) <= ord('9')) or str[0] =='+':
            s += str[0]
        if len(s) == 0:
            return 0
        for i in range(1,len(str)):
            if '0' <= str[i] <= '9':
                s += str[i]
            else:
                break

        if s[0] == '-':
            s = s[1:]
            if not s:
                return 0
            s = self.intt(s)
            if s > 2147483648:
                return -2147483648
            else:
                return -s
        elif s[0] == '+':
            s = s[1:]
            if not s:
                return 0
            s = self.intt(s)
            if s > 2147483648:
                return 2147483647
            else:
                return s
        else :
            s = self.intt(s)
            if s >= 2147483648:
                return 2147483647
            elif s < -2147483648:
                return -2147483648
            else:
                return s

