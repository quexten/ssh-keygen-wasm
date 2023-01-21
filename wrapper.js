import "./wasm_exec.js"
import {readFileSync} from "fs"

var getGlobal = function () {
    if (typeof self !== 'undefined') { return self; }
    if (typeof window !== 'undefined') { return window; }
    if (typeof global !== 'undefined') { return global; }
    throw new Error('unable to locate global object');
};
  
var go;
async function start() {
    go = new Go();
    var result;
    try {
        result = await WebAssembly.instantiateStreaming(fetch("./ssh-keygen.wasm"), go.importObject)
    } catch (e) {
        const wasmBuffer = readFileSync('./ssh-keygen.wasm');
        result = await WebAssembly.instantiate(wasmBuffer, go.importObject)
    }
    go.run(result.instance)
}

async function run(callback) {
    if (go === undefined) {
        await start()
    }

    return await callback()
}

export async function generatePrivateKey(type) {
    return run((async() => {return await getGlobal().generatePrivateKey(type)}))
}

export async function convertPrivateKeyToPublicKey(type) {
    return run((async() => {return await getGlobal().convertPrivateKeyToPublicKey(type)}))
}

export async function isPrivateKeyValid(key) {
    return run((async() => {return await getGlobal().isPrivateKeyValid(key)}))
}

export async function isEncryptedPemBlock(key) {
    return run((async() => {return await getGlobal().isPEMEncrypted(key)}))
}