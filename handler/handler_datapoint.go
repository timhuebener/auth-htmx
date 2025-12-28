package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Darkness4/auth-htmx/database/datapoint"
	"github.com/Darkness4/auth-htmx/jwt"
	"github.com/Darkness4/auth-htmx/security/csrf"
)

func CreateDatapoint(repo *datapoint.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		name := r.FormValue("name")
		if name == "" {
			http.Error(w, "name required", http.StatusBadRequest)
			return
		}
		item, err := repo.Create(r.Context(), []byte(claims.ID), name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintf(w, "%d", item.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListDatapoints(repo *datapoint.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		items, err := repo.ListByUser(r.Context(), []byte(claims.ID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl, err := template.ParseFiles("components/datapoints_list_item.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i, d := range items {
			if i > 0 {
				if _, err := fmt.Fprint(w, "\n"); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			if err := tpl.ExecuteTemplate(w, "DataPointsListItem", struct {
				ID        int64
				Name      string
				CSRFToken string
			}{
				ID:        d.ID,
				Name:      d.Name,
				CSRFToken: csrf.Token(r),
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func UpdateDatapointName(repo *datapoint.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		_ = claims
		idStr := r.FormValue("id")
		name := r.FormValue("name")
		if idStr == "" || name == "" {
			http.Error(w, "id and name required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		item, err := repo.UpdateName(r.Context(), id, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintf(w, "%d", item.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeleteDatapoint(repo *datapoint.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		_ = claims
		idStr := r.FormValue("id")
		if idStr == "" {
			http.Error(w, "id required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		if err := repo.Delete(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprint(w, "ok"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
