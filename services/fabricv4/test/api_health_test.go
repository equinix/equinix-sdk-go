/*
Equinix Fabric API v4

Testing HealthApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fabricv4

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fabricv4_HealthApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HealthApiService GetStatus", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HealthApi.GetStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
