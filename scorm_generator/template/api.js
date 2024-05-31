let findAPITries = 0;
const MAX_TRIES = 7;

function findAPI(win) {
    while (win.API == null && win.parent != null && win.parent != win) {
        findAPITries++;
        if (findAPITries > MAX_TRIES) {
            alert("Error finding API -- too deeply nested.");
            return null;
        }
        win = win.parent;
    }
    return win.API;
}

function getAPI() {
    let theAPI = findAPI(window);

    if (theAPI == null && window.opener != null && typeof(window.opener) != "undefined") {
        theAPI = findAPI(window.opener);
    }

    if (theAPI == null) {
        console.error("Unable to find an API adapter");
        return null
    }

    return theAPI;
}

const API = getAPI();

if (API != null) {
    API.LMSInitialize("");
    window.API = API;
}


