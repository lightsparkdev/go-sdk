
// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package objects




type Secret struct {

    
    EncryptedValue string `json:"secret_encrypted_value"`

    
    Cipher string `json:"secret_cipher"`

}

const (
    SecretFragment = `
fragment SecretFragment on Secret {
    __typename
    secret_encrypted_value: encrypted_value
    secret_cipher: cipher
}
`
)







