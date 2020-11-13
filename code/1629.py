from typing import List


class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        max_time = releaseTimes[0]
        max_key = keysPressed[0]
        for i in range(1, len(keysPressed)):
            cur_time = releaseTimes[i] - releaseTimes[i - 1]
            cur_key = keysPressed[i]
            if cur_time > max_time \
               or (cur_time == max_time and cur_key > max_key):
                max_time = cur_time
                max_key = cur_key
        return max_key
