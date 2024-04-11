package rpc

import (
	"testing"
)

func TestEncodeMesssage(t *testing.T) {
	type args struct {
		msg any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.

		{
			name: "TestEncodeMesssage",
			args: args{
				msg: BaseMessage{Method: "lol"},
			},
			want: "Content-Length: 16\r\n\r\n{\"method\":\"lol\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMessage(tt.args.msg); got != tt.want {
				t.Errorf("EncodeMesssage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeMessage(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestDecodeMessageSucess",
			args: args{
				content: []byte("Content-Length: 16\r\n\r\n{\"method\":\"lol\"}"),
			},
			want:    "lol",
			want1:   []byte("{\"method\":\"lol\"}"),
			wantErr: false,
		},
		{
			name: "TestDecodeMessageFail",
			args: args{
				content: []byte("Content-Length: hghg\r\n\r\n{\"method\":\"lol\"}"),
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method, content, err := DecodeMessage(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if method != tt.want {
				t.Errorf("DecodeMessage() got1 = %v, want %v", method, tt.want)
			}
			if string(content) != string(tt.want1) {
				t.Errorf("DecodeMessage() got2 = %v, want %v", string(content), string(tt.want1))
			}

		})
	}
}
