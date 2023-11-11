package main

func main() {

}

type AuthenticationManager struct {
	store map[string]int
	ttl   int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		store: make(map[string]int),
		ttl:   timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.store[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if val, ok := this.store[tokenId]; ok && val > currentTime {
		this.Generate(tokenId, currentTime)
		return
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var result int
	for _, v := range this.store {
		if v > currentTime {
			result++
		}
	}

	return result
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
