package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	max := new(big.Int).Sub(p, big.NewInt(2))
	r := new(big.Int).Rand(rand, max)
	return new(big.Int).Add(r, big.NewInt(2))
}

func PublicKey(priv, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), priv, p)
}

func NewPair(p *big.Int, g int64) (priv, pub *big.Int) {
	priv = PrivateKey(p)
	pub = PublicKey(priv, p, g)
	return priv, pub
}

func SecretKey(priv, pub, p *big.Int) *big.Int {
	return new(big.Int).Exp(pub, priv, p)
}
