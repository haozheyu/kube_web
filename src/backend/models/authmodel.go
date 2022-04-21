package models

const (
	AuthTypeDB   = "db"
	AuthTypeLDAP = "ldap"
)

// AuthModel holds information used to authenticate.
type AuthModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"usertype"`
}
