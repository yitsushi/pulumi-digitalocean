// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.DigitalOcean
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("digitalocean");
        /// <summary>
        /// The URL to use for the DigitalOcean API.
        /// </summary>
        public static string? ApiEndpoint { get; set; } = __config.Get("apiEndpoint") ?? Utilities.GetEnv("DIGITALOCEAN_API_URL") ?? "https://api.digitalocean.com";

        /// <summary>
        /// The access key ID for Spaces API operations.
        /// </summary>
        public static string? SpacesAccessId { get; set; } = __config.Get("spacesAccessId");

        /// <summary>
        /// The URL to use for the DigitalOcean Spaces API.
        /// </summary>
        public static string? SpacesEndpoint { get; set; } = __config.Get("spacesEndpoint") ?? Utilities.GetEnv("SPACES_ENDPOINT_URL");

        /// <summary>
        /// The secret access key for Spaces API operations.
        /// </summary>
        public static string? SpacesSecretKey { get; set; } = __config.Get("spacesSecretKey");

        /// <summary>
        /// The token key for API operations.
        /// </summary>
        public static string? Token { get; set; } = __config.Get("token");

    }
}
