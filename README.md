cubrid
======

CUBRID Database driver for Go


build
=====
```
$ export CGO_CFLAGS=-I$CUBRID/include
$ export CGO_LDFLAGS="-L$CUBRID/lib -lcascci -lnsl"
$ go install
  ```

;;or
```
$ make install
```
