/*
REST api to TON blockchain explorer

Testing EventsAPIService

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

func Test_tonapi_EventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsAPIService GetEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventId string

		resp, httpRes, err := apiClient.EventsAPI.GetEvent(context.Background(), eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}