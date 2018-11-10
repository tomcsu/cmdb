// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/seizadi/cmdb/pkg/pb/cmdb.proto

package pb // import "github.com/seizadi/cmdb/pkg/pb"

import options "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
import query "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
import _ "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.

var CmdbMethodsRequireFilteringValidation = map[string]map[string]options.FilteringOption{
	"/api.cmdb.AwsServices/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Containers/List": map[string]options.FilteringOption{
		"id":                options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":              options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description":       options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"container_name":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"image_repo":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"image_tag":         options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"image_pull_policy": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"digest":            options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Manifests/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.VersionTags/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"version":     options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"repo":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"commit":      options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Artifacts/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.AwsRdsInstances/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Deployments/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Environments/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.KubeClusters/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Regions/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Secrets/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"type":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"key":         options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"vault_id":    options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Vaults/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"path":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
	"/api.cmdb.Applications/List": map[string]options.FilteringOption{
		"id":          options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"name":        options.FilteringOption{ValueType: options.QueryValidate_STRING},
		"description": options.FilteringOption{ValueType: options.QueryValidate_STRING},
	},
}
var CmdbMethodsRequireSortingValidation = map[string][]string{
	"/api.cmdb.AwsServices/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Containers/List": []string{
		"id",
		"name",
		"description",
		"container_name",
		"image_repo",
		"image_tag",
		"image_pull_policy",
		"digest",
	},
	"/api.cmdb.Manifests/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.VersionTags/List": []string{
		"id",
		"name",
		"description",
		"version",
		"repo",
		"commit",
	},
	"/api.cmdb.Artifacts/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.AwsRdsInstances/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Deployments/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Environments/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.KubeClusters/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Regions/List": []string{
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Secrets/List": []string{
		"id",
		"name",
		"description",
		"type",
		"key",
		"vault_id",
	},
	"/api.cmdb.Vaults/List": []string{
		"id",
		"name",
		"description",
		"path",
	},
	"/api.cmdb.Applications/List": []string{
		"id",
		"name",
		"description",
	},
}
var CmdbMethodsRequireFieldSelectionValidation = map[string][]string{
	"/api.cmdb.AwsServices/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.AwsServices/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Containers/Read": {
		"id",
		"name",
		"description",
		"container_name",
		"image_repo",
		"image_tag",
		"image_pull_policy",
		"digest",
	},
	"/api.cmdb.Containers/List": {
		"id",
		"name",
		"description",
		"container_name",
		"image_repo",
		"image_tag",
		"image_pull_policy",
		"digest",
	},
	"/api.cmdb.Manifests/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Manifests/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.VersionTags/Read": {
		"id",
		"name",
		"description",
		"version",
		"repo",
		"commit",
	},
	"/api.cmdb.VersionTags/List": {
		"id",
		"name",
		"description",
		"version",
		"repo",
		"commit",
	},
	"/api.cmdb.Artifacts/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Artifacts/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.AwsRdsInstances/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.AwsRdsInstances/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Deployments/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Deployments/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Environments/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Environments/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.KubeClusters/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.KubeClusters/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Regions/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Regions/List": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Secrets/Read": {
		"id",
		"name",
		"description",
		"type",
		"key",
		"vault_id",
	},
	"/api.cmdb.Secrets/List": {
		"id",
		"name",
		"description",
		"type",
		"key",
		"vault_id",
	},
	"/api.cmdb.Vaults/Read": {
		"id",
		"name",
		"description",
		"path",
		"secrets.id",
		"secrets.name",
		"secrets.description",
		"secrets.type",
		"secrets.key",
		"secrets.vault_id",
		"secrets",
	},
	"/api.cmdb.Vaults/List": {
		"id",
		"name",
		"description",
		"path",
		"secrets.id",
		"secrets.name",
		"secrets.description",
		"secrets.type",
		"secrets.key",
		"secrets.vault_id",
		"secrets",
	},
	"/api.cmdb.Applications/Read": {
		"id",
		"name",
		"description",
	},
	"/api.cmdb.Applications/List": {
		"id",
		"name",
		"description",
	},
}

func CmdbValidateFiltering(methodName string, f *query.Filtering) error {
	info, ok := CmdbMethodsRequireFilteringValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFiltering(f, info)
}
func CmdbValidateSorting(methodName string, s *query.Sorting) error {
	info, ok := CmdbMethodsRequireSortingValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateSorting(s, info)
}
func CmdbValidateFieldSelection(methodName string, s *query.FieldSelection) error {
	info, ok := CmdbMethodsRequireFieldSelectionValidation[methodName]
	if !ok {
		return nil
	}
	return options.ValidateFieldSelection(s, info)
}
