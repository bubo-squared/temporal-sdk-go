package internal

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bubo-squared/temporal-sdk-go/converter"
)

func TestDecodeArg(t *testing.T) {
	t.Parallel()
	dc := converter.GetDefaultDataConverter()

	b, err := encodeArg(dc, testErrorDetails3)
	require.NoError(t, err)
	var r testStruct
	err = decodeArg(dc, b, &r)
	require.NoError(t, err)
	require.Equal(t, testErrorDetails3, r)

	// test mismatch of multi arguments
	b, err = encodeArgs(dc, []interface{}{testErrorDetails1, testErrorDetails2})
	require.NoError(t, err)
	require.Error(t, decodeArg(dc, b, &r))
}
