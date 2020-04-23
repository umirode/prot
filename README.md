# Prot - protobuf files manager.

It application can help your manage protobuf files and generate code based on him.

##Commands
### `prot help` - Shows a list of commands or help for one command

### `prot init` - Generate config for Prot application
Options:
* `--output` Output path for config (default: current directory)

### `prot install` - Install dependencies from prot.yml config
Options:
* `--config` Path to configuration file (default: prot.yaml)
* `--output` Output path (default: current directory)

##Configuration
`Lang` - now it`s `go` or `dart`. If yor want more - Welcome to the issue! :).

`ModuleName` or `ModuleNameWithAuth` - name for module.

`Repository` - link for your repository (ssh or https, depends on the type of authorization).

`Auth` - optional parameter, if your want auth for git.

```
Lang: go
Modules:
  ModuleName:
    Repository: https://github.com/umirode/prot.git
  ModuleNameWithAuth:
    Repository: git@github.com:umirode/prot.git
    Auth:
      Type: PublicKeys
      Config:
        PemFile: id_rsa.pem
        IgnoreHostKey: true

```

##Auth Configuration

###PublicKeys - represent ssh auth via public keys (type ssh).
```
Auth:
    Type: PublicKeys
    Config:
      PemFile: id_rsa.pem
      User: umirode # OPTIONAL DEFAULTS git
      PemFilePassword: password # OPTIONAL
      IgnoreHostKey: true # OPTIONAL DEFAULTS false
```

###Password - represent ssh auth via password (type ssh).
```
Auth:
    Type: Password
    Config:
      User: user
      Password: pass
```

###BasicAuth - represent a HTTP basic auth (type https).
```
Auth:
    Type: BasicAuth
    Config:
      Username: user
      Password: pass
```

###Token - represent authenticate with HTTP token authentication (also known as bearer authentication) (type https).

IMPORTANT: If you are looking to use OAuth tokens with popular servers (e.g.
GitHub, Bitbucket, GitLab) you should use BasicAuth instead. These servers
use basic HTTP authentication, with the OAuth token as user or password.
Check the documentation of your git server for details.

```
Auth:
    Type: Token
    Config:
      Token: token
```

