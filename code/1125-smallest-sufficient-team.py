from typing import List
from functools import lru_cache


class Solution:
    def smallestSufficientTeam(
        self, req_skills: List[str], people: List[List[str]]
    ) -> List[int]:

        masks = {s: (1 << i) for i, s in enumerate(req_skills)}

        @lru_cache(None)
        def dfs(idx, skills):
            if skills == 0:
                return []
            if idx >= len(people):
                return None
            res = dfs(idx + 1, skills)
            cur_skills = 0
            for s in people[idx]:
                cur_skills |= masks.get(s, 0)
            if cur_skills & skills == skills:
                return [idx]
            if cur_skills & skills != 0:
                rest = dfs(idx + 1, skills & (~cur_skills))
                if rest is not None and (
                    res is None or len(rest) + 1 < len(res)
                ):
                    res = [idx] + rest
            return res

        skills = 0
        for s in req_skills:
            skills |= masks.get(s, 0)
        return dfs(0, skills)


if __name__ == "__main__":
    # req_skills = ["java", "nodejs", "reactjs"]
    # people = [["java"], ["nodejs"], ["nodejs", "reactjs"]]
    req_skills = ["algorithms", "math", "java", "reactjs", "csharp", "aws"]
    people = [
        ["algorithms", "math", "java"],
        ["algorithms", "math", "reactjs"],
        ["java", "csharp", "aws"],
        ["reactjs", "csharp"],
        ["csharp", "math"],
        ["aws", "java"],
    ]
    print(Solution().smallestSufficientTeam(req_skills, people))
