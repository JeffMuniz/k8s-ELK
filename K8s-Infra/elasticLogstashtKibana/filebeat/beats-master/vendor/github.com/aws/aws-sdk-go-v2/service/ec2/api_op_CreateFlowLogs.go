// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateFlowLogsRequest
type CreateFlowLogsInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to
	// a CloudWatch Logs log group in your account.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	DeliverLogsPermissionArn *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Specifies the destination to which the flow log data is to be published.
	// Flow log data can be published to a CloudWatch Logs log group or an Amazon
	// S3 bucket. The value specified for this parameter depends on the value specified
	// for LogDestinationType.
	//
	// If LogDestinationType is not specified or cloud-watch-logs, specify the Amazon
	// Resource Name (ARN) of the CloudWatch Logs log group.
	//
	// If LogDestinationType is s3, specify the ARN of the Amazon S3 bucket. You
	// can also specify a subfolder in the bucket. To specify a subfolder in the
	// bucket, use the following ARN format: bucket_ARN/subfolder_name/. For example,
	// to specify a subfolder named my-logs in a bucket named my-bucket, use the
	// following ARN: arn:aws:s3:::my-bucket/my-logs/. You cannot use AWSLogs as
	// a subfolder name. This is a reserved term.
	LogDestination *string `type:"string"`

	// Specifies the type of destination to which the flow log data is to be published.
	// Flow log data can be published to CloudWatch Logs or Amazon S3. To publish
	// flow log data to CloudWatch Logs, specify cloud-watch-logs. To publish flow
	// log data to Amazon S3, specify s3.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	//
	// Default: cloud-watch-logs
	LogDestinationType LogDestinationType `type:"string" enum:"true"`

	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	LogGroupName *string `type:"string"`

	// The ID of the subnet, network interface, or VPC for which you want to create
	// a flow log.
	//
	// Constraints: Maximum of 1000 resources
	//
	// ResourceIds is a required field
	ResourceIds []string `locationName:"ResourceId" locationNameList:"item" type:"list" required:"true"`

	// The type of resource for which to create the flow log. For example, if you
	// specified a VPC ID for the ResourceId property, specify VPC for this property.
	//
	// ResourceType is a required field
	ResourceType FlowLogsResourceType `type:"string" required:"true" enum:"true"`

	// The type of traffic to log. You can log traffic that the resource accepts
	// or rejects, or all traffic.
	//
	// TrafficType is a required field
	TrafficType TrafficType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateFlowLogsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFlowLogsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFlowLogsInput"}

	if s.ResourceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceIds"))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if len(s.TrafficType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TrafficType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateFlowLogsResult
type CreateFlowLogsOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The IDs of the flow logs.
	FlowLogIds []string `locationName:"flowLogIdSet" locationNameList:"item" type:"list"`

	// Information about the flow logs that could not be created successfully.
	Unsuccessful []UnsuccessfulItem `locationName:"unsuccessful" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateFlowLogsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFlowLogs = "CreateFlowLogs"

// CreateFlowLogsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates one or more flow logs to capture information about IP traffic for
// a specific network interface, subnet, or VPC.
//
// Flow log data for a monitored network interface is recorded as flow log records,
// which are log events consisting of fields that describe the traffic flow.
// For more information, see Flow Log Records (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/flow-logs.html#flow-log-records)
// in the Amazon Virtual Private Cloud User Guide.
//
// When publishing to CloudWatch Logs, flow log records are published to a log
// group, and each network interface has a unique log stream in the log group.
// When publishing to Amazon S3, flow log records for all of the monitored network
// interfaces are published to a single log file object that is stored in the
// specified bucket.
//
// For more information, see VPC Flow Logs (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/flow-logs.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateFlowLogsRequest.
//    req := client.CreateFlowLogsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateFlowLogs
func (c *Client) CreateFlowLogsRequest(input *CreateFlowLogsInput) CreateFlowLogsRequest {
	op := &aws.Operation{
		Name:       opCreateFlowLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFlowLogsInput{}
	}

	req := c.newRequest(op, input, &CreateFlowLogsOutput{})
	return CreateFlowLogsRequest{Request: req, Input: input, Copy: c.CreateFlowLogsRequest}
}

// CreateFlowLogsRequest is the request type for the
// CreateFlowLogs API operation.
type CreateFlowLogsRequest struct {
	*aws.Request
	Input *CreateFlowLogsInput
	Copy  func(*CreateFlowLogsInput) CreateFlowLogsRequest
}

// Send marshals and sends the CreateFlowLogs API request.
func (r CreateFlowLogsRequest) Send(ctx context.Context) (*CreateFlowLogsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFlowLogsResponse{
		CreateFlowLogsOutput: r.Request.Data.(*CreateFlowLogsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFlowLogsResponse is the response type for the
// CreateFlowLogs API operation.
type CreateFlowLogsResponse struct {
	*CreateFlowLogsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFlowLogs request.
func (r *CreateFlowLogsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
