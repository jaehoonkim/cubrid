
CGO_CFLAGS=-I$(CUBRID)/include
CGO_LDFLAGS=-L$(CUBRID)/lib -lcascci -lnsl
export CGO_CFLAGS
export CGO_LDFLAGS

test :
	go test
build :
	go build
install :
	go install
clean :
	go clean
