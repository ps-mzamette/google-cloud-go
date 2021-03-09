// Copyright 2021 Google LLC
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

package tables

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	tablespb "google.golang.org/genproto/googleapis/area120/tables/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetTable        []gax.CallOption
	ListTables      []gax.CallOption
	GetWorkspace    []gax.CallOption
	ListWorkspaces  []gax.CallOption
	GetRow          []gax.CallOption
	ListRows        []gax.CallOption
	CreateRow       []gax.CallOption
	BatchCreateRows []gax.CallOption
	UpdateRow       []gax.CallOption
	BatchUpdateRows []gax.CallOption
	DeleteRow       []gax.CallOption
	BatchDeleteRows []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("area120tables.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("area120tables.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://area120tables.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetTable:        []gax.CallOption{},
		ListTables:      []gax.CallOption{},
		GetWorkspace:    []gax.CallOption{},
		ListWorkspaces:  []gax.CallOption{},
		GetRow:          []gax.CallOption{},
		ListRows:        []gax.CallOption{},
		CreateRow:       []gax.CallOption{},
		BatchCreateRows: []gax.CallOption{},
		UpdateRow:       []gax.CallOption{},
		BatchUpdateRows: []gax.CallOption{},
		DeleteRow:       []gax.CallOption{},
		BatchDeleteRows: []gax.CallOption{},
	}
}

// Client is a client for interacting with Area120 Tables API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	client tablespb.TablesServiceClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new tables service client.
//
// The Tables Service provides an API for reading and updating tables.
// It defines the following resource model:
//
//   The API has a collection of [Table][google.area120.tables.v1alpha1.Table]
//   resources, named tables/*
//
//   Each Table has a collection of [Row][google.area120.tables.v1alpha1.Row]
//   resources, named tables/*/rows/*
//
//   The API has a collection of
//   [Workspace][google.area120.tables.v1alpha1.Workspace]
//   resources, named workspaces/*.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultClientOptions()

	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultCallOptions(),

		client: tablespb.NewTablesServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// GetTable gets a table. Returns NOT_FOUND if the table does not exist.
