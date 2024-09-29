import { writable } from "svelte/store";
import { RULE_REDIRECT } from "./constants";
import { GetMany } from "../wailsjs/go/main/App";
import { HOME } from "./constants";


export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
export const cancels = writable([]);
export const delays = writable([]);
export const modifyHeaders = writable([]);
export const serverRunning = writable(false);
export const currentPage = writable(HOME);

GetMany("redirect").then((res) => {
    redirects.set(res.redirects);
});
GetMany("cancel").then((res) => {
    cancels.set(res.cancels);
});
GetMany("delay").then((res) => {
    delays.set(res.delays);
});
GetMany("modifyHeader").then((res) => {
    modifyHeaders.set(res.modifyHeaders);
});