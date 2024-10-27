export function validatePort(port) {
    let portNumber = parseInt(port);
    if (isNaN(portNumber)) {
        return "Invalid number";
    }

    if (portNumber.toString().length != 4) {
        return "Port must be of length 4";
    }
    return "";
}