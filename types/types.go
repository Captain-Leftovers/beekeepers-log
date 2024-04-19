package types



var UserKey = "user"


type AuthenticatedUser struct {
    Email string
    Username string
    IsLoggedIn bool
}

