package statsig_go_core_macos_aarch64

import (
	_ "embed"
)

//go:embed libstatsig_ffi.dylib
var binaryData []byte

//go:embed libstatsig_ffi.dylib.sig
var signatureData []byte

func GetBinaryData() []byte {
	return binaryData
}

func GetSignatureData() []byte {
	return signatureData
}
