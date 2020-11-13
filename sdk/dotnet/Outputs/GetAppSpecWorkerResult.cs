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
    public sealed class GetAppSpecWorkerResult
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
        public readonly ImmutableArray<Outputs.GetAppSpecWorkerEnvResult> Envs;
        /// <summary>
        /// A Git repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        public readonly Outputs.GetAppSpecWorkerGitResult? Git;
        /// <summary>
        /// A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        public readonly Outputs.GetAppSpecWorkerGithubResult? Github;
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
        public readonly Outputs.GetAppSpecWorkerRoutesResult Routes;
        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        public readonly string? RunCommand;
        /// <summary>
        /// An optional path to the working directory to use for the build.
        /// </summary>
        public readonly string? SourceDir;

        [OutputConstructor]
        private GetAppSpecWorkerResult(
            string? buildCommand,

            string? dockerfilePath,

            string? environmentSlug,

            ImmutableArray<Outputs.GetAppSpecWorkerEnvResult> envs,

            Outputs.GetAppSpecWorkerGitResult? git,

            Outputs.GetAppSpecWorkerGithubResult? github,

            int? instanceCount,

            string? instanceSizeSlug,

            string name,

            Outputs.GetAppSpecWorkerRoutesResult routes,

            string? runCommand,

            string? sourceDir)
        {
            BuildCommand = buildCommand;
            DockerfilePath = dockerfilePath;
            EnvironmentSlug = environmentSlug;
            Envs = envs;
            Git = git;
            Github = github;
            InstanceCount = instanceCount;
            InstanceSizeSlug = instanceSizeSlug;
            Name = name;
            Routes = routes;
            RunCommand = runCommand;
            SourceDir = sourceDir;
        }
    }
}
