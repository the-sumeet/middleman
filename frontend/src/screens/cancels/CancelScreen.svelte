<script>
    import { onDestroy, onMount } from "svelte";
    import { cancels } from "../../../src/stores";
    import Cancel from "./Cancel.svelte";
    import { GetMany } from "../../../wailsjs/go/main/App";

    let cancelList = [];

    // Subs
    const unSubRedirects = cancels.subscribe((value) => {
        console.log(value);
        cancelList = value;
    });

    onDestroy(() => {
        unSubRedirects();
    });
</script>

{#if cancelList.length > 0}
    {#each cancelList as cancel, i (crypto.randomUUID())}
        <div class="mb-4">
            <Cancel cancelId={i} {cancel} />
        </div>
    {/each}
{/if}
