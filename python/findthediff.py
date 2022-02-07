class Solution(object):
    def findTheDifference(self, s, t):
        putmap = {}
        for x in t:
            if x in putmap:
                putmap[x] += 1
            else:
                putmap[x] = 1
        
        
        for x in s:
                putmap[x] -= 1
        
        
        for y in putmap:
            if putmap[y] == 1:
                return y
