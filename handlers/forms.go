package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/y29-h/Formgram/db"
)

// generateID генерирует случайный уникальный ID для формы
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// CreateFormHandler создаёт новую форму для пользователя
func CreateFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Берём ID пользователя из куки
	cookie, err := r.Cookie("session_user_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userID, _ := strconv.Atoi(cookie.Value)

	type request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, "name and email are required", http.StatusBadRequest)
		return
	}

	id := generateID()

	_, err = db.DB.Exec(
		"INSERT INTO forms (id, user_id, name, email) VALUES (?, ?, ?, ?)",
		id, userID, req.Name, req.Email,
	)
	if err != nil {
		http.Error(w, "failed to create form", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id":       id,
		"endpoint": "/f/" + id,
	})
}

// ListFormsHandler возвращает список форм пользователя
func ListFormsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("session_user_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	userID, _ := strconv.Atoi(cookie.Value)

	rows, err := db.DB.Query(
		"SELECT id, name, email, created_at FROM forms WHERE user_id = ?", userID,
	)
	if err != nil {
		http.Error(w, "failed to fetch forms", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type form struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}

	var forms []form
	for rows.Next() {
		var f form
		rows.Scan(&f.ID, &f.Name, &f.Email, &f.CreatedAt)
		forms = append(forms, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forms)
}
