package wasmtime

// #include <instrument.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func Instrument(rawModule []byte) ([]byte, error) {
	cRawModule := (*C.uchar)(unsafe.Pointer(&rawModule[0]))
	len := C.size_t(len(rawModule))

	cResult := C.wasm_instrument(cRawModule, len)
	if cResult.ptr == nil {
		return nil, fmt.Errorf("wasm_instrument failed")
	}

	output := C.GoBytes(cResult.ptr, C.int(cResult.len))
	C.wasm_instrument_free(cResult.ptr)

	return output, nil
}
