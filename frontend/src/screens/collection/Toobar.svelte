<script>
    export let list;
    import { REDIRECT_TYPE } from "../../constants";
    import { GetCollection, NewRecord } from "../../../wailsjs/go/main/App";
    import { currentCollection } from "../../stores";
    import { onDestroy } from "svelte";
    import { scrollToBottom } from "../../utils";

    let addDropdown = false;
    let currentCollectionId;
    let currentColl;

    function addRecord(type) {
        if (currentCollectionId >= 0) {
            NewRecord(currentCollectionId, type).then(() => {
                currentCollection.set({ collectionId: currentCollectionId });
            });
        }
        addDropdown = false;
        scrollToBottom(list);
    }

    function toggleAddDropdown() {
        addDropdown = !addDropdown;
    }

    // Go Functions

    // Subscriptions
    const unSubCurrentCollection = currentCollection.subscribe((value) => {
        currentCollectionId = value.collectionId;

        GetCollection(currentCollectionId).then((value) => {
            console.log(currentColl);
            currentColl = value.collection;
        });
    });

    onDestroy(() => {
        unSubCurrentCollection();
    });
</script>

<div class="flex bg-gray-800 rounded-md justify-between w-full">

    <!-- Title -->
    {#if currentColl}
    <h1 class="p-2 text-white">{currentColl.name}</h1>
    {/if}

    <!-- Buttons -->
     <div class="flex items-end">
        <!-- Add Button -->
        <div class="relative inline-block">
            <!-- Dropdown toggle button -->
            <button on:click={toggleAddDropdown}  class="relative z-10 block p-2 text-gray-700 bg-white border border-transparent rounded-md dark:text-white focus:border-blue-500 focus:ring-opacity-40 dark:focus:ring-opacity-40 focus:ring-blue-300 dark:focus:ring-blue-400 focus:ring dark:bg-gray-800 focus:outline-none">
                <i class="bi bi-plus-lg"></i>
            </button>
        
            <!-- Dropdown menu -->
            {#if addDropdown}
            <div 
                class="absolute right-0 z-20 w-48 py-2 mt-2 origin-top-right bg-white rounded-md shadow-xl dark:bg-gray-800"
            >
                <a href="#!" on:click={() => addRecord(REDIRECT_TYPE)} class="block px-4 py-3 text-sm text-gray-600 capitalize transition-colors duration-300 transform dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 dark:hover:text-white">Redirect</a>
     
            </div>
            {/if}
        </div>
     </div>
</div>