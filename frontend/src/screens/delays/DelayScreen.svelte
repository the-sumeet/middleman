<script>
    import { onDestroy } from "svelte";
    import { delays } from "../../stores";
    import Delay from "./Delay.svelte";

    let delayList = [];

    // Subs
    const unSubDelays = delays.subscribe((value) => {
        console.log(value);
        delayList = value;
    });

    onDestroy(() => {
        unSubDelays();
    });
</script>

{#if delayList.length > 0}
    {#each delayList as delay, i (crypto.randomUUID())}
    <div class="mb-4">
        <Delay delayId={i} delay={delay} />
    </div>
    {/each}
{/if}
