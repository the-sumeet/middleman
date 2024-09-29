

export const REDIRECT_TYPE = "redirect";

// Rules
export const RULE_REDIRECT = "redirect";
export const RULE_MOD_BODY = "redirect";
export const RULE_MOD_HEADER = "modifyHeader";
export const RULE_CANCEL = "cancel";
export const RULE_DELAY= "delay";

// Tabs
export const HOME = "home";
export const REQUESTS = "requests";
export const SERVER = "server";
export const SETTINGS = "settings";

// Rule info
export const RULE_INFO = {
    'redirect': "Redirects the request to the given URL. Status 307, and 'Location' header is added with value of given URL.",
    'cancel': "Request will be cancelled with status 418.",
    'delay': "Adds extra delay of given seconds on top of the original response time.",
};