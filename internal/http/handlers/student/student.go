package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/abhijit1859/REST_API_GOLANG/internal/types"
	"github.com/abhijit1859/REST_API_GOLANG/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){

		var student types.Student

		err:=json.NewDecoder(r.Body).Decode(&student)
		
		if errors.Is(err,io.EOF){
			//returning json reponse
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return

		}

		//validating request
		//dont trust the client
		//jo chaiye tha whi mil rha na ?
		//for validating it will use go-playground/validator

		if err:=validator.New().Struct(student);err!=nil{
			validateErrs:=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return
		}



		slog.Info("creating a student")
		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"ok"})
	}
}