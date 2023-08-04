// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/network_name.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NetworkName int32

const (
	// BTC 0xxx
	NetworkName_BITCOIN_MAINNET_MAINNET NetworkName = 0
	NetworkName_BITCOIN_TESTNET_TESTNET NetworkName = 1
	// EVM 1xxxx
	// Etheruem
	NetworkName_ETHEREUM_MAINNET_MAINNET NetworkName = 10000
	NetworkName_ETHEREUM_TESTNET_GOERLI  NetworkName = 10001
	NetworkName_ETHEREUM_TESTNET_SEPOLIA NetworkName = 10002
	// Polygon
	NetworkName_POLYGON_MAINNET_MAINNET NetworkName = 10003
	NetworkName_POLYGON_TESTNET_MUMBAI  NetworkName = 10004
	// BNB Chain
	NetworkName_BNB_MAINNET_MAINNET NetworkName = 10005
	NetworkName_BNB_TESTNET_TESTNET NetworkName = 10006
	// Avalanche
	NetworkName_AVALANCHE_MAINNET_CCHAIN NetworkName = 10007
	NetworkName_AVAX_TESTNET_FUJI        NetworkName = 10008
	// Gnosis
	NetworkName_GNOSIS_MAINNET_MAINNET NetworkName = 10009
	NetworkName_GNOSIS_TESTNET_CHIADO  NetworkName = 10010
	// Optimism
	NetworkName_Optimism_MAINNET_MAINNET NetworkName = 10011
	NetworkName_Optimism_TESTNET_GOERLI  NetworkName = 10012
	// Arbitrum
	NetworkName_ARBITRUM_MAINNET_MAINNET NetworkName = 10013
	NetworkName_ARBITRUM_TESTNET_GOERLI  NetworkName = 10014
	// MOVE 2xxxx
	//Aptos
	NetworkName_APTOS_MAINNET_MAINNET NetworkName = 20000
	NetworkName_APTOS_TESTNET_TESTNET NetworkName = 20001
	// Sui
	NetworkName_SUI_MAINNET_MAINNET NetworkName = 20002
	NetworkName_SUI_TESTNET_TESTNET NetworkName = 20003
	// SOLANA 3xxxx
	NetworkName_SOLANA_MAINNET_MAINNET NetworkName = 30000
	NetworkName_SOLANA_TESTNET_TESTNET NetworkName = 30001
)

var NetworkName_name = map[int32]string{
	0:     "BITCOIN_MAINNET_MAINNET",
	1:     "BITCOIN_TESTNET_TESTNET",
	10000: "ETHEREUM_MAINNET_MAINNET",
	10001: "ETHEREUM_TESTNET_GOERLI",
	10002: "ETHEREUM_TESTNET_SEPOLIA",
	10003: "POLYGON_MAINNET_MAINNET",
	10004: "POLYGON_TESTNET_MUMBAI",
	10005: "BNB_MAINNET_MAINNET",
	10006: "BNB_TESTNET_TESTNET",
	10007: "AVALANCHE_MAINNET_CCHAIN",
	10008: "AVAX_TESTNET_FUJI",
	10009: "GNOSIS_MAINNET_MAINNET",
	10010: "GNOSIS_TESTNET_CHIADO",
	10011: "Optimism_MAINNET_MAINNET",
	10012: "Optimism_TESTNET_GOERLI",
	10013: "ARBITRUM_MAINNET_MAINNET",
	10014: "ARBITRUM_TESTNET_GOERLI",
	20000: "APTOS_MAINNET_MAINNET",
	20001: "APTOS_TESTNET_TESTNET",
	20002: "SUI_MAINNET_MAINNET",
	20003: "SUI_TESTNET_TESTNET",
	30000: "SOLANA_MAINNET_MAINNET",
	30001: "SOLANA_TESTNET_TESTNET",
}

var NetworkName_value = map[string]int32{
	"BITCOIN_MAINNET_MAINNET":  0,
	"BITCOIN_TESTNET_TESTNET":  1,
	"ETHEREUM_MAINNET_MAINNET": 10000,
	"ETHEREUM_TESTNET_GOERLI":  10001,
	"ETHEREUM_TESTNET_SEPOLIA": 10002,
	"POLYGON_MAINNET_MAINNET":  10003,
	"POLYGON_TESTNET_MUMBAI":   10004,
	"BNB_MAINNET_MAINNET":      10005,
	"BNB_TESTNET_TESTNET":      10006,
	"AVALANCHE_MAINNET_CCHAIN": 10007,
	"AVAX_TESTNET_FUJI":        10008,
	"GNOSIS_MAINNET_MAINNET":   10009,
	"GNOSIS_TESTNET_CHIADO":    10010,
	"Optimism_MAINNET_MAINNET": 10011,
	"Optimism_TESTNET_GOERLI":  10012,
	"ARBITRUM_MAINNET_MAINNET": 10013,
	"ARBITRUM_TESTNET_GOERLI":  10014,
	"APTOS_MAINNET_MAINNET":    20000,
	"APTOS_TESTNET_TESTNET":    20001,
	"SUI_MAINNET_MAINNET":      20002,
	"SUI_TESTNET_TESTNET":      20003,
	"SOLANA_MAINNET_MAINNET":   30000,
	"SOLANA_TESTNET_TESTNET":   30001,
}

func (x NetworkName) String() string {
	return proto.EnumName(NetworkName_name, int32(x))
}

func (NetworkName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c4aafa085c4431e, []int{0}
}

func init() {
	proto.RegisterEnum("mycel.registry.NetworkName", NetworkName_name, NetworkName_value)
}

func init() { proto.RegisterFile("mycel/registry/network_name.proto", fileDescriptor_9c4aafa085c4431e) }

