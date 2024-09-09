<script>
    import RequestRedirectScreen from "../request_redirect/RequestRedirectScreen.svelte";
    import { currentCollection } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectRequest from "./RedirectRequest.svelte";
    import { GetCollection } from "../../../wailsjs/go/main/App";
    import Toobar from "./Toobar.svelte";
    
    let collectionId;
    let currColl;
    let list;


    const unSubCurrentCollection = currentCollection.subscribe((value) => {
        if (value.collectionId >= 0) {
            collectionId = value.collectionId;
            GetCollection(collectionId).then((result) => {
                if (result.error == "") {
                    currColl = result.collection;
                    console.log(currColl);
                } else {
                    console.error(`collection '${collectionId}' not found`);
                }
            });
        }
    });

    onDestroy(() => {
        unSubCurrentCollection();
    });
</script>

<Toobar bind:list={list} />

<div bind:this={list} class="mt-4 h-full overflow-y-scroll">
    <div class=" flex flex-col space-y-2">
        {#if currColl && currColl.records}
            {#each currColl.records as record, i (crypto.randomUUID())}
                {#if record.type == "redirect"}
                <RedirectRequest collectionId={collectionId} recordId={i} {record} />
                {/if}
            {/each}
        {/if}
    </div>
</div>
