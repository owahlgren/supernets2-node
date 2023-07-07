package event

import "github.com/0xPolygon/supernets2-node/db"

// Config for event
type Config struct {
	// DB is the database configuration
	DB db.Config `mapstructure:"DB"`
}
