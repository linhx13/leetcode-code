from collections import deque


class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        queue = deque()
        people = [0, 0]
        ban = [0, 0]

        for p in senate:
            x = int(p == 'R')
            people[x] += 1
            queue.append(x)

        while all(people):
            x = queue.popleft()
            if ban[x] > 0:
                ban[x] -= 1
                people[x] -= 1
            else:
                ban[1 - x] += 1
                queue.append(x)
        return "Radiant" if people[1] else "Dire"
