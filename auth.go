package seabotserver

type AuthResponse struct {
	OK        bool   `json:"ok"`
	Error     string `json:"error",omitempty`
	PlayersID int64  `json:"id"`
}
