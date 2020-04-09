// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean DNS record resource.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/record.html.markdown.
type DnsRecord struct {
	pulumi.CustomResourceState

	// The domain to add the record to.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
	Flags pulumi.IntPtrOutput `pulumi:"flags"`
	// The FQDN of the record
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the record. Use `@` for records on domain's name itself.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
	// The time to live for the record, in seconds. Must be at least 0.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the record.
	Value pulumi.StringOutput `pulumi:"value"`
	// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewDnsRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsRecord(ctx *pulumi.Context,
	name string, args *DnsRecordArgs, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &DnsRecordArgs{}
	}
	var resource DnsRecord
	err := ctx.RegisterResource("digitalocean:index/dnsRecord:DnsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecord gets an existing DnsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordState, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	var resource DnsRecord
	err := ctx.ReadResource("digitalocean:index/dnsRecord:DnsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecord resources.
type dnsRecordState struct {
	// The domain to add the record to.
	Domain *string `pulumi:"domain"`
	// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
	Flags *int `pulumi:"flags"`
	// The FQDN of the record
	Fqdn *string `pulumi:"fqdn"`
	// The name of the record. Use `@` for records on domain's name itself.
	Name *string `pulumi:"name"`
	// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
	Port *int `pulumi:"port"`
	// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
	Priority *int `pulumi:"priority"`
	// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
	Tag *string `pulumi:"tag"`
	// The time to live for the record, in seconds. Must be at least 0.
	Ttl *int `pulumi:"ttl"`
	// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
	Type *string `pulumi:"type"`
	// The value of the record.
	Value *string `pulumi:"value"`
	// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
	Weight *int `pulumi:"weight"`
}

type DnsRecordState struct {
	// The domain to add the record to.
	Domain pulumi.StringPtrInput
	// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
	Flags pulumi.IntPtrInput
	// The FQDN of the record
	Fqdn pulumi.StringPtrInput
	// The name of the record. Use `@` for records on domain's name itself.
	Name pulumi.StringPtrInput
	// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
	Port pulumi.IntPtrInput
	// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
	Priority pulumi.IntPtrInput
	// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
	Tag pulumi.StringPtrInput
	// The time to live for the record, in seconds. Must be at least 0.
	Ttl pulumi.IntPtrInput
	// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
	Type pulumi.StringPtrInput
	// The value of the record.
	Value pulumi.StringPtrInput
	// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
	Weight pulumi.IntPtrInput
}

func (DnsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordState)(nil)).Elem()
}

type dnsRecordArgs struct {
	// The domain to add the record to.
	Domain string `pulumi:"domain"`
	// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
	Flags *int `pulumi:"flags"`
	// The name of the record. Use `@` for records on domain's name itself.
	Name *string `pulumi:"name"`
	// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
	Port *int `pulumi:"port"`
	// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
	Priority *int `pulumi:"priority"`
	// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
	Tag *string `pulumi:"tag"`
	// The time to live for the record, in seconds. Must be at least 0.
	Ttl *int `pulumi:"ttl"`
	// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
	Type string `pulumi:"type"`
	// The value of the record.
	Value string `pulumi:"value"`
	// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a DnsRecord resource.
type DnsRecordArgs struct {
	// The domain to add the record to.
	Domain pulumi.StringInput
	// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
	Flags pulumi.IntPtrInput
	// The name of the record. Use `@` for records on domain's name itself.
	Name pulumi.StringPtrInput
	// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
	Port pulumi.IntPtrInput
	// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
	Priority pulumi.IntPtrInput
	// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
	Tag pulumi.StringPtrInput
	// The time to live for the record, in seconds. Must be at least 0.
	Ttl pulumi.IntPtrInput
	// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
	Type pulumi.StringInput
	// The value of the record.
	Value pulumi.StringInput
	// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
	Weight pulumi.IntPtrInput
}

func (DnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordArgs)(nil)).Elem()
}
