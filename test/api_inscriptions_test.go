/*
REST api to TON blockchain explorer

Testing InscriptionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tonapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func Test_tonapi_InscriptionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InscriptionsAPIService GetAccountInscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.InscriptionsAPI.GetAccountInscriptions(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InscriptionsAPIService GetAccountInscriptionsHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.InscriptionsAPI.GetAccountInscriptionsHistory(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InscriptionsAPIService GetAccountInscriptionsHistoryByTicker", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string
		var ticker string

		resp, httpRes, err := apiClient.InscriptionsAPI.GetAccountInscriptionsHistoryByTicker(context.Background(), accountId, ticker).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InscriptionsAPIService GetInscriptionOpTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InscriptionsAPI.GetInscriptionOpTemplate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}