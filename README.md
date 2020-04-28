# Prot - protobuf files manager.

It application can help your manage protobuf files and generate code based on him.

!!! Before use prot you must install your lang dependencies (section Config/Lang).

1. Create repository for service protobuf files.
2. Run `prot init` command config in your server application.
3. Add your repository with protobuf files into prot.yml and run `prot install` command.
4. Write code for server.
5. Run `prot init` command config in your server application.
6. Add your repository with protobuf files into prot.yml and run `prot install` command.
7. Use generated grpc client in your client service.
8. Now after update protobuf files you can update generated code for grpc server and client by one command - `prot install`

## Install

`go get -u github.com/umirode/prot`

## Commands

`prot help` - Shows a list of commands or help for one command

```
❯ prot help
NAME:
   prot - protobuf files manager

USAGE:
   prot [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   install  Install dependencies from config file
   init     Generate config for Prot application
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

------

`prot init` - Generate config for Prot application
Options:
* `--output` Output path for config (default: current directory)

```
❯ prot init --help
NAME:
   prot init - Generate config for Prot application

USAGE:
   prot init [command options] [arguments...]

OPTIONS:
   --output value, -o value  Output path for config (default: current directory)
   --help, -h                show help (default: false)
```

`prot install` - Install dependencies from prot.yml config
Options:
* `--config` Path to configuration file (default: prot.yaml)
* `--output` Output path (default: current directory)

------

```
❯ prot install --help                                                                                                          
NAME:
   prot install - Install dependencies from config file

USAGE:
   prot install [command options] [arguments...]

OPTIONS:
   --config value, -c value  Path to configuration file (default: prot.yaml)
   --output value, -o value  Output path (default: current directory)
   --help, -h                show help (default: false)
```

## Configuration
`Lang` If yor want more - welcome to the issue! :)

|  Lang  | Info |
| ------ | ---- |
| go     | Follow https://grpc.io/docs/quickstart/go/ section Protocol Buffers      |
| dart   | Follow https://grpc.io/docs/quickstart/dart/ section Protocol Buffers v3 |

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

## Auth Configuration

### PublicKeys - represent ssh auth via public keys (type ssh).
```
Auth:
    Type: PublicKeys
    Config:
      PemFile: id_rsa.pem
      User: umirode # OPTIONAL DEFAULTS git
      PemFilePassword: password # OPTIONAL
      IgnoreHostKey: true # OPTIONAL DEFAULTS false
```

### Password - represent ssh auth via password (type ssh).
```
Auth:
    Type: Password
    Config:
      User: user
      Password: pass
```

### BasicAuth - represent a HTTP basic auth (type https).
```
Auth:
    Type: BasicAuth
    Config:
      Username: user
      Password: pass
```

### Token - represent authenticate with HTTP token authentication (also known as bearer authentication) (type https).

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

