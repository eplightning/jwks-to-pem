# JWKS to PEM

Simple tool that converts public keys from JWKS to PEM format. Specified JWKS can either be local JSON file or HTTP(s) URL.
Useful for applications that require public key but don't have native JWKS support.

## Usage

``
./jwks-to-pem https://www.googleapis.com/oauth2/v3/certs > public-key.pem
./jwks-to-pem -format pkcs1 https://www.googleapis.com/oauth2/v3/certs > rsa-public-key.pem
./jwks-to-pem -output public-key.pem https://www.googleapis.com/oauth2/v3/certs
``