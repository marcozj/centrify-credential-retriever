# Centrify Vault Credential Retrieval Tool

This is a command line utility that retrieves vaulted password/sshkey/secret from Centrify Vault.

## Usage

```sh
$ ./centrify-credential-retriever 
  -appid string
        OAuth application ID. Required if auth = oauth
  -auth string
        Authentication type <oauth|unpw|dmc> (default "dmc")
  -credpath string
        Path of the secret/pasword to be retrieved.
  -password string
        User password. You will be prompted to enter password if this isn\'t provided
  -scope string
        OAuth or DMC scope definition. Required if auth = oauth or dmc
  -skipcert
        Ignore certification verification
  -token string
        OAuth token. Optional if auth = oauth or dmc
  -url string
        Centrify tenant URL (Required)
  -user string
        Authorized user to login to tenant. Required if auth = unpw. Optional if auth = oauth
```
