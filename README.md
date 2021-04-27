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

### Examples

To retrieve password for a vaulted account named "account1" in system named "system1" using DMC authentication:

```sh
./centrify-credential-retriever -url https://tenantid.my.centrify.net -scope <yourscope> -credpath system/system1/account1
```

To retrieve password for a vaulted account named "account2" in database named "mydb" using Oauth2 authentication:

```sh
./centrify-credential-retriever -url https://tenantid.my.centrify.net -auth oauth -appid <yourappid> -scope <yourscope> -username <yourusername> -password <yourpassword> -credpath database/mydb/account2
```

To retrieve secret text for a secret named "testsecret" in folder "folder1/folder2" using oauth2 token authentication:

```sh
./centrify-credential-retriever -url https://tenantid.my.centrify.net -auth oauth -token <yourtoken> -credpath -credpath secret/folder1\\folder2/testsecret
```

To retrieve secret text for a secret named "testsecret" in folder "folder1" using username and password authentication:

```sh
./centrify-credential-retriever -url https://tenantid.my.centrify.net -auth unpw -user <yourusername> -credpath -credpath secret/folder1/testsecret
```
