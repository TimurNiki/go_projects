package v1
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Teacher string `json:"teacher"`
}