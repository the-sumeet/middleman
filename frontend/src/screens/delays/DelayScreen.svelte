<script>
    import { onMount } from "svelte";
    import { delays, refreshList } from "../../stores";
    import Delay from "./Delay.svelte";
    import { RULE_DELAY } from "../../../src/constants";

    onMount(() => {
        console.debug("Updating delays on mount");
        refreshList(RULE_DELAY);
    });
</script>

{#if $delays.length > 0}
    {#each $delays as delay, i (crypto.randomUUID())}
    <div class="mb-4">
        <Delay rule={delay} />
    </div>
    {/each}
{/if}
