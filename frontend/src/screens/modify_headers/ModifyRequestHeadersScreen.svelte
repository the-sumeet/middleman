<script>
    import { onMount } from "svelte";
    import { modifyRequestHeaders } from "../../stores";
    import ModifyHeader from "./ModifyHeader.svelte";
    import { refreshList } from "../../stores";
    import { RULE_MOD_REQUEST_HEADER } from "../../constants";

    onMount(() => {
        console.debug(`Updating ${RULE_MOD_REQUEST_HEADER} on mount`);
        refreshList(RULE_MOD_REQUEST_HEADER);
    });
</script>

{#if $modifyRequestHeaders.length > 0}
    {#each $modifyRequestHeaders as modifyHeader, i (crypto.randomUUID())}
        <div class="mb-4">
            <ModifyHeader
                headerType={RULE_MOD_REQUEST_HEADER}
                rule={modifyHeader}
            />
        </div>
    {/each}
{/if}
