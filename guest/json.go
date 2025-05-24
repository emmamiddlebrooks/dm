package guest

type Request struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	WifiPassword string `json:"wifi_password"`
	ClientID     string `json:"id"`
}

type Response struct {
	RedirectUrl string `json:"redirect_url"`
}
