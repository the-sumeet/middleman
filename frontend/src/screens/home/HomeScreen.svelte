<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import Rules from "../../screens/rules/Rules.svelte";
    import {cancels, redirects, currentRule, delays } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectList from "../../screens/redirects/redirect_list.svelte";
    import { RULE_CANCEL, RULE_INFO, RULE_REDIRECT, RULE_DELAY } from "../../constants";
    import { Add, GetMany } from "../../../wailsjs/go/main/App";
    import { main } from "../../../wailsjs/go/models";
    import { currentPage } from "../../stores";
    import CancelScreen from "../cancels/CancelScreen.svelte";
    import DelayScreen from "../delays/DelayScreen.svelte";
    
    let selectedRule;
    let currPage;

    $: ruleInfo = RULE_INFO[selectedRule];


    const unSubCurrentRule = currentRule.subscribe((value) => {
        selectedRule = value;
    });

    const unSubServerRunning = currentPage.subscribe((value) => {
        currPage = value;
    });

    function add() {
        if (selectedRule == RULE_REDIRECT) {
            Add('redirect', new main.Redirect()).then(() => {
                GetMany('redirect').then((res) => {
                    redirects.set(res.redirects);
                });
            });
        } else if (selectedRule == RULE_CANCEL) {
            Add('cancel', new main.Cancel()).then(() => {
                GetMany('cancel').then((res) => {
                    cancels.set(res.cancels);
                });
            });
        } else if (selectedRule == RULE_DELAY) {
            Add('delay', new main.Delay()).then(() => {
                GetMany('delay').then((res) => {
                    delays.set(res.delays);
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

    <!-- Info -->
     {#if ruleInfo}
    <div>
        <p class="text-white"><i class="bi bi-info-circle"></i> {ruleInfo}</p>
    </div>
    {/if}

    <!-- Selected RUle -->
    <div class="mt-4 overflow-y-scroll">
        {#if selectedRule == RULE_REDIRECT}
            <RedirectList />
        {:else if selectedRule == RULE_CANCEL}
            <CancelScreen />
        {:else if selectedRule == RULE_DELAY}
            <DelayScreen />
        {/if}
    </div>

    <!-- {#if currentColl !== undefined && currentColl >= 0}
    <CollectionPage />
  {/if} -->
</div>
