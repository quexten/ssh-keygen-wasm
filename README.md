# Go-based WASM module to read, decrypt, generate SSH keys in the browser
This module is a small wrapper around Go's ssh library.

Compile the wasm module using `GOOS=js GOARCH=wasm go build -o ssh-keygen.wasm main.go functions.go`

Example usage in typescrypt:
```ts
import { generatePrivateKey } from "ssh-keygen-wasm";

// in an async func:
  var privateKey = generatePrivateKey("ed25519") // will be returned encoded in PKCS#8, ready for openssh usage
  console.log(privateKey)

```