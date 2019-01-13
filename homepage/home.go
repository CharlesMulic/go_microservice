package homepage

import (
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

const message = "Hello World"

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	//h.db.ExecContext(r.Context(), "") // you'll be able to cancel request if using context
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message)) // ignoring returned int,error
}

// middleware example
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		time.Now()
		defer h.logger.Printf("Request Processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
		//h.logger.Printf("Request Processed in %s\n", time.Now().Sub(startTime))
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db: db,
	}
}
