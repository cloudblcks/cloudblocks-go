package control

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var username string
var password string
var url string = API_URL

func TestMain(m *testing.M) {
	username = os.Getenv("CLODUBLOCKS_USERNAME")
	password = os.Getenv("CLODUBLOCKS_PASSWORD")
	rand.Seed(time.Now().UnixNano())

	if username == "" {
		panic("CLODUBLOCKS_USERNAME not set")
	}
	if password == "" {
		panic("CLODUBLOCKS_PASSWORD not set")
	}

	if os.Getenv("CLOUDBLOCKS_URL") != "" {
		url = os.Getenv("CLOUDBLOCKS_URL")
	}

	code := m.Run()
	os.Exit(code)
}

func newTestClient(t *testing.T) (Client, error) {
	client, err := NewClientWithURL(username, password, url)
	assert.NoError(t, err)
	return client, err
}
