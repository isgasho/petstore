package internal

const (
	// DefaultBind is the default interace and port to bind the server to
	DefaultBind = ":8000"

	// DefaultDBPath is the default database storage path
	DefaultDBPath = "./petstore.db"
)

// Option is a function that takes a config struct and modifies it
type Option func(*config) error

type config struct {
	bind   string
	dbPath string
}

func newDefaultConfig() *config {
	return &config{
		bind:   DefaultBind,
		dbPath: DefaultDBPath,
	}
}

// WithBind sets port and interface to bind the server to
func WithBind(bind string) Option {
	return func(cfg *config) error {
		cfg.bind = bind
		return nil
	}
}

// WithDBPath sets the database storage path
func WithDBPath(path string) Option {
	return func(cfg *config) error {
		cfg.dbPath = path
		return nil
	}
}
