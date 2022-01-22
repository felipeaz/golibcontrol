package login

const (
	SuccessMessage         = "Welcome, %s!"
	FailMessage            = "Login Failed"
	AccountNotFoundMessage = "Account not found"
	InvalidPasswordMessage = "Invalid Password"
	LogoutSuccessMessage   = "Logout Successfully."
)

// Data will be used as custom messages for login routes
type Data struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Token   string `json:"-"`
}
