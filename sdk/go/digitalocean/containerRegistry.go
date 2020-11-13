// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Container Registry resource. A Container Registry is
// a secure, private location to store your containers for rapid deployment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewContainerRegistry(ctx, "foobar", &digitalocean.ContainerRegistryArgs{
// 			SubscriptionTierSlug: pulumi.String("starter"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ContainerRegistry struct {
	pulumi.CustomResourceState

	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the container_registry
	Name      pulumi.StringOutput `pulumi:"name"`
	ServerUrl pulumi.StringOutput `pulumi:"serverUrl"`
	// The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
	SubscriptionTierSlug pulumi.StringOutput `pulumi:"subscriptionTierSlug"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	if args == nil || args.SubscriptionTierSlug == nil {
		return nil, errors.New("missing required argument 'SubscriptionTierSlug'")
	}
	if args == nil {
		args = &ContainerRegistryArgs{}
	}
	var resource ContainerRegistry
	err := ctx.RegisterResource("digitalocean:index/containerRegistry:ContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	err := ctx.ReadResource("digitalocean:index/containerRegistry:ContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
	Endpoint *string `pulumi:"endpoint"`
	// The name of the container_registry
	Name      *string `pulumi:"name"`
	ServerUrl *string `pulumi:"serverUrl"`
	// The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
	SubscriptionTierSlug *string `pulumi:"subscriptionTierSlug"`
}

type ContainerRegistryState struct {
	Endpoint pulumi.StringPtrInput
	// The name of the container_registry
	Name      pulumi.StringPtrInput
	ServerUrl pulumi.StringPtrInput
	// The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
	SubscriptionTierSlug pulumi.StringPtrInput
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	// The name of the container_registry
	Name *string `pulumi:"name"`
	// The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
	SubscriptionTierSlug string `pulumi:"subscriptionTierSlug"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	// The name of the container_registry
	Name pulumi.StringPtrInput
	// The slug identifier for the subscription tier to use (`starter`, `basic`, or `professional`)
	SubscriptionTierSlug pulumi.StringInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}
