<script>
    import { onDestroy } from "svelte";
    import { cancels } from "../../../src/stores";
    import Cancel from "./Cancel.svelte";
    import { refreshList } from "../../../src/stores";
    import { RULE_CANCEL } from "../../../src/constants";
    import { onMount } from "svelte";

    onMount(() => {
        console.debug("Updating cancels on mount");
        refreshList(RULE_CANCEL);
    });

</script>

{#if $cancels.length > 0}
    {#each $cancels as cancel, i (crypto.randomUUID())}
        <div class="mb-4">
            <Cancel rule={cancel} />
        </div>
    {/each}
{/if}
