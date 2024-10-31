<script>
    export let rule;

    import { main } from "../../../wailsjs/go/models";
    import {
        UpdateRule,
        GetOneRule,
        GetManyRules,
    } from "../../../wailsjs/go/main/App";
    import { delays, errorMessage } from "../../stores";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { removeAndRefresh, updateRule } from "../../../src/utils";
    import { RULE_DELAY } from "../../../src/constants";
    import { onMount } from "svelte";

    let changed = false;
    let error = "";
    let entity;
    let op;
    let value;
    let delaySec;

    function fromDelay() {
        entity = rule.entity;
        op = rule.op;
        value = rule.value;
        delaySec = rule.delaySec;
    }

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

    function save() {
        const newDelay = new main.Rule({
            type: RULE_DELAY,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
            delaySec: parseInt(delaySec),
        });

        updateRule(rule.id, newDelay).then(async (result) => {
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
        fromDelay();
        changed = false;
    }

    function enableDisable() {
        rule.enabled = !rule.enabled;
        save();
    }

    onMount(() => {
        fromDelay();
    });
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
        {changed}
        {save}
        {cancelSave}
        remove={() => removeAndRefresh(RULE_DELAY, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
