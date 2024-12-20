/*
Equinix Fabric API v4

Testing RouteAggregationsApiService

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

func Test_fabricv4_RouteAggregationsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RouteAggregationsApiService AttachConnectionRouteAggregation", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string
		var connectionId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.AttachConnectionRouteAggregation(context.Background(), routeAggregationId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService CreateRouteAggregation", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RouteAggregationsApi.CreateRouteAggregation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService DeleteRouteAggregationByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.DeleteRouteAggregationByUuid(context.Background(), routeAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService DetachConnectionRouteAggregation", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string
		var connectionId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.DetachConnectionRouteAggregation(context.Background(), routeAggregationId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetConnectionRouteAggregationByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string
		var connectionId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetConnectionRouteAggregationByUuid(context.Background(), routeAggregationId, connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetConnectionRouteAggregations", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetConnectionRouteAggregations(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetRouteAggregationByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetRouteAggregationByUuid(context.Background(), routeAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetRouteAggregationChangeByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string
		var changeId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetRouteAggregationChangeByUuid(context.Background(), routeAggregationId, changeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetRouteAggregationChanges", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetRouteAggregationChanges(context.Background(), routeAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService GetRouteAggregationConnections", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.GetRouteAggregationConnections(context.Background(), routeAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService PatchRouteAggregationByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeAggregationId string

		resp, httpRes, err := apiClient.RouteAggregationsApi.PatchRouteAggregationByUuid(context.Background(), routeAggregationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteAggregationsApiService SearchRouteAggregations", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RouteAggregationsApi.SearchRouteAggregations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
