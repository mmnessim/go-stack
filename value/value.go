package value

import "fmt"

type Value interface {
	TypeOf() string
	String() string
}

type Number struct{ V int64 }
type Str struct{ V string }
type Boolean struct{ V bool }

func (n Number) TypeOf() string { return "number" }
func (n Number) String() string { return fmt.Sprintf("%d", n.V) }

func (s Str) TypeOf() string { return "string" }
func (s Str) String() string { return s.V }

func (b Boolean) TypeOf() string { return "boolean" }
func (b Boolean) String() string { return fmt.Sprintf("%t", b.V) }
