<script>
  export let log;
  import { onMount } from "svelte";
  import * as ace from "brace";

  import "brace/mode/html";
  import "brace/mode/json";
  import "brace/mode/text";
  import "brace/mode/xml";
  import "brace/mode/yaml";
  import "brace/theme/dracula";
  import RuleIcons from "./RuleIcons.svelte";
  import { tabSelectedStyle, tabUnselectedStyle } from "../../../src/constants";

  let expanded = false;
  let tab = "headers"; // header, body
  let requestEditor;
  let responseEditor;
  const random = crypto.randomUUID();
  const requestEditorId = `requestBody${random}`;
  const responseEditorId = `responseBody${random}`;


  function getUrl() {
    return `${log.scheme}://${log.host}${log.path}`;
  }

  function getHeadersString(headerObj) {
    return Object.entries(headerObj)
      .map(([key, value]) => `${key}: ${value}`)
      .join("\n");
  }

  function toggleExpand() {
    if (expanded) {
      requestEditor.destroy();
      responseEditor.destroy();
      expanded = false;
    } else {
      requestEditor = ace.edit(requestEditorId);
      requestEditor.setTheme("ace/theme/dracula");
      if (tab == "headers") {
        requestEditor.setValue(getHeadersString(log.requestHeaders), 1);
      } else {
        requestEditor.setValue(log.requestBody, 1);
      }
      requestEditor.setValue(log.requestBody, 1);
      requestEditor.setReadOnly(true);
      requestEditor.setFontSize("13px");
      requestEditor.setShowPrintMargin(false);

      responseEditor = ace.edit(responseEditorId);
      responseEditor.setTheme("ace/theme/dracula");
      if (tab == "headers") {
        requestEditor.setValue(getHeadersString(log.responseHeaders), 1);
      } else {
        requestEditor.setValue(log.responseBody, 1);
      }
      responseEditor.setReadOnly(true);
      responseEditor.setFontSize("13px");
      responseEditor.setShowPrintMargin(false);

      expanded = true;
    }
  }
</script>

<div class="mt-4">
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div
    class="cursor-pointer flex flex-col bg-gray-800  p-2 rounded border-2 border-gray-700 hover:border-gray-600 items-start overflow-x-auto"
  >
    <div on:click={toggleExpand} class="w-full">
      <!-- Datetime -->
      <p class="text-white text-xs">{new Date(log.timestamp)}</p>

      <!-- Icons -->
      {#if log.cancelled || log.redirected || log.requestHeaderModified || log.responseHeaderModified || log.requestBodyModified || log.responseBodyModified || log.delayed != 0}
        <RuleIcons {log} />
      {/if}

      <!-- Method and URL -->
      <div class="mt-2 flex gap-2 items-center">
        <!-- Text -->
        <div class="flex gap-2 text-md text-blue-400">
          <div>{log.method}</div>
          <div class="break-all">{getUrl()}</div>
        </div>
      </div>
    </div>

    <!-- Content on expant -->
    <div
      class="{expanded == true
        ? 'flex'
        : 'hidden'} flex-col w-full mt-4 bg-gray-900 p-2 rounded border-2 border-gray-700"
    >
      <div
        class="flex overflow-x-auto overflow-y-hidden border-b border-gray-200 whitespace-nowrap dark:border-gray-700"
      >
        <button
          class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {tab ===
          'headers'
            ? tabSelectedStyle
            : tabUnselectedStyle}"
        >
          Headers
        </button>

        <!-- <button
          class="inline-flex whitespace-nowrap items-center h-10 px-2 py-2 -mb-px text-center bg-transparent border-b-2 sm:px-4 -px-1 {tab ===
          'body'
            ? tabSelectedStyle
            : tabUnselectedStyle}"
        >
          Body
        </button> -->
      </div>

      <!-- Request and response headers / body -->
      <div class="flex gap-2">
        <!-- Request headers / Body -->
        <div class="flex flex-col w-full mt-2">
          <p class="text-white text-sm">
            {tab === "headers" ? "Request Headers" : "Request Body"}
          </p>
          <div
            id={requestEditorId}
            class="border border-gray-700 mt-2 w-full rounded"
            style="height: 128px;"
          ></div>
        </div>
        <!-- Response headers / Body -->
        <div class="flex flex-col w-full mt-2">
          <p class="text-white text-sm">
            {tab === "headers" ? "Request Headers" : "Request Body"}
          </p>
          <div
            id={responseEditorId}
            class="border border-gray-700 mt-2 w-full rounded"
            style="height: 128px;"
          ></div>
        </div>
      </div>
    </div>
  </div>
</div>
