// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// The provider type for the digitalocean package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/index.html.markdown.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, ResourceOptions? options = null)
            : base("digitalocean", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private static ResourceOptions MakeResourceOptions(ResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL to use for the DigitalOcean API.
        /// </summary>
        [Input("apiEndpoint")]
        public Input<string>? ApiEndpoint { get; set; }

        /// <summary>
        /// The access key ID for Spaces API operations.
        /// </summary>
        [Input("spacesAccessId")]
        public Input<string>? SpacesAccessId { get; set; }

        /// <summary>
        /// The secret access key for Spaces API operations.
        /// </summary>
        [Input("spacesSecretKey")]
        public Input<string>? SpacesSecretKey { get; set; }

        /// <summary>
        /// The token key for API operations.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public ProviderArgs()
        {
            ApiEndpoint = Utilities.GetEnv("DIGITALOCEAN_API_URL") ?? "https://api.digitalocean.com";
            SpacesAccessId = Utilities.GetEnv("SPACES_ACCESS_KEY_ID");
            SpacesSecretKey = Utilities.GetEnv("SPACES_SECRET_ACCESS_KEY");
            Token = Utilities.GetEnv("DIGITALOCEAN_TOKEN");
        }
    }
}
