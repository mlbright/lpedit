# lpedit
> Import/export your LastPass secrets

## Usage

```bash
# output your secrets to a JSON file
./lpedit --email 'first-account@example.com' --out > secrets.json

# make a temporary backup of your secrets
cp secrets.json edited.json

# edit your secrets
vim edited.json

# import your secrets into another account
./lpedit --email 'second-account@example.com' --in < edited.json

# then if your LastPass vaults are in order, ...
rm secrets.json edited.json
```
## Build

```bash
go build lpedit.go
```

## Notes

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
