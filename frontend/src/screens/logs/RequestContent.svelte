<script>
    
    import { onDestroy, onMount } from "svelte";
    import * as ace from "brace";
    import { selectedRequest } from "../../../src/stores";
    import "brace/mode/html";
    import "brace/mode/json";
    import "brace/mode/text";
    import "brace/mode/xml";
    import "brace/mode/yaml";
    import "brace/theme/dracula";
    import 'brace/ext/searchbox'; // This is to show the search box in the editor
    import { getStatusColor, getMethodColor, formatDate } from "../../../src/utils";
    import Badge from "../../../src/widgets/Badge.svelte";
    import {
        tabSelectedStyle,
        tabUnselectedStyle,
    } from "../../../src/constants";

    let expanded = false;
    let tab = "headers"; // header, body
    let requestEditor;
    let responseEditor;
    let currentRequest;
    let status;

    $: completeUrl = `${currentRequest.scheme}://${currentRequest.host}${currentRequest.path}`;
    $: if (currentRequest && requestEditor && responseEditor) {
        if (tab == "headers") {
            requestEditor.setValue(getHeadersString(currentRequest.requestHeaders), 1);
            responseEditor.setValue(getHeadersString(currentRequest.responseHeaders), 1);
        } else {
            requestEditor.setValue(currentRequest.requestBody, 1);
            responseEditor.setValue(currentRequest.responseBody, 1);
        }

        

        requestEditor.selection.moveTo(0, 0);
        responseEditor.selection.moveTo(0, 0);
    }

    const unSubSelectedRequest = selectedRequest.subscribe((value) => {
        console.debug("Selected request", value);
        currentRequest = value;
    });

    $: status = currentRequest.status;
    $: statucColor = getStatusColor(status);
    $: console.log(statucColor)
    function getHeadersString(headerObj) {
        return Object.entries(headerObj)
            .map(([key, value]) => `${key}: ${value}`)
            .join("\n");
    }

    onMount(() => {
        requestEditor = ace.edit("requestEditor");
        requestEditor.setTheme("ace/theme/dracula");
        // requestEditor.setValue(currentRequest.requestBody, 1);
        requestEditor.setReadOnly(true);
        requestEditor.setFontSize("13px");
        requestEditor.setShowPrintMargin(false);
        requestEditor.$blockScrolling = Infinity;
        
        responseEditor = ace.edit("responseEditor");
        responseEditor.setTheme("ace/theme/dracula");
        responseEditor.setReadOnly(true);
        responseEditor.setFontSize("13px");
        responseEditor.setShowPrintMargin(false);
        responseEditor.$blockScrolling = Infinity;

        expanded = true;
    });
    
    function close() {
        selectedRequest.set(null);
    }

    onDestroy(() => {
        unSubSelectedRequest();
    });
</script>

{#if currentRequest}
<div
    class="flex flex-col w-full mt-4 bg-gray-900 p-2 rounded border border-gray-700"
>   
    <div class="flex w-full justify-between">
        <!-- Status and method -->
        <div class="flex gap-2">
            <Badge text={"GET"} color={getMethodColor("GET")} />
            <Badge text={status} color={statucColor} />
            <p class="text-white">{formatDate(currentRequest.timestamp)}</p>

        </div>

        <!-- Action buttons -->
        <button on:click={close}>
            <i class="text-white text-sm hover:bg-gray-700 p-1 rounded bi bi-x-lg"></i>
        </button>

    </div>
    <!-- <input type="overflow-x-scroll text" readonly value={getUrl()+getUrl()+getUrl()+getUrl()+getUrl()} /> -->
    <textarea class="mt-2 text-sm bg-gray-800 rounded text-white p-2" rows="2" name="" id="">{completeUrl}</textarea>
    
    <!-- Tabs -->
    <div
        class="flex overflow-x-auto overflow-y-hidden border-b whitespace-nowrap border-gray-700"
    >
        <button
            on:click={() => (tab = "headers")}
            class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {tab ===
            'headers'
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            Headers
        </button>

        <button
            on:click={() => (tab = "body")}
            class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {tab ===
            'body'
                ? tabSelectedStyle
                : tabUnselectedStyle}"
        >
            Body
        </button>
    </div>

    <!-- Request and response headers / body -->
    <div class="flex gap-2">
        <!-- Request headers / Body -->
        <div class="flex flex-col w-full mt-2">
            <p class="text-white text-sm">
                {tab === "headers" ? "Request Headers" : "Request Body"}
            </p>
            <div
                id="requestEditor"
                class="border border-gray-700 mt-2 w-full rounded"
                style="height: 128px;"
            ></div>
        </div>
        <!-- Response headers / Body -->
        <div class="flex flex-col w-full mt-2">
            <p class="text-white text-sm">
                {tab === "headers" ? "Response Headers" : "Request Body"}
            </p>
            <div
                id="responseEditor"
                class="border border-gray-700 mt-2 w-full rounded"
                style="height: 128px;"
            ></div>
        </div>
    </div>
</div>
{/if}