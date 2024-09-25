package middleware

import (
	"NestJsStyle/data"
	"NestJsStyle/helper"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"net/http"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			helper.WriteResponseBody(w, &data.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Authorization token is missing!",
				Data:   nil,
			}, http.StatusUnauthorized)
			return
		}
		claims, err := helper.ValidateToken(bearerToken)
		if err != nil {
			helper.WriteResponseBody(w, &data.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: err.Error(),
				Data:   nil,
			}, http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), "userId", claims.Id)
		next(w, r.WithContext(ctx), params)
	}
}
