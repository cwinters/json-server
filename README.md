json-server
===========

A little json file server with a touch of configurability.

## configuration

The port and data directory may be given as command line arguments:

``` bash
json-server -d ./data/dir -p 9922
```

Alternatively, they may be specified as env vars.  The default content
type may alse be given via the env.

- `PORT` - port number on which to listen (default `7878`)
- `DATADIR` - root directory for files (default current working dir)
- `CONTENT_TYPE` - content type to set for responses