func (c *Client) GetTable(ctx context.Context, req *tablespb.GetTableRequest, opts ...gax.CallOption) (*tablespb.Table, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetTable[0:len(c.CallOptions.GetTable):len(c.CallOptions.GetTable)], opts...)
	var resp *tablespb.Table
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetTable(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTables lists tables for the user.
func (c *Client) ListTables(ctx context.Context, req *tablespb.ListTablesRequest, opts ...gax.CallOption) *TableIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListTables[0:len(c.CallOptions.ListTables):len(c.CallOptions.ListTables)], opts...)
	it := &TableIterator{}
	req = proto.Clone(req).(*tablespb.ListTablesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tablespb.Table, string, error) {
		var resp *tablespb.ListTablesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListTables(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTables(), resp.GetNextPageToken(), nil
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

// GetWorkspace gets a workspace. Returns NOT_FOUND if the workspace does not exist.
func (c *Client) GetWorkspace(ctx context.Context, req *tablespb.GetWorkspaceRequest, opts ...gax.CallOption) (*tablespb.Workspace, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetWorkspace[0:len(c.CallOptions.GetWorkspace):len(c.CallOptions.GetWorkspace)], opts...)
	var resp *tablespb.Workspace
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetWorkspace(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListWorkspaces lists workspaces for the user.
func (c *Client) ListWorkspaces(ctx context.Context, req *tablespb.ListWorkspacesRequest, opts ...gax.CallOption) *WorkspaceIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ListWorkspaces[0:len(c.CallOptions.ListWorkspaces):len(c.CallOptions.ListWorkspaces)], opts...)
	it := &WorkspaceIterator{}
	req = proto.Clone(req).(*tablespb.ListWorkspacesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tablespb.Workspace, string, error) {
		var resp *tablespb.ListWorkspacesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListWorkspaces(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetWorkspaces(), resp.GetNextPageToken(), nil
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

// GetRow gets a row. Returns NOT_FOUND if the row does not exist in the table.
func (c *Client) GetRow(ctx context.Context, req *tablespb.GetRowRequest, opts ...gax.CallOption) (*tablespb.Row, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetRow[0:len(c.CallOptions.GetRow):len(c.CallOptions.GetRow)], opts...)
	var resp *tablespb.Row
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetRow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListRows lists rows in a table. Returns NOT_FOUND if the table does not exist.
func (c *Client) ListRows(ctx context.Context, req *tablespb.ListRowsRequest, opts ...gax.CallOption) *RowIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListRows[0:len(c.CallOptions.ListRows):len(c.CallOptions.ListRows)], opts...)
	it := &RowIterator{}
	req = proto.Clone(req).(*tablespb.ListRowsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*tablespb.Row, string, error) {
		var resp *tablespb.ListRowsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListRows(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRows(), resp.GetNextPageToken(), nil
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

// CreateRow creates a row.
func (c *Client) CreateRow(ctx context.Context, req *tablespb.CreateRowRequest, opts ...gax.CallOption) (*tablespb.Row, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateRow[0:len(c.CallOptions.CreateRow):len(c.CallOptions.CreateRow)], opts...)
	var resp *tablespb.Row
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateRow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchCreateRows creates multiple rows.
func (c *Client) BatchCreateRows(ctx context.Context, req *tablespb.BatchCreateRowsRequest, opts ...gax.CallOption) (*tablespb.BatchCreateRowsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchCreateRows[0:len(c.CallOptions.BatchCreateRows):len(c.CallOptions.BatchCreateRows)], opts...)
	var resp *tablespb.BatchCreateRowsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BatchCreateRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateRow updates a row.
func (c *Client) UpdateRow(ctx context.Context, req *tablespb.UpdateRowRequest, opts ...gax.CallOption) (*tablespb.Row, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "row.name", url.QueryEscape(req.GetRow().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateRow[0:len(c.CallOptions.UpdateRow):len(c.CallOptions.UpdateRow)], opts...)
	var resp *tablespb.Row
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateRow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchUpdateRows updates multiple rows.
func (c *Client) BatchUpdateRows(ctx context.Context, req *tablespb.BatchUpdateRowsRequest, opts ...gax.CallOption) (*tablespb.BatchUpdateRowsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchUpdateRows[0:len(c.CallOptions.BatchUpdateRows):len(c.CallOptions.BatchUpdateRows)], opts...)
	var resp *tablespb.BatchUpdateRowsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BatchUpdateRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteRow deletes a row.
func (c *Client) DeleteRow(ctx context.Context, req *tablespb.DeleteRowRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteRow[0:len(c.CallOptions.DeleteRow):len(c.CallOptions.DeleteRow)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteRow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// BatchDeleteRows deletes multiple rows.
func (c *Client) BatchDeleteRows(ctx context.Context, req *tablespb.BatchDeleteRowsRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchDeleteRows[0:len(c.CallOptions.BatchDeleteRows):len(c.CallOptions.BatchDeleteRows)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.BatchDeleteRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// RowIterator manages a stream of *tablespb.Row.
type RowIterator struct {
	items    []*tablespb.Row
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
	InternalFetch func(pageSize int, pageToken string) (results []*tablespb.Row, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RowIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RowIterator) Next() (*tablespb.Row, error) {
	var item *tablespb.Row
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RowIterator) bufLen() int {
	return len(it.items)
}

func (it *RowIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TableIterator manages a stream of *tablespb.Table.
type TableIterator struct {
	items    []*tablespb.Table
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
	InternalFetch func(pageSize int, pageToken string) (results []*tablespb.Table, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TableIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TableIterator) Next() (*tablespb.Table, error) {
	var item *tablespb.Table
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TableIterator) bufLen() int {
	return len(it.items)
}

func (it *TableIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// WorkspaceIterator manages a stream of *tablespb.Workspace.
type WorkspaceIterator struct {
	items    []*tablespb.Workspace
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
	InternalFetch func(pageSize int, pageToken string) (results []*tablespb.Workspace, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *WorkspaceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *WorkspaceIterator) Next() (*tablespb.Workspace, error) {
	var item *tablespb.Workspace
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *WorkspaceIterator) bufLen() int {
	return len(it.items)
}

func (it *WorkspaceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
