// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package ledgerwriter

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver/examples/bankofanthos/model"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"github.com/hooto/weaver"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/examples/bankofanthos/ledgerwriter/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, addTransactionMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/bankofanthos/ledgerwriter/T", Method: "AddTransaction", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, addTransactionMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/bankofanthos/ledgerwriter/T", Method: "AddTransaction", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "⟦7237a6f4:wEaVeReDgE:github.com/ServiceWeaver/weaver/examples/bankofanthos/ledgerwriter/T→github.com/ServiceWeaver/weaver/examples/bankofanthos/balancereader/T⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl                  T
	tracer                trace.Tracer
	addTransactionMetrics *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) AddTransaction(ctx context.Context, a0 string, a1 string, a2 model.Transaction) (err error) {
	// Update metrics.
	begin := s.addTransactionMetrics.Begin()
	defer func() { s.addTransactionMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "ledgerwriter.T.AddTransaction", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AddTransaction(ctx, a0, a1, a2)
}

// Client stub implementations.

type t_client_stub struct {
	stub                  codegen.Stub
	addTransactionMetrics *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) AddTransaction(ctx context.Context, a0 string, a1 string, a2 model.Transaction) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.addTransactionMetrics.Begin()
	defer func() { s.addTransactionMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "ledgerwriter.T.AddTransaction", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.String(a0)
	enc.String(a1)
	(a2).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' (devel) (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "AddTransaction":
		return s.addTransaction
	default:
		return nil
	}
}

func (s t_server_stub) addTransaction(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()
	var a2 model.Transaction
	(&a2).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.AddTransaction(ctx, a0, a1, a2)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) AddTransaction(ctx context.Context, a0 string, a1 string, a2 model.Transaction) (err error) {
	err = s.caller("AddTransaction", ctx, []any{a0, a1, a2}, []any{})
	return
}
