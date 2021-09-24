package response

// Error is a basic error struct with code and msg.
type Error struct {
	Code int    `json:"errorCode"`
	Msg  string `json:"msg"`
}
