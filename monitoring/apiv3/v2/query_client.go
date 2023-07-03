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

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newQueryClientHook clientHook

// QueryCallOptions contains the retry settings for each method of QueryClient.
type QueryCallOptions struct {
	QueryTimeSeries []gax.CallOption
}

func defaultQueryGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultQueryCallOptions() *QueryCallOptions {
	return &QueryCallOptions{
		QueryTimeSeries: []gax.CallOption{},
	}
}

// internalQueryClient is an interface that defines the methods available from Cloud Monitoring API.
type internalQueryClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	QueryTimeSeries(context.Context, *monitoringpb.QueryTimeSeriesRequest, ...gax.CallOption) *TimeSeriesDataIterator
}

// QueryClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The QueryService API is used to manage time series data in Stackdriver
// Monitoring. Time series data is a collection of data points that describes
// the time-varying values of a metric.
type QueryClient struct {
	// The internal transport-dependent client.
	internalClient internalQueryClient

	// The call options for this service.
	CallOptions *QueryCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *QueryClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *QueryClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *QueryClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// QueryTimeSeries queries time series using Monitoring Query Language. This method does not require a Workspace.
func (c *QueryClient) QueryTimeSeries(ctx context.Context, req *monitoringpb.QueryTimeSeriesRequest, opts ...gax.CallOption) *TimeSeriesDataIterator {
	return c.internalClient.QueryTimeSeries(ctx, req, opts...)
}

// queryGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type queryGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing QueryClient
	CallOptions **QueryCallOptions

	// The gRPC API client.
	queryClient monitoringpb.QueryServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewQueryClient creates a new query service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The QueryService API is used to manage time series data in Stackdriver
// Monitoring. Time series data is a collection of data points that describes
// the time-varying values of a metric.
func NewQueryClient(ctx context.Context, opts ...option.ClientOption) (*QueryClient, error) {
	clientOpts := defaultQueryGRPCClientOptions()
	if newQueryClientHook != nil {
		hookOpts, err := newQueryClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := QueryClient{CallOptions: defaultQueryCallOptions()}

	c := &queryGRPCClient{
		connPool:    connPool,
		queryClient: monitoringpb.NewQueryServiceClient(connPool),
		CallOptions: &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *queryGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *queryGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *queryGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *queryGRPCClient) QueryTimeSeries(ctx context.Context, req *monitoringpb.QueryTimeSeriesRequest, opts ...gax.CallOption) *TimeSeriesDataIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).QueryTimeSeries[0:len((*c.CallOptions).QueryTimeSeries):len((*c.CallOptions).QueryTimeSeries)], opts...)
	it := &TimeSeriesDataIterator{}
	req = proto.Clone(req).(*monitoringpb.QueryTimeSeriesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.TimeSeriesData, string, error) {
		resp := &monitoringpb.QueryTimeSeriesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.queryClient.QueryTimeSeries(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTimeSeriesData(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// TimeSeriesDataIterator manages a stream of *monitoringpb.TimeSeriesData.
type TimeSeriesDataIterator struct {
	items    []*monitoringpb.TimeSeriesData
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.TimeSeriesData, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TimeSeriesDataIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TimeSeriesDataIterator) Next() (*monitoringpb.TimeSeriesData, error) {
	var item *monitoringpb.TimeSeriesData
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TimeSeriesDataIterator) bufLen() int {
	return len(it.items)
}

func (it *TimeSeriesDataIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
