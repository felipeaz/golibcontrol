package login

const (
	SuccessMessage         = "Login Successfully. Welcome, %s!"
	FailMessage            = "Login Failed"
	AccountNotFoundMessage = "Account not found. Sign In First to access the system"
	InvalidPasswordMessage = "Invalid Password"
)

// Message will be used as custom messages for login routes
type Message struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Token   string `json:"token,omitempty"`
}
