<script>
    import { onMount } from "svelte";
    import { modifyRequestBody } from "../../stores";
    import ModifyRequestBody from "./ModifyRequestBody.svelte";
    import { refreshList } from "../../stores";
    import { RULE_MODIFY_REQUEST_BODY } from "../../../src/constants";

    onMount(() => {
        console.debug(`Updating ${RULE_MODIFY_REQUEST_BODY} on mount`);
        refreshList(RULE_MODIFY_REQUEST_BODY);
    });
    

</script>

{#if $modifyRequestBody.length > 0}
    {#each $modifyRequestBody as modifyBody, i (crypto.randomUUID())}
    <div class="mb-4">
        <ModifyRequestBody rule={modifyBody} />
    </div>
    {/each}
{/if}
