
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
	cd driver; $(G) build
	cd gci; $(G) build
install :
	cd driver; $(G) install
	cd gci; $(G) install
clean :
	$(G) clean
	rm -rf CUBRID
