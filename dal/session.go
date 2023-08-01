package dal

import "github.com/mingkid/jtt808-gateway/domain"

var DefaultSessionCache domain.SessionCache = domain.NewMemorySessionCache()
