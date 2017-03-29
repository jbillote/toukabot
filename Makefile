ROOT = main.go
MODULES := $(wildcard modules/*)
UTIL := $(wildcard util/*)

toukabot: $(ROOT) $(MODULES) $(UTIL)
	go build

clean:
	rm toukabot