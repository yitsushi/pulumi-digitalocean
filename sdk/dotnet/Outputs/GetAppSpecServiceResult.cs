// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Outputs
{

    [OutputType]
    public sealed class GetAppSpecServiceResult
    {
        /// <summary>
        /// An optional build command to run while building this component from source.
        /// </summary>
        public readonly string? BuildCommand;
        /// <summary>
        /// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
        /// </summary>
        public readonly string? DockerfilePath;
        /// <summary>
        /// An environment slug describing the type of this app.
        /// </summary>
        public readonly string? EnvironmentSlug;
        /// <summary>
        /// Describes an environment variable made available to an app competent.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppSpecServiceEnvResult> Envs;
        /// <summary>
        /// A Git repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        public readonly Outputs.GetAppSpecServiceGitResult? Git;
        /// <summary>
        /// A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        public readonly Outputs.GetAppSpecServiceGithubResult? Github;
        public readonly Outputs.GetAppSpecServiceGitlabResult? Gitlab;
        /// <summary>
        /// A health check to determine the availability of this component.
        /// </summary>
        public readonly Outputs.GetAppSpecServiceHealthCheckResult? HealthCheck;
        /// <summary>
        /// The internal port on which this service's run command will listen.
        /// </summary>
        public readonly int HttpPort;
        /// <summary>
        /// The amount of instances that this component should be scaled to.
        /// </summary>
        public readonly int? InstanceCount;
        /// <summary>
        /// The instance size to use for this component.
        /// </summary>
        public readonly string? InstanceSizeSlug;
        /// <summary>
        /// The name of the component.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.GetAppSpecServiceRoutesResult Routes;
        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        public readonly string RunCommand;
        /// <summary>
        /// An optional path to the working directory to use for the build.
        /// </summary>
        public readonly string? SourceDir;

        [OutputConstructor]
        private GetAppSpecServiceResult(
            string? buildCommand,

            string? dockerfilePath,

            string? environmentSlug,

            ImmutableArray<Outputs.GetAppSpecServiceEnvResult> envs,

            Outputs.GetAppSpecServiceGitResult? git,

            Outputs.GetAppSpecServiceGithubResult? github,

            Outputs.GetAppSpecServiceGitlabResult? gitlab,

            Outputs.GetAppSpecServiceHealthCheckResult? healthCheck,

            int httpPort,

            int? instanceCount,

            string? instanceSizeSlug,

            string name,

            Outputs.GetAppSpecServiceRoutesResult routes,

            string runCommand,

            string? sourceDir)
        {
            BuildCommand = buildCommand;
            DockerfilePath = dockerfilePath;
            EnvironmentSlug = environmentSlug;
            Envs = envs;
            Git = git;
            Github = github;
            Gitlab = gitlab;
            HealthCheck = healthCheck;
            HttpPort = httpPort;
            InstanceCount = instanceCount;
            InstanceSizeSlug = instanceSizeSlug;
            Name = name;
            Routes = routes;
            RunCommand = runCommand;
            SourceDir = sourceDir;
        }
    }
}
