// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	storagepb "cloud.google.com/go/bigquery/storage/apiv1beta2/storagepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newBigQueryWriteClientHook clientHook

// BigQueryWriteCallOptions contains the retry settings for each method of BigQueryWriteClient.
type BigQueryWriteCallOptions struct {
	CreateWriteStream       []gax.CallOption
	AppendRows              []gax.CallOption
	GetWriteStream          []gax.CallOption
	FinalizeWriteStream     []gax.CallOption
	BatchCommitWriteStreams []gax.CallOption
	FlushRows               []gax.CallOption
}

func defaultBigQueryWriteGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerystorage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerystorage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBigQueryWriteCallOptions() *BigQueryWriteCallOptions {
	return &BigQueryWriteCallOptions{
		CreateWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AppendRows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		FinalizeWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchCommitWriteStreams: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		FlushRows: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultBigQueryWriteRESTCallOptions() *BigQueryWriteCallOptions {
	return &BigQueryWriteCallOptions{
		CreateWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		AppendRows: []gax.CallOption{
			gax.WithTimeout(86400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		GetWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		FinalizeWriteStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		BatchCommitWriteStreams: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		FlushRows: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalBigQueryWriteClient is an interface that defines the methods available from BigQuery Storage API.
type internalBigQueryWriteClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateWriteStream(context.Context, *storagepb.CreateWriteStreamRequest, ...gax.CallOption) (*storagepb.WriteStream, error)
	AppendRows(context.Context, ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error)
	GetWriteStream(context.Context, *storagepb.GetWriteStreamRequest, ...gax.CallOption) (*storagepb.WriteStream, error)
	FinalizeWriteStream(context.Context, *storagepb.FinalizeWriteStreamRequest, ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error)
	BatchCommitWriteStreams(context.Context, *storagepb.BatchCommitWriteStreamsRequest, ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error)
	FlushRows(context.Context, *storagepb.FlushRowsRequest, ...gax.CallOption) (*storagepb.FlushRowsResponse, error)
}

// BigQueryWriteClient is a client for interacting with BigQuery Storage API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BigQuery Write API.
//
// The Write API can be used to write data to BigQuery.
type BigQueryWriteClient struct {
	// The internal transport-dependent client.
	internalClient internalBigQueryWriteClient

	// The call options for this service.
	CallOptions *BigQueryWriteCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BigQueryWriteClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BigQueryWriteClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BigQueryWriteClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateWriteStream creates a write stream to the given table.
// Additionally, every table has a special COMMITTED stream named ‘_default’
// to which data can be written. This stream doesn’t need to be created using
// CreateWriteStream. It is a stream that can be used simultaneously by any
// number of clients. Data written to this stream is considered committed as
// soon as an acknowledgement is received.
func (c *BigQueryWriteClient) CreateWriteStream(ctx context.Context, req *storagepb.CreateWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	return c.internalClient.CreateWriteStream(ctx, req, opts...)
}

// AppendRows appends data to the given stream.
//
// If offset is specified, the offset is checked against the end of
// stream. The server returns OUT_OF_RANGE in AppendRowsResponse if an
// attempt is made to append to an offset beyond the current end of the stream
// or ALREADY_EXISTS if user provids an offset that has already been
// written to. User can retry with adjusted offset within the same RPC
// stream. If offset is not specified, append happens at the end of the
// stream.
//
// The response contains the offset at which the append happened. Responses
// are received in the same order in which requests are sent. There will be
// one response for each successful request. If the offset is not set in
// response, it means append didn’t happen due to some errors. If one request
// fails, all the subsequent requests will also fail until a success request
// is made again.
//
// If the stream is of PENDING type, data will only be available for read
// operations after the stream is committed.
//
// This method is not supported for the REST transport.
func (c *BigQueryWriteClient) AppendRows(ctx context.Context, opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error) {
	return c.internalClient.AppendRows(ctx, opts...)
}

// GetWriteStream gets a write stream.
func (c *BigQueryWriteClient) GetWriteStream(ctx context.Context, req *storagepb.GetWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	return c.internalClient.GetWriteStream(ctx, req, opts...)
}

// FinalizeWriteStream finalize a write stream so that no new data can be appended to the
// stream. Finalize is not supported on the ‘_default’ stream.
func (c *BigQueryWriteClient) FinalizeWriteStream(ctx context.Context, req *storagepb.FinalizeWriteStreamRequest, opts ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error) {
	return c.internalClient.FinalizeWriteStream(ctx, req, opts...)
}

// BatchCommitWriteStreams atomically commits a group of PENDING streams that belong to the same
// parent table.
// Streams must be finalized before commit and cannot be committed multiple
// times. Once a stream is committed, data in the stream becomes available
// for read operations.
func (c *BigQueryWriteClient) BatchCommitWriteStreams(ctx context.Context, req *storagepb.BatchCommitWriteStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error) {
	return c.internalClient.BatchCommitWriteStreams(ctx, req, opts...)
}

// FlushRows flushes rows to a BUFFERED stream.
// If users are appending rows to BUFFERED stream, flush operation is
// required in order for the rows to become available for reading. A
// Flush operation flushes up to any previously flushed offset in a BUFFERED
// stream, to the offset specified in the request.
// Flush is not supported on the _default stream, since it is not BUFFERED.
func (c *BigQueryWriteClient) FlushRows(ctx context.Context, req *storagepb.FlushRowsRequest, opts ...gax.CallOption) (*storagepb.FlushRowsResponse, error) {
	return c.internalClient.FlushRows(ctx, req, opts...)
}

// bigQueryWriteGRPCClient is a client for interacting with BigQuery Storage API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryWriteGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BigQueryWriteClient
	CallOptions **BigQueryWriteCallOptions

	// The gRPC API client.
	bigQueryWriteClient storagepb.BigQueryWriteClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBigQueryWriteClient creates a new big query write client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BigQuery Write API.
//
// The Write API can be used to write data to BigQuery.
func NewBigQueryWriteClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryWriteClient, error) {
	clientOpts := defaultBigQueryWriteGRPCClientOptions()
	if newBigQueryWriteClientHook != nil {
		hookOpts, err := newBigQueryWriteClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BigQueryWriteClient{CallOptions: defaultBigQueryWriteCallOptions()}

	c := &bigQueryWriteGRPCClient{
		connPool:            connPool,
		bigQueryWriteClient: storagepb.NewBigQueryWriteClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *bigQueryWriteGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryWriteGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryWriteGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryWriteRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing BigQueryWriteClient
	CallOptions **BigQueryWriteCallOptions
}

// NewBigQueryWriteRESTClient creates a new big query write rest client.
//
// BigQuery Write API.
//
// The Write API can be used to write data to BigQuery.
func NewBigQueryWriteRESTClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryWriteClient, error) {
	clientOpts := append(defaultBigQueryWriteRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBigQueryWriteRESTCallOptions()
	c := &bigQueryWriteRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &BigQueryWriteClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBigQueryWriteRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://bigquerystorage.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://bigquerystorage.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryWriteRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryWriteRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *bigQueryWriteRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *bigQueryWriteGRPCClient) CreateWriteStream(ctx context.Context, req *storagepb.CreateWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateWriteStream[0:len((*c.CallOptions).CreateWriteStream):len((*c.CallOptions).CreateWriteStream)], opts...)
	var resp *storagepb.WriteStream
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.CreateWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) AppendRows(ctx context.Context, opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp storagepb.BigQueryWrite_AppendRowsClient
	opts = append((*c.CallOptions).AppendRows[0:len((*c.CallOptions).AppendRows):len((*c.CallOptions).AppendRows)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.AppendRows(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) GetWriteStream(ctx context.Context, req *storagepb.GetWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetWriteStream[0:len((*c.CallOptions).GetWriteStream):len((*c.CallOptions).GetWriteStream)], opts...)
	var resp *storagepb.WriteStream
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.GetWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) FinalizeWriteStream(ctx context.Context, req *storagepb.FinalizeWriteStreamRequest, opts ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FinalizeWriteStream[0:len((*c.CallOptions).FinalizeWriteStream):len((*c.CallOptions).FinalizeWriteStream)], opts...)
	var resp *storagepb.FinalizeWriteStreamResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.FinalizeWriteStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) BatchCommitWriteStreams(ctx context.Context, req *storagepb.BatchCommitWriteStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCommitWriteStreams[0:len((*c.CallOptions).BatchCommitWriteStreams):len((*c.CallOptions).BatchCommitWriteStreams)], opts...)
	var resp *storagepb.BatchCommitWriteStreamsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.BatchCommitWriteStreams(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryWriteGRPCClient) FlushRows(ctx context.Context, req *storagepb.FlushRowsRequest, opts ...gax.CallOption) (*storagepb.FlushRowsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "write_stream", url.QueryEscape(req.GetWriteStream())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FlushRows[0:len((*c.CallOptions).FlushRows):len((*c.CallOptions).FlushRows)], opts...)
	var resp *storagepb.FlushRowsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryWriteClient.FlushRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateWriteStream creates a write stream to the given table.
// Additionally, every table has a special COMMITTED stream named ‘_default’
// to which data can be written. This stream doesn’t need to be created using
// CreateWriteStream. It is a stream that can be used simultaneously by any
// number of clients. Data written to this stream is considered committed as
// soon as an acknowledgement is received.
func (c *bigQueryWriteRESTClient) CreateWriteStream(ctx context.Context, req *storagepb.CreateWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetWriteStream()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateWriteStream[0:len((*c.CallOptions).CreateWriteStream):len((*c.CallOptions).CreateWriteStream)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.WriteStream{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// AppendRows appends data to the given stream.
//
// If offset is specified, the offset is checked against the end of
// stream. The server returns OUT_OF_RANGE in AppendRowsResponse if an
// attempt is made to append to an offset beyond the current end of the stream
// or ALREADY_EXISTS if user provids an offset that has already been
// written to. User can retry with adjusted offset within the same RPC
// stream. If offset is not specified, append happens at the end of the
// stream.
//
// The response contains the offset at which the append happened. Responses
// are received in the same order in which requests are sent. There will be
// one response for each successful request. If the offset is not set in
// response, it means append didn’t happen due to some errors. If one request
// fails, all the subsequent requests will also fail until a success request
// is made again.
//
// If the stream is of PENDING type, data will only be available for read
// operations after the stream is committed.
//
// This method is not supported for the REST transport.
func (c *bigQueryWriteRESTClient) AppendRows(ctx context.Context, opts ...gax.CallOption) (storagepb.BigQueryWrite_AppendRowsClient, error) {
	return nil, fmt.Errorf("AppendRows not yet supported for REST clients")
}

// GetWriteStream gets a write stream.
func (c *bigQueryWriteRESTClient) GetWriteStream(ctx context.Context, req *storagepb.GetWriteStreamRequest, opts ...gax.CallOption) (*storagepb.WriteStream, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetWriteStream[0:len((*c.CallOptions).GetWriteStream):len((*c.CallOptions).GetWriteStream)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.WriteStream{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// FinalizeWriteStream finalize a write stream so that no new data can be appended to the
// stream. Finalize is not supported on the ‘_default’ stream.
func (c *bigQueryWriteRESTClient) FinalizeWriteStream(ctx context.Context, req *storagepb.FinalizeWriteStreamRequest, opts ...gax.CallOption) (*storagepb.FinalizeWriteStreamResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).FinalizeWriteStream[0:len((*c.CallOptions).FinalizeWriteStream):len((*c.CallOptions).FinalizeWriteStream)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.FinalizeWriteStreamResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchCommitWriteStreams atomically commits a group of PENDING streams that belong to the same
// parent table.
// Streams must be finalized before commit and cannot be committed multiple
// times. Once a stream is committed, data in the stream becomes available
// for read operations.
func (c *bigQueryWriteRESTClient) BatchCommitWriteStreams(ctx context.Context, req *storagepb.BatchCommitWriteStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCommitWriteStreamsResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetParent())

	params := url.Values{}
	if items := req.GetWriteStreams(); len(items) > 0 {
		for _, item := range items {
			params.Add("writeStreams", fmt.Sprintf("%v", item))
		}
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).BatchCommitWriteStreams[0:len((*c.CallOptions).BatchCommitWriteStreams):len((*c.CallOptions).BatchCommitWriteStreams)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.BatchCommitWriteStreamsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// FlushRows flushes rows to a BUFFERED stream.
// If users are appending rows to BUFFERED stream, flush operation is
// required in order for the rows to become available for reading. A
// Flush operation flushes up to any previously flushed offset in a BUFFERED
// stream, to the offset specified in the request.
// Flush is not supported on the _default stream, since it is not BUFFERED.
func (c *bigQueryWriteRESTClient) FlushRows(ctx context.Context, req *storagepb.FlushRowsRequest, opts ...gax.CallOption) (*storagepb.FlushRowsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v", req.GetWriteStream())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "write_stream", url.QueryEscape(req.GetWriteStream())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).FlushRows[0:len((*c.CallOptions).FlushRows):len((*c.CallOptions).FlushRows)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.FlushRowsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
