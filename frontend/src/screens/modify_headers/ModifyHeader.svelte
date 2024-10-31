<script>
    export let rule;

    import { main } from "../../../wailsjs/go/models";
    import { UpdateRule } from "../../../wailsjs/go/main/App";
    import { errorMessage } from "../../stores";
    import { GetManyRules } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { removeAndRefresh, updateRule } from "../../../src/utils";
    import {
        RULE_MOD_HEADER,
        tabSelectedStyle,
        tabUnselectedStyle,
    } from "../../../src/constants";

    let requestTab = true;
    let changed = false;
    let entity;
    let op;
    let value;

    fromModifyHeader();

    function fromModifyHeader() {
        entity = rule.entity;
        op = rule.op;
        value = rule.value;
    }

    function setChanged() {
        changed = true;
    }

    function addMod() {
        const newMod = new main.Header({
            isRequest: requestTab,
            action: "remove",
            name: "name",
            value: "value",
        });
        if (requestTab == true) {
            const currentRequestMods =
                rule.requestHeaderMods != undefined ||
                rule.requestHeaderMods != null
                    ? rule.requestHeaderMods
                    : [];
            rule.requestHeaderMods = [...currentRequestMods, newMod];
        } else {
            const currentResponseMods =
                rule.responseHeaderMods != undefined ||
                rule.responseHeaderMods != null
                    ? rule.responseHeaderMods
                    : [];
            rule.responseHeaderMods = [...currentResponseMods, newMod];
        }
        changed = true;
    }

    function removeMod(index) {
        if (requestTab) {
            rule.requestHeaderMods = rule.requestHeaderMods.filter((_, i) => i !== index);
        } else {
            rule.responseHeaderMods = rule.responseHeaderMods.filter((_, i) => i !== index);
        }
        
        changed = true;
    }

    function setAction(index, action) {        
        if (requestTab) {
            rule.requestHeaderMods[index].action = action;
        } else {
            rule.responseHeaderMods[index].action = action;
        }
        changed = true;
    }

    function setName(e, i) {
        const name = e.target.value;
        if (requestTab) {
            rule.requestHeaderMods[i].name = name;
        } else {
            rule.responseHeaderMods[i].name = name;
        }
        changed = true;
    }

    function setValue(e, i) {
        const value = e.target.value;
        if (requestTab) {
            rule.requestHeaderMods[i].value = value;
        } else {
            rule.responseHeaderMods[i].value = value;
        }
        changed = true;
    }

    function save() {
        const modifyHeaderRecord = new main.Rule({
            type: RULE_MOD_HEADER,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
            requestHeaderMods: rule.requestHeaderMods,
            responseHeaderMods: rule.responseHeaderMods,
        });

        updateRule(rule.id, modifyHeaderRecord).then(async (result) => {
            if (result.error === "") {
                console.debug("Updated Rule", result.rules[0]);
                rule = result.rules[0];
                changed = false;
            } else {
                errorMessage.set(result.error);
            }
        });
    }

    function cancelSave() {
        fromModifyHeader();
        changed = false;
    }

    function enableDisable() {
        rule.enabled = !rule.enabled;
        save();
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-4">
        <EntitySelect bind:entity {setChanged} />

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
            class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {requestTab ===
            true
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            <span class="mx-1 text-sm sm:text-base"> Request Headers </span>
        </button>

        <button
            on:click={() => (requestTab = false)}
            class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {requestTab ===
            false
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            <span class="mx-1 text-sm sm:text-base"> Response Headers </span>
        </button>
    </div>

    <div class="mt-4 border-2 rounded-md border-gray-700">
        {#if (requestTab && rule.requestHeaderMods) || (!requestTab && rule.responseHeaderMods)}
            {#each requestTab ? rule.requestHeaderMods : rule.responseHeaderMods as mod, i}
                <div class="flex space-x-2 p-4">
                    <select
                        on:change={() => changed = true}
                        bind:value={mod.action}
                        class="p-2 rounded text-whites"
                        name=""
                        id=""
                    >
                        <option class="text-white" value="add">Add</option>
                        <option class="text-white" value="remove">Remove</option
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
            {/each}
        {/if}
        <button
            on:click={addMod}
            class="m-4 w-max px-6 py-2 font-medium tracking-wide text-white border-2 border-dashed rounded border-gray-700"
        >
            Add Modification
        </button>
    </div>

    <!-- Bottom buttons -->
    <BottomButtons
        {changed}
        {save}
        {cancelSave}
        remove={() => removeAndRefresh(RULE_MOD_HEADER, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
