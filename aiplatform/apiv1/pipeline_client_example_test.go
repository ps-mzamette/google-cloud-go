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

package aiplatform_test

import (
	"context"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleNewPipelineClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExamplePipelineClient_CreateTrainingPipeline() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.CreateTrainingPipelineRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#CreateTrainingPipelineRequest.
	}
	resp, err := c.CreateTrainingPipeline(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_GetTrainingPipeline() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.GetTrainingPipelineRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#GetTrainingPipelineRequest.
	}
	resp, err := c.GetTrainingPipeline(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_ListTrainingPipelines() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.ListTrainingPipelinesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#ListTrainingPipelinesRequest.
	}
	it := c.ListTrainingPipelines(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExamplePipelineClient_DeleteTrainingPipeline() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.DeleteTrainingPipelineRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#DeleteTrainingPipelineRequest.
	}
	op, err := c.DeleteTrainingPipeline(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_CancelTrainingPipeline() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.CancelTrainingPipelineRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#CancelTrainingPipelineRequest.
	}
	err = c.CancelTrainingPipeline(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_CreatePipelineJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.CreatePipelineJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#CreatePipelineJobRequest.
	}
	resp, err := c.CreatePipelineJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_GetPipelineJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.GetPipelineJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#GetPipelineJobRequest.
	}
	resp, err := c.GetPipelineJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_ListPipelineJobs() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.ListPipelineJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#ListPipelineJobsRequest.
	}
	it := c.ListPipelineJobs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExamplePipelineClient_DeletePipelineJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.DeletePipelineJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#DeletePipelineJobRequest.
	}
	op, err := c.DeletePipelineJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_CancelPipelineJob() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &aiplatformpb.CancelPipelineJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/aiplatform/apiv1/aiplatformpb#CancelPipelineJobRequest.
	}
	err = c.CancelPipelineJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExamplePipelineClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/iam/apiv1/iampb#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_DeleteOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.DeleteOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#DeleteOperationRequest.
	}
	err = c.DeleteOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePipelineClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePipelineClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExamplePipelineClient_WaitOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := aiplatform.NewPipelineClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.WaitOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#WaitOperationRequest.
	}
	resp, err := c.WaitOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}