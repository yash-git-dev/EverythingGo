package Controllers

// import (
// 	"crypto/rand"
// 	"net/http"
// 	"time"
// )

// var urls map[string]string

// func main() {
// 	http.HandleFunc("/", redirectToOriginalURL)
// 	http.HandleFunc("/shortenURL" shortenURL)
// 	http.ListenAndServe(":8080", nil)
// }

// func redirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
// 	shortURL := r.URL.Path[1:]
// 	if originalURL, ok := urls[shortURL]; ok {
// 		http.Redirect(w, r, originalURL, http.StatusOK)
// 	} else {
// 		http.NotFound(w,r)
// 	}
// }

// func shortenURL(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method not allowed", http.StatusBadRequest)
// 		return
// 	}

// 	originalURL := r.FormValue("url")
// 	shortURL := generateShortURL(originalURL)
// 	urls[shortURL] = originalURL

// }

// func generateShortURL(originalURL string) string {

// 	var shortURL string
// 	u := uuid.New()
// 	shortURL = "http://shorturl/" + u.String()

// 	return shortURL
// }
