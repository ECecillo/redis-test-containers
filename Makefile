.PHONY: test-counter

COUNTER_RELATIVE_PATH=./counter/.
HELPER_RELATIVE_PATH=./helper

test-counter:
	@go test -v ${COUNTER_RELATIVE_PATH}

benchmark-helper:
	@go test ${HELPER_RELATIVE_PATH} -bench=.
