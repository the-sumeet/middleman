<script>
    import { onMount } from "svelte";
    import { modifyResponseHeaders } from "../../stores";
    import ModifyHeader from "./ModifyHeader.svelte";
    import { refreshList } from "../../stores";
    import { RULE_MOD_RESPONSE_HEADER } from "../../constants";

    onMount(() => {
        console.debug(`Updating ${RULE_MOD_RESPONSE_HEADER} on mount`);
        refreshList(RULE_MOD_RESPONSE_HEADER);
    });
    
</script>

{#if $modifyResponseHeaders.length > 0}
    {#each $modifyResponseHeaders as modifyHeader, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyHeader headerType={RULE_MOD_RESPONSE_HEADER} rule={modifyHeader} />
    </div>
    {/each}
{/if}
