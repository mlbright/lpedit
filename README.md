# lpedit
> Import/export your LastPass secrets

## Usage

```bash
# output your secrets to a JSON file
./lpedit --email 'first-account@example.com' --password "correcthorsebatterystaple" --out > secrets.json

# make a temporary backup of your secrets
cp secrets.json edited.json

# edit your secrets
vim edited.json

# import your secrets into another account
./lpedit --email 'second-account@example.com' --password "someotherinsanepassword" --in < edited.json

# then if your LastPass vaults are in order, ...
rm secrets.json edited.json
```

:warning: Note that supplying your master password like this can be insecure.

You could put it in a file and do:

```bash
./lpedit --email 'first-account@example.com' --password "$(cat $HOME/temporary-password-file.txt)" --out

# ... followed by ...
rm $HOME/temporary-password-file.txt
```

## Build

```bash
go build lpedit.go
```

## Notes

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
