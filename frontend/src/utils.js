


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
    return "yellow";
  } else if (status >= 200) {
    return "green";
  } else {
    return "blue";
  }
}

export const getMethodColor = (method) => {
  if (method.toLowerCase() === "get") {
    return "green";
  } else if (method.toLowerCase() === "post" || method.toLowerCase() === "put" || method.toLowerCase() === "patch") {
    return "yellow";
  } else if (method.toLowerCase() === "delete") {
    return "red";
  }

  return "blue";
}


export function formatDate(isoDate) {
  const inputDate = new Date(isoDate);
  const now = new Date();

  // Helper functions to format dates
  const formatTime = (date) =>
    date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
  const formatDateOnly = (date) =>
    date.toLocaleDateString([], {
      year: "numeric",
      month: "short",
      day: "numeric",
    });

  // Check if the input date is today
  const isToday = inputDate.toDateString() === now.toDateString();

  // Check if the input date is within the same month and year
  const isSameMonth =
    inputDate.getFullYear() === now.getFullYear() &&
    inputDate.getMonth() === now.getMonth();

  // Formatting logic
  if (isToday) {
    // Only show time if today
    return formatTime(inputDate);
  } else if (isSameMonth) {
    // Show time and date if this month
    return `${formatDateOnly(inputDate)} ${formatTime(inputDate)}`;
  } else {
    // Show full date and year if in a different month or year
    return `${formatDateOnly(inputDate)} ${formatTime(inputDate)}`;
  }
}