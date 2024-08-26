package api

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
	"waves/model"
	"waves/utilities"
)

type RegisterFieldsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

const errorString = "ERR"
const okString = "OK"

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()
	filePath, err := utilities.SaveProfilePicture(file)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	user, err := model.NewUser(username, password, email, filePath)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User registered successfully: %+v", user)
}

func isValidEmail(email string) bool {
	// Regular expression for basic email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	hasDigit := false
	hasSpecialChar := false
	for _, ch := range password {
		if ch >= '0' && ch <= '9' {
			hasDigit = true
		} else if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')) {
			hasSpecialChar = true
		}
	}
	return hasDigit && hasSpecialChar
}

func isValidImageFile(header *multipart.FileHeader) bool {
	// Get the file extension
	ext := strings.ToLower(strings.TrimPrefix(header.Filename[strings.LastIndex(header.Filename, "."):], "."))
	allowedExts := map[string]bool{"png": true, "jpg": true, "jpeg": true, "webp": true}

	// Check the extension
	if !allowedExts[ext] {
		return false
	}

	// Check the MIME type (simple check based on file header)
	contentType := header.Header.Get("Content-Type")
	return (contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/webp")
}

func CheckRegisterFields(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	username := r.FormValue("username")
	if utf8.RuneCountInString(username) > 30 {
		if err := encoder.Encode(RegisterFieldsResponse{Status: errorString, Message: "Username must be shorter than 30 characters", Field: "username"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	password := r.FormValue("password")
	if !isValidPassword(password) {
		if err := encoder.Encode(RegisterFieldsResponse{Status: errorString, Message: "Password must contain numbers and special characters", Field: "password"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	email := r.FormValue("email")
	if !isValidEmail(email) {
		if err := encoder.Encode(RegisterFieldsResponse{Status: errorString, Message: "Invalid email address", Field: "email"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	_, header, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}

	if !isValidImageFile(header) {
		if err := encoder.Encode(RegisterFieldsResponse{Status: errorString, Message: "Invalid file type. Only PNG, JPG, JPEG, and WEBP files are allowed.", Field: "profile_picture"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if err := encoder.Encode(RegisterFieldsResponse{Status: okString, Message: "All fields are okay"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
