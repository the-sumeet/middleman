<script>
    export let modifyHeader;
    export let modifyHeaderId;

    import { main } from "../../../wailsjs/go/models";
    import { Save } from "../../../wailsjs/go/main/App";
    import { Remove } from "../../../wailsjs/go/main/App";
    import { modifyHeaders } from "../../stores";
    import { GetMany } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";

    let requestTab = true;
    const tabSelectedStyle =
        "border-blue-400 text-blue-300 whitespace-nowrap focus:outline-none";
    const tabUnselectedStyle =
        "border-transparent  text-white whitespace-nowrap cursor-base focus:outline-none hover:border-gray-400";
    let changed = false;
    let entity = modifyHeader.entity;
    let op = modifyHeader.op;
    let value = modifyHeader.value;

    function setChanged() {
        changed = true;
    }

    function addRecord() {
        modifyHeader.mods = [
            ...(modifyHeader.mods === null ? [] : modifyHeader.mods),
            new main.Header({
                isRequest: requestTab,
                action: "remove",
                name: "name",
                value: "value",
            }),
        ];
        changed = true;
    }

    function removeMod(index) {
        modifyHeader.mods = modifyHeader.mods.filter((_, i) => i !== index);
        changed = true;
    }

    function setAction(index, action) {
        modifyHeader.mods[index].action = action;
        changed = true;
    }

    function setName(e, i) {
        const name = e.target.value;
        modifyHeader.mods[i].name = name;
        changed = true;
    }

    function setValue(e, i) {
        const value = e.target.value;
        modifyHeader.mods[i].value = value;
        changed = true;
    }

    function save() {
        const modifyHeaderRecord = new main.ModifyHeader({
            enabled: modifyHeader.enabled,
            entity: entity,
            op: op,
            value: value,
            mods: [
                ...modifyHeader.mods.map((mod) => {
                    return new main.Header({
                        isRequest: mod.isRequest,
                        action: mod.action,
                        name: mod.name,
                        value: mod.value,
                    });
                }),
            ],
        });

        const input = new main.InValue({
            modifyHeader: modifyHeaderRecord,
        });

        Save("modifyHeader", modifyHeaderId, input).then(async () => {
            const result = await GetMany("modifyHeader");
            modifyHeaders.set(result.modifyHeaders);
            changed = false;
        });
    }

    function enableDisable() {
        modifyHeader.enabled = !modifyHeader.enabled;
        save();
    }

    function remove() {
        Remove("modifyHeader", modifyHeaderId).then(async () => {
            const result = await GetMany("modifyHeader");
            modifyHeaders.set(result.modifyHeaders);
            changed = false;
        });
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-4">
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

    <h1 class="mt-4 text-md text-white">Then Modify the Headers</h1>

    <div
        class="mt-4 flex overflow-x-auto overflow-y-hidden border-b whitespace-nowrap border-gray-700"
    >
        <button
            on:click={() => (requestTab = true)}
            class="inline-flex items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {requestTab ===
            true
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            <span class="mx-1 text-sm sm:text-base"> Request Headers </span>
        </button>

        <button
            on:click={() => (requestTab = false)}
            class="inline-flex items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {requestTab ===
            false
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            <span class="mx-1 text-sm sm:text-base"> Response Headers </span>
        </button>
    </div>

    <div class="mt-4 border-2 rounded-md border-gray-700">
        {#if modifyHeader.mods}
            {#each modifyHeader.mods as mod, i}
                {#if mod.isRequest === requestTab}
                    <div class="flex space-x-2 p-4">
                        <select
                            bind:value={mod.action}
                            class="p-2 rounded text-whites"
                            name=""
                            id=""
                        >
                            <option class="text-white" value="add">Add</option>
                            <option class="text-white" value="remove"
                                >Remove</option
                            >
                            <option class="text-white" value="override"
                                >Override</option
                            >
                        </select>

                        <input
                            on:input={(e) => setName(e, i)}
                            value={mod.name}
                            autocapitalize="off"
                            autocorrect="off"
                            type="text"
                            placeholder="example"
                            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
                        />

                        <input
                            on:input={(e) => setValue(e, i)}
                            value={mod.value}
                            autocapitalize="off"
                            autocorrect="off"
                            type="text"
                            placeholder="example"
                            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
                        />

                        <button
                            on:click={() => removeMod(i)}
                            class="w-max px-2 py-1 font-medium tracking-wide text-white hover:text-red-500"
                        >
                            <i class="bi bi-trash"></i>
                        </button>
                    </div>
                {/if}
            {/each}
        {/if}
        <button
            on:click={addRecord}
            class="m-4 w-max px-6 py-2 font-medium tracking-wide text-white border-2 border-dashed rounded border-gray-700"
        >
            Add Modification
        </button>
    </div>

    <!-- Bottom buttons -->
    <BottomButtons
        changed={changed}
        save={save}
        remove={remove}
        enableDisable={enableDisable}
        enabled={modifyHeader.enabled}
    />
   
</div>
