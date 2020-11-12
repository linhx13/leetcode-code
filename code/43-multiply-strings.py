class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        res = "0"
        base = 0
        for i in range(len(num2) - 1, -1, -1):
            cur = self.single_mul(num1, num2[i])
            cur += '0' * base
            base += 1
            res = self.add(res, cur)
        return res

    def single_mul(self, x, y):
        if x == '0' or y == '0':
            return '0'
        x = x[::-1]
        res = []
        carry = 0
        y = ord(y[0]) - ord('0')
        for i in range(len(x)):
            cur = (ord(x[i]) - ord('0')) * y + carry
            res.append(chr(cur % 10 + ord('0')))
            carry = cur // 10
        if carry != 0:
            res.append(chr(carry + ord('0')))
        res.reverse()
        return "".join(res)

    def add(self, x, y):
        if x == "0":
            return y
        if y == "0":
            return x
        if len(x) < len(y):
            x, y = y, x
        res = []
        carry = 0
        x = x[::-1]
        y = y[::-1]
        i = 0
        while i < len(y):
            cur = ord(x[i]) - ord('0') + ord(y[i]) - ord('0') + carry
            res.append(chr(cur % 10 + ord('0')))
            carry = cur // 10
            i += 1
        while i < len(x):
            cur = ord(x[i]) - ord('0') + carry
            res.append(chr(cur % 10 + ord('0')))
            carry = cur // 10
            i += 1
        if carry != 0:
            res.append(chr(carry + ord('0')))
        res.reverse()
        return "".join(res)


if __name__ == "__main__":
    num1 = "999999999"
    num2 = "999999999"
    print(Solution().multiply(num1, num2))
