## ioctl bc bucketlist

Get bucket list with method and arg(s) on IoTeX blockchain

### Synopsis

Read bucket list
Valid methods: [voter, cand]

```
ioctl bc bucketlist <method> [arguments] [flags]
```

### Examples

```
ioctl bc bucketlist voter [VOTER_ADDRESS] [OFFSET] [LIMIT]
ioctl bc bucketlist cand [CANDIDATE_NAME] [OFFSET] [LIMIT]
```

### Options

```
  -h, --help   help for bucketlist
```

### Options inherited from parent commands

```
      --endpoint string        set endpoint for once
      --insecure               insecure connection for once
  -o, --output-format string   output format
```

### SEE ALSO

* [ioctl bc](ioctl_bc.md)	 - Deal with block chain of IoTeX blockchain

###### Auto generated by docgen on 7-Mar-2022