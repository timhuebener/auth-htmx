package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Darkness4/auth-htmx/database/dataentry"
	"github.com/Darkness4/auth-htmx/jwt"
)

func CreateDataentry(repo *dataentry.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		_ = claims
		dpStr := r.FormValue("datapoint_id")
		typ := r.FormValue("type")
		text := r.FormValue("text_value")
		intStr := r.FormValue("int_value")
		if dpStr == "" || typ == "" {
			http.Error(w, "datapoint_id and type required", http.StatusBadRequest)
			return
		}
		dpID, err := strconv.ParseInt(dpStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid datapoint_id", http.StatusBadRequest)
			return
		}
		var textVal sql.NullString
		if text != "" {
			textVal = sql.NullString{String: text, Valid: true}
		}
		var intVal sql.NullInt64
		if intStr != "" {
			v, err := strconv.ParseInt(intStr, 10, 64)
			if err != nil {
				http.Error(w, "invalid int_value", http.StatusBadRequest)
				return
			}
			intVal = sql.NullInt64{Int64: v, Valid: true}
		}
		item, err := repo.Create(r.Context(), dpID, typ, textVal, intVal)
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

func GetDataentry(repo *dataentry.Repository) http.HandlerFunc {
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
		item, err := repo.Get(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintf(w, "%d\t%d\t%s", item.ID, item.DatapointID, item.Type); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListDataentriesByDatapoint(repo *dataentry.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		_ = claims
		dpStr := r.FormValue("datapoint_id")
		if dpStr == "" {
			http.Error(w, "datapoint_id required", http.StatusBadRequest)
			return
		}
		dpID, err := strconv.ParseInt(dpStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid datapoint_id", http.StatusBadRequest)
			return
		}
		items, err := repo.ListByDatapoint(r.Context(), dpID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i, it := range items {
			if i > 0 {
				if _, err := fmt.Fprint(w, "\n"); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			if _, err := fmt.Fprintf(w, "%d\t%s", it.ID, it.Type); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func UpdateDataentry(repo *dataentry.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := jwt.GetClaimsFromRequest(r)
		if !ok {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		_ = claims
		idStr := r.FormValue("id")
		typ := r.FormValue("type")
		text := r.FormValue("text_value")
		intStr := r.FormValue("int_value")
		if idStr == "" || typ == "" {
			http.Error(w, "id and type required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		var textVal sql.NullString
		if text != "" {
			textVal = sql.NullString{String: text, Valid: true}
		}
		var intVal sql.NullInt64
		if intStr != "" {
			v, err := strconv.ParseInt(intStr, 10, 64)
			if err != nil {
				http.Error(w, "invalid int_value", http.StatusBadRequest)
				return
			}
			intVal = sql.NullInt64{Int64: v, Valid: true}
		}
		item, err := repo.Update(r.Context(), id, typ, textVal, intVal)
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

func DeleteDataentry(repo *dataentry.Repository) http.HandlerFunc {
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
