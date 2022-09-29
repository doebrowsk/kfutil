# Keyfactor Util
General go-lang CLI utility for the Keyfactor API.

## Quickstart
```bash
make install
kfutil --help
````

## Commands

### Bulk operations

#### Bulk create cert stores
`# TODO: Not implemented`  
This will attempt to process a CSV input file of certificate stores to create. The template can be generated by running: `kfutil generate-template --type bulk-certstore` command.
```bash
kfutil bulk create certstores --file <path to csv file>
```

#### Bulk create cert store types
`# TODO: Not implemented` 
This will attempt to process a CSV input file of certificate store types to create. The template can be generated by running: `kfutil generate-template --type bulk-certstore-types` command.
```bash
kfutil bulk create certstores --file <path to csv file>
```

### Root of Trust

#### Generate Certificate List Template
This will write the file `certs_template.csv` to the current directory.
```bash
kfutil stores generate-template-rot --type certs
```

#### Generate Certificate Store List Template
This will write the file `certs_template.csv` to the current directory.
```bash
kfutil stores generate-template-rot --type stores
```

#### Run Root of Trust Check
This will read the file `certs.csv` from the current directory or the absolute path, and generate a report of the certificate stores that contain the specified certificates.
```bash
kfutil stores rot --stores stores.csv --certs certs.csv
```

### Development
This CLI developed using [cobra](https://umarcor.github.io/cobra/)
#### Adding a new command
```bash
cobra-cli add <my-new-command>
```
alternatively you can specify the parent command
```bash
cobra-cli add <my-new-command> -p '<parent>Cmd'
```
