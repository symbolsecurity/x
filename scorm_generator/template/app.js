import "./api.js";
import video from "./video.js";
import "./quiz.js";

if (window.API) {
    window.API.LMSSetValue("cmi.core.score.min", "0");
    window.API.LMSSetValue("cmi.core.score.max", "100");
    window.API.LMSSetValue("cmi.core.lesson_status", "incomplete");
    window.API.LMSCommit("");
}


video.updateDuration();
