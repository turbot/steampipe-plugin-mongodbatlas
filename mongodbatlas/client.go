package mongodbatlas

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/mongodb-forks/digest"
	"github.com/turbot/steampipe-plugin-mongodbatlas/constants"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"go.mongodb.org/atlas/mongodbatlas"
)

// getMongoDBAtlasClient :: returns a mongodbatlas client to perform API requests.
func getMongoDBAtlasClient(ctx context.Context, d *plugin.QueryData) (*mongodbatlas.Client, error) {
	// Try to load client from cache
	if cachedData, ok := d.ConnectionManager.Cache.Get(constants.CacheKeyMongoDbAtlasClient); ok {
		return cachedData.(*mongodbatlas.Client), nil
	}

	// Get mongodbatlas keys
	publicKey, privateKey, err := getKeysFromConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// create the mongodbatlas client
	client := createClient(ctx, publicKey, privateKey)

	// save client in cache
	d.ConnectionManager.Cache.Set(constants.CacheKeyMongoDbAtlasClient, client)

	return client, nil
}

// getKeysFromConfig fetches the public and private keys from the connection config
// falls back to the environment variables if it cannot find one in the config
// returns an error if both keys could not be resolved
func getKeysFromConfig(ctx context.Context, d *plugin.QueryData) (publicKey string, privateKey string, _ error) {
	config := GetConfig(d.Connection)

	// Get the authorization publicKey
	publicKey = os.Getenv(constants.EnvKeyClientPublicKey)
	if config.PublicKey != nil {
		publicKey = *config.PublicKey
	}

	// Get the authorization publicKey
	privateKey = os.Getenv(constants.EnvKeyClientPrivateKey)
	if config.PrivateKey != nil {
		privateKey = *config.PrivateKey
	}

	if len(publicKey) == 0 || len(privateKey) == 0 {
		return "", "", fmt.Errorf("public key and private key must be configured")
	}

	return publicKey, privateKey, nil
}

func createClient(ctx context.Context, publicKey string, privateKey string) *mongodbatlas.Client {
	t := digest.NewTransport(publicKey, privateKey)
	tc, err := t.Client()
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	return mongodbatlas.NewClient(&http.Client{
		Transport: loggingRoundTripper{tc.Transport},
	})
}

// This type implements the http.RoundTripper interface
type loggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt loggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	plugin.Logger(req.Context()).Trace("Sending request to:", req.URL)
	res, e = lrt.Proxied.RoundTrip(req)
	if e != nil {
		plugin.Logger(req.Context()).Error("Error: %v", e)
	}
	response, _ := httputil.DumpResponse(res, true)
	plugin.Logger(req.Context()).Trace("Response: %v", string(response))
	return res, e
}
