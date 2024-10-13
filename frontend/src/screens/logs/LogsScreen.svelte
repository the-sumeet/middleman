<script>
  import { GetLogs } from "../../../wailsjs/go/main/App";
  import { RULE_CANCEL } from "../../../src/constants";
  let logs = [];

  GetLogs().then((res) => {
    logs = res.httpRequests
  });
</script>

<section class="mt-4 h-screen container px-4 mx-auto">
  <div class="flex flex-col h-full">
    <h2 class="text-2xl font-medium text-gray-800 dark:text-white">Logs</h2>

    <!-- <p class="mt-1 text-sm text-gray-500 dark:text-gray-300">
      These companies have purchased in the last 12 months.
    </p> -->

    <div class="max-h-full overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
        <div
          class="overflow-hidden border border-gray-200 dark:border-gray-700 md:rounded-lg"
        >
          <table
            class="min-w-full divide-y divide-gray-200 dark:divide-gray-700"
          >
            <tbody
              class="bg-white divide-y divide-gray-200 dark:divide-gray-700 dark:bg-gray-900"
            >
              {#each logs as log}
                <tr>
                  <td class="px-4 py-4 text-sm font-medium whitespace-nowrap">
                    <div class="flex space-x-2 items-center">
                      {#if log.method && log.method != ""}
                          <div
                            class="inline px-3 py-1 text-sm font-normal rounded-full text-emerald-500 gap-x-2 bg-emerald-100/60 dark:bg-gray-800"
                          >
                            {log.method}
                          </div>
                        {/if}
                        {#if log.host && log.host != ""}
                          <div
                            class="inline px-3 py-1 text-sm font-normal rounded-full text-emerald-500 gap-x-2 bg-emerald-100/60 dark:bg-gray-800"
                          >
                            {log.host}
                          </div>
                        {/if}
                        {#if log.path && log.path != ""}
                          <div
                            class="inline px-3 py-1 text-sm font-normal rounded-full text-emerald-500 gap-x-2 bg-emerald-100/60 dark:bg-gray-800"
                          >
                            {log.path}
                          </div>
                        {/if}
                    </div>
                  </td>

                  <!-- <td class="px-4 py-4 text-sm whitespace-nowrap">
                  <button
                    class="px-1 py-1 text-gray-500 transition-colors duration-200 rounded-lg dark:text-gray-300 hover:bg-gray-100"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke-width="1.5"
                      stroke="currentColor"
                      class="w-6 h-6"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
                      />
                    </svg>
                  </button>
                </td> -->
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</section>
