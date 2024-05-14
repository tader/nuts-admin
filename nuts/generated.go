// Package nuts provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package nuts

// CredentialStatus Object enabling the discovery of information related to the status of a verifiable credential, such as whether it is suspended or revoked.
// Interpretation of the credentialStatus is defined by its 'type' property.
type CredentialStatus = interface{}

// CredentialSubject Subject of a Verifiable Credential identifying the holder and expressing claims.
type CredentialSubject = interface{}

// DIDDocumentMetadata The DID document metadata.
type DIDDocumentMetadata struct {
	// Created Time when DID document was created in rfc3339 form.
	Created string `json:"created"`

	// Deactivated Whether the DID document has been deactivated.
	Deactivated bool `json:"deactivated"`

	// Hash Sha256 in hex form of the DID document contents.
	Hash string `json:"hash"`

	// PreviousHash Sha256 in hex form of the previous version of this DID document.
	PreviousHash *string `json:"previousHash,omitempty"`

	// Txs txs lists the transaction(s) that created the current version of this DID Document.
	// If multiple transactions are listed, the DID Document is conflicted
	Txs []string `json:"txs"`

	// Updated Time when DID document was updated in rfc3339 form.
	Updated *string `json:"updated,omitempty"`
}

// EmbeddedProof Cryptographic proofs that can be used to detect tampering and verify the authorship of a
// credential or presentation. An embedded proof is a mechanism where the proof is included in
// the data, such as a Linked Data Signature.
type EmbeddedProof struct {
	// Challenge A random or pseudo-random value, provided by the verifier, used by some authentication protocols to
	// mitigate replay attacks.
	Challenge *string `json:"challenge,omitempty"`

	// Created Date and time at which proof has been created.
	Created string `json:"created"`

	// Domain A string value that specifies the operational domain of a digital proof. This could be an Internet domain
	// name like example.com, an ad-hoc value such as mycorp-level3-access, or a very specific transaction value
	// like 8zF6T$mqP. A signer could include a domain in its digital proof to restrict its use to particular
	// target, identified by the specified domain.
	Domain *string `json:"domain,omitempty"`

	// Jws JSON Web Signature
	Jws string `json:"jws"`

	// Nonce A unique string value generated by the holder, MUST only be used once for a particular domain
	// and window of time. This value can be used to mitigate replay attacks.
	Nonce *string `json:"nonce,omitempty"`

	// ProofPurpose It expresses the purpose of the proof and ensures the information is protected by the
	// signature.
	ProofPurpose string `json:"proofPurpose"`

	// Type Type of the object or the datatype of the typed value. Currently only supported value is "JsonWebSignature2020".
	Type string `json:"type"`

	// VerificationMethod Specifies the public key that can be used to verify the digital signature.
	// Dereferencing a public key URL reveals information about the controller of the key,
	// which can be checked against the issuer of the credential.
	VerificationMethod string `json:"verificationMethod"`
}

// Revocation Credential revocation record
type Revocation struct {
	// Date date is a rfc3339 formatted datetime.
	Date string `json:"date"`

	// Issuer DID according to Nuts specification
	Issuer DID `json:"issuer"`

	// Proof Proof contains the cryptographic proof(s).
	Proof *map[string]interface{} `json:"proof,omitempty"`

	// Reason reason describes why the VC has been revoked
	Reason *string `json:"reason,omitempty"`

	// Subject subject refers to the credential identifier that is revoked (not the credential subject)
	Subject string `json:"subject"`
}

// Service A service supported by a DID subject.
type Service struct {
	// Id ID of the service.
	Id string `json:"id"`

	// ServiceEndpoint Either a URI or a complex object.
	ServiceEndpoint interface{} `json:"serviceEndpoint"`

	// Type The type of the endpoint.
	Type string `json:"type"`
}

// VerificationMethod A public key in JWK form.
type VerificationMethod struct {
	// Controller The DID subject this key belongs to.
	Controller string `json:"controller"`

	// Id The ID of the key, used as KID in various JWX technologies.
	Id string `json:"id"`

	// PublicKeyJwk The public key formatted according rfc7517.
	PublicKeyJwk map[string]interface{} `json:"publicKeyJwk"`

	// Type The type of the key.
	Type string `json:"type"`
}
