<script>
    export let collectionId;
    export let recordId;
    export let record;

    import { onMount } from "svelte";
    import { main } from "../../../wailsjs/go/models";
    import { DeleteRecord, SaveRecord } from "../../../wailsjs/go/main/App";
    import { currentCollection } from "../../stores";

    let ifValue;
    let thenValue;
    let editing = false;
    let saving = false;

    function toggleEditing() {
        if (editing === true) {
            save();
        }
        editing = !editing;
    }

    function save() {
        saving = true;
        const newRecord = new main.Record({
            id: "",
            type: "redirect",
            if: ifValue,
            then: thenValue,
        });
        SaveRecord(collectionId, recordId, newRecord).then(() => {
            saving = false;
        });
    }

    function remove() {
        console.log(`deleting ${recordId}`);
        DeleteRecord(collectionId, recordId);
        // This is to refresh collections.
        currentCollection.set({collectionId: collectionId});
    }

    onMount(() => {
        ifValue = record.if;
        thenValue = record.then;
    });
</script>

<div class="w-full p-6 mx-auto rounded-md shadow-md bg-gray-800">
    <h2 class="text-lg font-semibold capitalize text-white">If</h2>

    <div class="mt-4 mb-4 pl-6">
        <textarea
            readonly={!editing}
            bind:value={ifValue}
            rows="3"
            placeholder="if url contains &#34;example.com&#34;"
            class="block w-full px-4 py-2 border rounded-md bg-gray-800 text-gray-300 border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 focus:border-blue-300 focus:outline-none focus:ring"
        ></textarea>
    </div>

    <h2 class="text-lg font-semibold capitalize text-white">
        Then Redirect To
    </h2>

    <div class="mt-4 mb-4 pl-6">
        <input
            readonly={!editing}
            bind:value={thenValue}
            type="text"
            placeholder="if url contains &#34;example.com&#34;"
            class="block w-full px-4 py-2 border rounded-md bg-gray-800 text-gray-300 border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 focus:border-blue-300 focus:outline-none focus:ring"
        />
    </div>

    <div class="flex justify-end gap-2 mt-6">
        <button
            on:click={toggleEditing}
            class="px-4 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600"
        >
            {#if editing}
                <i class="bi bi-floppy"></i>
            {:else}
                <i class="bi bi-pencil"></i>
            {/if}
        </button>
        {#if !editing}
        <button
            on:click={remove}
            class="px-4 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600"
        >
            <i class="bi bi-trash"></i>
        </button>
        {/if}
    </div>
</div>
