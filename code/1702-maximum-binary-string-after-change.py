class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        first = binary.find("0")
        if first == -1:
            return binary
        cnt = binary.count("1", first)
        n = len(binary)
<<<<<<< HEAD
        return "1" * (n - cnt - 1) + "0" + "1" * cnt 
=======
        return "1" * (n - cnt - 1) + "0" + "1" * cnt
>>>>>>> 7322206f0ae1ca425e04817e3fca5c4fc7e2333e
