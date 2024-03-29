/*
Equinix Fabric API v4

Testing RoutingProtocolsApiService

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

func Test_fabricv4_RoutingProtocolsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoutingProtocolsApiService CreateConnectionRoutingProtocol", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.CreateConnectionRoutingProtocol(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService CreateConnectionRoutingProtocolsInBulk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.CreateConnectionRoutingProtocolsInBulk(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService DeleteConnectionRoutingProtocolByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.DeleteConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocolAllBgpActions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolAllBgpActions(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocolByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocols", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocols(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocolsBgpActionByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string
		var routingProtocolId string
		var actionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsBgpActionByUuid(context.Background(), connectionId, routingProtocolId, actionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocolsChangeByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string
		var routingProtocolId string
		var changeId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsChangeByUuid(context.Background(), connectionId, routingProtocolId, changeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService GetConnectionRoutingProtocolsChanges", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string
		var routingProtocolId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.GetConnectionRoutingProtocolsChanges(context.Background(), connectionId, routingProtocolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService PatchConnectionRoutingProtocolByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.PatchConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService PostConnectionRoutingProtocolBgpActionByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.PostConnectionRoutingProtocolBgpActionByUuid(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService ReplaceConnectionRoutingProtocolByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routingProtocolId string
		var connectionId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.ReplaceConnectionRoutingProtocolByUuid(context.Background(), routingProtocolId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RoutingProtocolsApiService ValidateRoutingProtocol", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.RoutingProtocolsApi.ValidateRoutingProtocol(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
