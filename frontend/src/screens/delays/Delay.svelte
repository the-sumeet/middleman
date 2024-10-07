<script>
    export let delay;
    export let delayId;

    import { main } from "../../../wailsjs/go/models";
    import { Save, Remove, GetMany } from "../../../wailsjs/go/main/App";
    import { delays } from "../../stores";
    import { onDestroy } from "svelte";
    import { parse } from "svelte/compiler";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    let changed = false;
    let entity = delay.entity;
    let op = delay.op;
    let value = delay.value;
    let delaySec = delay.delaySec;
    let error = "";

    function setChanged() {
        changed = true;
    }

    function validate() {
        error = "";

        // Check if delaySec is number
        if (isNaN(delaySec)) {
            error = "Delay seconds must be a number";
            return false;
        }
        
    }
    // Go functions
    function save() {
        
        if (validate() === false) {
            return;
        }

        const newDelay = new main.Delay({
            enabled: delay.enabled,
            entity: entity,
            op: op,
            value: value,
            delaySec: parseInt(delaySec),
        });

        const input = new main.InValue({ delay: newDelay });

        Save("delay", delayId, input).then(async () => {
            const result = await GetMany("delay");
            delays.set(result.delays);
            changed = false;
        });
    }

    function enableDisable() {
        delay.enabled = !delay.enabled;
        save();
    }

    function remove() {
        Remove("delay", delayId).then(() => {
            GetMany("delay").then((res) => {
                delays.set(res.delays);
            });
        });
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-2">
\        <EntitySelect bind:entity={entity} {setChanged} />

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
            bind:value={value}
            on:input={setChanged}
            autocapitalize="off"
            autocorrect="off"
            type="text"
            placeholder="example"
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

    <h1 class="mt-2 text-md text-white">Then Delay the Request by Seconds:</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-2">
        <input
            bind:value={delaySec}
            on:input={setChanged}
            autocapitalize="off"
            autocorrect="off"
            type="text"
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

    {#if error != ""}
        <p class="px-4 text-red-500">{error}</p>
    {/if}

    <!-- Bottom buttons -->
     <BottomButtons 
        changed={changed} 
        save={save} 
        remove={remove} 
        enableDisable={enableDisable} 
        enabled={delay.enabled}
    />
</div>
