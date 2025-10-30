package rfinger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/datasektionen/dock/pkg/config"
	"github.com/datasektionen/dock/pkg/dao"
)

func Listen(cfg *config.Config, dao *dao.Dao) {
	db := dao.Db.Rfinger

	h := http.NewServeMux()

	h.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("missing.svg")

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(content)
	});

	h.HandleFunc("GET /api/{kthid}", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		query := r.URL.Query()

		id := r.PathValue("kthid")
		quality := query["quality"]

		path := db.Default

		if len(quality) > 0 && quality[0] == "true" {
			path = db.Pictures[id].Regular
		} else {
			path = db.Pictures[id].Small
		}

		if path == "" {
			path = db.Default
		}

		fmt.Fprint(w, path)
	})

	h.HandleFunc("GET /api/batch", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var data []string

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		query := r.URL.Query()

		quality := query["quality"]

		resp := make(map[string]string)

		for _, kthid := range data {
			if len(quality) > 0 && quality[0] == "true" {
				resp[kthid] = db.Pictures[kthid].Regular
			} else {
				resp[kthid] = db.Pictures[kthid].Small
			}

			if resp[kthid] == "" {
				resp[kthid] = db.Default
			}
		}

		json.NewEncoder(w).Encode(resp)
	})

	h.HandleFunc("POST /api/{kthid}", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		image := r.FormValue("image")

		if image == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	h.HandleFunc("POST /api/nollan/{kthid}", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		image := r.FormValue("image")

		if image == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Printf("rfinger listening on http://localhost:%s\n", cfg.RfingerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.RfingerPort), h); err != nil {
		panic(err)
	}
}
