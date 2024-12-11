// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bookstore2

import (
	"github.com/nguyenhoaibao/stainless-bookstore-demo-go/internal/apijson"
)

type EchoParams struct {
	Body string `json:"body,required"`
}

func (r EchoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
