package mongodbatlas

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb-forks/digest"
	"github.com/turbot/steampipe-plugin-mongodbatlas/constants"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"go.mongodb.org/atlas/mongodbatlas"
)

// getMongodbAtlasClient :: returns a mongodbatlas client to perform API requests
func getMongodbAtlasClient(ctx context.Context, d *plugin.QueryData) (*mongodbatlas.Client, error) {
	// Try to load client from cache
	if cachedData, ok := d.ConnectionManager.Cache.Get(constants.CacheKeyMongodbAtlasClient); ok {
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
	d.ConnectionManager.Cache.Set(constants.CacheKeyMongodbAtlasClient, client)

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
		log.Fatalf(err.Error())
	}

	return mongodbatlas.NewClient(tc)
}
