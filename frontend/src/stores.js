import { writable } from "svelte/store";
import { RULE_CANCEL, RULE_DELAY, RULE_MOD_HEADER, RULE_REDIRECT } from "./constants";
import { GetMany } from "../wailsjs/go/main/App";
import { HOME } from "./constants";


export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
export const cancels = writable([]);
export const delays = writable([]);
export const modifyHeaders = writable([]);
export const serverRunning = writable(false);
export const currentPage = writable(HOME);

GetMany(RULE_REDIRECT).then((res) => {
    redirects.set(res.redirects);
});
GetMany(RULE_CANCEL).then((res) => {
    cancels.set(res.cancels);
});
GetMany(RULE_DELAY).then((res) => {
    delays.set(res.delays);
});
GetMany(RULE_MOD_HEADER).then((res) => {
    modifyHeaders.set(res.modifyHeaders);
});