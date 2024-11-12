<script>
    export let rule;

    import { main } from "../../../wailsjs/go/models";
    import { UpdateRule } from "../../../wailsjs/go/main/App";
    import { errorMessage } from "../../stores";
    import { GetManyRules } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../widgets/BottomButtons.svelte";
    import EntitySelect from "../../widgets/EntitySelect.svelte";
    import { updateRule, removeAndRefresh } from "../../utils";
    import { RULE_REDIRECT } from "../../constants";
    import { onMount } from "svelte";

    let changed = false;
    let entity;
    let op;
    let value;
    let redirectTo;


    function fromRedirect() {
        entity = rule.entity;
        op = rule.op;
        value = rule.value;
        redirectTo = rule.redirectTo;
    }

    function setChanged() {
        changed = true;
    }

    function save() {
        const redictRecord = new main.Rule({
            type: RULE_REDIRECT,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
            redirectTo: redirectTo,
        });

        updateRule(rule.id, redictRecord).then(async (result) => {
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
        fromRedirect();
        changed = false;
    }

    function enableDisable() {
        rule.enabled = !rule.enabled;
        save();
    }

    onMount(() => {
        fromRedirect();
    });
</script>

<div class="rule">
    <h1 class="font-bold text-sm text-white">If</h1>

    <div class="flex items-center justify-center gap-2 mx-2 rounded-md mt-2">
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

    <h1 class="mt-2 font-bold text-sm text-white">Then Redirect To</h1>

    <div class="flex items-center justify-center gap-2 mx-2 rounded-md mt-2">
        <input
            bind:value={redirectTo}
            on:input={setChanged}
            autocapitalize="off"
            autocorrect="off"
            type="text"
            placeholder="https://example.com"
            class="w-full placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
        />
    </div>

    <!-- Bottom buttons -->
    <BottomButtons
        {changed}
        {save}
        {cancelSave}
        remove={() => removeAndRefresh(RULE_REDIRECT, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
