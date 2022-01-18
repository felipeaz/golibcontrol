package login

const (
	SuccessMessage         = "Welcome, %s!"
	FailMessage            = "Login Failed"
	AccountNotFoundMessage = "Account not found"
	InvalidPasswordMessage = "Invalid Password"

	LogoutSuccessMessage = "Logout Successfully."
)

// Message will be used as custom messages for login routes
type Message struct {
	Status     int    `json:"-"`
	Message    string `json:"message"`
	UserId     uint   `json:"userId,omitempty"`
	ConsumerId string `json:"consumerId,omitempty"`
	Reason     string `json:"reason,omitempty"`
	Token      string `json:"token,omitempty"`
}
