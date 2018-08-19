// +build !solution

package main

/*
Task 7: Password Manager

This task focuses on building a password manager that stores
fictional passwords for websites. These passwords should be
handled somewhat securely, which is why only a hash will be stored.

The password manager should be implemented through a command line
application; allowing the user to execute all the functions in the
PasswordManager interface.

*/

type PasswordManager struct {
	// TODO(student): Add necessary field(s)
}

// NewPasswordManager returns an initialized instance of PasswordManager
func NewPasswordManager() *PasswordManager {
	// TODO(student): Initialize PasswordManager struct
	return nil
}

// Set creates or updates the password associated with the given site.
// The stored password should be hashed and salted, which can be accomplished
// with the bcrypt package (https://godoc.org/golang.org/x/crypto/bcrypt).
// Use the command "go get golang.org/x/crypto/bcrypt" if you do not have it installed.
func (m *PasswordManager) Set(site, password string) error {
	return nil
}

// Get returns the stored hash for the given site.
func (m *PasswordManager) Get(site string) []byte {
	return nil
}

// Verify checks whether the password given for a site matches the stored password.
// If the returned error value is nil, the passwords match.
// Hint: The bcrypt package may be of use here
func (m *PasswordManager) Verify(site, password string) error {
	return nil
}

// Remove deletes the password for a given site from the password manager.
func (m *PasswordManager) Remove(site string) {

}

// Save stores all the passwords in a file of the given name.
// Read the ioutil documentation for inforamtion on reading/writing files
// https://golang.org/pkg/io/ioutil/
//
// The file contents should be serialized in some way, for example using JSON
// (https://golang.org/pkg/encoding/json/) or XML (https://golang.org/pkg/encoding/xml/) etc.
func (m *PasswordManager) Save(fileName string) error {
	return nil
}

// Load reads the given file and decodes the values to replace the state of the
// PasswordManager.
func (m *PasswordManager) Load(fileName string) error {
	return nil
}

// Encrypt encrypts the input data using a cipher of your choice.
// This should be performed before saving passwords to disk.
// Go supports the AES(https://golang.org/pkg/crypto/aes/) and
// DES(https://golang.org/pkg/crypto/des/) ciphers. Note that
// TripleDES should be used if you choose DES.
//
// https://golang.org/pkg/crypto/cipher contains examples of different
// cipher modes. The simplest mode is block mode, although you are free
// to use stream mode if you prefer.
//
// It is also important to use keys of the right size depending on which
// cipher you chose. For example the AES key size is 16 bytes, or regular
// latin characters. You can either force the given key to 16 bytes by repeating
// or cutting off the key, or you can just rely on the user inputting 16 bytes.
func (m *PasswordManager) Encrypt(plaintext []byte, keystring string) []byte {
	// NOTE: This implementation is OPTIONAL
	return nil
}

// Decrypt reverses the operations performed by Encrypt, and is
// used when loading the password file from storage.
func (m *PasswordManager) Decrypt(ciphertext []byte, keystring string) []byte {
	// NOTE: This implementation is OPTIONAL
	return nil
}

// This is the main function of the application.
// User input should be continuously read and checked for commands
// for all the defined operations.
// See https://golang.org/pkg/bufio/#Reader and especially the ReadLine
// function.
func main() {

}
