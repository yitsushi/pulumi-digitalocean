// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the digitalocean package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.ApiEndpoint == nil {
		args.ApiEndpoint = pulumi.StringPtr(getEnvOrDefault("https://api.digitalocean.com", nil, "DIGITALOCEAN_API_URL").(string))
	}
	if args.SpacesAccessId == nil {
		args.SpacesAccessId = pulumi.StringPtr(getEnvOrDefault("", nil, "SPACES_ACCESS_KEY_ID").(string))
	}
	if args.SpacesSecretKey == nil {
		args.SpacesSecretKey = pulumi.StringPtr(getEnvOrDefault("", nil, "SPACES_SECRET_ACCESS_KEY").(string))
	}
	if args.Token == nil {
		args.Token = pulumi.StringPtr(getEnvOrDefault("", nil, "DIGITALOCEAN_TOKEN").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:digitalocean", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The URL to use for the DigitalOcean API.
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// The access key ID for Spaces API operations.
	SpacesAccessId *string `pulumi:"spacesAccessId"`
	// The secret access key for Spaces API operations.
	SpacesSecretKey *string `pulumi:"spacesSecretKey"`
	// The token key for API operations.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The URL to use for the DigitalOcean API.
	ApiEndpoint pulumi.StringPtrInput
	// The access key ID for Spaces API operations.
	SpacesAccessId pulumi.StringPtrInput
	// The secret access key for Spaces API operations.
	SpacesSecretKey pulumi.StringPtrInput
	// The token key for API operations.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
