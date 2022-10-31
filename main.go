/*
Purpose:
- obtain Discourse User-API-Key

Description:
- This program obtains a User-API-Key for a Discourse forum.

Releases:
- v1.0.0 - 2022/10/31: initial release

Author:
- Klaus Tockloth

Copyright:
- Copyright (c) 2022 Klaus Tockloth

Contact (eMail):
- freizeitkarte@googlemail.com

License (MIT):
Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Remarks:
- Lint: golangci-lint run --no-config --enable gocritic
- Vulnerability detection: govulncheck ./...

Links:
- https://meta.discourse.org/t/user-api-keys-specification/48536
- https://meta.discourse.org/t/generate-user-api-keys-for-testing/145744
- https://stackoverflow.com/questions/13555085/save-and-load-crypto-rsa-privatekey-to-and-from-the-disk
- https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
*/

package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofrs/uuid"
)

// general program info
var (
	progName    = filepath.Base(os.Args[0])
	progVersion = "v1.0.0"
	progDate    = "2022/10/31"
	progPurpose = "obtain Discourse User-API-Key"
	progInfo    = "This program obtains a User-API-Key for a Discourse forum."
)

// DiscourseUserApiKey represents User-API-Key structure returned from Discourse service
type DiscourseUserApiKey struct {
	Key   string `json:"key"`
	Nonce string `json:"nonce"`
	Push  bool   `json:"push"`
	API   int    `json:"api"`
}

/*
main starts this program.
*/
func main() {
	fmt.Printf("\nProgram:\n"+
		"  Name    : %s\n"+
		"  Release : %s - %s\n"+
		"  Purpose : %s\n"+
		"  Info    : %s\n", progName, progVersion, progDate, progPurpose, progInfo)

	forum := flag.String("forum", "", "Discourse forum URL")
	application := flag.String("application", "GenericDiscourseReader", "name of application shown on forum site")
	client := flag.String("client", "", "client ID (default [generated unique UUID4])")
	scopes := flag.String("scopes", "read", "comma-separated list of access scopes allowed for the key")
	nonce := flag.String("nonce", "", "random string generated once (default [generated URL-safe random string])")
	verbose := flag.Bool("verbose", false, "verbose output (maybe helpful in case of problems)")

	flag.Usage = printUsage
	flag.Parse()

	if *forum == "" {
		fmt.Printf("\nError: mandatory option '-forum=string' missing\n")
		printUsage()
	}

	fmt.Printf("%s\n", workflow)

	if *verbose {
		fmt.Printf("\nParameters from command line:\n"+
			"  forum       : %s\n"+
			"  application : %s\n"+
			"  client      : %s\n"+
			"  scopes      : %s\n"+
			"  nonce       : %s\n", *forum, *application, *client, *scopes, *nonce)
	}

	// generate RSA private key
	DiscoursePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		handleFatalError(fmt.Sprintf("error [%v] at rsa.GenerateKey()", err))
	}
	if *verbose {
		fmt.Printf("\nPrivate Key = %v", exportRSAPrivateKeyAsPEM(DiscoursePrivateKey))
	}

	// encode public key to PEM
	DiscoursePublicKeyPEM, err := exportRSAPublicKeyAsPEM(&DiscoursePrivateKey.PublicKey)
	if err != nil {
		handleFatalError(fmt.Sprintf("error [%v] at exportRSAPublicKeyAsPEM()", err))
	}
	if *verbose {
		fmt.Printf("\nPublic Key = %v", DiscoursePublicKeyPEM)
	}

	// create request parameters
	if *client == "" {
		uuidTemp, err := uuid.NewV4()
		if err != nil {
			handleFatalError(fmt.Sprintf("error [%v] at uuid.NewV4()", err))
		}
		*client = uuidTemp.String()
	}
	if *nonce == "" {
		randomStringTemp, err := generateRandomStringURLSafe(20)
		if err != nil {
			handleFatalError(fmt.Sprintf("error [%v] at generateRandomStringURLSafe()", err))
		}
		*nonce = randomStringTemp
	}

	if *verbose {
		fmt.Printf("\nRequest parameters for URL:\n"+
			"  forum       : %s\n"+
			"  application : %s\n"+
			"  client      : %s\n"+
			"  scopes      : %s\n"+
			"  nonce       : %s\n", *forum, *application, *client, *scopes, *nonce)
	}

	// build URL
	baseURL := fmt.Sprintf("https://%s/user-api-key/new", *forum)
	values := url.Values{}
	values.Set("application_name", *application)
	values.Set("client_id", *client)
	values.Set("scopes", *scopes)
	values.Set("public_key", DiscoursePublicKeyPEM)
	values.Set("nonce", *nonce)

	performURL := baseURL + "?" + values.Encode()

	fmt.Printf("\nStep 1: copy forum URL into your browser ...\n\n%s\n", performURL)

	fmt.Printf("\nStep 2: authorize application access on forum site ...\n")

	// read encrypted data from stdin (copied from browser)
	fmt.Print("\nStep 3: copy encrypted User-API-Key data from forum site in here (and press Enter) ...\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// verify, that something was copied in
	if scanner.Text() == "" {
		handleFatalError("User-API-Key data from forum site missing")
	}

	// process encoded User-API-Key data from discourse service
	encodedUserApiKey := strings.ReplaceAll(scanner.Text(), " ", "")
	textUserApiKey, err := base64.StdEncoding.DecodeString(encodedUserApiKey)
	if err != nil {
		handleFatalError(fmt.Sprintf("error [%v] at base64.StdEncoding.DecodeString()", err))
	}
	if *verbose {
		fmt.Printf("\nEncrypted User-API-Key data = %s\n", encodedUserApiKey)
	}

	decryptedUserApiKey, err := DiscoursePrivateKey.Decrypt(rand.Reader, textUserApiKey, nil)
	if err != nil {
		handleFatalError(fmt.Sprintf("error [%v] at DiscoursePrivateKey.Decrypt()", err))
	}
	fmt.Printf("\nDecrypted User-API-Key data = %s\n", decryptedUserApiKey)

	var discourseUserApiKey DiscourseUserApiKey
	err = json.Unmarshal(decryptedUserApiKey, &discourseUserApiKey)
	if err != nil {
		handleFatalError(fmt.Sprintf("error [%v] at json.Unmarshal()", err))
	}

	fmt.Printf("\nUser-API-Key = %s\n-----------------------------------------------\n", discourseUserApiKey.Key)

	fmt.Printf("\nStep 4: save User-API-Key into your key vault\n\n")
}

