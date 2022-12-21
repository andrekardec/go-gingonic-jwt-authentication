package user

type CreateUserRequestDto struct {
	Id                   string `json:"id" binding:"required"`
	FirstName            string `json:"firstName" binding:"required"`
	LastName             string `json:"lastName" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"password"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

type CreateUserDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"password"`
}
