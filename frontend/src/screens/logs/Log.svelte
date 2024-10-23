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
  // Randon id

  let expanded = false;
  let requestEditor;
  let responseEditor;
  const random = crypto.randomUUID();
  const requestEditorId = `requestBody${random}`;
  const responseEditorId = `responseBody${random}`;

  console.log(log);

  function getUrl() {
    return `${log.scheme}://${log.host}${log.path}`;
  }

  function toggleExpand() {
    if (expanded) {
      requestEditor.destroy();
      responseEditor.destroy();
      expanded = false;
    } else {
      requestEditor = ace.edit(requestEditorId);
      requestEditor.setTheme("ace/theme/dracula");
      requestEditor.setValue(log.requestBody, 1);
      requestEditor.setReadOnly(true);
      requestEditor.setFontSize("13px");
      requestEditor.setShowPrintMargin(false);

      responseEditor = ace.edit(responseEditorId);
      responseEditor.setTheme("ace/theme/dracula");
      responseEditor.setValue(log.responseBody, 1);
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
    class="cursor-pointer flex flex-col bg-gray-800 hover:border-blue-600 p-2 rounded border-2 border-gray-700 items-start overflow-x-auto"
  >
    <!-- Datetime -->
    <p class="text-white text-xs">{new Date(log.timestamp)}</p>

    {#if log.cancelled || log.redirected || log.requestHeaderModified || log.responseHeaderModified || log.requestBodyModified || log.responseBodyModified || log.delayed != 0}
      <RuleIcons {log} />
    {/if}

    <div class="mt-2 flex gap-2 items-center">
      <!-- Text -->
      <div class="flex gap-2 text-md text-emerald-500">
        <div>{log.method}</div>
        <div class="break-all">{getUrl()}</div>
      </div>
    </div>

    <!-- Content on expant -->
      <div class="{expanded == true ? 'flex' : 'hidden'} flex-col w-full">
        <!-- Request and response body -->
        <div class="flex gap-2">
          <!-- Request Body -->
          <div class="flex flex-col w-full mt-4">
            <p class="text-white">Request Body</p>
            <div
              id={requestEditorId}
              class="border border-gray-700 mt-2 w-full rounded"
              style="height: 128px;"
            ></div>
          </div>
          <!-- Response Body -->
          <div class="flex flex-col w-full mt-4">
            <p class="text-white">Response Body</p>
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
