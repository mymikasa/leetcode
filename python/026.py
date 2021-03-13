class Solution:
    def removeDuplicates(self, a):

        """
        :type nums: List[int]
        :rtype: int
        """
        i = 1
        j = 0
        while (i < len(a)):
            if a[i] == a[j]:
                del a[i]
                # i += 2
            else:
                j = i
                i += 1
        print(len(a))