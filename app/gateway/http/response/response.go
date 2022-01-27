package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomResponse struct {
	Writer http.ResponseWriter
}

func (r CustomResponse) InternalError(body interface{}) {
	r.Writer.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(r.Writer).Encode(body); err != nil {
		log.Fatal(err)
	}
}

func (r CustomResponse) BadRequest(body interface{}) {
	r.Writer.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(r.Writer).Encode(body); err != nil {
		log.Fatal(err)
	}
}

func (r CustomResponse) Ok(body interface{}) {
	r.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(r.Writer).Encode(body); err != nil {
		log.Fatal(err)
	}
}

func (r CustomResponse) Created(body interface{}) {
	r.Writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(r.Writer).Encode(body); err != nil {
		log.Fatal(err)
	}
}

func (r CustomResponse) Unauthorized(body interface{}) {
	r.Writer.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(r.Writer).Encode(body); err != nil {
		log.Fatal(err)
	}
}
