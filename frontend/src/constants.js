

export const REDIRECT_TYPE = "redirect";

// Rules
export const RULE_REDIRECT = "redirect";
export const RULE_MOD_REQUEST_HEADER = "modifyRequestHeader";
export const RULE_MOD_RESPONSE_HEADER = "modifyResponseHeader";
export const RULE_CANCEL = "cancel";
export const RULE_DELAY= "delay";
export const RULE_MODIFY_REQUEST_BODY= "modifyRequestBody";
export const RULE_MODIFY_RESPONSE_BODY= "modifyResponseBody";

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
"border-transparent text-white cursor-base focus:outline-none hover:border-gray-400";

// Rule info
export const RULE_INFO = {
    'redirect': "Redirects the request to the given URL. Status 307, and 'Location' header is added with value of given URL.",
    'cancel': "Request will be cancelled with status 418.",
    'delay': "Adds extra delay of given seconds on top of the original response time.",
    'modifyRequestHeader': "Modifies the request header. You can add, remove or modify the header.",
    'modifyResponseHeader': "Modifies the Response header. You can add, remove or modify the header.",
    'modifyRequestBody': "Replace the request body",
    'modifyResponseBody': "Replace the response body"
};

// Alert types
export const ALERT_INFO = "info";
export const ALERT_SUCCESS = "success";
export const ALERT_ERROR = "error";
export const ALERT_WARNING = "warning";

export const ALERT_BG_COLOR = {
    'info': 'bg-blue-500',
    'success': 'bg-emerald-500',
    'error': 'bg-red-500',
    "warning": 'bg-yellow-400'
};