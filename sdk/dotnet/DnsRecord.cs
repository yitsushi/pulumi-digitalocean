// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a DigitalOcean DNS record resource.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/record.html.markdown.
    /// </summary>
    public partial class DnsRecord : Pulumi.CustomResource
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Output("flags")]
        public Output<int?> Flags { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The name of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/dnsRecord:DnsRecord", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// The name of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DnsRecordArgs()
        {
        }
    }

    public sealed class DnsRecordState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The name of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DnsRecordState()
        {
        }
    }
}
