package localize

import (
	"os"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/api"
	"github.com/stretchr/testify/assert"
)

func TestLocalizeSuccess(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	db, err := api.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestLocalizeFail(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	db, err := api.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
