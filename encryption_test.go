package ravepay

import (
	"testing"
)

func Test_getEncryptionKey(t *testing.T) {
	tests := []struct {
		name   string
		seckey string
		want   string
	}{
		{
			name: "return empty if string is empty",
		},
		{
			name:   "return empty if string is less than 12",
			seckey: "1234567",
		},
		{
			name:   "returns the expected encryption key - 1",
			seckey: "FLWSECK-bb971402072265fb156e90a3578fe5e6-X",
			want:   "bb9714020722eb4cf7a169f2",
		},
		{
			name:   "returns the expected encryption key - 2",
			seckey: "FLWSECK-6b32914d4d60c10d0ef72bdad734134a-X",
			want:   "6b32914d4d60cb85d8eb73db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncryptionKey(tt.seckey); got != tt.want {
				t.Errorf("getKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tripleDESEncrypt(t *testing.T) {
	type args struct {
		payload []byte
		key     []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns the expected 3DES encrypted payload - 1",
			args: args{
				payload: []byte("A 16 byte string"),
				key:     []byte("6b32914d4d60cb85d8eb73db"),
			},
			want: "9fx+9uGjG+Oikq8syKpfeg==",
		},
		{
			name: "returns the expected 3DES encrypted payload - 2",
			args: args{
				payload: []byte("Hello world"),
				key:     []byte("bb9714020722eb4cf7a169f2"),
			},
			want: "Lgk7z/IvTT9mx3t9vOzHmg==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tripleDESEncrypt(tt.args.payload, tt.args.key); got != tt.want {
				t.Errorf("tripleDESEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
