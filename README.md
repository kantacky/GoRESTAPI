# Kantacky API

## Generate Key for JWT Signing
```
#!/bin/bash
openssl ecparam -genkey -name secp521r1 -noout -out ecdsa-p521-private.pem
chmod 600 ecdsa-p521-private.pem
openssl ec -in ecdsa-p521-private.pem -pubout -out ecdsa-p521-public.pem
```

## Set Environment Variables
- DB_HOST
- DB_PORT
- DB_USER
- DB_PASS
- HOST
- PORT
- JWT_PRIVATE_KEY


&copy; 2023 Kanta Oikawa
