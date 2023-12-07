package src

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func NewAppConfigTest() (*AppConfig, error) {
	// load .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("cannot load env :%v", err)
	}

	// load env
	return NewAppConfig()
}

func TestNewAppConfig(t *testing.T) {
	t.Run("Should Load Correct Env", func(t *testing.T) {
		// load env
		appConfig, err := NewAppConfigTest()

		assert.Nil(t, err)
		assert.Equal(t, appConfig.Name, "Test Order Service")
	})

	t.Run("Should return error get env port", func(t *testing.T) {
		// set PORT to empty
		os.Setenv("PORT", "")

		// load env
		_, err := NewAppConfig()
		assert.NotNil(t, err)
	})
}
