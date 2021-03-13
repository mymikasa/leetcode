class Solution:
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        init_list = s.split(' ')
        ans = ""
        end_list = []

        for i in init_list:
            ans += i[::-1]
            ans += " "

        ans = ans[0:-1]
        return ans