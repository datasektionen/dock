package spam

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "strings"

	"github.com/datasektionen/dock/pkg/config"
	"github.com/datasektionen/dock/pkg/dao"
)

func Listen(cfg *config.Config, dao *dao.Dao) {
	// db := dao.Db.Rfinger

	h := http.NewServeMux()

	// h.HandleFunc("GET /api/{kthid}", func(w http.ResponseWriter, r *http.Request) {
	//
	// }

	fmt.Printf("rfinger listening on http://localhost:%s\n", cfg.SpamPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.SpamPort), h); err != nil {
		panic(err)
	}
}
