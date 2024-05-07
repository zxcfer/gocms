package roleservice_test

import (
	"net/http/httptest"
	"testing"

	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRoleServiceAuth(t *testing.T) {
	testApp := createRoleTest()
	t.Run("Test_ParseUser", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/testuser", nil)
		resp := utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode)
		response := utils.Must(utils.ReadCloserToString(resp.Body))
		assert.Equal(t, `{"data":null}`, response)

		req = httptest.NewRequest("GET", "/testuser", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.adminToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode)
		response = utils.Must(utils.ReadCloserToString(resp.Body))
		assert.Contains(t, response, `"username":"adminuser"`)

		req = httptest.NewRequest("GET", "/testuser", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.normalUserToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode)
		response = utils.Must(utils.ReadCloserToString(resp.Body))
		assert.Contains(t, response, `"username":"normaluser"`)
	})

	t.Run("Test_Authorize", func(t *testing.T) {
		// Admin user should have access to any resource without white list or permission set
		req := httptest.NewRequest("GET", "/content/blog", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.adminToken)
		resp := utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode, "Admin user should have access to any resource without white list or permission set")
		assert.Equal(t, `{"data":"blog list"}`, utils.Must(utils.ReadCloserToString(resp.Body)))

		// Guest user should have access to white listed resource
		req = httptest.NewRequest("GET", "/test", nil)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode, "Guest user should have access to white listed resource")
		assert.Equal(t, `{"data":"test response"}`, utils.Must(utils.ReadCloserToString(resp.Body)))

		// Guest user should not have access to non white listed resource
		req = httptest.NewRequest("GET", "/content/blog", nil)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 401, resp.StatusCode, "Guest user should not have access to non white listed resource")
		assert.Contains(t, utils.Must(utils.ReadCloserToString(resp.Body)), `Unauthorized`)

		// Inactive user should not have access to any non white listed resource
		req = httptest.NewRequest("GET", "/content/blog", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.inactiveUserToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 403, resp.StatusCode, "Inactive user should not have access to any non white listed resource")
		assert.Contains(t, utils.Must(utils.ReadCloserToString(resp.Body)), `User is inactive`)

		// Active user has these permissions:
		// - content.list: allow
		// - content.detail: deny
		// - content.meta: no permission set
		// Expectation: user should have access to content.list but not content.detail and content.meta
		req = httptest.NewRequest("GET", "/content/blog", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.normalUserToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 200, resp.StatusCode, "User should have access to content.blog.list")
		assert.Equal(t, `{"data":"blog list"}`, utils.Must(utils.ReadCloserToString(resp.Body)))

		req = httptest.NewRequest("GET", "/content/1", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.normalUserToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 403, resp.StatusCode, "User should not have access to content.blog.detail")
		assert.Contains(t, utils.Must(utils.ReadCloserToString(resp.Body)), `Forbidden`)

		req = httptest.NewRequest("GET", "/content/blog/meta", nil)
		req.Header.Set("Authorization", "Bearer "+testApp.normalUserToken)
		resp = utils.Must(testApp.server.Test(req))
		defer func() { assert.NoError(t, resp.Body.Close()) }()
		assert.Equal(t, 403, resp.StatusCode, "User should not have access to content.blog.meta")
		assert.Contains(t, utils.Must(utils.ReadCloserToString(resp.Body)), `Forbidden`)
	})
}
