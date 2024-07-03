#ifndef ASPECT_H
#define ASPECT_H

#include <wasm.h>
#include <wasmtime/error.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void* ptr;
    size_t len;
} WasmInstrumentResult;

WasmInstrumentResult wasm_instrument(const unsigned char* raw_module, size_t len);
void wasm_instrument_free(void* ptr);

/**
 * \brief Validate a WebAssembly binary.
 *
 * This function will validate the provided byte sequence to determine if it is
 * a valid WebAssembly binary within the context of the engine provided.
 *
 * This function does not take ownership of its arguments but the caller is
 * expected to deallocate the returned error if it is non-`NULL`.
 *
 * If the binary validates then `NULL` is returned, otherwise the error returned
 * describes why the binary did not validate.
 */
WASM_API_EXTERN wasmtime_error_t *
aspect_validate(const uint8_t *wasm, size_t wasm_len);

#ifdef __cplusplus
}
#endif

#endif // #ifdef ASPECT_H
