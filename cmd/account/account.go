package account

type Persona int

const (
	Fish Persona = iota
	Slowpoke
	Speculator
	Whale
)

type Account struct {
	ID        string
	Name      string
	Type      Persona
	Balance   float64
	Portfolio map[string]int
}
