package middleware

import (
	"context"
	"database/sql"
	"errors"
	"example/controller"
	"example/helpers"
	"example/model"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Authentication(next httprouter.Handle, db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("MASUK AUTHENTICATION")
		// cek token nya ada apa ga?
		token := r.Header.Get("access-token")

		// token valid apa ga?
		claims, err := helpers.DecodeToken(token)
		if err != nil {
			controller.WriteJSON(w, http.StatusUnauthorized, controller.ErrorResponse{
				Message: "invalid token, please provide valid access token",
				Detail:  err,
			})
			return
		}

		// cek informasi dalam token nya beneran ada apa ga?
		userId := int(claims["userId"].(float64))
		loggedInUser, err := findUserById(userId, db)
		if err != nil {
			controller.WriteJSON(w, http.StatusUnauthorized, controller.ErrorResponse{
				Message: "invalid token, please provide valid access token",
				Detail:  err,
			})
			return
		}

		ctx := context.WithValue(r.Context(), "loggedInUser", loggedInUser)

		next(w, r.WithContext(ctx), p)
	}
}

func AuthorizeSuperAdmin(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("AUTHORIZATION")
		userVal := r.Context().Value("loggedInUser")
		loggedInUser := userVal.(model.User)

		if loggedInUser.Role != "superadmin" {
			controller.WriteJSON(w, http.StatusForbidden, controller.ErrorResponse{
				Message: "forbidden action. only superadmin can access this endpoint",
			})
			return
		}

		next(w, r, p)
	}
}

func findUserById(id int, db *sql.DB) (model.User, error) {
	query := `
		SELECT * FROM users WHERE id = ?
	`
	rows, err := db.QueryContext(context.Background(), query, id)
	if err != nil {
		panic(err)
	}

	userFound := model.User{}

	if !rows.Next() {
		return userFound, errors.New("user not found")
	}

	rows.Scan(&userFound.Id, &userFound.Username, &userFound.Password, &userFound.Role)
	return userFound, err
}
