class Solution:
    def validIPAddress(self, IP: str) -> str:
        if IP.find(".") != -1:
            if self.check_ipv4(IP):
                return "IPv4"
        if IP.find(":") != -1:
            if self.check_ipv6(IP):
                return "IPv6"
        return "Neither"

    def check_ipv4(self, ip):
        arr = ip.split(".")
        if len(arr) != 4:
            return False
        for x in arr:
            try:
                n = int(x)
            except:
                return False
            if not (0 <= n <= 255):
                return False
            if (n == 0 and len(x) > 1) or (n != 0 and x[0] == "0"):
                return False
        return True

    def check_ipv6(self, ip):
        arr = ip.split(":")
        if len(arr) != 8:
            return False
        for x in arr:
            if not (1 <= len(x) <= 4):
                return False
            for c in x.lower():
                if not ("0" <= c <= "9" or "a" <= c <= "f"):
                    return False
        return True


if __name__ == "__main__":
    IP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
    sol = Solution()
    print(sol.validIPAddress(IP))