/*
handleFatalError handles fatal error.
*/
func handleFatalError(message string) {
	fmt.Printf("Fatal Error: %s\n\n", message)
	os.Exit(1)
}

/*
exportRSAPrivateKeyAsPEM exports RSA private key as PEM-encoded string.
*/
func exportRSAPrivateKeyAsPEM(privkey *rsa.PrivateKey) string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privkey)

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes})

	return string(privateKeyPEM)
}

/*
exportRSAPublicKeyAsPEM exports RSA public key as PEM-encoded string.
*/
func exportRSAPublicKeyAsPEM(pubkey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})

	return string(publicKeyPEM), nil
}

/*
generateRandomBytes returns securely generated random bytes.
*/
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

/*
generateRandomStringURLSafe returns a URL-safe, base64 encoded securely generated random string.
*/
func generateRandomStringURLSafe(n int) (string, error) {
	b, err := generateRandomBytes(n)

	return base64.URLEncoding.EncodeToString(b), err
}

/*
printUsage prints the usage of this program.
*/
func printUsage() {
	fmt.Printf("\nUsage:\n")
	fmt.Printf("  %s -forum=string [-application=string] [-client=string] [-scopes=list] [-nonce=string] [-verbose]\n", os.Args[0])

	fmt.Printf("\nExamples:\n")
	fmt.Printf("  %s -forum=community.openstreetmap.org\n", os.Args[0])
	fmt.Printf("  %s -forum=meta.discourse.org -application=UltimateReaderWriter -scopes=read,write\n", os.Args[0])

	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	fmt.Printf("%s\n", workflow)

	fmt.Printf("\n")
	os.Exit(1)
}

var workflow = `
Workflow for getting an User-API-Key:
  Step 1: copy forum URL into your browser
  Step 2: authorize application access on forum site
  Step 3: copy encrypted User-API-Key data from forum site in here
  Step 4: save User-API-Key into your key vault`
