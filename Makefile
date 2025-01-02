.PHONY: test-redis-counter

COUNTER_RELATIVE_PATH=./internals/counter/.
HELPER_RELATIVE_PATH=./internals/helper

test-redis-counter:
	@go test -v ${COUNTER_RELATIVE_PATH}

benchmark-helper:
	@go test ${HELPER_RELATIVE_PATH} -bench=.
