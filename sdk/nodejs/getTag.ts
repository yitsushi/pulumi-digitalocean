// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Get information on a tag. This data source provides the name as configured on
 * your DigitalOcean account. This is useful if the tag name in question is not
 * managed by the provider or you need validate if the tag exists in the account.
 *
 * An error is triggered if the provided tag name does not exist.
 *
 * ## Example Usage
 *
 * Get the tag:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const exampleTag = digitalocean.getTag({
 *     name: "example",
 * });
 * const exampleDroplet = new digitalocean.Droplet("exampleDroplet", {
 *     image: "ubuntu-18-04-x64",
 *     region: "nyc2",
 *     size: "s-1vcpu-1gb",
 *     tags: [exampleTag.then(exampleTag => exampleTag.name)],
 * });
 * ```
 */
export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<GetTagResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getTag:getTag", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTag.
 */
export interface GetTagArgs {
    /**
     * The name of the tag.
     */
    name: string;
}

/**
 * A collection of values returned by getTag.
 */
export interface GetTagResult {
    /**
     * A count of the database clusters that the tag is applied to.
     */
    readonly databasesCount: number;
    /**
     * A count of the Droplets the tag is applied to.
     */
    readonly dropletsCount: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A count of the images that the tag is applied to.
     */
    readonly imagesCount: number;
    readonly name: string;
    /**
     * A count of the total number of resources that the tag is applied to.
     */
    readonly totalResourceCount: number;
    /**
     * A count of the volume snapshots that the tag is applied to.
     */
    readonly volumeSnapshotsCount: number;
    /**
     * A count of the volumes that the tag is applied to.
     */
    readonly volumesCount: number;
}
