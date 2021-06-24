package http

type UserInfo struct {
	Sub         string   `json:"sub"`
	Authorities []string `json:"authorities"`
}

