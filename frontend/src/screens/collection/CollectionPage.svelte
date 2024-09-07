<script>
    import RequestRedirectScreen from "../request_redirect/RequestRedirectScreen.svelte";
    import { currentCollection } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectRequest from "./RedirectRequest.svelte";
    import { GetCollection } from "../../../wailsjs/go/main/App";

    let collectionId;
    let currColl;

    const unSubCurrentCollection = currentCollection.subscribe((value) => {
        if (value >= 0) {
            collectionId = value;
            GetCollection(collectionId).then((result) => {
                if (result.error == "") {
                    currColl = result.collection;
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

<div class="flex">
    <button class="text-white">Add</button>
</div>

<div class="h-full overflow-y-scroll">
    <div class=" flex flex-col space-y-2">
        {#if currColl && currColl.records}
            {#each currColl.records as record, i}

                {#if record.type == "redirect"}
                <RedirectRequest collectionId={collectionId} recordId={i} {record} />
                {/if}
            {/each}
        {/if}
    </div>
</div>
