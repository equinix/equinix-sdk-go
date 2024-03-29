/*
Equinix Fabric API v4

Testing RouteFilterRulesApiService

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

func Test_fabricv4_RouteFilterRulesApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RouteFilterRulesApiService CreateRouteFilterRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.CreateRouteFilterRule(context.Background(), routeFilterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService CreateRouteFilterRulesInBulk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.CreateRouteFilterRulesInBulk(context.Background(), routeFilterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService DeleteRouteFilterRuleByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.DeleteRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService GetRouteFilterRuleByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService GetRouteFilterRuleChangeByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string
		var changeId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleChangeByUuid(context.Background(), routeFilterId, routeFilterRuleId, changeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService GetRouteFilterRuleChanges", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.GetRouteFilterRuleChanges(context.Background(), routeFilterId, routeFilterRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService GetRouteFilterRules", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.GetRouteFilterRules(context.Background(), routeFilterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService PatchRouteFilterRuleByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.PatchRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test RouteFilterRulesApiService ReplaceRouteFilterRuleByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routeFilterId string
		var routeFilterRuleId string

		resp, httpRes, err := apiClient.RouteFilterRulesApi.ReplaceRouteFilterRuleByUuid(context.Background(), routeFilterId, routeFilterRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
