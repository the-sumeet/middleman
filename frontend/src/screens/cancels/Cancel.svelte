<script>
    export let cancel;
    export let cancelId;

    import { main } from "../../../wailsjs/go/models";
    import { Save } from "../../../wailsjs/go/main/App";
    import { Remove } from "../../../wailsjs/go/main/App";
    import { GetMany } from "../../../wailsjs/go/main/App";
    import { cancels } from "../../../src/stores";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";
    import EntitySelect from "../../../src/widgets/EntitySelect.svelte";
    import { RULE_CANCEL } from "src/constants";
    import { remove } from "../../../src/utils";

    let changed = false;
    let entity;
    let op;
    let value;

    fromCancel();

    function fromCancel() {
        entity = cancel.entity;
        op = cancel.op;
        value = cancel.value;
    }

    function setChanged() {
        changed = true;
    }

    // Go functions
    function save() {
        const newCancel = new main.Cancel({
            enabled: cancel.enabled,
            entity: entity,
            op: op,
            value: value,
        });

        const input = new main.InValue({ cancel: newCancel });

        Save("cancel", cancelId, input).then(async () => {
            const result = await GetMany("cancel");
            cancels.set(result.cancels);
            changed = false;
        });
    }

    function cancelSave() {
        fromCancel();
        changed = false;
    }

    function enableDisable() {
        cancel.enabled = !cancel.enabled;
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
        remove={() => remove(RULE_CANCEL, cancelId)}
        {enableDisable}
        enabled={cancel.enabled}
    />
</div>
