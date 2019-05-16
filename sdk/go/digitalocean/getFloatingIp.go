// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information on a floating ip. This data source provides the region and Droplet id
// as configured on your DigitalOcean account. This is useful if the floating IP
// in question is not managed by Terraform or you need to find the Droplet the IP is
// attached to.
// 
// An error is triggered if the provided floating IP does not exist.
func LookupFloatingIp(ctx *pulumi.Context, args *GetFloatingIpArgs) (*GetFloatingIpResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["ipAddress"] = args.IpAddress
	}
	outputs, err := ctx.Invoke("digitalocean:index/getFloatingIp:getFloatingIp", inputs)
	if err != nil {
		return nil, err
	}
	return &GetFloatingIpResult{
		DropletId: outputs["dropletId"],
		IpAddress: outputs["ipAddress"],
		Region: outputs["region"],
		Urn: outputs["urn"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getFloatingIp.
type GetFloatingIpArgs struct {
	// The allocated IP address of the specific floating IP to retrieve.
	IpAddress interface{}
}

// A collection of values returned by getFloatingIp.
type GetFloatingIpResult struct {
	DropletId interface{}
	IpAddress interface{}
	Region interface{}
	Urn interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
