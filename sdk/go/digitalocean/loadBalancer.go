// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Load Balancer resource. This can be used to create,
// modify, and delete Load Balancers.
//
// ## Import
//
// Load Balancers can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/loadBalancer:LoadBalancer myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either `roundRobin`
	// or `leastConnections`. The default value is `roundRobin`.
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds pulumi.IntArrayOutput `pulumi:"dropletIds"`
	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag pulumi.StringPtrOutput `pulumi:"dropletTag"`
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
	EnableBackendKeepalive pulumi.BoolPtrOutput `pulumi:"enableBackendKeepalive"`
	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is `false`.
	EnableProxyProtocol pulumi.BoolPtrOutput `pulumi:"enableProxyProtocol"`
	// A list of `forwardingRule` to be assigned to the
	// Load Balancer. The `forwardingRule` block is documented below.
	ForwardingRules LoadBalancerForwardingRuleArrayOutput `pulumi:"forwardingRules"`
	// A `healthcheck` block to be assigned to the
	// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
	Healthcheck LoadBalancerHealthcheckOutput `pulumi:"healthcheck"`
	Ip          pulumi.StringOutput           `pulumi:"ip"`
	// The uniform resource name for the Load Balancer
	LoadBalancerUrn pulumi.StringOutput `pulumi:"loadBalancerUrn"`
	// The Load Balancer name
	Name pulumi.StringOutput `pulumi:"name"`
	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is `false`.
	RedirectHttpToHttps pulumi.BoolPtrOutput `pulumi:"redirectHttpToHttps"`
	// The region to start in
	Region pulumi.StringOutput `pulumi:"region"`
	Status pulumi.StringOutput `pulumi:"status"`
	// A `stickySessions` block to be assigned to the
	// Load Balancer. The `stickySessions` block is documented below. Only 1 stickySessions block is allowed.
	StickySessions LoadBalancerStickySessionsOutput `pulumi:"stickySessions"`
	// The ID of the VPC where the load balancer will be located.
	VpcUuid pulumi.StringOutput `pulumi:"vpcUuid"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil || args.ForwardingRules == nil {
		return nil, errors.New("missing required argument 'ForwardingRules'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil {
		args = &LoadBalancerArgs{}
	}
	var resource LoadBalancer
	err := ctx.RegisterResource("digitalocean:index/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("digitalocean:index/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either `roundRobin`
	// or `leastConnections`. The default value is `roundRobin`.
	Algorithm *string `pulumi:"algorithm"`
	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds []int `pulumi:"dropletIds"`
	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag *string `pulumi:"dropletTag"`
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
	EnableBackendKeepalive *bool `pulumi:"enableBackendKeepalive"`
	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is `false`.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// A list of `forwardingRule` to be assigned to the
	// Load Balancer. The `forwardingRule` block is documented below.
	ForwardingRules []LoadBalancerForwardingRule `pulumi:"forwardingRules"`
	// A `healthcheck` block to be assigned to the
	// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
	Healthcheck *LoadBalancerHealthcheck `pulumi:"healthcheck"`
	Ip          *string                  `pulumi:"ip"`
	// The uniform resource name for the Load Balancer
	LoadBalancerUrn *string `pulumi:"loadBalancerUrn"`
	// The Load Balancer name
	Name *string `pulumi:"name"`
	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is `false`.
	RedirectHttpToHttps *bool `pulumi:"redirectHttpToHttps"`
	// The region to start in
	Region *string `pulumi:"region"`
	Status *string `pulumi:"status"`
	// A `stickySessions` block to be assigned to the
	// Load Balancer. The `stickySessions` block is documented below. Only 1 stickySessions block is allowed.
	StickySessions *LoadBalancerStickySessions `pulumi:"stickySessions"`
	// The ID of the VPC where the load balancer will be located.
	VpcUuid *string `pulumi:"vpcUuid"`
}

type LoadBalancerState struct {
	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either `roundRobin`
	// or `leastConnections`. The default value is `roundRobin`.
	Algorithm pulumi.StringPtrInput
	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds pulumi.IntArrayInput
	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag pulumi.StringPtrInput
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
	EnableBackendKeepalive pulumi.BoolPtrInput
	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is `false`.
	EnableProxyProtocol pulumi.BoolPtrInput
	// A list of `forwardingRule` to be assigned to the
	// Load Balancer. The `forwardingRule` block is documented below.
	ForwardingRules LoadBalancerForwardingRuleArrayInput
	// A `healthcheck` block to be assigned to the
	// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
	Healthcheck LoadBalancerHealthcheckPtrInput
	Ip          pulumi.StringPtrInput
	// The uniform resource name for the Load Balancer
	LoadBalancerUrn pulumi.StringPtrInput
	// The Load Balancer name
	Name pulumi.StringPtrInput
	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is `false`.
	RedirectHttpToHttps pulumi.BoolPtrInput
	// The region to start in
	Region pulumi.StringPtrInput
	Status pulumi.StringPtrInput
	// A `stickySessions` block to be assigned to the
	// Load Balancer. The `stickySessions` block is documented below. Only 1 stickySessions block is allowed.
	StickySessions LoadBalancerStickySessionsPtrInput
	// The ID of the VPC where the load balancer will be located.
	VpcUuid pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either `roundRobin`
	// or `leastConnections`. The default value is `roundRobin`.
	Algorithm *string `pulumi:"algorithm"`
	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds []int `pulumi:"dropletIds"`
	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag *string `pulumi:"dropletTag"`
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
	EnableBackendKeepalive *bool `pulumi:"enableBackendKeepalive"`
	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is `false`.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// A list of `forwardingRule` to be assigned to the
	// Load Balancer. The `forwardingRule` block is documented below.
	ForwardingRules []LoadBalancerForwardingRule `pulumi:"forwardingRules"`
	// A `healthcheck` block to be assigned to the
	// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
	Healthcheck *LoadBalancerHealthcheck `pulumi:"healthcheck"`
	// The Load Balancer name
	Name *string `pulumi:"name"`
	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is `false`.
	RedirectHttpToHttps *bool `pulumi:"redirectHttpToHttps"`
	// The region to start in
	Region string `pulumi:"region"`
	// A `stickySessions` block to be assigned to the
	// Load Balancer. The `stickySessions` block is documented below. Only 1 stickySessions block is allowed.
	StickySessions *LoadBalancerStickySessions `pulumi:"stickySessions"`
	// The ID of the VPC where the load balancer will be located.
	VpcUuid *string `pulumi:"vpcUuid"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// The load balancing algorithm used to determine
	// which backend Droplet will be selected by a client. It must be either `roundRobin`
	// or `leastConnections`. The default value is `roundRobin`.
	Algorithm pulumi.StringPtrInput
	// A list of the IDs of each droplet to be attached to the Load Balancer.
	DropletIds pulumi.IntArrayInput
	// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
	DropletTag pulumi.StringPtrInput
	// A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
	EnableBackendKeepalive pulumi.BoolPtrInput
	// A boolean value indicating whether PROXY
	// Protocol should be used to pass information from connecting client requests to
	// the backend service. Default value is `false`.
	EnableProxyProtocol pulumi.BoolPtrInput
	// A list of `forwardingRule` to be assigned to the
	// Load Balancer. The `forwardingRule` block is documented below.
	ForwardingRules LoadBalancerForwardingRuleArrayInput
	// A `healthcheck` block to be assigned to the
	// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
	Healthcheck LoadBalancerHealthcheckPtrInput
	// The Load Balancer name
	Name pulumi.StringPtrInput
	// A boolean value indicating whether
	// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
	// Default value is `false`.
	RedirectHttpToHttps pulumi.BoolPtrInput
	// The region to start in
	Region pulumi.StringInput
	// A `stickySessions` block to be assigned to the
	// Load Balancer. The `stickySessions` block is documented below. Only 1 stickySessions block is allowed.
	StickySessions LoadBalancerStickySessionsPtrInput
	// The ID of the VPC where the load balancer will be located.
	VpcUuid pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

type LoadBalancerOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerOutput)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerOutput{})
}
