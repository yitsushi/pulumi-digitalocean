// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Retrieve information about all DigitalOcean projects associated with an account, with
 * the ability to filter and sort the results. If no filters are specified, all projects
 * will be returned.
 *
 * Note: You can use the `digitalocean.Project` data source to
 * obtain metadata about a single project if you already know the `id` to retrieve or the unique
 * `name` of the project.
 *
 * ## Example Usage
 *
 * Use the `filter` block with a `key` string and `values` list to filter projects.
 *
 * For example to find all staging environment projects:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const staging = pulumi.output(digitalocean.getProjects({
 *     filters: [{
 *         key: "environment",
 *         values: ["Staging"],
 *     }],
 * }, { async: true }));
 * ```
 *
 * You can filter on multiple fields and sort the results as well:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const non_default_production = pulumi.output(digitalocean.getProjects({
 *     filters: [
 *         {
 *             key: "environment",
 *             values: ["Production"],
 *         },
 *         {
 *             key: "is_default",
 *             values: ["false"],
 *         },
 *     ],
 *     sorts: [{
 *         direction: "asc",
 *         key: "name",
 *     }],
 * }, { async: true }));
 * ```
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getProjects:getProjects", {
        "filters": args.filters,
        "sorts": args.sorts,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * Filter the results.
     * The `filter` block is documented below.
     */
    readonly filters?: inputs.GetProjectsFilter[];
    /**
     * Sort the results.
     * The `sort` block is documented below.
     */
    readonly sorts?: inputs.GetProjectsSort[];
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    readonly filters?: outputs.GetProjectsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of projects satisfying any `filter` and `sort` criteria. Each project has
     * the following attributes:
     */
    readonly projects: outputs.GetProjectsProject[];
    readonly sorts?: outputs.GetProjectsSort[];
}
