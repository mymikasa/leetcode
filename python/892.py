class Solution:
    def surfaceArea(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """

        l1 = len(grid)
        l2 = len(grid[0])
        area1 = 0
        area2 = 0
        area3 = 0

        for i in range(l1):
            maxx = 0
            area2 += max(grid[i])
            for j in range(l2):
                area1 += 1 if (grid[i][j]) else 0
                maxx = max(maxx, grid[j][i])
            area3 += maxx

        return (area1*2 + area2*2 + area3*2)


print(Solution().surfaceArea([[2,2,2],[2,1,2],[2,2,2]]))