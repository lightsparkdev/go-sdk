package uma

import "time"

// PublicKeyCache is an interface for a cache of public keys for other VASPs.
//
// Implementations of this interface should be thread-safe.
type PublicKeyCache interface {
	// FetchPublicKeyForVasp fetches the public key entry for a VASP if in the cache, otherwise returns nil.
	FetchPublicKeyForVasp(vaspDomain string) *PubKeyResponse

	// AddPublicKeyForVasp adds a public key entry for a VASP to the cache.
	AddPublicKeyForVasp(vaspDomain string, pubKey *PubKeyResponse)

	// RemovePublicKeyForVasp removes a public key for a VASP from the cache.
	RemovePublicKeyForVasp(vaspDomain string)

	// Clear clears the cache.
	Clear()
}

type InMemoryPublicKeyCache struct {
	cache map[string]*PubKeyResponse
}

func NewInMemoryPublicKeyCache() *InMemoryPublicKeyCache {
	return &InMemoryPublicKeyCache{
		cache: make(map[string]*PubKeyResponse),
	}
}

func (c *InMemoryPublicKeyCache) FetchPublicKeyForVasp(vaspDomain string) *PubKeyResponse {
	entry := c.cache[vaspDomain]
	if entry == nil || (entry.ExpirationTimestamp != nil && time.Unix(*entry.ExpirationTimestamp, 0).Before(time.Now())) {
		return nil
	}
	return entry
}

func (c *InMemoryPublicKeyCache) AddPublicKeyForVasp(vaspDomain string, pubKey *PubKeyResponse) {
	c.cache[vaspDomain] = pubKey
}

func (c *InMemoryPublicKeyCache) RemovePublicKeyForVasp(vaspDomain string) {
	delete(c.cache, vaspDomain)
}

func (c *InMemoryPublicKeyCache) Clear() {
	c.cache = make(map[string]*PubKeyResponse)
}
