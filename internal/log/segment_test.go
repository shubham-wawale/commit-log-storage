package log

import (
	"io"
	"os"
	"testing"

	api "github.com/shubham-wawale/dswgo/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, _ := os.MkdirTemp("", "segment-test")
	defer os.RemoveAll(dir)
	want := &api.Record{
		Value: []byte("hello world"),
	}
	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)
	require.Equal(t, false, s.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)

		rec, err := s.Read(i)
		require.NoError(t, err)
		require.Equal(t, want.Value, rec.Value)
	}

	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)

	require.True(t, s.IsMaxed())

	c.Segment.MaxStoreBytes = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexBytes = 1024

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	// maxed store
	require.True(t, s.IsMaxed())

	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())

}
