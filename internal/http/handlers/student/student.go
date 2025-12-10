package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/abhijit1859/REST_API_GOLANG/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){

		var student types.Student

		err:=json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err,io.EOF{
			
		})
		slog.Info("Creating a student")
		w.Write([]byte("hello ji"))
	}
}