// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbmesh/v2beta1/gateway_class.proto

package meshv2beta1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *GatewayClass) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *GatewayClass) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *ParametersReference) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *ParametersReference) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
