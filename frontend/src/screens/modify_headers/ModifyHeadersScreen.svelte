<script>
    import { onDestroy } from "svelte";
    import { modifyHeaders } from "../../stores";
    import ModifyHeader from "./ModifyHeader.svelte";
    let modifyHeadersList = [];

    // Subs
    const unSubModifyHeaders = modifyHeaders.subscribe((value) => {
        modifyHeadersList = value;
    });

    onDestroy(() => {
        unSubModifyHeaders();
    });
</script>

{#if modifyHeadersList.length > 0}
    {#each modifyHeadersList as modifyHeader, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyHeader modifyHeaderId={i} {modifyHeader} />
    </div>
    {/each}
{/if}
