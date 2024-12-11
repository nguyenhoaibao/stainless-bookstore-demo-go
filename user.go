// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bookstore2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/apijson"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/apiquery"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/param"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/requestconfig"
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/option"
)

// UserService contains methods and other services that help with interacting with
// the bookstore-2 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

// Some description of the operation. You can use `markdown` here.
func (r *UserService) Get(ctx context.Context, username string, query UserGetParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("users/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This can only be done by the logged in user.
func (r *UserService) Update(ctx context.Context, username string, params UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("users/%s", username)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

type User struct {
	// User email address
	Email string `json:"email" format:"test"`
	// User first name
	FirstName string `json:"firstName"`
	// User last name
	LastName string `json:"lastName"`
	// User supplied username
	Username string   `json:"username"`
	JSON     userJSON `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserParam struct {
	// User email address
	Email param.Field[string] `json:"email" format:"test"`
	// User first name
	FirstName param.Field[string] `json:"firstName"`
	// User last name
	LastName param.Field[string] `json:"lastName"`
	// User supplied username
	Username param.Field[string] `json:"username"`
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGetParams struct {
	// Pretty print response
	PrettyPrint param.Field[bool] `query:"pretty_print"`
	// Filter users without email
	WithEmail param.Field[bool] `query:"with_email"`
}

// URLQuery serializes [UserGetParams]'s query parameters as `url.Values`.
func (r UserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserUpdateParams struct {
	User UserParam `json:"user,required"`
	// Pretty print response
	PrettyPrint param.Field[bool] `query:"pretty_print"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.User)
}

// URLQuery serializes [UserUpdateParams]'s query parameters as `url.Values`.
func (r UserUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
