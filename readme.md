# tfupdate

terraform wapper application  

## Usage

```bash
 ./tfupdate help
A brief description of your application

Usage:
  tfupdate [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print tfupdate version
  wapper      Run terraform command

Flags:
  -h, --help     help for tfupdate
  -t, --toggle   Help message for toggle

Use "tfupdate [command] --help" for more information about a command.
```

## sample

```bash
./tfupdate wapper plan --dirs=test/nullresource
```

## Build  

```bash
make build
```
