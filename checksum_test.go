package ravepay

import "testing"

func TestCalculateChecksum(t *testing.T) {
	type args struct {
		payload interface{}
		prefix  []byte
		suffix  []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Calculates the correct checksum val",
			args: args{
				payload: Payment{
					Amount:            20,
					PaymentMethod:     "both",
					CustomDescription: "Pay Internet",
					CustomLogo:        "http://localhost/payporte-3/skin/frontend/ultimo/shoppy/custom/images/logo.svg",
					CustomTitle:       "Shoppy Global systems",
					Country:           "NG",
					Currency:          "NGN",
					CustomerEmail:     "user@example.com",
					CustomerFirstname: "Temi",
					CustomerLastname:  "Adelewa",
					CustomerPhone:     "234099940409",
					TxRef:             "MG-1500041286295",
				},
				prefix: []byte("FLWPUBK-e634d14d9ded04eaf05d5b63a0a06d2f-X"),
				suffix: []byte("FLWSECK-bb971402072265fb156e90a3578fe5e6-X"),
			},
			want: "a14ac4eba0902e8fd6b5fdf542f46d6efc18885a63c3d5f100c26715c7c8d8f4",
		},
		{
			name: "Calculates the correct checksum val ignoring unexported fields",
			args: args{
				payload: struct{ A, b string }{A: "Hello", b: "world"},
				prefix:  []byte("x"),
				suffix:  []byte("x"),
			},
			want: CalculateChecksum(struct{ A string }{"Hello"}, []byte("x"), []byte("x")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateChecksum(tt.args.payload, tt.args.prefix, tt.args.suffix); got != tt.want {
				t.Errorf("CalculateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
