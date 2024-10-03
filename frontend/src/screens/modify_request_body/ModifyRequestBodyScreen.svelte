<script>
    import { onDestroy } from "svelte";
    import { modifyRequestBody } from "../../stores";
    import ModifyRequestBody from "./ModifyRequestBody.svelte";

    let modifyBodyList = [];

    // Subs
    const unSubRequestBody = modifyRequestBody.subscribe((value) => {
        modifyBodyList = value;
    });

    onDestroy(() => {
        unSubRequestBody();
    });
</script>

{#if modifyBodyList.length > 0}
    {#each modifyBodyList as modifyBody, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyRequestBody modifyBodyId={i} {modifyBody} />
    </div>
    {/each}
{/if}
