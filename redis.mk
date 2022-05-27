# Call `make redis`, which runs redis binary after downloading/installing.

REDIS = redis-2.6.12

redis: bin/redis-server
	# bin/redis-server path/to/redis.conf
	bin/redis-server

bin/redis-server: src/$(REDIS)/src/redis-server
	mkdir -p bin
	cp $< $@

src/$(REDIS)/src/redis-server: src/$(REDIS)/README
	cd src/$(REDIS) && make

src/$(REDIS)/README: src/$(REDIS).tar.gz
	cd src && tar -xvf $(REDIS).tar.gz
	@touch $@ # Ensure we do not untar every time, by updating README time.

src/$(REDIS).tar.gz:
	mkdir -p src
	cd src && wget http://redis.googlecode.com/files/$(REDIS).tar.gz

clean:
	rm -fr bin/redis-server src/$(REDIS)*
