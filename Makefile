
PROG = gojournal

.PHONY: run clean validate default build

default: run

validate:
	@command -v fswatch --version >/dev/null 2>&1 || { printf >&2 "fswatch is not installed, please run: brew install fswatch\n"; exit 1; }

build: vendor
	go fmt
	go build -o $(PROG)

vendor:
	dep ensure

clean:
	rm $(PROG)

kill:
	-@killall -9 $(PROG) 2>/dev/null || true

restart:
	@make kill
	@make build; (if [ "$$?" -eq 0 ]; then (./$(PROG) &); fi)

serve: validate
	@make restart
	@fswatch -r -o ./*.go pkg/* tmpl/* | xargs -n1 -I{} make restart || make kill
