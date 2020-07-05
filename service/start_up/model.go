package start_up

type status int

const (
	ACCEPTED status = 1
	DENIED   status = -1
)

type RegToken struct {
	Status status `json:"status"`
	Msg    string `json:"msg"`
	Token  string `json:"token"`
}
