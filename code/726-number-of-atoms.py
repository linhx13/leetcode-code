from collections import Counter


class Solution:
    def countOfAtoms(self, formula: str) -> str:
        def parse(start, end) -> Counter:
            counter = Counter()
            if end <= start:
                return counter
            idx = start
            while idx < end:
                if formula[idx].isupper():
                    atom_start = idx
                    idx += 1
                    while idx < end and formula[idx].islower():
                        idx += 1
                    atom = formula[atom_start:idx]
                    cnt_start = idx
                    while idx < end and formula[idx].isdigit():
                        idx += 1
                    if cnt_start == idx:
                        cnt = 1
                    else:
                        cnt = int(formula[cnt_start:idx])
                    counter[atom] += cnt
                elif formula[idx] == "(":
                    stack = ["("]
                    for j in range(idx + 1, end):
                        if formula[j] == "(":
                            stack.append("(")
                        elif formula[j] == ")":
                            stack.pop()
                            if len(stack) == 0:
                                break
                    cur = parse(idx + 1, j)
                    cnt_start = j + 1
                    j += 1
                    while j < end and formula[j].isdigit():
                        j += 1
                    if cnt_start == j:
                        cnt = 1
                    else:
                        cnt = int(formula[cnt_start:j])
                    for k, v in cur.items():
                        counter[k] += v * cnt
                    idx = j
                else:
                    idx += 1
            return counter

        counter = parse(0, len(formula))
        res = []
        for k in sorted(counter.keys()):
            res.append(k)
            v = counter[k]
            if v > 1:
                res.append(str(v))
        return "".join(res)


if __name__ == "__main__":
    # formula = "H2O"
    # formula = "Mg(OH)2"
    # formula = "K4(ON(SO3)2)2"
    formula = "((N42)24(OB40Li30CHe3O48LiNN26)33(C12Li48N30H13HBe31)21(BHN30Li26BCBe47N40)15(H5)16)14"
    print(Solution().countOfAtoms(formula))
