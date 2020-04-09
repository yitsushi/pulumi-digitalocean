// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Droplet snapshots are saved instances of a Droplet. Use this data
 * source to retrieve the ID of a DigitalOcean Droplet snapshot for use in other
 * resources.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const webSnapshot = digitalocean.getDropletSnapshot({
 *     mostRecent: true,
 *     nameRegex: "^web",
 *     region: "nyc3",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/droplet_snapshot.html.md.
 */
export function getDropletSnapshot(args?: GetDropletSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetDropletSnapshotResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getDropletSnapshot:getDropletSnapshot", {
        "mostRecent": args.mostRecent,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDropletSnapshot.
 */
export interface GetDropletSnapshotArgs {
    /**
     * If more than one result is returned, use the most recent Droplet snapshot.
     */
    readonly mostRecent?: boolean;
    /**
     * The name of the Droplet snapshot.
     */
    readonly name?: string;
    /**
     * A regex string to apply to the Droplet snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
     */
    readonly nameRegex?: string;
    /**
     * A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only Droplet snapshots available in the region will be returned.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getDropletSnapshot.
 */
export interface GetDropletSnapshotResult {
    /**
     * The date and time the Droplet snapshot was created.
     */
    readonly createdAt: string;
    /**
     * The ID of the Droplet from which the Droplet snapshot originated.
     */
    readonly dropletId: string;
    /**
     * The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.
     */
    readonly minDiskSize: number;
    readonly mostRecent?: boolean;
    readonly name?: string;
    readonly nameRegex?: string;
    readonly region?: string;
    /**
     * A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.
     */
    readonly regions: string[];
    /**
     * The billable size of the Droplet snapshot in gigabytes.
     */
    readonly size: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
