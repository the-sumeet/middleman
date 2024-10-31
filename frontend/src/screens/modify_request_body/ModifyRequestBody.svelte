<script>
    export let rule;

    import { main } from "../../../wailsjs/go/models";
    import { UpdateRule } from "../../../wailsjs/go/main/App";
    import { errorMessage } from "../../stores";
    import { GetManyRules } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../widgets/BottomButtons.svelte";
    import { RULE_MODIFY_REQUEST_BODY } from "../../../src/constants";
    import * as ace from "brace";
    import { onMount } from "svelte";
    import "brace/mode/html";
    import "brace/mode/json";
    import "brace/mode/text";
    import "brace/mode/xml";
    import "brace/mode/yaml";
    import "brace/theme/dracula";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { removeAndRefresh, updateRule } from "../../../src/utils";

    let editor;
    const editorId = `editor${rule.id}`;
    let changed = false;
    let entity;
    let op;
    let value;
    let body;

    function fromModifyRequestBody() {
        entity = rule.entity;
        op = rule.op;
        value = rule.value;
        body = rule.requestBody;
    }

    function setChanged() {
        changed = true;
    }

    function save() {
        const modifyRequestBodyRecord = new main.Rule({
            type: RULE_MODIFY_REQUEST_BODY,
            enabled: rule.enabled,
            entity: entity,
            op: op,
            value: value,
            requestBody: body,
        });

        updateRule(rule.id, modifyRequestBodyRecord).then(async (result) => {
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
        fromModifyRequestBody();
        changed = false;
    }

    function enableDisable() {
        rule.enabled = !rule.enabled;
        save();
    }
    onMount(() => {

        fromModifyRequestBody();

        editor = ace.edit(editorId);
        editor.setTheme("ace/theme/dracula");
        editor.setValue(body, 1);
        // editor.setReadOnly(true);
        // editor.session.setMode("ace/mode/" + responseType);
        editor.setFontSize("16px");
        editor.getSession().on("change", function (delta) {
            body = editor.getValue();
            changed = true;
        });
        editor.setShowPrintMargin(false);
    });
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

    <h1 class="mt-4 text-md text-white">Then use the following body</h1>

    <div
        id={editorId}
        class="border border-gray-700 mt-4 w-full rounded"
        style="height: 256px;"
    ></div>

    <!-- Bottom buttons -->
    <BottomButtons
        {changed}
        {save}
        {cancelSave}
        remove={() => removeAndRefresh(RULE_MODIFY_REQUEST_BODY, rule.id)}
        {enableDisable}
        enabled={rule.enabled}
    />
</div>
