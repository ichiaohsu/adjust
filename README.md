# Adjust

The tool provides the required MD5 response list to desired URLs and simple error report if any

## How to run

It provides a flag setting `parellel` to set up how many parallel request user could have. Its default is set to 10. Set it in following formats will both set the concurrent limit to 3:
- `-parallel 3`
- `-parallel=3`

It can be run directly with simple `go run` method. Here is an example to request two URLs with `go run`:
```bash
go run ./... http://www.adjust.com http://google.com
```
This will create two requests to these URLs and print out corresponding MD5 of their response.

Recommended way is to build the executive and run it:
```bash
go build -o myhttp .
./myhttp -parallel 3 https://adjust.com http://google.com facebook.com yahoo.com
```

URLs could have `http`, `https` scheme, or without a scheme. URLs with different schemes could be mixed in the execution. Those URLs without scheme will be added the default `http://` scheme. Those URLs will be shown in modified format in result.
```bash
./myhttp -parallel=3 adjust.com http://google.com https://www.facebook.com
```
will actually request http://adjust.com instead. The result looks like this:
```bash
http://google.com 69fa1abbdbbdfc1a3436522b3068db98
https://www.facebook.com 879060e4b58ab6f5061000c71da7479b
http://adjust.com 413bfe0cdb886cd843edf67d6410380e
```

### Error report
There is an extra error handling. Whenever there is an error during the request, it will be collected in the process and printed out after the requests are done. Errors are listed in the format of `(url) (error message)` and printed at the bottom of the result separated under a line `-----`. For example, following request
```bash
./myhttp -parallel 3 adjust.com http://google.com foo
```
will create a report like this:
```bash
http://adjust.com 413bfe0cdb886cd843edf67d6410380e
http://google.com 74f1904494c1aae794919f27080f4920
-----
http://foo error:failed to request url: http://foo, err: Get "http://foo": dial tcp: lookup foo: no such host
```
It could be seen clearly after the separation line is the error we get from the request.