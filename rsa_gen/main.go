package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating RSA key pair:", err)
		return
	}

	// Message to be signed
	message := []byte("Your message goes here")

	// Sign the message
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	} else {
		signatureBase64 := base64.StdEncoding.EncodeToString(signature)
		fmt.Println("Sign message:", signatureBase64)
	}

	// Verify the signature
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Signature verification failed. Message may have been tampered with.")
		return
	}

	fmt.Println("Signature verified. Message is authentic.")

	////

	// Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Convert PEM block to PEM format
	// privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)

	// Encode PEM block to Base64
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEM.Bytes)

	// Add BEGIN and END headers and footers to private key
	finalPrivateKey := "-----BEGIN RSA PRIVATE KEY-----\n" + privateKeyBase64 + "\n-----END RSA PRIVATE KEY-----"

	// Encode public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}

	// Convert PEM block to PEM format
	// publicKeyPEMBytes := pem.EncodeToMemory(publicKeyPEM)

	// Encode PEM block to Base64
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyPEM.Bytes)

	// Add BEGIN and END headers and footers to public key
	finalPublicKey := "-----BEGIN RSA PUBLIC KEY-----\n" + publicKeyBase64 + "\n-----END RSA PUBLIC KEY-----"

	fmt.Println("Private Key (Base64 with Headers and Footers):")
	fmt.Println(finalPrivateKey)

	fmt.Println("\nPublic Key (Base64 with Headers and Footers):")
	fmt.Println(finalPublicKey)

	///
	decodePublicKey()
}

func decodePublicKey() {
	// PEM encoded public key
	// 	var pemPublicKey = `-----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoZ67dtUTLxoXnNEzRBFB
	// mwukEJGC+y69cGgpNbtElQj3m4Aft/7cu9qYbTNguTSnCDt7uovZNb21u1vpZwKH
	// yVgFEGO4SA8RNnjhJt2D7z8RDMWX3saody7jo9TKlrPABLZGo2o8vadW8Dly/v+I
	// d0YDheCkVCoCEeUjQ8koXZhTwhYkGPu+vkdiqX5cUaiVTu1uzt591aO5Vw/hV4DI
	// hFKnOTnYXnpXiwRwtPyYoGTa64yWfi2t0bv99qz0BgDjQjD0civCe8LRXGGhyB1U
	// 1aHjDDGEnulTYJyEqCzNGwBpzEHUjqIOXElFjt55AFGpCHAuyuoXoP3gQvoSj6RC
	// sQIDAQAB
	// -----END PUBLIC KEY-----`

	//Works also:
	var pemPublicKey = `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAv/udjSyDH7RcbBybzELwtMWZfDjnndJy2YlzPCR9aJCORpbXwS2zKmroWJ9CyaHY5p11KXIfADlW0WamhpNebMVXYj9lZi/l3AdL09XhE4SrPo9lm6+UQoPEtPMCW8BtC1d+xsBaXE/PG28TKeQ8Hg1oAW7EefmybPlmyjQIM8Nmkn9uqC6tyera3defIX250/oXPiaXCj4CyTkmMNUvoBtvP3B/VdnGLvh5DDOKWAIVitclCmzAPKrt4qevDar3zjRQVI1mrcd8DQHytGyNwb2tFUSL++Zc4g+70CpSF5JO9jH1XdU3/4nHB4L6KY18p1lBJcS9h67xs3LoOQZFqQIDAQAB
-----END RSA PUBLIC KEY-----
`

	// Signature
	signature := `oMRyloPHBG9Uq8ecJqENx601OvepVbIhwKnBNrc34SQwob4qg/kdMpt52QFa6Wd0xyzdnfSvFWBTYi6EvdvDlnmoxF73kfzUhRLsirpgNS3sbdJr3zqn14y8TXdb+t6YwOlxUqvS2XoBfDQ/oQmARsKGiYyQbM0G89km82Y/TGYbw/8p6/Kz8IeKAX3keDJnjMLecTlXa6jIRwKVFhXD70S4klJZLXFDfEGnon9vbVDEFfL14rrDvCmMaigYY0LzlANH4Qb1jFV3NYdOkyv4Iw2Kr9RyBnkaNTjtreUf5306NtxwonXsdQfx1mJQLM5MkNV45zZWcdiPEUIDtagmHg==`

	// Message
	message := []byte("Your message goes here")

	// Validate signature
	if validateSignature(pemPublicKey, message, signature) {
		fmt.Println("Signature is valid")
	} else {
		fmt.Println("Signature is invalid")
	}
}

func validateSignature(pemPublicKey string, message []byte, signatureBase64 string) bool {
	// Parse PEM block
	rsaPubKey, err := parseRSAPublicKeyFromPEM(pemPublicKey)
	if err != nil {
		log.Fatal("Error parsing PKIX public key:", err)
	}

	// Decode Base64 signature
	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		log.Fatal("Error decoding signature:", err)
	}

	// Hash the message
	hashed := sha256.Sum256(message)

	// Verify the signature
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Error verifying signature:", err)
		return false
	}

	return true
}

// parseRSAPublicKeyFromPEM parses RSA public key from PEM data
func parseRSAPublicKeyFromPEM(pemData string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	// Attempt to parse as PKCS1 public key
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return pubKey, nil
	}

	// If parsing as PKCS1 fails, try parsing as PKIX public key
	pubKeyPKIX, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pubKeyPKIX.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not RSA public key")
	}

	return rsaPubKey, nil
}
