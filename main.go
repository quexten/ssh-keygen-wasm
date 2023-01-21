package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("generatePrivateKey", js.FuncOf(jsGeneratePrivateKey))
	js.Global().Set("verifyPrivateKey", js.FuncOf(jsVerifyPrivateKey))
	js.Global().Set("convertPrivateKeyToPublicKey", js.FuncOf(jsConvertPrivateKeyToPublicKey))
	<-done
}

func jsGeneratePrivateKey(this js.Value, args []js.Value) interface{} {
	keyType := args[0].String()
	privateKeyPem, err := generatePrivateKey(keyType)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return privateKeyPem
}

func jsVerifyPrivateKey(this js.Value, args []js.Value) interface{} {
	key := args[0].String()
	return verifyPrivateKey(key, nil)
}

func jsConvertPrivateKeyToPublicKey(this js.Value, args []js.Value) interface{} {
	keyPem := args[0].String()

	publicKey, err := privateKeyToPublicKey(keyPem, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return publicKey
}
