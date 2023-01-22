export function generatePrivateKey(keyType: string): Promise<string>;
export function convertPrivateKeyToPublicKey(key: string): Promise<string>;
export function isPrivateKeyValid(key: string): Promise<boolean>;
export function isEncryptedPemBlock(key: string): Promise<boolean>;
export function start(wasmPath?: string): Promise<void>; 