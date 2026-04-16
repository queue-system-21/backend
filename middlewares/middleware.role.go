package middlewares

import (
	"net/http"
	"queue/utils"
	"slices"
)

type roleMiddleware struct {
	next         http.Handler
	allowedRoles []string
}

func NewRole(roles []string, next http.Handler) http.Handler {
	return &roleMiddleware{
		next:         next,
		allowedRoles: roles,
	}
}

func (m *roleMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	if !slices.Contains(m.allowedRoles, role) {
		utils.SendErrMsg(w, "Forbidden", 403)
		return
	}
	m.next.ServeHTTP(w, r)
}
