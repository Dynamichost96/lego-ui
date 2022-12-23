package config

import(	
	"time"
	"github.com/go-acme/lego/v4/registration"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log")

// The Config struct contains parameters that can be customized by the User through the Settings Page
type Config struct {
	// E-Mail Adress needed to regstier with LetsEncrypt
	Email        string

	//Is only set, if a user is registered with LetsEncrypt
	Registration *registration.Resource
	
	//Private Key used for certificates
	Key          crypto.PrivateKey
	
	//Set to true, if the user accepts LetsEncrypt's TOS
	TOSAgreed	 bool

	//After how many seconds, should we cancel the DNS verification request
	PropagationTimeout time.Duration

	//What interval should we check DNS Servers, if the verification succeeded
	PollingIntervall time.Duration
}

// Randomly generates a new private Key, which will be used for future certificates
func NewPrivateKey() crypto.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func (u *Config) GetEmail() string {
	return u.Email
}
func (u Config) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *Config) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}