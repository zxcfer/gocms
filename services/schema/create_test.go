package schemaservice_test

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fastschema/fastschema/pkg/utils"
	"github.com/fastschema/fastschema/schema"
	"github.com/stretchr/testify/assert"
)

func TestSchemaServiceCreate(t *testing.T) {
	testApp, _, server := createSchemaService(t, nil)

	// Case 1: schema already exists
	req := httptest.NewRequest("POST", "/schema", bytes.NewReader([]byte(`{"name": "category"}`)))
	resp := utils.Must(server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 400, resp.StatusCode)
	response := utils.Must(utils.ReadCloserToString(resp.Body))
	assert.Contains(t, response, `schema already exists`)

	// Case 2: schema validation failed
	req = httptest.NewRequest("POST", "/schema", bytes.NewReader([]byte(`{"name": "blog"}`)))
	resp = utils.Must(server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 422, resp.StatusCode)
	response = utils.Must(utils.ReadCloserToString(resp.Body))
	assert.Contains(t, response, `schema validation error`)
	assert.Contains(t, response, `label_field is required`)
	assert.Contains(t, response, `namespace is required`)

	// Case 3: invalid relation schema
	newBlogJSON := strings.ReplaceAll(
		testBlogJSON,
		`"fields": [`,
		`"fields": [{
			"type": "relation",
			"name": "categories",
			"label": "Categories",
			"relation": {
				"schema": "cat",
				"field": "blogs",
				"type": "m2m",
				"owner": true,
				"optional": false
			}
		},`,
	)
	req = httptest.NewRequest("POST", "/schema", bytes.NewReader([]byte(newBlogJSON)))
	resp = utils.Must(server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 400, resp.StatusCode)
	response = utils.Must(utils.ReadCloserToString(resp.Body))
	assert.Contains(t, response, `Invalid field 'blog.categories'. Target schema 'cat' not found`)

	// Case 4: target relation field existed
	newBlogJSON = strings.ReplaceAll(
		testBlogJSON,
		`"fields": [`,
		`"fields": [{
			"type": "relation",
			"name": "categories",
			"label": "Categories",
			"relation": {
				"schema": "category",
				"field": "name",
				"type": "m2m",
				"owner": true,
				"optional": false
			}
		},`,
	)
	req = httptest.NewRequest("POST", "/schema", bytes.NewReader([]byte(newBlogJSON)))
	resp = utils.Must(server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 400, resp.StatusCode)
	response = utils.Must(utils.ReadCloserToString(resp.Body))
	assert.Contains(t, response, `Invalid field 'blog.categories'. Target schema 'category' already has field 'name'`)

	// Case 4: create schema successfully
	newBlogJSON = strings.ReplaceAll(
		testBlogJSON,
		`"fields": [`,
		`"fields": [{
			"type": "relation",
			"name": "categories",
			"label": "Categories",
			"relation": {
				"schema": "category",
				"field": "blogs",
				"type": "m2m",
				"owner": true,
				"optional": false
			}
		},`,
	)
	req = httptest.NewRequest("POST", "/schema", bytes.NewReader([]byte(newBlogJSON)))
	resp = utils.Must(server.Test(req))
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	assert.Equal(t, 200, resp.StatusCode)
	response = utils.Must(utils.ReadCloserToString(resp.Body))
	assert.NotEmpty(t, response)

	blogSchema := utils.Must(testApp.SchemaBuilder().Schema("blog"))
	blogCategoriesField, err := blogSchema.Field("categories")
	assert.NoError(t, err)
	assert.NotNil(t, blogCategoriesField)
	assert.Equal(t, "relation", blogCategoriesField.Type.String())
	assert.Equal(t, schema.M2M, blogCategoriesField.Relation.Type)
	assert.Equal(t, "category", blogCategoriesField.Relation.TargetSchemaName)
	assert.Equal(t, "blogs", blogCategoriesField.Relation.TargetFieldName)

	categorySchema := utils.Must(testApp.SchemaBuilder().Schema("category"))
	categoryBlogsField, err := categorySchema.Field("blogs")
	assert.NoError(t, err)
	assert.NotNil(t, categoryBlogsField)
	assert.Equal(t, "relation", categoryBlogsField.Type.String())
	assert.Equal(t, schema.M2M, categoryBlogsField.Relation.Type)
	assert.Equal(t, "blog", categoryBlogsField.Relation.TargetSchemaName)
	assert.Equal(t, "categories", categoryBlogsField.Relation.TargetFieldName)
}
