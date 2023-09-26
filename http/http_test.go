package http

import (
	"github.com/golang/mock/gomock"

	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"

	"inventory/services"
)

func TestHandler_Index(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := services.NewMockOrderSvc(ctrl)

	svc := New(serviceMock)

	tests := []struct {
		desc        string
		url         string
		item        string
		expectedOp  interface{}
		expectedErr error
		calls       []*gomock.Call
	}{
		{
			desc:        "Success: item is present",
			url:         "/inventory?item=Earbuds",
			item:        "Earbuds",
			expectedOp:  7,
			expectedErr: nil,
			calls: []*gomock.Call{
				serviceMock.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), "Earbuds").
					Return(7, nil),
			},
		},
	}

	for i, tc := range tests {
		app := gofr.New()
		r := httptest.NewRequest(http.MethodGet, tc.url, nil)
		req := request.NewHTTPRequest(r)
		ctx := gofr.NewContext(nil, req, app)

		resp, err := svc.Index(ctx)

		if !reflect.DeepEqual(resp, tc.expectedOp) {
			t.Errorf("Test: %v, Expected: %v, Got: %v", i, tc.expectedOp, resp)
		}

		if !reflect.DeepEqual(err, tc.expectedErr) {
			t.Errorf("Test: %v, Expected: %v, Got: %v", i, tc.expectedErr, err)
		}
	}
}
