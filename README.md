# Keyfactor Command Utility (kfutil)

`kfutil` is a go-lang CLI wrapper for Keyfactor Command API. It also includes other utility/helper functions around automating common Keyfactor Command operations.

#### Integration status: Pilot - Ready for use in test environments. Not for use in production.

## About the Keyfactor API Client

This API client allows for programmatic management of Keyfactor resources.



## Support for Keyfactor Command Utility (kfutil)

Keyfactor Command Utility (kfutil) is open source and there is **no SLA** for this tool/library/client. Keyfactor will address issues as resources become available. Keyfactor customers may request escalation by opening up a support ticket through their Keyfactor representative.

###### To report a problem or suggest a new feature, use the **[Issues](../../issues)** tab. If you want to contribute actual bug fixes or proposed enhancements, use the **[Pull requests](../../pulls)** tab.
___




## Quickstart

### Prerequisites:
- [Github CLI](https://cli.github.com/)
- [zip](https://linuxize.com/post/how-to-unzip-files-in-linux/#installing-unzip) CLI tool, used to unzip the release files. 
- `$HOME/.local/bin` in your `$PATH` and exists.

### Installation:
```bash
bash <(curl -s https://raw.githubusercontent.com/Keyfactor/kfutil/main/gh-dl-release.sh)
````

### Environmental Variables

All the variables listed below need to be set in your environment. The `kfutil` command will look for these variables
and use them if they are set. If they are not set, the utility will fail to connect to Keyfactor.

Linux/MacOS:
```bash
export KEYFACTOR_HOSTNAME="<mykeyfactorhost.mydomain.com>"
export KEYFACTOR_USERNAME="<myusername>" # Do not include domain
export KEYFACTOR_PASSWORD="<mypassword>"
export KEYFACTOR_DOMAIN="<mykeyfactordomain>"
export KEYFACTOR_API_PATH="KeyfactorAPI/" # Optional, defaults to KeyfactorAPI/ only use this if you have a custom API path
```

Windows Powershell
```powershell
$env:KEYFACTOR_HOSTNAME="<mykeyfactorhost.mydomain.com>"
$env:KEYFACTOR_USERNAME="<myusername>" # Do not include domain
$env:KEYFACTOR_PASSWORD="<mypassword>"
$env:KEYFACTOR_DOMAIN="<mykeyfactordomain>"
$env:KEYFACTOR_API_PATH="KeyfactorAPI/" # Optional, defaults to KeyfactorAPI/ only use this if you have a custom API path
```

## Commands

### Login
For full documentation on the `login` command, see the [login](docs/kfutil_login.md) documentation.

*WARNING* - The `login` command will store your Keyfactor credentials in a file on your local machine. This file is not
encrypted and is not secure. It is recommended that you use the `login` command only on your local machine and not on a
shared machine. Instead of using the `login` command, you can set the environmental variables listed above.

```bash
kfutil login
```

#### Example Config File v1
```json
{
  "api_path": "KeyfactorAPI",
  "domain": "command",
  "host": "lab.mydomain.com",
  "password": "dontusethispassword",
  "username": "myusername"
}
```

#### Example Config File v2
**Must use kfutil version v1.0.0 or higher.**
```json
{
  "servers": { 
    "lab1" : {
      "api_path": "KeyfactorAPI",
      "domain": "command",
      "host": "lab.mydomain.com",
      "password": "dontusethispassword",
      "username": "myusername"
    },
    "lab2" : {
      "api_path": "KeyfactorAPI",
      "domain": "command",
      "host": "lab2.mydomain.com",
      "password": "dontusethispassword",
      "username": "myusername"
    }
  }
}
```

### Logout
For full documentation on the `logout` command, see the [logout](docs/kfutil_logout.md) documentation.

*WARNING* - This will delete the file containing your Keyfactor credentials at `$HOME/.keyfactor/command_config.json`.

```bash
kfutil logout
```

## Commands

### Bulk operations

#### Bulk create cert stores

For full documentation, see [stores import](docs/kfutil_stores_import.md).

This will attempt to process a CSV input file of certificate stores to create. The template can be generated by
running: `kfutil stores import generate-template` command.

```bash
kfutil stores import create --file <file name to import> --store-type-id <store type id> --store-type-name <store type name> --results-path <filepath for results> --dry-run <check fields only> [flags]
```

```bash
kfutil stores import --help
Tool for generating import templates and importing certificate stores

Usage:
  kfutil stores import [command]

Available Commands:
  create            Create certificate stores
  generate-template For generating a CSV template with headers for bulk store creation.

Flags:
  -h, --help   help for import

Use "kfutil stores import [command] --help" for more information about a command.
```

#### Bulk create cert store types

For full documentation, see [store-types](docs/kfutil_store-types.md).

This will attempt to process a CSV input file of certificate store types to create. The template can be generated by
running: `kfutil generate-template --type bulk-certstore-types` command.

```bash
kfutil store-types create --name $STORE_TYPE_NAME
```

```bash
kfutil store-types --help             
A collections of APIs and utilities for interacting with Keyfactor Command certificate store types.

Usage:
  kfutil store-types [command]

Available Commands:
  create          Create a new certificate store type in Keyfactor Command.
  delete          Delete a specific store type by ID.
  get             Get a specific store type by either name or ID.
  list            List certificate store types.
  templates-fetch Fetches store type templates from Keyfactor's Github.
  update          Update a certificate store type in Keyfactor.

Flags:
  -h, --help   help for store-types

Use "kfutil store-types [command] --help" for more information about a command.
```

### Root of Trust

For full documentation, see [stores rot](docs/kfutil_stores_rot.md).

The root of trust (rot) utility is a tool that allows you to bulk manage Keyfactor certificate stores and ensure that a
set of defined certificates are present in each store that meets a certain set of criteria or no criteria at all.

### Root of Trust Quickstart

```bash
echo "Generating cert template file certs_template.csv"
kfutil stores rot generate-template-rot --type certs
# edit the certs_template.csv file
echo "Generating stores template file stores_template.csv"
kfutil stores rot generate-template-rot --type stores
# edit the stores_template.csv file
kfutil stores rot audit --add-certs certs_template.csv --stores stores_template.csv #This will audit the stores and generate a report file
# review/edit the report file generated `rot_audit.csv`
kfutil stores rot reconcile --import-csv
# Alternatively this can be done in one step
kfutil stores rot reconcile --add-certs certs_template.csv --stores stores_template.csv
```

#### Generate Certificate List Template

For full documentation, see [stores rot generate template](docs/kfutil_stores_rot_generate-template.md).

This will write the file `certs_template.csv` to the current directory.

```bash
kfutil stores generate-template-rot --type certs
```

#### Generate Certificate Store List Template

For full documentation, see [stores rot generate template](docs/kfutil_stores_rot_generate-template.md).

This will write the file `stores_template.csv` to the current directory. For full documentation

```bash
kfutil stores generate-template-rot --type stores
```

#### Run Root of Trust Audit

For full documentation, see [stores rot audit](docs/kfutil_stores_rot_audit.md).

Audit will take in a list of certificates and a list of certificate stores and check that the certificate store's
inventory either contains the certificate or does not contain the certificate based on the `--add-certs` and
`--remove-certs` flags. These flags can be used together or separately. The aforementioned flags take in a path to CSV
files containing a list of certificate thumbprints. To generate a template for these files, run the following command:

```bash
kfutil stores rot generate-template --type certs
```

To prepopulate the template file you can provide `--cn` multiple times.

```bash
kfutil stores rot generate-template --type certs \
  --cn <cert subject name> \
  --cn <additional cert subject name>
```

In addition, you must provide a list of stores you wish to audit. To generate a template for this file, run the
following
command:

```bash
kfutil stores rot generate-template --type stores
```

To prepopulate the template file you can provide `--store-type` and `--container-type` multiple times.

```bash
kfutil stores rot generate-template --type stores \
  --store-type <store type name> \
  --store-type <additional store type name> \
  --container-type <container type name> \
  --container-type <additional container type name>
```

With all the files generated and populated, you can now run the audit command:

```bash
kfutil stores rot audit \
  --stores stores_template.csv \
  --add-certs certs_template.csv \
  --remove-certs certs_template2.csv
```

This will generate an audit file that contains the results of the audit and actions that will be taken if `reconcile` is
executed. By default, the audit file will be named `rot_audit.csv` and will be written to the current directory. To
output
the audit file to a different location, use the `--output` flag:

```bash
kfutil stores rot audit \
  --stores stores.csv \
  --add-certs addCerts.csv \
  --remove-certs removeCerts.csv \
  --output /path/to/output/autdit_file.csv
```

#### Run Root of Trust Reconcile

For full documentation, see [stores rot](docs/kfutil_stores_rot_reconcile.md).

Reconcile will take in a list of certificates and a list of certificate stores and check that the certificate store's
inventory either contains the certificate or does not contain the certificate based on the `--add-certs` and
`--remove-certs` flags. These flags can be used together or separately. The aforementioned flags take in a path to CSV
files containing a list of certificate thumbprints. To generate a template for these files, run the following command:

```bash
kfutil stores rot generate-template --type certs
```

To pre-populate the template file you can provide `--cn` multiple times.

```bash
kfutil stores rot generate-template --type certs \
  --cn <cert subject name> \
  --cn <additional cert subject name>
```

In addition, you must provide a list of stores you wish to reconcile. To generate a template for this file, run the
following
command:

```bash
kfutil stores rot generate-template --type stores
```

To pre-populate the stores template file you can provide multiple values in any combination of the following flags:

```bash
kfutil stores rot generate-template --type stores \
  --store-type <store type name> \
  --store-type <additional store type name> \
  --container-type <container type name> \
  --container-type <additional container type name>
```

With all the files generated and populated, you can now run the reconcile command:

```bash
kfutil stores rot reconcile \
  --stores stores_template.csv \
  --add-certs certs_template.csv \
  --remove-certs certs_template2.csv
```

This will generate an audit file that contains the results of the audit and actions will immediately execute those
actions.
By default, the reconcile file will be named `rot_audit.csv` and will be written to the current directory. To output
the reconcile file to a different location, use the `--output` flag:

```bash
kfutil stores rot reconcile \
  --stores stores.csv \
  --add-certs addCerts.csv \
  --remove-certs removeCerts.csv \
  --output /path/to/output/audit_file.csv
```

Alternatively you can provide an audit CSV file as an input to the reconcile command using the `--import-csv` flag:

```bash
kfutil stores rot reconcile \
  --import-csv /path/to/audit_file.csv
```

### Certificate Store Inventory

For full documentation, see [stores inventory](docs/kfutil_stores_inventory.md).

#### Show the inventory of a certificate store

For full documentation, see [stores inventory show](docs/kfutil_stores_inventory_show.md).

```bash
# Show by store ID:
```bash
kfutil stores inventory show --sid <store id>

# Nested command lookup: shows inventory of first cert store found
kfutil stores inventory show \
  --sid $(kfutil stores list | jq -r ".[0].Id")
```

Show by client machine name:

```bash
kfutil stores inventory show --client <machine name>

# Nested command lookup: shows inventory of first cert store found
kfutil stores inventory show \
  --client $(kfutil orchs list | jq -r ".[0].ClientMachine")
```

#### Add certificates to certificate stores

For full documentation, see [stores inventory add](docs/kfutil_stores_inventory_add.md).

```bash
# Add 2 certs to 2 certificate stores
kfutil stores inventory add \
  --sid <store id> \
  --sid <additional store id> \
  --cn <cert subject name> \
  --cn <additional cert subject name>
```

#### Remove certificates from certificate stores

For full documentation, see [stores inventory remove](docs/kfutil_stores_inventory_remove.md).

```bash
# Remove 2 certs from all stores associated with a client machine
kfutil stores inventory remove \
  --client <machine name> \
  --cn <cert subject name> \
  --cn <additional cert subject name>
```

### Export Instance Data

For full documentation, see [export](docs/kfutil_export.md).

Export select instance data to JSON file:

```bash
# export only collections, metadata, and roles
kfutil export --collections --metadata --roles --file <path to JSON file>
```

Export all exportable instance data to JSON file:

```bash
kfutil export --all --file <path to JSON file>
```

### Import Instance Data

For full documentation, see [import](docs/kfutil_import.md).

Import select instance data from exported JSON file:

```bash
# export only collections, metadata, and roles
kfutil import --collections --metadata --roles --file <path to JSON file>
```

Import all importable instance data from exported JSON file:

```bash
kfutil import --all --file <path to JSON file>
```

## Development

This CLI developed using [cobra](https://umarcor.github.io/cobra/)

### Adding a new command

```bash
cobra-cli add <my-new-command>
```

alternatively you can specify the parent command

```bash
cobra-cli add <my-new-command> -p '<parent>Cmd'
```
