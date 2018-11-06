// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/seizadi/cmdb/pkg/pb/cmdb.proto

package pb // import "github.com/seizadi/cmdb/pkg/pb"

import fmt "fmt"
import http "net/http"
import json "encoding/json"
import ioutil "io/ioutil"
import bytes "bytes"
import context "golang.org/x/net/context"
import metadata "google.golang.org/grpc/metadata"
import runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
import validate_runtime "github.com/infobloxopen/protoc-gen-atlas-validate/runtime"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// validate_Cmdb_GetVersion_0 is an entrypoint for validating "GET" HTTP request
// that match *.pb.gw.go/pattern_Cmdb_GetVersion_0.
func validate_Cmdb_GetVersion_0(r json.RawMessage) (err error) {
	if len(r) != 0 {
		return fmt.Errorf("Body is not allowed")
	}
	return nil
}

// validate_Object_VersionResponse function validates a JSON for a given object.
func validate_Object_VersionResponse(r json.RawMessage, path string, allowUnknown bool) (err error) {
	obj := &VersionResponse{}
	if hook, ok := interface{}(obj).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	var v map[string]json.RawMessage
	if err = json.Unmarshal(r, &v); err != nil {
		return fmt.Errorf("Invalid value for %q: expected object.", path)
	}
	for k, _ := range v {
		switch k {
		case "version":
		default:
			if !allowUnknown {
				return fmt.Errorf("Unknown field %q", validate_runtime.JoinPath(path, k))
			}
		}
	}
	return nil
}

// AtlasValidateJSON function validates a JSON for object VersionResponse.
func (o *VersionResponse) AtlasValidateJSON(r json.RawMessage, path string, allowUnknown bool) (err error) {
	if hook, ok := interface{}(o).(interface {
		AtlasJSONValidate(json.RawMessage, string, bool) (json.RawMessage, error)
	}); ok {
		if r, err = hook.AtlasJSONValidate(r, path, allowUnknown); err != nil {
			return err
		}
	}
	return validate_Object_VersionResponse(r, path, allowUnknown)
}

var validate_Patterns = []struct {
	pattern    runtime.Pattern
	httpMethod string
	validator  func(json.RawMessage) error
	// Included for introspection purpose.
	allowUnknown bool
}{
	{
		pattern:      pattern_Cmdb_GetVersion_0,
		httpMethod:   "GET",
		validator:    validate_Cmdb_GetVersion_0,
		allowUnknown: false,
	},
}

// AtlasValidateAnnotator parses JSON input and validates unknown fields
// based on 'allow_unknown_fields' options specified in proto file.
func AtlasValidateAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)
	for _, v := range validate_Patterns {
		if r.Method == v.httpMethod && validate_runtime.PatternMatch(v.pattern, r.URL.Path) {
			var b []byte
			var err error
			if b, err = ioutil.ReadAll(r.Body); err != nil {
				md.Set("Atlas-Validation-Error", "Invalid value: unable to parse body")
				return md
			}
			r.Body = ioutil.NopCloser(bytes.NewReader(b))
			if err = v.validator(b); err != nil {
				md.Set("Atlas-Validation-Error", err.Error())
			}
			break
		}
	}
	return md
}
