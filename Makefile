
G=go
CUBRID_PATH=$(CUBRID)

copy :
	-mkdir CUBRID
	-mkdir CUBRID/lib
	-mkdir CUBRID/include
	cp $(CUBRID_PATH)/lib/libcascci.so ./CUBRID/lib
	cp $(CUBRID_PATH)/include/cas_cci.h ./CUBRID/include
	cp $(CUBRID_PATH)/include/cas_error.h ./CUBRID/include
test :
	$(G) test
build :
	$(G) build
install :
	$(G) install
clean :
	$(G) clean
	rm -rf CUBRID
