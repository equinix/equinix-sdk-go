/*
Metal API

Testing PortsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package metalv1

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_metalv1_PortsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PortsApiService AssignNativeVlan", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.AssignNativeVlan(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService AssignPort", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.AssignPort(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService BondPort", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.BondPort(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService ConvertLayer2", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.ConvertLayer2(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService ConvertLayer3", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.ConvertLayer3(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService CreatePortVlanAssignmentBatch", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.CreatePortVlanAssignmentBatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService DeleteNativeVlan", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.DeleteNativeVlan(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService DisbondPort", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.DisbondPort(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService FindPortById", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.FindPortById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService FindPortVlanAssignmentBatchByPortIdAndBatchId", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string
		var batchId string

		resp, httpRes, err := apiClient.PortsApi.FindPortVlanAssignmentBatchByPortIdAndBatchId(context.Background(), id, batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService FindPortVlanAssignmentBatches", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.FindPortVlanAssignmentBatches(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService FindPortVlanAssignmentByPortIdAndAssignmentId", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string
		var assignmentId string

		resp, httpRes, err := apiClient.PortsApi.FindPortVlanAssignmentByPortIdAndAssignmentId(context.Background(), id, assignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService FindPortVlanAssignments", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.FindPortVlanAssignments(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PortsApiService UnassignPort", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.PortsApi.UnassignPort(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}