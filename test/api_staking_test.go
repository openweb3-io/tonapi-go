/*
REST api to TON blockchain explorer

Testing StakingAPIService

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

func Test_tonapi_StakingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StakingAPIService GetAccountNominatorsPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.StakingAPI.GetAccountNominatorsPools(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StakingAPIService GetStakingPoolHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.StakingAPI.GetStakingPoolHistory(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StakingAPIService GetStakingPoolInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.StakingAPI.GetStakingPoolInfo(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StakingAPIService GetStakingPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StakingAPI.GetStakingPools(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
