package wasmtime

// #include <aspect.h>
import "C"
import (
	"fmt"
	"runtime"
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

// AspectValidate validates whether `wasm` would be a valid wasm module according to the
// configuration in `store`
func AspectValidate(wasm []byte) error {
	var wasmPtr *C.uint8_t
	if len(wasm) > 0 {
		wasmPtr = (*C.uint8_t)(unsafe.Pointer(&wasm[0]))
	}
	err := C.aspect_validate(wasmPtr, C.size_t(len(wasm)))
	runtime.KeepAlive(wasm)
	if err == nil {
		return nil
	}

	return mkError(err)
}
