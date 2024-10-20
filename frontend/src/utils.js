


import { Remove, GetMany } from "../wailsjs/go/main/App";
import { delays, redirects, modifyHeaders, modifyRequestBody, modifyResponseBody, cancels } from "./stores";
import { RULE_DELAY, MODIFY_RESPONSE_BODY, RULE_CANCEL, RULE_INFO, RULE_MOD_HEADER, RULE_REDIRECT, MODIFY_REQUEST_BODY } from "./constants";

export const scrollToBottom = async (el) => {
  el.scroll({ top: el.scrollHeight, behavior: 'smooth' });
};


export const remove = (type, id) => {
  if (type === RULE_DELAY) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        delays.set(res.delays);
      });
    });
  } else if (type === RULE_REDIRECT) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        redirects.set(res.redirects);
      });
    });
  } else if (type === RULE_MOD_HEADER) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        modifyHeaders.set(res.modifyHeaders);
      });
    });
  } else if (type === MODIFY_RESPONSE_BODY) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        modifyResponseBody.set(res.modifyResponseBody);
      });
    });
    
  } else if (type === RULE_CANCEL) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        cancels.set(res.cancels);
      });
    });
  } else if (type === MODIFY_REQUEST_BODY) {
    Remove(type, id).then(() => {
      GetMany(type).then((res) => {
        modifyRequestBody.set(res.modifyRequestBody);
      });
    });
  }

}