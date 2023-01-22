import "./wasm_exec.js"
var getGlobal = function () {
    if (typeof self !== 'undefined') { return self; }
    if (typeof window !== 'undefined') { return window; }
    if (typeof global !== 'undefined') { return global; }
    throw new Error('unable to locate global object');
};
if (fs !== undefined) {
    import("fs").then((fs) => {
        getGlobal().readFileSync = fs.readFileSync
    })
}

async function loadWasmBinary(path) {
    const wasmPath = path || 'scripts/ssh-keygen.wasm';
    if (typeof require === 'function') {
        return Promise.resolve(require(wasmPath).then(
            (wasmModule) => {
                return decodeWasmBinary(wasmModule);
            }
        ))
    }

    try {
        return await WebAssembly.instantiateStreaming(fetch(wasmPath), go.importObject)
    } catch (e) {
        const wasmBuffer = getGlobal().readFileSync(wasmPath);
        return await WebAssembly.instantiate(wasmBuffer, go.importObject)
    }
}

var go;
export async function start(wasmPath) {
    go = new Go();

    var result = await loadWasmBinary(wasmPath);
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