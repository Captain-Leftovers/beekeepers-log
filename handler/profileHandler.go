package handler

import (
	"net/http"
	"time"

	"github.com/Captain-Leftovers/beekeepers-log/internal/auth"
	"github.com/Captain-Leftovers/beekeepers-log/internal/database"
	"github.com/Captain-Leftovers/beekeepers-log/util"
	"github.com/Captain-Leftovers/beekeepers-log/view/profile"
)

func ProfileIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := util.GetUserFromContext(r.Context())

		err := profile.ProfileIndex(user).Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}

func HandleProfilePost(DBQ *database.Queries) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			logError(r, err)
			return
		}

		formUsername := r.FormValue("username")
		formEmail := r.FormValue("email")

		formCurrentPassword := r.FormValue("current-password")
		formNewPassword := r.FormValue("new-password")
		formRepeatNewPassword := r.FormValue("repeat-new-password")

		userFromCtx := util.GetUserFromContext(r.Context())

		userFromDb, err := DBQ.GetUserById(r.Context(), userFromCtx.ID)
		if err != nil {
			logError(r, err)
			return
		}

		var pass string

		pass, _ = changePass(userFromDb, formCurrentPassword, formNewPassword, formRepeatNewPassword)

		userForUpdate := database.UpdateUserParams{
			Username:  formUsername,
			Email:     formEmail,
			ID:        userFromDb.ID,
			Password:  pass,
			UpdatedAt: time.Now().UTC(),
		}

		_, err = DBQ.UpdateUser(r.Context(), userForUpdate)
		if err != nil {
			logError(r, err)
			return
		}

		w.Header().Set("hx-redirect", "/sign-out")

	})
}

func ChangePasswordFields() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := profile.ChangePasswordFields().Render(r.Context(), w)
		if err != nil {
			logError(r, err)
		}

	})

}

func changePass(userFromDb database.User, formCurrentPassword string, formNewPassword string, formRepeatNewPassword string) (pass string, ok bool) {
	if formCurrentPassword == "" || formNewPassword == "" || formRepeatNewPassword == "" {

		return userFromDb.Password, false
	}

	if formNewPassword != formRepeatNewPassword {

		return userFromDb.Password, false
	}

	if err := auth.CompareHashToPass(userFromDb.Password, formCurrentPassword); err != nil {
		return userFromDb.Password, false
	}

	hashedPass, err := auth.HashNewPassword(formNewPassword)

	if err != nil {
		return userFromDb.Password, false
	}

	return hashedPass, true

}
