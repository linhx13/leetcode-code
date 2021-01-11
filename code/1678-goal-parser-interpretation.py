import re


class Solution:
    def interpret(self, command: str) -> str:
        command = re.sub("\(\)", "o", command)
        command = re.sub("\(al\)", "al", command)
        return command
