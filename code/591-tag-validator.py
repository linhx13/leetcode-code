class Solution:
    def isValid(self, code: str) -> bool:
        stack = []
        contains_tag = False

        def is_valid_tag(s, ending):
            nonlocal stack, contains_tag
            if not (1 <= len(s) <= 9):
                return False
            for ch in s:
                if not ch.isupper():
                    return False
            if ending:
                return stack and stack[-1] == s
            return True

        def is_valid_cdata(s: str):
            return s.find("[CDATA[") == 0

        if code[0] != "<" or code[-1] != ">":
            return False

        idx = 0
        while idx < len(code):
            ending = False
            close_idx = 0
            if len(stack) == 0 and contains_tag:
                return False
            if code[idx] == "<":
                if len(stack) != 0 and code[idx + 1] == "!":
                    close_idx = code.find("]]>", idx + 1)
                    if close_idx < 0 or not is_valid_cdata(
                        code[idx + 2 : close_idx]
                    ):
                        return False
                else:
                    if code[idx + 1] == "/":
                        ending = True
                        idx += 1
                    close_idx = code.find(">", idx + 1)
                    tag_name = code[idx + 1 : close_idx]
                    if close_idx < 0 or not is_valid_tag(tag_name, ending):
                        return False
                    else:
                        if ending:
                            stack.pop()
                        else:
                            stack.append(tag_name)
                            contains_tag = True
                idx = close_idx
            idx += 1
        return len(stack) == 0 and contains_tag
