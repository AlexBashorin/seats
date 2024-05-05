package types

type Amount struct {
	Date string  `json:"date"`
	Sum  float32 `json:"sum"`
	Mov  string  `json:"mov"`
	Id   int     `json:"id"`
}
