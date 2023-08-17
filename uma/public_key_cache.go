package uma

import "crypto/rsa"

// PublicKeyCache is an interface for a cache of public keys for other VASPs.
//
// Implementations of this interface should be thread-safe.
type PublicKeyCache interface {
	// FetchPublicKeyForVasp fetches the public key for a VASP if in the cache, otherwise returns nil.
	FetchPublicKeyForVasp(vaspDomain string) *rsa.PublicKey

	// AddPublicKeyForVasp adds a public key for a VASP to the cache.
	AddPublicKeyForVasp(vaspDomain string, pubKey *rsa.PublicKey)

	// RemovePublicKeyForVasp removes a public key for a VASP from the cache.
	RemovePublicKeyForVasp(vaspDomain string)

	// Clear clears the cache.
	Clear()
}

type InMemoryPublicKeyCache struct {
	cache map[string]*rsa.PublicKey
}

func NewInMemoryPublicKeyCache() *InMemoryPublicKeyCache {
	return &InMemoryPublicKeyCache{
		cache: make(map[string]*rsa.PublicKey),
	}
}

func (c *InMemoryPublicKeyCache) FetchPublicKeyForVasp(vaspDomain string) *rsa.PublicKey {
	return c.cache[vaspDomain]
}

func (c *InMemoryPublicKeyCache) AddPublicKeyForVasp(vaspDomain string, pubKey *rsa.PublicKey) {
	c.cache[vaspDomain] = pubKey
}

func (c *InMemoryPublicKeyCache) RemovePublicKeyForVasp(vaspDomain string) {
	delete(c.cache, vaspDomain)
}

func (c *InMemoryPublicKeyCache) Clear() {
	c.cache = make(map[string]*rsa.PublicKey)
}
