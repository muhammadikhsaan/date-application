package pensiel

type Error struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
	Origin     error  `json:"origin"`
}
