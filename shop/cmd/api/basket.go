package main

import (
	"errors"
	"fainal.net/internal/data"
	"fainal.net/internal/validator"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) createBasketHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Items []int64 `json:"items"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	w.Header().Add("Vary", "Authorization")
	authorizationHeader := r.Header.Get("Authorization")

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	token := headerParts[1]
	v := validator.New()

	if data.ValidateTokenPlaintext(v, token); !v.Valid() {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidAuthenticationTokenResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	basket := &data.Basket{
		Items:   input.Items,
		User_id: user.ID,
	}

	//if role.Role_name != ("admin") && role.Role_name != ("user") {
	//	app.invalidRoleResponse(w)
	//	return
	//}
	err = app.models.Baskets.BasketInsert(basket)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	err = app.writeJSON(w, http.StatusCreated, envelope{"basket": basket}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showBasketHandler(w http.ResponseWriter, r *http.Request) {
	//id, err := app.readIDParam(r)
	//if err != nil {
	//	app.notFoundResponse(w, r)
	//	return
	//}
	w.Header().Add("Vary", "Authorization")
	authorizationHeader := r.Header.Get("Authorization")

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	token := headerParts[1]
	v := validator.New()

	if data.ValidateTokenPlaintext(v, token); !v.Valid() {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidAuthenticationTokenResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	basket, err := app.models.Baskets.GetBasket(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"basket": basket}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) updateBasketHandler(w http.ResponseWriter, r *http.Request) {

	//id, err := app.readIDParam(r)
	//if err != nil {
	//	app.notFoundResponse(w, r)
	//	return
	//}

	w.Header().Add("Vary", "Authorization")
	authorizationHeader := r.Header.Get("Authorization")

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	token := headerParts[1]
	v := validator.New()

	if data.ValidateTokenPlaintext(v, token); !v.Valid() {
		app.invalidAuthenticationTokenResponse(w, r)
		return
	}

	user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidAuthenticationTokenResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	basket, err := app.models.Baskets.GetBasket(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Items []int64 `json:"items"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Items != nil {
		basket.Items = input.Items
	}
	fmt.Println(basket)
	err = app.models.Baskets.UpdateBasket(basket)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"basket": basket}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteBasketHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Baskets.DeleteBasket(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Basket deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
