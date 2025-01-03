.PHONY: test-counter test-counter-verbose

COUNTER_RELATIVE_PATH=./internals/counter
HELPER_RELATIVE_PATH=./internals/helper

test-counter:
	@go test ${COUNTER_RELATIVE_PATH}

test-counter-verbose:
	@go test -v ${COUNTER_RELATIVE_PATH}

benchmark-helper:
	@go test ${HELPER_RELATIVE_PATH} -bench=.
