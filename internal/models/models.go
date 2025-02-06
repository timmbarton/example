package models

type userRole string

const (
	RoleAdmin     = "admin"
	RoleWebmaster = "webmaster"
	RolePartner   = "partner"
)

type User struct {
	Id    int    `json:"id" db:"id"`
	Login string `json:"login" db:"login"`
	Role  string `json:"role" db:"role"`
}
