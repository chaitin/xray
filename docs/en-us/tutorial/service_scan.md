# Service scan in xray

The focus of xray is web scan, however, capabilities for service scan will be gradually opened. If you havd upgraded from older versions, please make sure the poc names of service scan is in you config file.

There are two arguments of service scan, `target` and `json-output`. The `target` argument is the address of your service, for example `127.0.0.1:8009`. The `json-output` argument is the file path to store the result.

```
NAME:
    servicescan - Run a service scan task

USAGE:
    servicescan [command options] [arguments...]

OPTIONS:
   --target value      specify the target, for example: host:8009
   --json-output FILE  output xray results to FILE in json format
```