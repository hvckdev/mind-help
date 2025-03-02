package user

type RegisterRequest struct {
	Name                 string `json:"name" binding:"required,min=2,max=50"`
	Email                string `json:"email" binding:"required,email,unique"`
	Password             string `json:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" binding:"confirmPassword"`
}
