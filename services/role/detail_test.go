package roleservice_test

import (
	"net/http/httptest"
	"testing"

	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRoleServiceDetail(t *testing.T) {
	testApp := createRoleTest()
	// Case 1: Invalid ID
	req := httptest.NewRequest("GET", "/role/9999", nil)
	req.Header.Set("Authorization", "Bearer "+testApp.adminToken)
	resp := utils.Must(testApp.server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 404, resp.StatusCode)

	// Case 2: Success
	req = httptest.NewRequest("GET", "/role/2", nil)
	req.Header.Set("Authorization", "Bearer "+testApp.adminToken)
	resp = utils.Must(testApp.server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 200, resp.StatusCode)
	response := utils.Must(utils.ReadCloserToString(resp.Body))
	assert.Contains(t, response, `"name":"User"`)
}
