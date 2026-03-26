package vm

type Dictionary struct {
	Words []map[string]func(*VM) error
}
