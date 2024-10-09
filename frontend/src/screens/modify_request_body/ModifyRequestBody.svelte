<script>
    export let modifyBody;
    export let modifyBodyId;
    import { main } from "../../../wailsjs/go/models";
    import { Save } from "../../../wailsjs/go/main/App";
    import { Remove } from "../../../wailsjs/go/main/App";
    import { modifyRequestBody } from "../../stores";
    import { GetMany } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../widgets/BottomButtons.svelte";
    import { MODIFY_REQUEST_BODY } from "../../../src/constants";
    import * as ace from "brace";
    import { onMount } from "svelte";
    import 'brace/mode/html';
    import 'brace/mode/json';
    import 'brace/mode/text';
    import 'brace/mode/xml';
    import 'brace/mode/yaml';
    import 'brace/theme/dracula';
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";

    let editor;
    const editorId = `editor${modifyBodyId}`;
    let changed = false;

    // modifyBody properties
    let entity = modifyBody.entity;
    let op = modifyBody.op;
    let value = modifyBody.value;
    let body = modifyBody.body;
    
    function setChanged() {
        changed = true;
    }

    function save() {
        const modifyHeaderRecord = new main.ModifyRequestBody({
            enabled: modifyBody.enabled,
            entity: entity,
            op: op,
            value: value,
            body: body,
        });

        const input = new main.InValue({
            modifyRequestBody: modifyHeaderRecord,
        });

        Save(MODIFY_REQUEST_BODY, modifyBodyId, input).then(async () => {
            const result = await GetMany(MODIFY_REQUEST_BODY);
            modifyRequestBody.set(result.modifyRequestBody);
            changed = false;
        });
    }

    function enableDisable() {
        modifyBody.enabled = !modifyBody.enabled;
        save();
    }

    function remove() {
        Remove(MODIFY_REQUEST_BODY, modifyBodyId).then(async () => {
            const result = await GetMany(MODIFY_REQUEST_BODY);
            modifyRequestBody.set(result.modifyRequestBody);
            changed = false;
        });
    }
    onMount(() => {
        editor = ace.edit(editorId);
        editor.setTheme("ace/theme/dracula");
        editor.setValue(body, 1);
        // editor.setReadOnly(true);
        // editor.session.setMode("ace/mode/" + responseType);
        editor.setFontSize('16px');
        editor.getSession().on('change', function(delta) {
            body = editor.getValue();
            changed = true;
        });
        editor.setShowPrintMargin(false);
    });
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div class="flex items-center justify-center gap-2 p-4 rounded-md mt-4">
        <EntitySelect bind:entity={entity} {setChanged} />

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

    <h1 class="mt-4 text-md text-white">Then user the following body</h1>

    <div id={editorId} class="border border-gray-700 mt-4 w-full rounded" style="height: 256px;"></div>

    <!-- Bottom buttons -->
    <BottomButtons
        {changed}
        {save}
        {remove}
        {enableDisable}
        enabled={modifyBody.enabled}
    />
</div>
