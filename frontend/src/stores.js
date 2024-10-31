import { writable } from "svelte/store";
import { RULE_MODIFY_REQUEST_BODY, RULE_MODIFY_RESPONSE_BODY, RULE_CANCEL, RULE_DELAY, RULE_MOD_HEADER, RULE_REDIRECT } from "./constants";
import { GetManyRules } from "../wailsjs/go/main/App";
import { HOME } from "./constants";

// Messages
export const successMessage = writable("");
export const errorMessage = writable("");
export const infoMessage = writable("");
export const warningMessage = writable("");

export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
export const cancels = writable([]);
export const delays = writable([]);
export const modifyHeaders = writable([]);
export const modifyRequestBody = writable([]);
export const modifyResponseBody = writable([]);
export const serverRunning = writable(false);
export const currentPage = writable(HOME);


// export async function refreshRedirects() {
//     const result = await GetManyRules(RULE_REDIRECT);
//     if (result.error != "") {
//         return result.error;
//     }
//     redirects.set(result.rules);
// }

// export async function refreshCancels() {
//     const result = await GetManyRules(RULE_CANCEL);
//     if (result.error != "") {
//         return result.error;
//     }
//     cancels.set(result.rules);
// }

// export async function refreshDelays() {
//     const result = await GetManyRules(RULE_DELAY);
//     if (result.error != "") {
//         return result.error;
//     }
//     delays.set(result.rules);
// }

// export async function refreshModifyHeaders() {
//     const result = await GetManyRules(RULE_MOD_HEADER);
//     if (result.error != "") {
//         return result.error;
//     }
//     modifyHeaders.set(result.rules);
// }

// export async function refreshModifyRequestBody() {
//     const result = await GetManyRules(RULE_MODIFY_REQUEST_BODY);
//     if (result.error != "") {
//         return result.error;
//     }
//     modifyRequestBody.set(result.rules);
// }

// export async function refreshModifyResponseBody() {
//     const result = await GetManyRules(RULE_MODIFY_RESPONSE_BODY);
//     if (result.error != "") {
//         return result.error;
//     }
//     modifyResponseBody.set(result.rules);
// }

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
        case RULE_MOD_HEADER:
            modifyHeaders.set(result.rules);
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
    await refreshList(RULE_MOD_HEADER);
    await refreshList(RULE_MODIFY_REQUEST_BODY);
    await refreshList(RULE_MODIFY_RESPONSE_BODY);
}