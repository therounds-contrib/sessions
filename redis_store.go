package sessions

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
)

// RedisStore is an interface that represents a Cookie based storage
// for Sessions.
type RediStore interface {
	// Store is an embedded interface so that RedisStore can be used
	// as a session store.
	Store
	// Options sets the default options for each session stored in this
	// CookieStore.
	Options(Options)
}

// NewCookieStore returns a new CookieStore.
//
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewRediStore(size int, network, address, password string, keyPairs ...[]byte) (RediStore, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &rediStore{store}, nil
}

type rediStore struct {
	*redistore.RediStore
}

func (c *rediStore) Options(options Options) {
	c.RediStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
		SameSite: options.SameSite,
	}
}
