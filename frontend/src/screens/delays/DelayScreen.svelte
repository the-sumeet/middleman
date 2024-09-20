<script>
    import { onDestroy } from "svelte";
    import { cancels } from "../../stores";
    import Cancel from "./Delay.svelte";

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
        <Cancel cancelId={i} cancel={cancel} />
    </div>
    {/each}
{/if}
