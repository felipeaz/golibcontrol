package login

const (
	SuccessMessage         = "Login Successfully. Welcome, %s!"
	FailMessage            = "Login Failed"
	AccountNotFoundMessage = "Account not found. Sign In First to access the system"
	InvalidPasswordMessage = "Invalid Password"
)

type LoginMessage struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Token   string `json:"token,omitempty"`
}
