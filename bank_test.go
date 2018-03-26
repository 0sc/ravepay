package rave

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListBanks(t *testing.T) {
	handler := &testServer{}
	server := httptest.NewServer(handler)
	defer server.Close()

	tests := []struct {
		name    string
		want    []Bank
		wantErr bool
	}{
		{
			name: "returns a list of banks",
			want: []Bank{
				Bank{
					Name:            "ACCESS BANK NIGERIA",
					Code:            "044",
					Internetbanking: false,
				},
				Bank{
					Name:            "ECOBANK NIGERIA PLC",
					Code:            "050",
					Internetbanking: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.resp = []byte(listBanksResp)
			banksURL = server.URL

			got, err := ListBanks()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBanks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBanks() = %v, want %v", got, tt.want)
			}
		})
	}
}

var listBanksResp = `[{"bankname":"ACCESS BANK NIGERIA","bankcode":"044","internetbanking":false},{"bankname":"ECOBANK NIGERIA PLC","bankcode":"050","internetbanking":false}]`
