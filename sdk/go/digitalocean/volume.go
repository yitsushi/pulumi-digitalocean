// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DigitalOcean Block Storage volume which can be attached to a Droplet in order to provide expanded storage.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/volume.html.markdown.
type Volume struct {
	pulumi.CustomResourceState

	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of associated droplet ids.
	DropletIds pulumi.IntArrayOutput `pulumi:"dropletIds"`
	// Filesystem label for the block storage volume.
	FilesystemLabel pulumi.StringOutput `pulumi:"filesystemLabel"`
	// Filesystem type (`xfs` or `ext4`) for the block storage volume.
	FilesystemType pulumi.StringOutput `pulumi:"filesystemType"`
	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel pulumi.StringPtrOutput `pulumi:"initialFilesystemLabel"`
	// Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
	InitialFilesystemType pulumi.StringPtrOutput `pulumi:"initialFilesystemType"`
	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region that the block storage volume will be created in.
	Region pulumi.StringOutput `pulumi:"region"`
	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size pulumi.IntOutput `pulumi:"size"`
	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// A list of the tags to be applied to this Volume.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// the uniform resource name for the volume.
	Urn pulumi.StringOutput `pulumi:"urn"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &VolumeArgs{}
	}
	var resource Volume
	err := ctx.RegisterResource("digitalocean:index/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("digitalocean:index/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description *string `pulumi:"description"`
	// A list of associated droplet ids.
	DropletIds []int `pulumi:"dropletIds"`
	// Filesystem label for the block storage volume.
	FilesystemLabel *string `pulumi:"filesystemLabel"`
	// Filesystem type (`xfs` or `ext4`) for the block storage volume.
	FilesystemType *string `pulumi:"filesystemType"`
	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel *string `pulumi:"initialFilesystemLabel"`
	// Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
	InitialFilesystemType *string `pulumi:"initialFilesystemType"`
	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
	Name *string `pulumi:"name"`
	// The region that the block storage volume will be created in.
	Region *string `pulumi:"region"`
	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size *int `pulumi:"size"`
	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of the tags to be applied to this Volume.
	Tags []string `pulumi:"tags"`
	// the uniform resource name for the volume.
	Urn *string `pulumi:"urn"`
}

type VolumeState struct {
	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description pulumi.StringPtrInput
	// A list of associated droplet ids.
	DropletIds pulumi.IntArrayInput
	// Filesystem label for the block storage volume.
	FilesystemLabel pulumi.StringPtrInput
	// Filesystem type (`xfs` or `ext4`) for the block storage volume.
	FilesystemType pulumi.StringPtrInput
	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel pulumi.StringPtrInput
	// Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
	InitialFilesystemType pulumi.StringPtrInput
	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
	Name pulumi.StringPtrInput
	// The region that the block storage volume will be created in.
	Region pulumi.StringPtrInput
	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size pulumi.IntPtrInput
	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
	SnapshotId pulumi.StringPtrInput
	// A list of the tags to be applied to this Volume.
	Tags pulumi.StringArrayInput
	// the uniform resource name for the volume.
	Urn pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description *string `pulumi:"description"`
	// Filesystem type (`xfs` or `ext4`) for the block storage volume.
	FilesystemType *string `pulumi:"filesystemType"`
	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel *string `pulumi:"initialFilesystemLabel"`
	// Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
	InitialFilesystemType *string `pulumi:"initialFilesystemType"`
	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
	Name *string `pulumi:"name"`
	// The region that the block storage volume will be created in.
	Region string `pulumi:"region"`
	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size int `pulumi:"size"`
	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of the tags to be applied to this Volume.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
	Description pulumi.StringPtrInput
	// Filesystem type (`xfs` or `ext4`) for the block storage volume.
	FilesystemType pulumi.StringPtrInput
	// Initial filesystem label for the block storage volume.
	InitialFilesystemLabel pulumi.StringPtrInput
	// Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
	InitialFilesystemType pulumi.StringPtrInput
	// A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
	Name pulumi.StringPtrInput
	// The region that the block storage volume will be created in.
	Region pulumi.StringInput
	// The size of the block storage volume in GiB. If updated, can only be expanded.
	Size pulumi.IntInput
	// The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
	SnapshotId pulumi.StringPtrInput
	// A list of the tags to be applied to this Volume.
	Tags pulumi.StringArrayInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

