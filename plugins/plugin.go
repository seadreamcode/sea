package plugin

type plugin interface {
	Exec([]byte) ([]byte, error)
}
