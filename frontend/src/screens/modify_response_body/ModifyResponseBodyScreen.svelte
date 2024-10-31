<script>
    import { onMount } from "svelte";
    import { modifyResponseBody } from "../../stores";
    import ModifyResponseBody from "./ModifyResponseBody.svelte";
    import { refreshList } from "../../stores";
    import { RULE_MODIFY_RESPONSE_BODY } from "../../../src/constants";

    onMount(() => {
        console.debug(`Updating ${RULE_MODIFY_RESPONSE_BODY} on mount`);
        refreshList(RULE_MODIFY_RESPONSE_BODY);
    });
</script>

{#if $modifyResponseBody.length > 0}
    {#each $modifyResponseBody as modifyBody, i (crypto.randomUUID())}
        <div class="mb-4">
            <ModifyResponseBody rule={modifyBody} />
        </div>
    {/each}
{/if}
