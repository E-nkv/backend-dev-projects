package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/E-nkv/backend-dev-projects/httpServer/errs"
	"github.com/E-nkv/backend-dev-projects/httpServer/types"
)

//option 1 is to have all the handlers directly in api.go (this becomes messy once there are a lot of handlers)
//option 2 is to have them all in a file like this one, handlers.go. this hits the sweet spot between "good enough structure" and "simple enough"
//option 3 is to have dedicated files per resource for the handlers. for example, userHandlers.go, whateverHandlers.go, and so on. Use this only once there are LOOOTS of endpoints, since it may be a bit overkill for small projects (~less than 10 endpoints)

//we're going for option 2.

func (app *App) HandleHome(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, "hello from home!", "message")
}
func (app *App) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := app.Service.GetUsers()
	if err != nil {
		switch err {
		default:
			WriteInternalServerError(w, err.Error())
			return
		}
	}
	WriteJSON(w, http.StatusOK, result, "users")
}
func (app *App) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("userID"), 10, 64)
	if err != nil {
		WriteBadRequestError(w, "id must be a number")
		return
	}
	result, err := app.Service.GetUser(id)
	if err != nil {
		switch err {
		case errs.ErrNotFound:
			WriteBadRequestError(w, "user not found")
			return
		default:
			WriteInternalServerError(w, err.Error())
			return
		}
	}
	WriteJSON(w, http.StatusOK, result, "user")
}
func (app *App) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var u types.UserCreate
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		WriteBadRequestError(w, "expected a request body")
		return
	}
	if err := json.Unmarshal(bs, &u); err != nil {
		WriteBadRequestError(w, err.Error())
		return
	}

	userID, err := app.Service.CreateUser(&u)
	if err != nil {
		switch err {
		default:
			WriteInternalServerError(w, err.Error())
			return
		}
	}

	response := map[string]any{
		"message": "user created succesfully",
		"user_id": userID,
	}
	WriteJSON(w, http.StatusCreated, response, "data")
}
func (app *App) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("userID"), 10, 64)
	if err != nil {
		WriteBadRequestError(w, "id must be a number")
		return
	}
	if err := app.Service.DeleteUser(id); err != nil {
		switch err {
		case errs.ErrNotFound:
			WriteBadRequestError(w, "user does not exist. maybe the account was deleted?")
			return
		default:
			WriteInternalServerError(w, err.Error())
			return
		}
	}
	WriteJSON(w, http.StatusOK, "user deleted", "message")
}
