<script>
    import { onDestroy } from "svelte";
    import { modifyResponseBody } from "../../stores";
    import ModifyResponseBody from "./ModifyResponseBody.svelte";
    let modifyBodyList = [];

    // Subs
    const unSubResponseBody = modifyResponseBody.subscribe((value) => {
        modifyBodyList = value;
    });

    onDestroy(() => {
        unSubResponseBody();
    });
</script>

{#if modifyBodyList.length > 0}
    {#each modifyBodyList as modifyBody, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyResponseBody modifyBodyId={i} {modifyBody} />
    </div>
    {/each}
{/if}
