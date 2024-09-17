<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import Rules from "../../screens/rules/Rules.svelte";
    import { cancels, currentRule } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectList from "../../screens/redirects/redirect_list.svelte";
    import { RULE_CANCEL, RULE_REDIRECT } from "../../constants";
    import { AddCancel, AddRedirect, GetCancels, GetRedirects } from "../../../wailsjs/go/main/App";
    import { main } from "../../../wailsjs/go/models";
    import { redirects } from "../../stores";
    import { currentPage } from "../../stores";
    import CancelScreen from "../cancels/CancelScreen.svelte";

    let selectedRule;
    let currPage;

    const unSubCurrentRule = currentRule.subscribe((value) => {
        selectedRule = value;
    });

    const unSubServerRunning = currentPage.subscribe((value) => {
        currPage = value;
    });

    function add() {
        if (selectedRule == RULE_REDIRECT) {
            AddRedirect(new main.Redirect()).then(() => {
                GetRedirects().then((res) => {
                    redirects.set(res.redirects);
                });
            });
        } else if (selectedRule == RULE_CANCEL) {
            AddCancel(new main.Cancel()).then(() => {
                GetCancels().then((res) => {
                    cancels.set(res.cancels);
                });
            });
        }
    }

    onDestroy(() => {
        unSubCurrentRule();
        unSubServerRunning();
    });
</script>

<Rules />

<div class="p-4 flex flex-col w-full h-screen">
    <!-- Toolbar -->
    <div class="flex justify-end rounded-md">
        <button
            on:click={add}
            class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >
            <i class="bi bi-plus-lg"></i>
        </button>
    </div>

    <!-- Selected RUle -->
    <div class="mt-4 overflow-y-scroll">
        {#if selectedRule == RULE_REDIRECT}
            <RedirectList />
        {:else if selectedRule == RULE_CANCEL}
            <CancelScreen />
        {/if}
    </div>

    <!-- {#if currentColl !== undefined && currentColl >= 0}
    <CollectionPage />
  {/if} -->
</div>
