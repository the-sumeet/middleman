<script>
    export let rule;
    import { onMount } from "svelte";
    import { main } from "../../../wailsjs/go/models";
    import { GetOneRule } from "../../../wailsjs/go/main/App";
    import { errorMessage, refreshList } from "../../../src/stores";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { RULE_CANCEL } from "../../../src/constants";
    import { removeAndRefresh, updateRule } from "../../../src/utils";

    let changed = false;
    let entity;
    let op;
    let value;

    fromCancel();

    function fromCancel() {
        entity = rule.entity;
        op = rule.op;
        value = rule.value;
    }

    function setChanged() {
        changed = true;
    }

    
    function save() {
        const newCancel = new main.Rule({
            type: RULE_CANCEL,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
        });

        updateRule(rule.id, newCancel).then(async (result) => {
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
        fromCancel();
        changed = false;
    }

    function enableDisable() {
        rule.enabled = !rule.enabled;
        save();
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-2">
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

    <h1 class="mt-2 text-md text-white">Then Cancel the Request</h1>

    <!-- Bottom buttons -->
    <BottomButtons
        {changed}
        {save}
        {cancelSave}
        remove={() => removeAndRefresh(RULE_CANCEL, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
