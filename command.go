package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/marcozj/golang-sdk/enum/authenticationtype"
)

// CentrifyCLI is data structure for injecting secret retrieved from vaults into environment variables
type CentrifyCLI struct {
	//secrets     []vaultObject
	//vaultClient *restapi.RestClient
	authtype    string
	url         string
	appid       string
	scope       string
	token       string
	user        string
	password    string
	credpath    string
	checkin 	bool
	skipcert    bool
}

// getCmdParms parse command line argument
func (c *CentrifyCLI) getCmdParms() {
	// Common arguments
	authTypePtr := flag.String("auth", "dmc", "Authentication type <oauth|unpw|dmc>")
	urlPtr := flag.String("url", "", "Centrify tenant URL (Required)")
	skipCertPtr := flag.Bool("skipcert", false, "Ignore certification verification")

	// Other arguments
	appIDPtr := flag.String("appid", "", "OAuth application ID. Required if auth = oauth")
	scopePtr := flag.String("scope", "", "OAuth or DMC scope definition. Required if auth = oauth or dmc")
	tokenPtr := flag.String("token", "", "OAuth token. Optional if auth = oauth or dmc")
	usernamePtr := flag.String("user", "", "Authorized user to login to tenant. Required if auth = unpw. Optional if auth = oauth")
	passwordPtr := flag.String("password", "", "User password. You will be prompted to enter password if this isn't provided")
	credPathPtr := flag.String("credpath", "", "Path of the secret/pasword to be retrieved")
	checkinPtr := flag.Bool("checkin", false, "Whether checkin immediately")

	flag.Usage = func() {
		fmt.Printf("Usage: centrifyvault-cli -auth dmc -url https://tenant.my.centrify.net -scope scope -credpath <system/systemname/accountname>\n")
		fmt.Printf("Usage: centrifyvault-cli -auth oauth -token <oauthtoken> -url https://tenant.my.centrify.net -credpath <secret/folder/secretname>\n")
		fmt.Printf("Usage: centrifyvault-cli -auth oauth -url https://tenant.my.centrify.net -appid <appid> -scope <scope> -user <username> -credpath <secret/folder/secretname>\n")
		fmt.Printf("Usage: centrifyvault-cli -auth unpw -url https://tenant.my.centrify.net -user <username> -credpath <secret/folder/secretname>\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Verify command argument length
	if len(os.Args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Verify authTypePtr value
	authChoices := map[string]bool{authenticationtype.OAuth2.String(): true, authenticationtype.UsernamePassword.String(): true, authenticationtype.DelegatedMachineCredential.String(): true}
	if _, validChoice := authChoices[*authTypePtr]; !validChoice {
		fmt.Println("Invalid authentication method")
		flag.Usage()
		os.Exit(1)
	}
	// Check required argument that do not have default value
	if *urlPtr == "" {
		fmt.Println("Missing url parameter")
		flag.Usage()
		os.Exit(1)
	}

	// Check secret path is provided
	if *credPathPtr == "" {
		fmt.Println("Missing vaulted credential path parameter")
		flag.Usage()
		os.Exit(1)
	}

	switch strings.ToLower(*authTypePtr) {
	case authenticationtype.OAuth2.String():
		if *tokenPtr == "" && (*appIDPtr == "" || *scopePtr == "") {
			fmt.Println("Missing appid and scope parameter")
			flag.Usage()
			os.Exit(1)
		}
		// Either token or username must be provided
		if *tokenPtr == "" && *usernamePtr == "" {
			fmt.Println("Missing token or user parameter")
			flag.Usage()
			os.Exit(1)
		}
	case authenticationtype.UsernamePassword.String():
		if *urlPtr == "" || *usernamePtr == "" {
			fmt.Println("Missing url and user parameter")
			flag.Usage()
			os.Exit(1)
		}
	case authenticationtype.DelegatedMachineCredential.String():
		if *tokenPtr == "" && *scopePtr == "" {
			fmt.Println("Missing token or scope parameter")
			flag.Usage()
			os.Exit(1)
		}
	}

	// Assign argument values to struct
	c.authtype = *authTypePtr
	c.url = *urlPtr
	c.appid = *appIDPtr
	c.scope = *scopePtr
	c.token = *tokenPtr
	c.user = *usernamePtr
	c.password = *passwordPtr
	c.skipcert = *skipCertPtr
	c.credpath = *credPathPtr
	c.checkin = *checkinPtr
}
