package user

type User struct {
	tableName struct{} `pg:"user"`
	Id        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
