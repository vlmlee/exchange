package account

type Persona int

const (
	Fish Persona = iota
	Slowpoke
	Speculator
	Whale
)

type Account struct {
	Name      string
	Type      Persona
	Balance   float32
	Portfolio map[string]int
}
