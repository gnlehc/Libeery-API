package output

type LoginResponseDTO struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	UserId     string `json:"userid"`
	Username   string `json:"username"`
}
