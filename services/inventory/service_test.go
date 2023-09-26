package inventory

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
)

func TestService_Get(t *testing.T) {
	tests := []struct {
		desc        string
		item        string
		expectedOp  int
		expectedErr error
	}{
		{
			desc:        "Success: item is present",
			item:        "Earbuds",
			expectedOp:  7,
			expectedErr: nil,
		},
	}

	for i, tc := range tests {
		app := gofr.New()
		r := httptest.NewRequest(http.MethodGet, "/inventory", nil)
		req := request.NewHTTPRequest(r)
		ctx := gofr.NewContext(nil, req, app)

		svc := New()
		resp, err := svc.Get(ctx, tc.item)

		if !reflect.DeepEqual(resp, tc.expectedOp) {
			t.Errorf("Test: %v, Expected: %v, Got: %v", i, tc.expectedOp, resp)
		}

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Test: %v, Expected: %v, Got: %v", i, tc.expectedErr, err)
		}
	}
}
