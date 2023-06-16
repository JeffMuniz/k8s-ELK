// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilLoadBalancerAvailable uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/x/net v0.7.0pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilLoadBalancerAvailable(ctx context.Context, input *DescribeLoadBalancersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilLoadBalancerAvailable",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "active",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathAnyWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "provisioning",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilLoadBalancerExists uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/x/net v0.7.0pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilLoadBalancerExists(ctx context.Context, input *DescribeLoadBalancersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilLoadBalancerExists",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilLoadBalancersDeleted uses the Elastic Load Balancing v2 API operation
// DescribeLoadBalancers to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/x/net v0.7.0pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilLoadBalancersDeleted(ctx context.Context, input *DescribeLoadBalancersInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilLoadBalancersDeleted",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "LoadBalancers[].State.Code",
				Expected: "active",
			},
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "LoadBalancerNotFound",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeLoadBalancersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeLoadBalancersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTargetDeregistered uses the Elastic Load Balancing v2 API operation
// DescribeTargetHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/x/net v0.7.0pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTargetDeregistered(ctx context.Context, input *DescribeTargetHealthInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTargetDeregistered",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "InvalidTarget",
			},
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "TargetHealthDescriptions[].TargetHealth.State",
				Expected: "unused",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTargetHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTargetHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilTargetInService uses the Elastic Load Balancing v2 API operation
// DescribeTargetHealth to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/x/net v0.7.0pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilTargetInService(ctx context.Context, input *DescribeTargetHealthInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilTargetInService",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "TargetHealthDescriptions[].TargetHealth.State",
				Expected: "healthy",
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "InvalidInstance",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeTargetHealthInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeTargetHealthRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
