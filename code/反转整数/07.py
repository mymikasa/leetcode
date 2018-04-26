class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        i = 2147483648
        res = 0
        cnt = False
        if x < 0:
            x = -x
            cnt = True
        while True:
            Remainder = x % 10
            res = Remainder + res * 10
            x = int(x / 10)
            if x < 1 and x > -1:
                break
        if cnt:
            res = -res
        return res if res > 0 - i and res < i-1 else 0
