package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/admin", allowRoles("admin")(http.HandlerFunc(adminHandler)))
	http.Handle("/staff", allowRoles("staff")(http.HandlerFunc(staffHandler)))
	http.Handle("/admin-staff", allowRoles("admin", "staff")(http.HandlerFunc(adminStaffHandler)))

	err := http.ListenAndServe(":8080", nil)
	log.Println(err)
}

type middleware func(http.Handler) http.Handler

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func allowRoles(roles ...string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			for _, role := range roles {
				if reqRole == role {
					h.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}

func chain(hs ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs) - 1; i >= 0; i-- {
			h = hs[i](h)
		}
		return h
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Admin"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello staff"))
}

func adminStaffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello admin or staff"))
}
