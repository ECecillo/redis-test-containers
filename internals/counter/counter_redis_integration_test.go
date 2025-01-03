package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementAndRetrieveCounter(t *testing.T) {
	testCases := []struct {
		desc    string
		counter *Counter
	}{
		{
			desc:    "using Redis",
			counter: setupCounterWithRedisCluster(t),
		},
		//TODO:
		// {
		// 	desc: "return the correct counter value using ClickHouse",
		// 	counter: setupCounterWithClickHouse(t),
		// },
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			counter := tC.counter

			counter.Increment()
			counter.Increment()
			counter.Increment()

			got, err := counter.Get()
			assert.NoError(t, err, "unexpected error while retrieving counter value")

			expected := 3

			assert.Equal(t, expected, got, "expected same counter value")
		})
	}
}

func TestDeleteCounterAfterIncrement(t *testing.T) {

	testCases := []struct {
		desc    string
		counter *Counter
	}{
		{
			desc:    "using Redis",
			counter: setupCounterWithRedisCluster(t),
		},
		//TODO:
		// {
		// 	desc: "return the correct counter value using ClickHouse",
		// 	counter: setupCounterWithClickHouse(t),
		// },
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			counter := tC.counter

			counter.Increment()
			counter.Increment()
			counter.Increment()

			ok, err := counter.Delete()
			assert.NoError(t, err)

			assert.True(t, ok, "delete should have return an ok with value true")

			_, err = counter.Get()
			assert.Error(t, err, "expected an error to be thrown")
		})
	}

}

func TestDeleteCounterWhenNotExisting(t *testing.T) {
	testCases := []struct {
		desc    string
		counter *Counter
	}{
		{
			desc:    "using Redis",
			counter: setupCounterWithRedisCluster(t),
		},
		//TODO:
		// {
		// 	desc: "return the correct counter value using ClickHouse",
		// 	counter: setupCounterWithClickHouse(t),
		// },
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			counter := tC.counter

			ok, err := counter.Delete()
			assert.NoError(t, err)
			assert.True(t, ok, "delete should have return an ok with value true")
		})
	}

}
