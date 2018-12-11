package config

type Mongo struct {
	// Enable
	Enable bool

	// Database
	Database string

	// Endpoint
	Endpoint string

	// Username
	Username string

	// Password
	Password string

	// Auth Database
	AuthDatabase string
}
