from typing import List


class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        res = []
        line_words = []
        line_len = 0
        for word in words:
            l = line_len + len(word)
            if len(line_words) > 0:
                l += 1
            if l > maxWidth:
                res.append(self.build_line(maxWidth, line_words))
                line_words = []
                line_len = 0
            line_words.append(word)
            line_len += len(word)
            if len(line_words) > 1:
                line_len += 1
        if line_words:
            line = " ".join(line_words)
            line += " " * (maxWidth - len(line))
            res.append(line)
        return res

    def build_line(self, max_width, words):
        if len(words) == 1:
            return words[0] + (" " * (max_width - len(words[0])))
        words_len = sum(len(w) for w in words)
        spaces = max_width - words_len
        extra_spaces = spaces % (len(words) - 1)
        same_spaces = " " * ((spaces - extra_spaces) // (len(words) - 1))
        res = []
        for i in range(len(words) - 1):
            res.append(words[i])
            res.append(same_spaces)
            if extra_spaces > 0:
                res.append(" ")
                extra_spaces -= 1
        res.append(words[-1])
        return ''.join(res)


if __name__ == "__main__":
    # words = ["This", "is", "an", "example", "of", "text", "justification."]
    # maxWidth = 16
    # words = ["What", "must", "be", "acknowledgment", "shall", "be"]
    # maxWidth = 16
    # words = ["Listen", "to", "many,", "speak", "to", "a", "few."]
    # maxWidth = 6
    words = ["Imagination", "is", "more", "important", "than", "knowledge."]
    maxWidth = 14
    print(Solution().fullJustify(words, maxWidth))
