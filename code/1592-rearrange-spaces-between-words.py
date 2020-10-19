class Solution:
    def reorderSpaces(self, text: str) -> str:
        len_all = len(text)
        words = text.strip().split()
        len_words = sum(len(w) for w in words)
        len_spaces = len_all - len_words
        if len(words) == 1:
            return words[0] + (" " * len_spaces)
        sep = " " * (len_spaces // (len(words) - 1))
        res = [words[0]]
        c = 0
        for i in range(1, len(words)):
            res.append(sep)
            res.append(words[i])
            c += len(sep)
        if len_spaces % (len(words) - 1) != 0:
            res.append(" " * (len_spaces - c))
        return "".join(res)
