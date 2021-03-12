// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class GetFirewallOutboundRuleArgs : Pulumi.InvokeArgs
    {
        [Input("destinationAddresses")]
        private List<string>? _destinationAddresses;

        /// <summary>
        /// An array of strings containing the IPv4
        /// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
        /// outbound traffic will be allowed.
        /// </summary>
        public List<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new List<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationDropletIds")]
        private List<int>? _destinationDropletIds;

        /// <summary>
        /// An array containing the IDs of
        /// the Droplets to which the outbound traffic will be allowed.
        /// </summary>
        public List<int> DestinationDropletIds
        {
            get => _destinationDropletIds ?? (_destinationDropletIds = new List<int>());
            set => _destinationDropletIds = value;
        }

        [Input("destinationLoadBalancerUids")]
        private List<string>? _destinationLoadBalancerUids;

        /// <summary>
        /// An array containing the IDs
        /// of the Load Balancers to which the outbound traffic will be allowed.
        /// </summary>
        public List<string> DestinationLoadBalancerUids
        {
            get => _destinationLoadBalancerUids ?? (_destinationLoadBalancerUids = new List<string>());
            set => _destinationLoadBalancerUids = value;
        }

        [Input("destinationTags")]
        private List<string>? _destinationTags;

        /// <summary>
        /// An array containing the names of Tags
        /// corresponding to groups of Droplets to which the outbound traffic will
        /// be allowed.
        /// traffic.
        /// </summary>
        public List<string> DestinationTags
        {
            get => _destinationTags ?? (_destinationTags = new List<string>());
            set => _destinationTags = value;
        }

        /// <summary>
        /// The ports on which traffic will be allowed
        /// specified as a string containing a single port, a range (e.g. "8000-9000"),
        /// or "1-65535" to open all ports for a protocol. Required for when protocol is
        /// `tcp` or `udp`.
        /// </summary>
        [Input("portRange")]
        public string? PortRange { get; set; }

        /// <summary>
        /// The type of traffic to be allowed.
        /// This may be one of "tcp", "udp", or "icmp".
        /// </summary>
        [Input("protocol", required: true)]
        public string Protocol { get; set; } = null!;

        public GetFirewallOutboundRuleArgs()
        {
        }
    }
}
