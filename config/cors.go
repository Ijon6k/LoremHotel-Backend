package config

import "github.com/rs/cors"

// SetupCORS mengatur konfigurasi CORS
func SetupCORS() *cors.Cors {
	return cors.AllowAll()
}
