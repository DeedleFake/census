package events

type Sub interface {
	name() string
	params() map[string][]string
}

type AllSub struct {
}

func (s AllSub) name() string {
	return "all"
}

func (s AllSub) params() map[string][]string {
	return map[string][]string{
		"characters": []string{"all"},
		"worlds":     []string{"all"},
	}
}

type CharSub struct {
	Event string
	Chars []string
}

func (s CharSub) name() string {
	return s.Event
}

func (s CharSub) params() map[string][]string {
	return map[string][]string{
		"characters": s.Chars,
	}
}

type WorldSub struct {
	Event  string
	Worlds []string
}

func (s WorldSub) name() string {
	return s.Event
}

func (s WorldSub) params() map[string][]string {
	return map[string][]string{
		"worlds": s.Worlds,
	}
}

type CharWorldSub struct {
	Event  string
	Chars  []string
	Worlds []string
}

func (s CharWorldSub) name() string {
	return s.Event
}

func (s CharWorldSub) params() map[string][]string {
	return map[string][]string{
		"characters": s.Chars,
		"worlds":     s.Worlds,
	}
}
