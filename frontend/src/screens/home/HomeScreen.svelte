<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import Rules from "../../screens/rules/Rules.svelte";
    import { currentRule, infoMessage, errorMessage } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectList from "../redirects/RequestRedirectScreen.svelte";
    import {
        RULE_CANCEL,
        RULE_MOD_REQUEST_HEADER,
        RULE_MOD_RESPONSE_HEADER,
        RULE_INFO,
        RULE_REDIRECT,
        RULE_DELAY,
        RULE_MODIFY_REQUEST_BODY,
        RULE_MODIFY_RESPONSE_BODY,
    } from "../../constants";
    import { AddRule } from "../../../wailsjs/go/main/App";
    import { main } from "../../../wailsjs/go/models";
    import { currentPage, refreshList } from "../../stores";
    import CancelScreen from "../cancels/CancelScreen.svelte";
    import DelayScreen from "../delays/DelayScreen.svelte";
    import ModifyRequestBodyScreen from "../modify_request_body/ModifyRequestBodyScreen.svelte";
    import ModifyResponseBodyScreen from "../modify_response_body/ModifyResponseBodyScreen.svelte";
    import ModifyRequestHeadersScreen from "../modify_headers/ModifyRequestHeadersScreen.svelte";
    import ModifyResponseHeadersScreen from "../modify_headers/ModifyResponseHeadersScreen.svelte";

    let selectedRule;
    let currPage;
    const initAttributes = { enabled: true, entity: "host", op: "contains" };
    
    $: ruleInfo = RULE_INFO[selectedRule];

    const unSubCurrentRule = currentRule.subscribe((value) => {
        selectedRule = value;
    });

    const unSubServerRunning = currentPage.subscribe((value) => {
        currPage = value;
    });

    async function add() {
        initAttributes.type = selectedRule;
        let inValue = new main.InValue({
            rule: new main.Rule(initAttributes),
        });
        const result = await AddRule(inValue);
        if (result.error != "") {
            errorMessage.set(result.error);
            return;
        }
        refreshList(selectedRule);
    }

    onDestroy(() => {
        unSubCurrentRule();
        unSubServerRunning();
    });
</script>

<div class="w-full flex">
    <Rules />

    <div class=" p-4 flex flex-col w-full h-screen">
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
                <p class="text-white">
                    <i class="bi bi-info-circle"></i>
                    {ruleInfo}
                </p>
            </div>
        {/if}

        <hr class="border mt-2 border-gray-800">

        <!-- Selected RUle -->
        <div class="mt-4 overflow-y-scroll">
            {#if selectedRule == RULE_REDIRECT}
                <RedirectList />
            {:else if selectedRule == RULE_CANCEL}
                <CancelScreen />
            {:else if selectedRule == RULE_DELAY}
                <DelayScreen />
            {:else if selectedRule == RULE_MOD_REQUEST_HEADER}
                <ModifyRequestHeadersScreen />
            {:else if selectedRule == RULE_MOD_RESPONSE_HEADER}
                <ModifyResponseHeadersScreen />
            {:else if selectedRule == RULE_MODIFY_REQUEST_BODY}
                <ModifyRequestBodyScreen />
            {:else if selectedRule == RULE_MODIFY_RESPONSE_BODY}
                <ModifyResponseBodyScreen />
            {/if}
        </div>

        <!-- {#if currentColl !== undefined && currentColl >= 0}
    <CollectionPage />
  {/if} -->
    </div>
</div>
