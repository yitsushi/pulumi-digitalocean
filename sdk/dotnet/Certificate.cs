// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a DigitalOcean Certificate resource that allows you to manage
    /// certificates for configuring TLS termination in Load Balancers.
    /// Certificates created with this resource can be referenced in your
    /// Load Balancer configuration via their ID. The certificate can either
    /// be a custom one provided by you or automatically generated one with
    /// Let's Encrypt.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/certificate.html.markdown.
    /// </summary>
    public partial class Certificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The full PEM-formatted trust chain
        /// between the certificate authority's certificate and your domain's TLS
        /// certificate. Only valid when type is `custom`.
        /// </summary>
        [Output("certificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// List of fully qualified domain names (FQDNs) for
        /// which the certificate will be issued. The domains must be managed using
        /// DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
        /// </summary>
        [Output("domains")]
        public Output<ImmutableArray<string>> Domains { get; private set; } = null!;

        /// <summary>
        /// The contents of a PEM-formatted public
        /// TLS certificate. Only valid when type is `custom`.
        /// </summary>
        [Output("leafCertificate")]
        public Output<string?> LeafCertificate { get; private set; } = null!;

        /// <summary>
        /// The name of the certificate for identification.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the certificate
        /// </summary>
        [Output("notAfter")]
        public Output<string> NotAfter { get; private set; } = null!;

        /// <summary>
        /// The contents of a PEM-formatted private-key
        /// corresponding to the SSL certificate. Only valid when type is `custom`.
        /// </summary>
        [Output("privateKey")]
        public Output<string?> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// The SHA-1 fingerprint of the certificate
        /// </summary>
        [Output("sha1Fingerprint")]
        public Output<string> Sha1Fingerprint { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of certificate to provision. Can be either
        /// `custom` or `lets_encrypt`. Defaults to `custom`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/certificate:Certificate", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/certificate:Certificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full PEM-formatted trust chain
        /// between the certificate authority's certificate and your domain's TLS
        /// certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// List of fully qualified domain names (FQDNs) for
        /// which the certificate will be issued. The domains must be managed using
        /// DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// The contents of a PEM-formatted public
        /// TLS certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("leafCertificate")]
        public Input<string>? LeafCertificate { get; set; }

        /// <summary>
        /// The name of the certificate for identification.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The contents of a PEM-formatted private-key
        /// corresponding to the SSL certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The type of certificate to provision. Can be either
        /// `custom` or `lets_encrypt`. Defaults to `custom`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateArgs()
        {
        }
    }

    public sealed class CertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full PEM-formatted trust chain
        /// between the certificate authority's certificate and your domain's TLS
        /// certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// List of fully qualified domain names (FQDNs) for
        /// which the certificate will be issued. The domains must be managed using
        /// DigitalOcean's DNS. Only valid when type is `lets_encrypt`.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        /// <summary>
        /// The contents of a PEM-formatted public
        /// TLS certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("leafCertificate")]
        public Input<string>? LeafCertificate { get; set; }

        /// <summary>
        /// The name of the certificate for identification.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The expiration date of the certificate
        /// </summary>
        [Input("notAfter")]
        public Input<string>? NotAfter { get; set; }

        /// <summary>
        /// The contents of a PEM-formatted private-key
        /// corresponding to the SSL certificate. Only valid when type is `custom`.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        /// <summary>
        /// The SHA-1 fingerprint of the certificate
        /// </summary>
        [Input("sha1Fingerprint")]
        public Input<string>? Sha1Fingerprint { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of certificate to provision. Can be either
        /// `custom` or `lets_encrypt`. Defaults to `custom`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public CertificateState()
        {
        }
    }
}
