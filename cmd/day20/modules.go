package main

type pulse bool

const (
	low  pulse = false
	high pulse = true
)

func (p pulse) String() string {
	switch p {
	case low:
		return "low"
	case high:
		return "high"
	}
	return "unknown"
}

type modtype int

const (
	ff modtype = iota
	con
	bc
)

type message struct {
	input       string
	pulse       pulse
	destination string
}

type module interface {
	handle_message(message) []message
	add_inputs([]string)
	get_name() string
	get_destinations() []string
	get_type() modtype
}

type flip_flop struct {
	name         string
	destinations []string
	state        bool
}

func (mod *flip_flop) handle_message(m message) []message {
	var messages []message
	if m.pulse == low {
		mod.state = !mod.state
		send := pulse(mod.state)
		for _, d := range mod.destinations {
			messages = append(messages, message{mod.name, send, d})
		}
	}
	return messages
}

func (mod *flip_flop) get_name() string {
	return mod.name
}

func (mod *flip_flop) get_destinations() []string {
	return mod.destinations
}

func (mod *flip_flop) get_type() modtype {
	return ff
}

func (mod *flip_flop) add_inputs(names []string) {}

type conjunction struct {
	name         string
	destinations []string
	inputs       map[string]pulse
}

func (mod *conjunction) handle_message(m message) []message {
	var messages []message
	mod.inputs[m.input] = m.pulse
	send := low
	for _, p := range mod.inputs {
		if p == low {
			send = high
			break
		}
	}
	for _, d := range mod.destinations {
		messages = append(messages, message{mod.name, send, d})
	}
	return messages
}

func (mod *conjunction) get_name() string {
	return mod.name
}

func (mod *conjunction) get_destinations() []string {
	return mod.destinations
}

func (mod *conjunction) get_type() modtype {
	return con
}

func (mod *conjunction) add_inputs(names []string) {
	for _, name := range names {
		mod.inputs[name] = low
	}
}

type broadcast struct {
	name         string
	destinations []string
}

func (mod *broadcast) handle_message(m message) []message {
	var messages []message
	for _, d := range mod.destinations {
		messages = append(messages, message{mod.name, m.pulse, d})
	}
	return messages
}

func (mod *broadcast) get_name() string {
	return mod.name
}

func (mod *broadcast) get_destinations() []string {
	return mod.destinations
}

func (mod *broadcast) get_type() modtype {
	return bc
}

func (mod *broadcast) add_inputs(names []string) {}
