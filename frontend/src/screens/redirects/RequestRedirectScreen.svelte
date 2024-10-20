<script>
    import { onDestroy } from "svelte";
    import Redirect from "./RedirectRequest.svelte";
    import { redirects } from "../../stores";
    
    let redirectsList = [];

    // Subs
    const unSubRedirects = redirects.subscribe((value) => {
        redirectsList = value;
        console.log(redirectsList);
    });

    onDestroy(() => {
        unSubRedirects();
    });
</script>

{#if redirectsList.length > 0}
    {#each redirectsList as redirect, i (crypto.randomUUID())}
    <div class="mb-4">
        <Redirect redirectId={i} {redirect} />
    </div>
    {/each}
{/if}
