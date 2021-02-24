class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        first = binary.find("0")
        if first == -1:
            return binary
        cnt = binary.count("1", first)
        n = len(binary)
        return "1" * (n - cnt - 1) + "0" + "1" * cnt 
