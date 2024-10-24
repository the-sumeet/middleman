

export const REDIRECT_TYPE = "redirect";

// Rules
export const RULE_REDIRECT = "redirect";
export const RULE_MOD_HEADER = "modifyHeader";
export const RULE_CANCEL = "cancel";
export const RULE_DELAY= "delay";
export const MODIFY_REQUEST_BODY= "modifyRequestBody";
export const MODIFY_RESPONSE_BODY= "modifyResponseBody";

// Tabs
export const HOME = "home";
export const REQUESTS = "requests";
export const SERVER = "server";
export const SETTINGS = "settings";
export const LOGS = "logs";

// Tabs CSS
export const tabSelectedStyle =
"border-blue-400 text-blue-300  focus:outline-none";
export const tabUnselectedStyle =
"border-transparent  text-white cursor-base focus:outline-none hover:border-gray-400";

// Rule info
export const RULE_INFO = {
    'redirect': "Redirects the request to the given URL. Status 307, and 'Location' header is added with value of given URL.",
    'cancel': "Request will be cancelled with status 418.",
    'delay': "Adds extra delay of given seconds on top of the original response time.",
    'modifyHeader': "Modifies the request header. You can add, remove or modify the header.",
    'modifyRequestBody': "Replace the request body",
    'modifyResponseBody': "Replace the response body"
};