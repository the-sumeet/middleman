<script>
    export let cancel;
    export let cancelId;

    import { main } from "../../../wailsjs/go/models";
    import { SaveCancel } from "../../../wailsjs/go/main/App";
    import { RemoveCancel } from "../../../wailsjs/go/main/App";
    import { GetMany } from "../../../wailsjs/go/main/App";
    import { cancels } from "../../../src/stores";

    let changed = false;
    let entity = cancel.entity;
    let op = cancel.op;
    let value = cancel.value;

    function setChanged() {
        changed = true;
    }

    // Go functions
    function save() {
        const cancelRecord = new main.Cancel({
            entity: entity,
            op: op,
            value: value,
        });

        SaveCancel(cancelId, cancelRecord).then(() => {
            changed = false;
        });
    }

    function remove() {
        RemoveCancel(cancelId).then(() => {
            GetMany('cancel').then((res) => {
                cancels.set(res.cancels);
            });
        });
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div
        class="flex items-center justify-center gap-2 p-4 rounded-md mt-2"
    >
        <!-- Entity -->
        <select
            bind:value={entity}
            on:change={setChanged}
            class="p-2 rounded text-whites"
            name=""
            id=""
        >
            <option class="text-white" value="url">URL</option>
            <option class="text-white" value="host">HOST</option>
            <option value="header">Header</option>
        </select>

        <!-- Op -->
        <select
            on:change={setChanged}
            bind:value={op}
            class="p-2 rounded text-whites"
            name=""
            id=""
        >
            <option class="text-white" value="eq">Equal</option>
            <option class="text-white" value="contains">Contains</option>
        </select>

        <input
            bind:value
            on:input={setChanged}
            autocapitalize="off"
            autocorrect="off"
            type="text"
            placeholder="example"
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

    <h1 class="mt-2 text-md text-white">Then Cancel the Request</h1>


    <!-- Bottom buttons -->
    <div class="flex mt-2 justify-end">
        {#if changed}
            <button
                on:click={save}
                class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
            >
                <i class="bi bi-floppy"></i>
            </button>
        {:else}
            <button
                on:click={remove}
                class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-red-600 rounded-lg hover:bg-red-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
            >
                <i class="bi bi-trash"></i>
            </button>
        {/if}
    </div>
</div>
