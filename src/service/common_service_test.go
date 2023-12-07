package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	service := NewServiceTest(t)

	health := service.GetHealth()

	assert.Equal(t, health.GetName(), service.Config.Name)
}
