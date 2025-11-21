package middleware

import (
	pkgLog "learner-model-service/pkg/log"
	pkgScope "learner-model-service/pkg/scope"
)

type Middleware struct {
	l          pkgLog.Logger
	jwtManager pkgScope.Manager
}

func New(l pkgLog.Logger, jwtManager pkgScope.Manager) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
	}
}
