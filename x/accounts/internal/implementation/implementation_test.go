package implementation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation(t *testing.T) {
	impl, err := NewImplementation(TestAccount{})
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("execute ok", func(t *testing.T) {
		resp, err := impl.Execute(ctx, &wrapperspb.StringValue{Value: "test"})
		require.NoError(t, err)
		require.Equal(t, "testexecute-echo", resp.(*wrapperspb.StringValue).Value)

		resp, err = impl.Execute(ctx, &wrapperspb.BytesValue{Value: []byte("test")})
		require.NoError(t, err)
		require.Equal(t, "testbytes-execute-echo", string(resp.(*wrapperspb.BytesValue).Value))
	})

	t.Run("execute - unknown message", func(t *testing.T) {
		_, err := impl.Execute(ctx, &wrapperspb.Int32Value{Value: 1})
		require.ErrorIs(t, err, errInvalidMessage)
	})

	t.Run("init ok", func(t *testing.T) {
		resp, err := impl.Init(ctx, &wrapperspb.StringValue{Value: "test"})
		require.NoError(t, err)
		require.Equal(t, "testinit-echo", resp.(*wrapperspb.StringValue).Value)
	})

	t.Run("init - unknown message", func(t *testing.T) {
		_, err := impl.Init(ctx, &wrapperspb.Int32Value{Value: 1})
		require.ErrorIs(t, err, errInvalidMessage)
	})

	t.Run("query ok", func(t *testing.T) {
		resp, err := impl.Query(ctx, &wrapperspb.StringValue{Value: "test"})
		require.NoError(t, err)
		require.Equal(t, "testquery-echo", resp.(*wrapperspb.StringValue).Value)

		resp, err = impl.Query(ctx, &wrapperspb.BytesValue{Value: []byte("test")})
		require.NoError(t, err)
		require.Equal(t, "testbytes-query-echo", string(resp.(*wrapperspb.BytesValue).Value))
	})

	t.Run("query - unknown message", func(t *testing.T) {
		_, err := impl.Query(ctx, &wrapperspb.Int32Value{Value: 1})
		require.ErrorIs(t, err, errInvalidMessage)
	})

	// schemas
	t.Run("decode init request - ok", func(t *testing.T) {
		want := &wrapperspb.StringValue{Value: "test"}
		req, err := impl.InitHandlerSchema.ResponseSchema.TxEncode(want)
		require.NoError(t, err)

		got, err := impl.InitHandlerSchema.RequestSchema.TxDecode(req)
		require.NoError(t, err)
		require.True(t, proto.Equal(want, got.(protoreflect.ProtoMessage)))
	})

	t.Run("encode init response - ok", func(t *testing.T) {
		want := &wrapperspb.StringValue{Value: "test"}

		gotBytes, err := impl.InitHandlerSchema.ResponseSchema.TxEncode(want)
		require.NoError(t, err)

		wantBytes, err := proto.Marshal(want)
		require.NoError(t, err)

		require.Equal(t, wantBytes, gotBytes)
	})
}
