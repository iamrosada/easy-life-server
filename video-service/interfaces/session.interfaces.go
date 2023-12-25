package interfaces

// Session interface
type Session struct {
	ID       string `json:"id"`
	Host     string `json:"host"`
	Title    string `json:"title"`
	Password string `json:"password"`
}
