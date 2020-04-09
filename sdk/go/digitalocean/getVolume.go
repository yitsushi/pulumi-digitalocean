// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("digitalocean:index/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	// Text describing a block storage volume.
	Description *string `pulumi:"description"`
	// The name of block storage volume.
	Name string `pulumi:"name"`
	// The region the block storage volume is provisioned in.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	// Text describing a block storage volume.
	Description *string `pulumi:"description"`
	// A list of associated Droplet ids.
	DropletIds []int `pulumi:"dropletIds"`
	// Filesystem label currently in-use on the block storage volume.
	FilesystemLabel string `pulumi:"filesystemLabel"`
	// Filesystem type currently in-use on the block storage volume.
	FilesystemType string `pulumi:"filesystemType"`
	// id is the provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The size of the block storage volume in GiB.
	Size int `pulumi:"size"`
	// A list of the tags associated to the Volume.
	Tags []string `pulumi:"tags"`
	Urn  string   `pulumi:"urn"`
}
