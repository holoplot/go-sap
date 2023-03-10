package sap

import (
	"net"
	"reflect"
	"testing"
)

func TestDecodePacket(t *testing.T) {
	tests := []struct {
		name    string
		raw     []byte
		want    *Packet
		wantErr bool
	}{
		{
			name: "1",
			raw: []byte{
				0x20, 0x00, 0x00, 0x01, 0xc0, 0xa8,
				0x64, 0xfe, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x64,
				0x70, 0x00, 0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
				0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
				0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
				0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
				0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
				0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
				0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
				0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
				0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
				0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
				0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
				0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
				0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
				0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
				0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
				0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
				0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
				0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
				0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
				0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
				0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
				0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
				0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
				0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
				0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
				0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
				0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
				0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
				0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
				0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
				0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
				0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
				0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
				0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
				0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
				0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
			},
			want: &Packet{
				Type:        MessageTypeAnnouncement,
				IDHash:      0x0001,
				Origin:      net.ParseIP("192.168.100.254"),
				PayloadType: SDPPayloadType,
				Payload: []byte{
					0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
					0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
					0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
					0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
					0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
					0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
					0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
					0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
					0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
					0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
					0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
					0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
					0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
					0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
					0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
					0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
					0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
					0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
					0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
					0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
					0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
					0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
					0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
					0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
					0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
					0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
					0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
					0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
					0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
					0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
					0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodePacket(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodePacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacket_Encode(t *testing.T) {
	type fields struct {
		Type               MessageType
		IDHash             uint16
		Origin             net.IP
		Encrypted          bool
		Compressed         bool
		PayloadType        string
		AuthenticationData []byte
		Payload            []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				Type:        MessageTypeAnnouncement,
				IDHash:      0x0001,
				Origin:      net.ParseIP("192.168.100.254"),
				PayloadType: SDPPayloadType,
				Payload: []byte{
					0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
					0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
					0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
					0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
					0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
					0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
					0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
					0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
					0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
					0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
					0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
					0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
					0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
					0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
					0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
					0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
					0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
					0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
					0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
					0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
					0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
					0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
					0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
					0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
					0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
					0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
					0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
					0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
					0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
					0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
					0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				},
			},
			want: []byte{
				0x20, 0x00, 0x00, 0x01, 0xc0, 0xa8,
				0x64, 0xfe, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x64,
				0x70, 0x00, 0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
				0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
				0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
				0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
				0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
				0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
				0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
				0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
				0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
				0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
				0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
				0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
				0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
				0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
				0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
				0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
				0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
				0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
				0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
				0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
				0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
				0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
				0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
				0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
				0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
				0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
				0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
				0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
				0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
				0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
				0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
				0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
				0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
				0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
				0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
				0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
				0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Packet{
				Type:               tt.fields.Type,
				IDHash:             tt.fields.IDHash,
				Origin:             tt.fields.Origin,
				Encrypted:          tt.fields.Encrypted,
				Compressed:         tt.fields.Compressed,
				PayloadType:        tt.fields.PayloadType,
				AuthenticationData: tt.fields.AuthenticationData,
				Payload:            tt.fields.Payload,
			}
			got, err := p.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Packet.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Packet.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPacket_Rencode(t *testing.T) {
	type fields struct {
		Type               MessageType
		IDHash             uint16
		Origin             net.IP
		Encrypted          bool
		Compressed         bool
		PayloadType        string
		AuthenticationData []byte
		Payload            []byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "uncompressed",
			fields: fields{
				Type:        MessageTypeAnnouncement,
				IDHash:      0x0001,
				Origin:      net.ParseIP("192.168.100.254"),
				PayloadType: SDPPayloadType,
				Payload: []byte{
					0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
					0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
					0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
					0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
					0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
					0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
					0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
					0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
					0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
					0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
					0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
					0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
					0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
					0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
					0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
					0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
					0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
					0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
					0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
					0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
					0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
					0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
					0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
					0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
					0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
					0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
					0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
					0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
					0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
					0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
					0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				},
			},
		},
		{
			name: "compressed",
			fields: fields{
				Type:        MessageTypeAnnouncement,
				IDHash:      0x0001,
				Origin:      net.ParseIP("192.168.100.254"),
				PayloadType: SDPPayloadType,
				Payload: []byte{
					0x76, 0x3d, 0x30, 0x0d, 0x0a, 0x6f, 0x3d, 0x2d, 0x20, 0x30, 0x32, 0x38, 0x34, 0x34,
					0x32, 0x34, 0x37, 0x31, 0x38, 0x30, 0x30, 0x30, 0x31, 0x20, 0x30, 0x20, 0x49, 0x4e, 0x20, 0x49,
					0x50, 0x34, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x73, 0x3d, 0x54, 0x58, 0x2d, 0x31, 0x2d, 0x44, 0x4e, 0x54, 0x31, 0x2d,
					0x38, 0x0d, 0x0a, 0x74, 0x3d, 0x30, 0x20, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x63, 0x6c, 0x6f, 0x63,
					0x6b, 0x2d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x50, 0x54, 0x50, 0x76, 0x32, 0x20, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x74, 0x73, 0x2d, 0x72, 0x65, 0x66, 0x63, 0x6c, 0x6b, 0x3a, 0x70, 0x74,
					0x70, 0x3d, 0x49, 0x45, 0x45, 0x45, 0x31, 0x35, 0x38, 0x38, 0x2d, 0x32, 0x30, 0x30, 0x38, 0x3a,
					0x43, 0x38, 0x2d, 0x30, 0x44, 0x2d, 0x33, 0x32, 0x2d, 0x46, 0x46, 0x2d, 0x46, 0x45, 0x2d, 0x34,
					0x43, 0x2d, 0x38, 0x35, 0x2d, 0x34, 0x32, 0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64,
					0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
					0x61, 0x3d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x44, 0x55, 0x50, 0x20, 0x72, 0x61, 0x30, 0x20,
					0x72, 0x61, 0x31, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35, 0x30, 0x30,
					0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a, 0x63, 0x3d,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
					0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20, 0x49, 0x4e,
					0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32, 0x35, 0x34,
					0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39, 0x38, 0x20,
					0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61, 0x3d, 0x6d,
					0x69, 0x64, 0x3a, 0x72, 0x61, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x63,
					0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76, 0x6f, 0x6e,
					0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e, 0x31, 0x32,
					0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30,
					0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64, 0x69, 0x72,
					0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a, 0x6d, 0x3d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x20, 0x35,
					0x30, 0x30, 0x34, 0x20, 0x52, 0x54, 0x50, 0x2f, 0x41, 0x56, 0x50, 0x20, 0x39, 0x38, 0x0d, 0x0a,
					0x63, 0x3d, 0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x2e, 0x31, 0x2f, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x6f, 0x75, 0x72,
					0x63, 0x65, 0x2d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x20,
					0x49, 0x4e, 0x20, 0x49, 0x50, 0x34, 0x20, 0x32, 0x33, 0x39, 0x2e, 0x32, 0x30, 0x30, 0x2e, 0x32,
					0x35, 0x34, 0x2e, 0x31, 0x20, 0x31, 0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x32, 0x30, 0x30,
					0x2e, 0x32, 0x35, 0x34, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x74, 0x70, 0x6d, 0x61, 0x70, 0x3a, 0x39,
					0x38, 0x20, 0x4c, 0x32, 0x34, 0x2f, 0x34, 0x38, 0x30, 0x30, 0x30, 0x2f, 0x38, 0x0d, 0x0a, 0x61,
					0x3d, 0x6d, 0x69, 0x64, 0x3a, 0x72, 0x61, 0x31, 0x0d, 0x0a, 0x61, 0x3d, 0x66, 0x72, 0x61, 0x6d,
					0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x36, 0x0d, 0x0a, 0x61, 0x3d, 0x72, 0x65, 0x63, 0x76,
					0x6f, 0x6e, 0x6c, 0x79, 0x0d, 0x0a, 0x61, 0x3d, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x30, 0x2e,
					0x31, 0x32, 0x35, 0x0d, 0x0a, 0x61, 0x3d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x74, 0x69, 0x6d, 0x65,
					0x3a, 0x30, 0x0d, 0x0a, 0x61, 0x3d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x63, 0x6c, 0x6b, 0x3a, 0x64,
					0x69, 0x72, 0x65, 0x63, 0x74, 0x3d, 0x30, 0x0d, 0x0a,
				},
				Compressed: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Packet{
				Type:               tt.fields.Type,
				IDHash:             tt.fields.IDHash,
				Origin:             tt.fields.Origin,
				Encrypted:          tt.fields.Encrypted,
				Compressed:         tt.fields.Compressed,
				PayloadType:        tt.fields.PayloadType,
				AuthenticationData: tt.fields.AuthenticationData,
				Payload:            tt.fields.Payload,
			}

			raw, err := p.Encode()
			if err != nil {
				t.Errorf("Packet.Encode() error = %v", err)
				return
			}

			back, err := DecodePacket(raw)
			if err != nil {
				t.Errorf("Packet.Encode() error = %v", err)
				return
			}

			if !reflect.DeepEqual(p, back) {
				t.Errorf("Packet.Encode() = %v, want %v", p, back)
			}
		})
	}
}