var fileDescriptor_9c4aafa085c4431e = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0x0b, 0x58, 0x0c, 0x12, 0x1a, 0xa6, 0x6a, 0x5a, 0x68, 0xb1, 0xc4, 0x16, 0x09,
	0x7b, 0xc1, 0x13, 0x8c, 0xcd, 0x10, 0x0f, 0xb2, 0xcf, 0x44, 0xbe, 0x54, 0xc0, 0xa6, 0x4a, 0x8b,
	0x55, 0x2c, 0x70, 0x1c, 0x39, 0x46, 0x90, 0xb7, 0x28, 0xf7, 0x9b, 0x41, 0x5c, 0x36, 0x2c, 0xe1,
	0x2d, 0x58, 0x76, 0xc9, 0x12, 0x25, 0xbb, 0x3e, 0x05, 0xea, 0xa4, 0x76, 0x82, 0xa7, 0xab, 0x91,
	0xf5, 0x9d, 0x4f, 0xe7, 0x3f, 0x47, 0x3e, 0xf8, 0x5a, 0x3e, 0xdd, 0x4f, 0x1f, 0xdb, 0x65, 0x7a,
	0x90, 0x4d, 0xaa, 0x72, 0x6a, 0x8f, 0xd2, 0xea, 0x69, 0x51, 0x3e, 0xda, 0x1d, 0x0d, 0xf3, 0xd4,
	0x1a, 0x97, 0x45, 0x55, 0xd0, 0x8b, 0xaa, 0xc4, 0x6a, 0x4a, 0xae, 0xff, 0x38, 0x87, 0x2f, 0xc0,
	0xa2, 0x0c, 0x86, 0x79, 0x4a, 0xb7, 0xf0, 0x86, 0x23, 0x62, 0x57, 0x0a, 0xd8, 0x0d, 0x98, 0x00,
	0xe0, 0x71, 0xf3, 0x12, 0x63, 0x15, 0xc6, 0x3c, 0x8a, 0x4f, 0xe0, 0xe9, 0x4b, 0x10, 0xbd, 0x8a,
	0x37, 0x79, 0xec, 0xf1, 0x90, 0x27, 0x81, 0xa6, 0x1e, 0x02, 0xdd, 0xc6, 0x1b, 0x2d, 0x6e, 0xe4,
	0xbe, 0xe4, 0xa1, 0x2f, 0xc8, 0x73, 0xf8, 0x4f, 0x6e, 0x68, 0xc4, 0x07, 0xd2, 0x17, 0x8c, 0xbc,
	0x50, 0xf2, 0x40, 0xfa, 0xf7, 0xfa, 0x52, 0x4f, 0xf5, 0x12, 0xe8, 0x16, 0xee, 0x35, 0xb4, 0x71,
	0x83, 0x24, 0x70, 0x98, 0x20, 0xaf, 0x80, 0x6e, 0xe2, 0x35, 0x07, 0x1c, 0x4d, 0x7b, 0xdd, 0x92,
	0xee, 0x24, 0x6f, 0x54, 0x1a, 0xb6, 0xc3, 0x7c, 0x06, 0xae, 0xc7, 0x5b, 0xd3, 0x75, 0x3d, 0x26,
	0x80, 0xbc, 0x05, 0xda, 0xc3, 0x97, 0xd8, 0x0e, 0xbb, 0xdb, 0x9a, 0xb7, 0x93, 0x3b, 0x82, 0xbc,
	0x53, 0x39, 0xfa, 0x20, 0x23, 0x11, 0x69, 0xdd, 0xde, 0x03, 0xbd, 0x82, 0xd7, 0x4f, 0x61, 0xa3,
	0xb9, 0x9e, 0x60, 0xb7, 0x24, 0xf9, 0xa0, 0xfa, 0xc9, 0x71, 0x95, 0xe5, 0xd9, 0x24, 0xd7, 0xd4,
	0x8f, 0x6a, 0xfa, 0x16, 0x77, 0x56, 0x57, 0x2f, 0xc2, 0x86, 0x8e, 0x88, 0xc3, 0x33, 0xf6, 0xfe,
	0x49, 0xc9, 0x2d, 0xee, 0xc8, 0x9f, 0x4f, 0x22, 0xaf, 0xb3, 0x41, 0x2c, 0xf5, 0xc4, 0x5f, 0x6a,
	0xb4, 0x84, 0xdd, 0x15, 0x7d, 0xad, 0x11, 0xbd, 0x8c, 0xd7, 0xa2, 0x44, 0x68, 0xde, 0xb7, 0x25,
	0xea, 0x5a, 0xdf, 0x6b, 0x44, 0xb7, 0x71, 0x2f, 0x92, 0x3e, 0x03, 0xa6, 0x89, 0x3f, 0x8f, 0x57,
	0x69, 0xd7, 0xfd, 0x75, 0x8c, 0x1c, 0xef, 0xf7, 0xcc, 0x44, 0x47, 0x33, 0x13, 0xfd, 0x9d, 0x99,
	0xe8, 0x70, 0x6e, 0x1a, 0x47, 0x73, 0xd3, 0xf8, 0x33, 0x37, 0x8d, 0xfb, 0xd6, 0x41, 0x56, 0x3d,
	0x7c, 0xb2, 0x67, 0xed, 0x17, 0xb9, 0xad, 0xfe, 0xef, 0x1b, 0x0f, 0x8a, 0x7c, 0x98, 0x8d, 0x16,
	0x1f, 0xf6, 0xb3, 0xe5, 0x45, 0x54, 0xd3, 0x71, 0x3a, 0xd9, 0x3b, 0xaf, 0x6e, 0xe1, 0xe6, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x47, 0xb0, 0x9d, 0x30, 0x03, 0x00, 0x00,
}
