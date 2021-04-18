// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on images for use in other resources (e.g. creating a Droplet
// based on a snapshot), with the ability to filter and sort the results. If no filters are specified,
// all images will be returned.
//
// This data source is useful if the image in question is not managed by the provider or you need to utilize any
// of the image's data.
//
// Note: You can use the `getImage` data source to obtain metadata
// about a single image if you already know the `slug`, unique `name`, or `id` to retrieve.
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("digitalocean:index/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	// Filter the results.
	// The `filter` block is documented below.
	Filters []GetImagesFilter `pulumi:"filters"`
	// Sort the results.
	// The `sort` block is documented below.
	Sorts []GetImagesSort `pulumi:"sorts"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	Filters []GetImagesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of images satisfying any `filter` and `sort` criteria. Each image has the following attributes:
	// - `slug`: Unique text identifier of the image.
	// - `id`: The ID of the image.
	// - `name`: The name of the image.
	// - `type`: Type of the image.
	Images []GetImagesImage `pulumi:"images"`
	Sorts  []GetImagesSort  `pulumi:"sorts"`
}
