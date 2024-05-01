package server

import (
	"gorepo/server/storage"
	"io"
	"log"
	"net/http"
)

func StartServer(address string, storagePath string) {
	storage := storage.NewStorage(storagePath)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path[len("/upload/"):]
		data, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := storage.SaveArtifact(path, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/upload/"):]
		data, err := storage.GetArtifact(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe(address, nil))
}
