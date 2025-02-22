// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get Docker credentials for your DigitalOcean container registry.
 *
 * An error is triggered if the provided container registry name does not exist.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * Get the container registry:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = new digitalocean.ContainerRegistryDockerCredentials("example", {
 *     registryName: "example",
 * });
 * ```
 * ### Docker Provider Example
 *
 * Use the `endpoint` and `dockerCredentials` with the Docker provider:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const exampleContainerRegistry = digitalocean.getContainerRegistry({
 *     name: "example",
 * });
 * const exampleContainerRegistryDockerCredentials = new digitalocean.ContainerRegistryDockerCredentials("exampleContainerRegistryDockerCredentials", {registryName: "example"});
 * ```
 */
export class ContainerRegistryDockerCredentials extends pulumi.CustomResource {
    /**
     * Get an existing ContainerRegistryDockerCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerRegistryDockerCredentialsState, opts?: pulumi.CustomResourceOptions): ContainerRegistryDockerCredentials {
        return new ContainerRegistryDockerCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials';

    /**
     * Returns true if the given object is an instance of ContainerRegistryDockerCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerRegistryDockerCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerRegistryDockerCredentials.__pulumiType;
    }

    public /*out*/ readonly credentialExpirationTime!: pulumi.Output<string>;
    public /*out*/ readonly dockerCredentials!: pulumi.Output<string>;
    /**
     * The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
     */
    public readonly expirySeconds!: pulumi.Output<number | undefined>;
    /**
     * The name of the container registry.
     */
    public readonly registryName!: pulumi.Output<string>;
    /**
     * Allow for write access to the container registry. Defaults to false.
     */
    public readonly write!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ContainerRegistryDockerCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerRegistryDockerCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerRegistryDockerCredentialsArgs | ContainerRegistryDockerCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerRegistryDockerCredentialsState | undefined;
            inputs["credentialExpirationTime"] = state ? state.credentialExpirationTime : undefined;
            inputs["dockerCredentials"] = state ? state.dockerCredentials : undefined;
            inputs["expirySeconds"] = state ? state.expirySeconds : undefined;
            inputs["registryName"] = state ? state.registryName : undefined;
            inputs["write"] = state ? state.write : undefined;
        } else {
            const args = argsOrState as ContainerRegistryDockerCredentialsArgs | undefined;
            if ((!args || args.registryName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryName'");
            }
            inputs["expirySeconds"] = args ? args.expirySeconds : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["write"] = args ? args.write : undefined;
            inputs["credentialExpirationTime"] = undefined /*out*/;
            inputs["dockerCredentials"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContainerRegistryDockerCredentials.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerRegistryDockerCredentials resources.
 */
export interface ContainerRegistryDockerCredentialsState {
    credentialExpirationTime?: pulumi.Input<string>;
    dockerCredentials?: pulumi.Input<string>;
    /**
     * The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
     */
    expirySeconds?: pulumi.Input<number>;
    /**
     * The name of the container registry.
     */
    registryName?: pulumi.Input<string>;
    /**
     * Allow for write access to the container registry. Defaults to false.
     */
    write?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ContainerRegistryDockerCredentials resource.
 */
export interface ContainerRegistryDockerCredentialsArgs {
    /**
     * The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
     */
    expirySeconds?: pulumi.Input<number>;
    /**
     * The name of the container registry.
     */
    registryName: pulumi.Input<string>;
    /**
     * Allow for write access to the container registry. Defaults to false.
     */
    write?: pulumi.Input<boolean>;
}
