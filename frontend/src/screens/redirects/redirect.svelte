<script>
    export let redirect;
    export let redirectId;

    import { main } from "../../../wailsjs/go/models";
    import { Save } from "../../../wailsjs/go/main/App";
    import { Remove } from "../../../wailsjs/go/main/App";
    import { redirects } from "../../../src/stores";
    import { GetMany } from "../../../wailsjs/go/main/App";

    let changed = false;
    let entity = redirect.entity;
    let op = redirect.op;
    let value = redirect.value;
    let toType = redirect.toType;
    let toValue = redirect.toValue;

    function setChanged() {
        changed = true;
    }

    // Go functions
    function save() {
        const redictRecord = new main.Redirect({
            entity: entity,
            op: op,
            value: value,
            toType: toType,
            toValue: toValue,
        });

        Save('redirect', redirectId, redictRecord).then(() => {
            changed = false;
        });
    }

    function remove() {
        Remove('redirect', redirectId).then(() => {
            GetMany('redirect').then((res) => {
                console.log(res.redirects);
                redirects.set(res.redirects);
            });
        });
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div
        class="flex items-center justify-center gap-2 p-4 rounded-md mt-4"
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

    <h1 class="mt-4 text-md text-white">Then Redirect To</h1>

    <div
        class="flex items-center justify-center gap-2 p-4  rounded-md mt-4"
    >
        <input
            bind:value={toValue}
            on:input={setChanged}
            autocapitalize="off"
            autocorrect="off"
            type="text"
            placeholder="https://example.com"
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

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
