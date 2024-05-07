package output

type LoginResponseDTO struct {
	UserId     string `json:"userid"`
	Username   string `json:"username"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}
