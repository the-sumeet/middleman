import { writable } from "svelte/store";
import { RULE_MODIFY_REQUEST_BODY, RULE_MODIFY_RESPONSE_BODY, RULE_CANCEL, RULE_DELAY, RULE_MOD_REQUEST_HEADER, RULE_MOD_RESPONSE_HEADER, RULE_REDIRECT } from "./constants";
import { GetManyRules } from "../wailsjs/go/main/App";
import { HOME } from "./constants";
import { IsProxyRunning, IsWebServerRunning } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";
import { GetConfig } from "../wailsjs/go/main/App";

// Messages
export const successMessage = writable("");
export const errorMessage = writable("");
export const infoMessage = writable("");
export const warningMessage = writable("");

// Rules list
export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
export const cancels = writable([]);
export const delays = writable([]);
export const modifyRequestHeaders = writable([]);
export const modifyResponseHeaders = writable([]);
export const modifyRequestBody = writable([]);
export const modifyResponseBody = writable([]);

// Misc
export const proxyServerRunning = writable(false);
export const webServerRunning = writable(false);
export const currentPage = writable(HOME);
export const selectedRequest = writable(null);
export const config = writable(new main.Config());


export async function refreshList(type) {
    
    console.debug("Refreshing list", type);
    const result = await GetManyRules(type);
    if (result.error != "") {
        return result.error;
    }

    switch (type) {
        case RULE_DELAY:
            delays.set(result.rules);
            break;
        case RULE_MODIFY_REQUEST_BODY:
            modifyRequestBody.set(result.rules);
            break;
        case RULE_CANCEL:
            console.debug("Setting cancels", result.rules);
            cancels.set(result.rules);
            break;
        case RULE_MODIFY_RESPONSE_BODY:
            modifyResponseBody.set(result.rules);
            break;
        case RULE_MOD_REQUEST_HEADER:
            modifyRequestHeaders.set(result.rules);
            break;
        case RULE_MOD_RESPONSE_HEADER:
            modifyResponseHeaders.set(result.rules);
            break;
        case RULE_REDIRECT:
            redirects.set(result.rules);
            break;
    }

    return "";
}

async function refreshAllLists() {
    await refreshList(RULE_REDIRECT);
    await refreshList(RULE_CANCEL);
    await refreshList(RULE_DELAY);
    await refreshList(RULE_MOD_REQUEST_HEADER);
    await refreshList(RULE_MOD_RESPONSE_HEADER);
    await refreshList(RULE_MODIFY_REQUEST_BODY);
    await refreshList(RULE_MODIFY_RESPONSE_BODY);
}

export async function refreshConfig() {
    const result = await GetConfig();
    config.set(result);
}

IsProxyRunning().then((result) => {
    proxyServerRunning.set(result);
})

IsWebServerRunning().then((result) => {
    webServerRunning.set(result);
})

refreshConfig();

