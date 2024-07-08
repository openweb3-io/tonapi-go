/*
REST api to TON blockchain explorer

Testing WalletAPIService

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

func Test_tonapi_WalletAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WalletAPIService GetAccountSeqno", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.WalletAPI.GetAccountSeqno(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletAPIService GetWalletBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletAPI.GetWalletBackup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletAPIService GetWalletsByPublicKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var publicKey string

		resp, httpRes, err := apiClient.WalletAPI.GetWalletsByPublicKey(context.Background(), publicKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletAPIService SetWalletBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.WalletAPI.SetWalletBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletAPIService TonConnectProof", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletAPI.TonConnectProof(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
