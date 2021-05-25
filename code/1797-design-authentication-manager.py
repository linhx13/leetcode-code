class AuthenticationManager:
    def __init__(self, timeToLive: int):
        self.ttl = timeToLive
        self.tokens = {}

    def generate(self, tokenId: str, currentTime: int) -> None:
        self.tokens[tokenId] = currentTime

    def renew(self, tokenId: str, currentTime: int) -> None:
        v = self.tokens.get(tokenId)
        if v is not None and v + self.ttl > currentTime:
            self.tokens[tokenId] = currentTime

    def countUnexpiredTokens(self, currentTime: int) -> int:
        self.tokens = {
            k: v for k, v in self.tokens.items() if v + self.ttl > currentTime
        }
        return len(self.tokens)
