<script>
    import { onMount } from "svelte";
    import Redirect from "./RedirectRequest.svelte";
    import { redirects } from "../../stores";
    import { refreshList } from "../../stores";
    import { REDIRECT_TYPE } from "../../../src/constants";

    onMount(() => {
        console.debug(`Updating ${REDIRECT_TYPE} on mount`);
        refreshList(REDIRECT_TYPE);
    });
    
</script>

{#if $redirects.length > 0}
    {#each $redirects as redirect, i (crypto.randomUUID())}
    <div class="mb-4">
        <Redirect rule={redirect} />
    </div>
    {/each}
{/if}
