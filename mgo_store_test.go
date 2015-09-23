package surl

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func haveMongoAvailable() (string, bool) {
	v := os.Getenv("MGO_STORE_TEST_INSTANCE")
	return v, len(v) > 0
}

func TestMgoStorePersistsUrls(t *testing.T) {
	mu, haveMongo := haveMongoAvailable()

	if !haveMongo {
		t.Skip("Skipping test as we have not Mongo DB testing instance available")
	}

	u, _ := url.Parse("http://www.google.com")
	ms, me := NewMgoStore(mu)
	assert.Nil(t, me)
	assert.Nil(t, ms.Put("lalelu", u))
	v, err := ms.Get("lalelu")
	assert.Nil(t, err)
	assert.Equal(t, *u, *v, "Mismatching URLs")
}

func TestMgoStoreGracefullyErrorsForUnknownKey(t *testing.T) {
	u, haveMongo := haveMongoAvailable()

	if !haveMongo {
		t.Skip("Skipping test as we have not Mongo DB testing instance available")
	}

	ms, me := NewMgoStore(u)
	assert.Nil(t, me)

	rand.Seed(time.Now().Unix())
	v, err := ms.Get(fmt.Sprint(rand.Int()))

	assert.Nil(t, v)
	assert.NotNil(t, err)
}
