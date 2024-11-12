<script>
    export let rule;
    export let headerType;

    import { main } from "../../../wailsjs/go/models";
    import { errorMessage } from "../../stores";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { removeAndRefresh, updateRule } from "../../../src/utils";
    import { RULE_MOD_REQUEST_HEADER } from "../../../src/constants";
    import RuleTile from "../rules/RuleTile.svelte";

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
            action: "add",
            name: "Foo",
            value: "Bar",
        });
        const currentRequestMods =
            rule.headerMods != undefined || rule.headerMods != null
                ? rule.headerMods
                : [];
        rule.headerMods = [...currentRequestMods, newMod];

        changed = true;
    }

    function removeMod(index) {
        rule.headerMods = rule.headerMods.filter((_, i) => i !== index);

        changed = true;
    }

    // function setAction(index, action) {
    //     if (headerType == RULE_MOD_REQUEST_HEADER) {
    //         rule.requestHeaderMods[index].action = action;
    //     } else {
    //         rule.responseHeaderMods[index].action = action;
    //     }
    //     changed = true;
    // }

    function setName(e, i) {
        rule.headerMods[i].name = e.target.value;
        changed = true;
    }

    function setValue(e, i) {
        rule.headerMods[i].value = e.target.value;
        changed = true;
    }

    function save() {
        const modifyHeaderRecord = new main.Rule({
            type: headerType,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
            headerMods: rule.headerMods,
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

<div class="rule">
    <h1 class="font-bold text-sm text-white">If</h1>

    <div class="mx-2 flex items-center justify-center gap-2 rounded-md mt-2">
        <EntitySelect bind:entity {setChanged} />

        <!-- Op -->
        <select
            on:change={setChanged}
            bind:value={op}
            class=""
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
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

    <h1 class="mt-2 font-bold text-sm text-white">Then Modify the Headers</h1>

    <div class="mx-2 mt-2 border-2 rounded-md border-gray-700">
        {#if rule.headerMods}
            {#each rule.headerMods as mod, i}
                <div class="flex space-x-2 p-4">
                    <select
                        on:change={() => (changed = true)}
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
        remove={() => removeAndRefresh(headerType, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
