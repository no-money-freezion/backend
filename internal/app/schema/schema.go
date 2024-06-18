package schema

type RegisterRequest struct {
	Email        string `json:"email" example:"user@example.com" validate:"required,email"`
	PasswordHash string `json:"passwordHash" example:"hashpassword" validate:"required"`
}

type LoginRequest struct {
	Email        string `json:"email" example:"user@example.com" validate:"required,email"`
	PasswordHash string `json:"passwordHash" example:"hashpassword" validate:"required"`
}

type PasswordRecoveryRequest struct {
	Email string `json:"email" example:"user@example.com" validate:"required,email"`
}

type PasswordResetRequest struct {
	Token           string `json:"token" example:"reset_token" validate:"required"`
	NewPasswordHash string `json:"newPasswordHash" example:"newhashpassword" validate:"required"`
}

type HealthcheckResponse struct {
	Status string `json:"status"`
}

type RouletteInfoResponse struct {
	Version      string                `json:"version" example:"1.0"`
	Combinations []RouletteCombination `json:"combinations"`
	MaxBet       int                   `json:"maxBet" example:"1000"`
	MinBet       int                   `json:"minBet" example:"10"`
}

type RouletteCombination struct {
	Combination string `json:"combination" example:"Triple 7"`
	Payout      int    `json:"payout" example:"100"`
}

type RouletteBetRequest struct {
	UserId       string `json:"userId" example:"user123" validate:"required"`
	Amount       int    `json:"amount" example:"100" validate:"required,min=1"`
	RulesVersion string `json:"rulesVersion" example:"1.0" validate:"required"`
}

type RouletteBetResponse struct {
	Result    string `json:"result" example:"Lose"`
	WinAmount int    `json:"winAmount" example:"0"`
	Message   string `json:"message" example:"Better luck next time!"`
}
