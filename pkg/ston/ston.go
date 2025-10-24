package ston

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datasektionen/dock/pkg/config"
	"github.com/datasektionen/dock/pkg/dao"
)

func Listen(cfg *config.Config, dao *dao.Dao) {
	db := dao.Db.Ston

	h := http.NewServeMux()

	h.HandleFunc("GET /api/pax", func(w http.ResponseWriter, r *http.Request) {
		if !r.URL.Query().Has("api_key") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(db.Nollan)
	})

	fmt.Printf("ston listening on http://localhost:%s\n", cfg.SpamPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.SpamPort), h); err != nil {
		panic(err)
	}
}
