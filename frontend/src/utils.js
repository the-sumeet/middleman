


import { RemoveRule, GetManyRules, UpdateRule } from "../wailsjs/go/main/App";
import { errorMessage, modifyResponseBody, refreshList } from "./stores";
import { RULE_DELAY, RULE_MODIFY_REQUEST_BODY, RULE_CANCEL, RULE_MOD_HEADER, RULE_REDIRECT, RULE_MODIFY_RESPONSE_BODY } from "./constants";
import { main } from "../wailsjs/go/models";

export const scrollToBottom = async (el) => {
  el.scroll({ top: el.scrollHeight, behavior: 'smooth' });
};

export const isAtBottom = async (el) => {
  return el.scrollHeight - el.scrollTop === el.clientHeight;
}

export const scrollToBottomIfAtBottom = async (el) => {
  if (await isAtBottom(el)) {
    scrollToBottom(el);
  }
}


export const remove = async (id) => {
  const result = await RemoveRule(id);
  return result;
}

export const removeAndRefresh = async (type, id) => {
  const result = await remove(id);
  if (result.error != "") {
    errorMessage.set(result.error);
    return;
  }
  const error = await refreshList(type);
  if (error != "") {
    errorMessage.set(error);
    return;
  }
}

export const updateRule = async (id, rule) => {
  console.debug("Updating rule", id, rule);
  const inValue = new main.InValue({ id: id, rule: rule });
  const result = await UpdateRule(inValue)
  return result;
}

export const updateAndRefresh = async (type, id, rule) => {
  let result = await updateRule(id, rule);
  if (result.error != "") {
    errorMessage.set(result.error);
    return;
  }

  const error = await refreshList(type);
  if (error != "") {
    errorMessage.set(error);
    return;
  }
}

export const getStatusColor = (status) => {
  if (status >= 400) {
    return "red";
  } else if (status >= 300) {
    return "gray";
  } else if (status >= 200) {
    return "yellow";
  } else {
    return "blue";
  }
}

export const getMethodColor = (method) => {
  if (method.toLowerCase() === "get") {
    return "green";
  } else if (method.toLowerCase() === "post" || method.toLowerCase() === "put" || method.toLowerCase() === "patch") {
    return "blue";
  } else if (method.toLowerCase() === "delete") {
    return "red";
  }

  return "blue";
}

