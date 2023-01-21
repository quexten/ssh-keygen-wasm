package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ssh"
)

func generatePrivateKey(keyType string) (string, error) {
	switch keyType {
	case "rsa-2048":
		rsaKey, err := generateRSAKey(2048)
		if err != nil {
			return "", err
		}
		return string(encodeRSAPrivateKeyToPEM(rsaKey)), nil
	case "rsa-4096":
		rsaKey, err := generateRSAKey(4096)
		if err != nil {
			return "", err
		}
		return string(encodeRSAPrivateKeyToPEM(rsaKey)), nil
	case "ed25519":
		ed25519Key, err := generateEd25519Key()
		if err != nil {
			return "", err
		}
		return string(encodeED25519PrivateKeyToPem(ed25519Key)), nil
	default:
		return "Invalid key type", nil
	}
}

func generateRSAKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func generateEd25519Key() (ed25519.PrivateKey, error) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func encodeRSAPrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	bytes, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	privBlock := pem.Block{
		Type:    "PRIVATE KEY",
		Headers: nil,
		Bytes:   bytes,
	}
	return pem.EncodeToMemory(&privBlock)
}

func encodeED25519PrivateKeyToPem(privateKey ed25519.PrivateKey) []byte {
	privBlock := pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privateKey),
	}
	return pem.EncodeToMemory(&privBlock)
}

func privateKeyToPublicKey(privateKeyPem string, password *string) (string, error) {
	var key interface{}
	var err error
	if password == nil {
		key, err = ssh.ParseRawPrivateKey([]byte(privateKeyPem))
	} else {
		key, err = ssh.ParseRawPrivateKeyWithPassphrase([]byte(privateKeyPem), []byte(*password))
	}

	if err != nil {
		return "", err
	}

	var pubKey ssh.PublicKey
	switch key.(type) {
	case (*rsa.PrivateKey):
		pubKey, err = ssh.NewPublicKey((key.(*rsa.PrivateKey)).Public())
		break
	case (*ed25519.PrivateKey):
		pubKey, err = ssh.NewPublicKey((key.(*ed25519.PrivateKey)).Public())
		break
	default:
		return "", errors.New("Unsupported private key type")
	}
	if err != nil {
		return "", err
	}

	return string(ssh.MarshalAuthorizedKey(pubKey)), nil
}

func verifyPrivateKey(privateKeyPem string, password *string) bool {
	var key interface{}
	var err error

	if password == nil {
		key, err = ssh.ParseRawPrivateKey([]byte(privateKeyPem))
	} else {
		key, err = ssh.ParseRawPrivateKeyWithPassphrase([]byte(privateKeyPem), []byte(*password))
	}
	if err != nil {
		return false
	}

	switch key.(type) {
	case (*rsa.PrivateKey):
		return (key.(*rsa.PrivateKey)).Validate() == nil
	}

	return true
}

func isEncryptedPemBlock(privateKeyPem string) bool {
	block, _ := pem.Decode([]byte(privateKeyPem))
	return strings.Contains(block.Headers["Proc-Type"], "ENCRYPTED")
}
