package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Создаем новый запрос к google.com
		req, err := http.NewRequest(r.Method, "https://www.google.com"+r.URL.Path, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		// Копируем заголовки клиента
		req.Header = r.Header.Clone()

		// Создаем HTTP-клиент
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Копируем заголовки и тело обратно клиенту
		for k, v := range resp.Header {
			for _, vv := range v {
				w.Header().Add(k, vv)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	log.Println("Localhost redirector proxy running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
