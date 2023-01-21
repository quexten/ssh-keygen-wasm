import "./wasm_exec.js"
    
const go = new Go();
WebAssembly.instantiateStreaming(fetch("ssh-keygen.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
}).then(async () => {
    document.getElementById("log").textContent += ("generating rsa 2048 key..\n\n.")
    var privateKey = await generatePrivateKey("rsa-2048")
    var publicKey = await convertPrivateKeyToPublicKey(privateKey)
    document.getElementById("log").textContent += ("rsa private key: " + privateKey + "\n\n")
    document.getElementById("log").textContent += ("rsa public key:" + publicKey + "\n\n")
    document.getElementById("log").textContent += ("verified: " + verifyPrivateKey(privateKey) + "\n\n")

    document.getElementById("log").textContent += ("generating ed25519 key..." + "\n\n")
    var ed25519PrivateKey = await generatePrivateKey("ed25519")
    var ed25519PublicKey = await convertPrivateKeyToPublicKey(ed25519PrivateKey)
    document.getElementById("log").textContent += ("ed25519 private key: " + ed25519PrivateKey + "\n\n")
    document.getElementById("log").textContent += ("ed25519 public key: "+ ed25519PublicKey + "\n\n")
})