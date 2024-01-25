package rollKeys

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestRotationalSlice_Get(t *testing.T) {
	givenAPIKeys := []string{"a", "b"}
	givenNow := time.Now()
	givenRequestTimes := 100
	givenRate := 5

	expectDuration := (givenRequestTimes / (len(givenAPIKeys) * givenRate)) - 1

	keys, err := NewRotationalSlice(givenAPIKeys, givenRate)
	require.NoError(t, err)

	eg := errgroup.Group{}

	for i := 0; i < givenRequestTimes; i++ {
		eg.Go(func() error {
			key, err := keys.Get(context.TODO())
			if err != nil {
				return err
			}

			assert.Contains(t, givenAPIKeys, key)

			return nil
		})
	}
	err = eg.Wait()
	require.NoError(t, err)

	assert.Equal(t, expectDuration, int(time.Since(givenNow).Seconds()))
}
