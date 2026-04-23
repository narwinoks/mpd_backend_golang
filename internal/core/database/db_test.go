package database

import (
	"backend-app/config"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDatabase(t *testing.T) {
	t.Run("Database Connection Failure - Invalid Host", func(t *testing.T) {
		// Chaos Rule: Generate randomized invalid data
		cfg := &config.Config{
			Database: config.DatabaseConfig{
				Host:     gofakeit.DomainName(), // Random domain that likely doesn't exist or doesn't have DB
				Port:     int(gofakeit.Uint16()),
				User:     gofakeit.Username(),
				Password: gofakeit.Password(true, true, true, true, false, 12),
				Name:     gofakeit.AppName(),
				SSLMode:  "disable",
			},
		}

		// Execution
		db, err := NewDatabase(cfg)

		// Assertions (Negative Scenario: Database Failure)
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("Buffer Stress Test - Long Input", func(t *testing.T) {
		// Buffer Stress: Extremely long strings
		cfg := &config.Config{
			Database: config.DatabaseConfig{
				Host:     gofakeit.LetterN(500),
				Port:     5432,
				User:     gofakeit.LetterN(500),
				Password: gofakeit.LetterN(500),
				Name:     gofakeit.LetterN(500),
				SSLMode:  "disable",
			},
		}

		// Execution
		db, err := NewDatabase(cfg)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("Security & Injection Test", func(t *testing.T) {
		// Injection & Security: SQL Injection & XSS strings
		injectionStrings := []string{
			"' OR 1=1 --",
			"\"; DROP TABLE users; --",
			"<script>alert('xss')</script>",
			"😊🚀UnicodeTest",
		}

		for _, injection := range injectionStrings {
			t.Run("Injection - "+injection, func(t *testing.T) {
				cfg := &config.Config{
					Database: config.DatabaseConfig{
						Host:     "localhost",
						Port:     5432,
						User:     injection,
						Password: injection,
						Name:     injection,
						SSLMode:  "disable",
					},
				}

				// Execution
				db, err := NewDatabase(cfg)

				// Assertions
				assert.Error(t, err)
				assert.Nil(t, db)
			})
		}
	})

	t.Run("Boundary Scenario - Zero Port", func(t *testing.T) {
		// Zero & Nil: Pass 0 to Port
		cfg := &config.Config{
			Database: config.DatabaseConfig{
				Host:     "localhost",
				Port:     0,
				User:     gofakeit.Username(),
				Password: gofakeit.Password(true, true, true, true, false, 10),
				Name:     gofakeit.AppName(),
				SSLMode:  "disable",
			},
		}

		// Execution
		db, err := NewDatabase(cfg)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("Requirement Check - SSLMode", func(t *testing.T) {
		// require for terminal failure check (if we can't even setup config)
		cfg := &config.Config{}
		require.NotNil(t, cfg)
		
		cfg.Database.SSLMode = "invalid-mode"
		
		db, err := NewDatabase(cfg)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}
