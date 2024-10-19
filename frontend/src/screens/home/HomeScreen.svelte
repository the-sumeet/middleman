<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import Rules from "../../screens/rules/Rules.svelte";
    import {
        cancels,
        redirects,
        currentRule,
        delays,
        modifyHeaders,
        modifyRequestBody,
        modifyResponseBody,
    } from "../../stores";
    import { onDestroy } from "svelte";
    import RedirectList from "../../screens/redirects/redirect_list.svelte";
    import {
        RULE_CANCEL,
        RULE_MOD_HEADER,
        RULE_INFO,
        RULE_REDIRECT,
        RULE_DELAY,
        MODIFY_REQUEST_BODY,
        MODIFY_RESPONSE_BODY,
    } from "../../constants";
    import { Add, GetMany } from "../../../wailsjs/go/main/App";
    import { main } from "../../../wailsjs/go/models";
    import { currentPage } from "../../stores";
    import CancelScreen from "../cancels/CancelScreen.svelte";
    import DelayScreen from "../delays/DelayScreen.svelte";
    import ModifyHeadersScreen from "../modify_headers/ModifyHeadersScreen.svelte";
    import ModifyRequestBodyScreen from "../modify_request_body/ModifyRequestBodyScreen.svelte";
    import ModifyResponseBodyScreen from "../modify_response_body/ModifyResponseBodyScreen.svelte";

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

    function add() {
        if (selectedRule == RULE_REDIRECT) {
            Add(
                RULE_REDIRECT,
                new main.InValue({
                    redirect: new main.Redirect(initAttributes),
                }),
            ).then(() => {
                GetMany(RULE_REDIRECT).then((res) => {
                    redirects.set(res.redirects);
                });
            });
        } else if (selectedRule == RULE_CANCEL) {
            Add(
                RULE_CANCEL,
                new main.InValue({ cancel: new main.Cancel(initAttributes) }),
            ).then(() => {
                GetMany(RULE_CANCEL).then((res) => {
                    cancels.set(res.cancels);
                });
            });
        } else if (selectedRule == RULE_DELAY) {
            Add(
                RULE_DELAY,
                new main.InValue({ delay: new main.Delay(initAttributes) }),
            ).then(() => {
                GetMany(RULE_DELAY).then((res) => {
                    delays.set(res.delays);
                });
            });
        } else if (selectedRule == RULE_MOD_HEADER) {
            Add(
                RULE_MOD_HEADER,
                new main.InValue({
                    modifyHeader: new main.ModifyHeader(initAttributes),
                }),
            ).then(() => {
                GetMany(RULE_MOD_HEADER).then((res) => {
                    modifyHeaders.set(res.modifyHeaders);
                });
            });
        } else if (selectedRule == MODIFY_REQUEST_BODY) {
            Add(
                MODIFY_REQUEST_BODY,
                new main.InValue({
                    modifyRequestBody: new main.ModifyRequestBody(
                        initAttributes,
                    ),
                }),
            ).then(() => {
                GetMany(MODIFY_REQUEST_BODY).then((res) => {
                    modifyRequestBody.set(res.modifyRequestBody);
                });
            });
        } else if (selectedRule == MODIFY_RESPONSE_BODY) {
            Add(
                MODIFY_RESPONSE_BODY,
                new main.InValue({
                    modifyResponseBody: new main.ModifyResponseBody(
                        initAttributes,
                    ),
                }),
            ).then(() => {
                GetMany(MODIFY_RESPONSE_BODY).then((res) => {
                    modifyResponseBody.set(res.modifyResponseBody);
                });
            });
        }
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

        <!-- Selected RUle -->
        <div class="mt-4 overflow-y-scroll">
            {#if selectedRule == RULE_REDIRECT}
                <RedirectList />
            {:else if selectedRule == RULE_CANCEL}
                <CancelScreen />
            {:else if selectedRule == RULE_DELAY}
                <DelayScreen />
            {:else if selectedRule == RULE_MOD_HEADER}
                <ModifyHeadersScreen />
            {:else if selectedRule == MODIFY_REQUEST_BODY}
                <ModifyRequestBodyScreen />
            {:else if selectedRule == MODIFY_RESPONSE_BODY}
                <ModifyResponseBodyScreen />
            {/if}
        </div>

        <!-- {#if currentColl !== undefined && currentColl >= 0}
    <CollectionPage />
  {/if} -->
    </div>
</div>
