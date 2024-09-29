<script>
    export let redirect;
    export let redirectId;

    import { main } from "../../../wailsjs/go/models";
    import { Save } from "../../../wailsjs/go/main/App";
    import { Remove } from "../../../wailsjs/go/main/App";
    import { redirects } from "../../../src/stores";
    import { GetMany } from "../../../wailsjs/go/main/App";
    import BottomButtons from "../../../src/widgets/BottomButtons.svelte";

    let changed = false;
    let entity = redirect.entity;
    let op = redirect.op;
    let value = redirect.value;
    let toType = redirect.toType;
    let toValue = redirect.toValue;

    function setChanged() {
        changed = true;
    }

    function save() {
        const redictRecord = new main.Redirect({
            enabled: redirect.enabled,
            entity: entity,
            op: op,
            value: value,
            toType: toType,
            toValue: toValue,
        });

        const input = new main.InValue({ redirect: redictRecord });

        Save('redirect', redirectId, input).then(async () => {
            const result = await GetMany("redirect");
            redirects.set(result.redirects);
            changed = false;
        });
    }

    function enableDisable() {
        redirect.enabled = !redirect.enabled;
        save();
    }

    function remove() {
        Remove('redirect', redirectId).then(async () => {
            const result = await GetMany("redirect");
            redirects.set(result.redirects);
            changed = false;
        });
    }
</script>

<div class="p-2 flex flex-col rounded-md bg-gray-800 border border-gray-700">
    <h1 class="text-md text-white">If</h1>

    <div
        class="flex items-center justify-center gap-2 p-4 rounded-md mt-4"
    >
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

    <h1 class="mt-4 text-md text-white">Then Redirect To</h1>

    <div
        class="flex items-center justify-center gap-2 p-4  rounded-md mt-4"
    >
        <input
            bind:value={toValue}
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
        changed={changed}
        save={save}
        remove={remove}
        enableDisable={enableDisable}
        enabled={redirect.enabled}
    />
</div>
