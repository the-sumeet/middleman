<script>
    import { onMount } from "svelte";
    import { modifyHeaders } from "../../stores";
    import ModifyHeader from "./ModifyHeader.svelte";
    import { refreshList } from "../../stores";
    import { RULE_MOD_HEADER } from "../../../src/constants";

    onMount(() => {
        console.debug(`Updating ${RULE_MOD_HEADER} on mount`);
        refreshList(RULE_MOD_HEADER);
    });
    
</script>

{#if $modifyHeaders.length > 0}
    {#each $modifyHeaders as modifyHeader, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyHeader rule={modifyHeader} />
    </div>
    {/each}
{/if}
