package getuser

type Mode int

const (
	ANONYMOUS Mode = iota
	SIGNED
	SELF
)

type GetUserQuery struct {
	Id   string
	Mode Mode
}
