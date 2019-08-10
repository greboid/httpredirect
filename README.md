# Go http redirect
Simple HTTP server to redirect requests according to the below config options.

|Setting|Explanation|
---|---
HOST|specifies what host to redirect to, if unset will use the requested host
PATH|Specifies the path to redirect to, if unset will use the requested path
EXCLUDEQUERY|If set, will ignore query params when redirecting
HTTPREDIRECT|If set, will redirect to http rather than https
TEMPREDIRECT|If set, will do a temporary redirect (307) rather than a permanent one (308)