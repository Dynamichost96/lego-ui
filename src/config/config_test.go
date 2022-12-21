package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var sample = Config {
	Email: "example@example.com",
}

func TestGeneratingPrivateKeysShouldNotBeEqual(t *testing.T) {
	assert := assert.New(t)
	assert.NotEqual(NewPrivateKey(), NewPrivateKey(), "Generating Private Keys should result in different results every time")
}

func GetAndPropertiesShouldBeEqual(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(sample.Email, sample.GetEmail(), "Getter Method GetEmail() does not produce the expected Result")
}