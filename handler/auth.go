package handler

import (
	"cloudstorage/v1/db"
	"log"
	"net/http"
)

// TokenHandler 权限拦截器
func TokenHandler(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			username := r.FormValue("username")
			token := r.FormValue("token")
			if len(username) < 3 {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			// 校验token
			ok := db.IsValidateToken(username, token)
			if !ok {
				log.Println("token validate failed")
				w.WriteHeader(http.StatusForbidden)
				return
			}

			h(w, r)
		})
}
