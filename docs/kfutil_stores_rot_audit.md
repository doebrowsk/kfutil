## kfutil stores rot audit

Audit generates a CSV report of what actions will be taken based on input CSV files.

### Synopsis

Root of Trust Audit: Will read and parse inputs to generate a report of certs that need to be added or removed from the "root of trust" stores.

```
kfutil stores rot audit [flags]
```

### Options

```
  -a, --add-certs string      CSV file containing cert(s) to enroll into the defined cert stores
  -d, --dry-run               Dry run mode
  -h, --help                  help for audit
  -k, --max-keys -1           The max number of private keys that should be in a store to be considered a 'root' store. If set to -1 then all stores will be considered. (default -1)
  -l, --max-leaf-certs -1     The max number of non-root-certs that should be in a store to be considered a 'root' store. If set to -1 then all stores will be considered. (default -1)
  -m, --min-certs -1          The minimum number of certs that should be in a store to be considered a 'root' store. If set to -1 then all stores will be considered. (default -1)
  -o, --outpath string        Path to write the audit report file to. If not specified, the file will be written to the current directory.
  -r, --remove-certs string   CSV file containing cert(s) to remove from the defined cert stores
  -s, --stores string         CSV file containing cert stores to enroll into
```

### SEE ALSO

* [kfutil stores rot](kfutil_stores_rot.md)	 - Root of trust utility

###### Auto generated by spf13/cobra on 27-Oct-2022