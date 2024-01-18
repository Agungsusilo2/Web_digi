package routes

import (
	"net/http"
)

func main() {

	// Buat router
	router := http.NewServeMux()

	// Tambahkan middleware untuk script
	router.HandleFunc("/script/tailwind.js", func(w http.ResponseWriter, r *http.Request) {
		// Arahkan permintaan ke URL script
		http.Redirect(w, r, "https://cdn.tailwindcss.com", http.StatusFound)
	})

	// Tambahkan middleware untuk style
	router.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		// Arahkan permintaan ke URL style
		http.Redirect(w, r, "view/css/style.css", http.StatusFound)
	})

}
