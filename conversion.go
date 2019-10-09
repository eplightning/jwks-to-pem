package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	jose "gopkg.in/square/go-jose.v2"
)

func convertToJWK(data []byte) (*jose.JSONWebKey, error) {
	jwks := &jose.JSONWebKeySet{}
	if err := json.Unmarshal(data, jwks); err != nil {
		return nil, fmt.Errorf("could not parse JWKS: %w", err)
	}

	if len(jwks.Keys) == 0 {
		return nil, errors.New("JWKS doesn't contain any JWK's")
	}

	return &jwks.Keys[0], nil
}

func convertToPEM(jwk *jose.JSONWebKey, format string) ([]byte, error) {
	if format == "pkcs1" {
		return convertToPKCS1(jwk)
	}

	return convertToPKIX(jwk)
}

func convertToPKIX(jwk *jose.JSONWebKey) ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(jwk.Key)
	if err != nil {
		return nil, err
	}

	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: data,
	})

	return pemData, nil
}

// TODO support private keys
func convertToPKCS1(jwk *jose.JSONWebKey) ([]byte, error) {
	rsa, casted := jwk.Key.(*rsa.PublicKey)
	if !casted {
		return nil, errors.New("PKCS#1 format can only be used for RSA public keys")
	}

	data := x509.MarshalPKCS1PublicKey(rsa)

	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: data,
	})

	return pemData, nil
}
