// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides access to the available DigitalOcean Kubernetes Service versions.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/kubernetes_versions.html.md.
        /// </summary>
        [Obsolete("Use GetKubernetesVersions.InvokeAsync() instead")]
        public static Task<GetKubernetesVersionsResult> GetKubernetesVersions(GetKubernetesVersionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesVersionsResult>("digitalocean:index/getKubernetesVersions:getKubernetesVersions", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetKubernetesVersions
    {
        /// <summary>
        /// Provides access to the available DigitalOcean Kubernetes Service versions.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/kubernetes_versions.html.md.
        /// </summary>
        public static Task<GetKubernetesVersionsResult> InvokeAsync(GetKubernetesVersionsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKubernetesVersionsResult>("digitalocean:index/getKubernetesVersions:getKubernetesVersions", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetKubernetesVersionsArgs : Pulumi.InvokeArgs
    {
        [Input("versionPrefix")]
        public string? VersionPrefix { get; set; }

        public GetKubernetesVersionsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKubernetesVersionsResult
    {
        /// <summary>
        /// The most recent version available.
        /// </summary>
        public readonly string LatestVersion;
        /// <summary>
        /// A list of available versions.
        /// </summary>
        public readonly ImmutableArray<string> ValidVersions;
        public readonly string? VersionPrefix;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKubernetesVersionsResult(
            string latestVersion,
            ImmutableArray<string> validVersions,
            string? versionPrefix,
            string id)
        {
            LatestVersion = latestVersion;
            ValidVersions = validVersions;
            VersionPrefix = versionPrefix;
            Id = id;
        }
    }
}
