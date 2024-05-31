const video = document.querySelector('video');

// Main controls
const playButton = document.querySelector('#play-pause');
const replayButton = document.querySelector('#replay');
const volumeButton = document.querySelector('#volume');
const progressIndicator = document.querySelector("#progress-indicator");
const duration = document.querySelector("#duration");

// Mask controls
const mask = document.querySelector("#mask");
const play = document.querySelector("#play");
const replay2 = document.querySelector("#replay2");

const continueButton = document.querySelector("#continue");


const updateDuration = () => {
    const formatTime = (time) => String(time).padStart(2, '0');

    const cmin = formatTime(Math.floor(video.currentTime / 60));
    const csec = formatTime(Math.floor(video.currentTime % 60));
    const tmin = formatTime(Math.floor(video.duration / 60));
    const tsec = formatTime(Math.floor(video.duration % 60));

    duration.textContent = `${cmin}:${csec} / ${tmin}:${tsec}`;
};

const togglePlayPause = () => {
    video.paused ? video.play() : video.pause();
};

const handleReplay = () => {
    mask.classList.add('hidden');
    video.currentTime = 0;
    video.play();
};

const toggleMute = () => {
    video.muted = !video.muted;

    volumeButton.children[0].classList.toggle('hidden');
    volumeButton.children[1].classList.toggle('hidden');
};

const updateProgress = () => {
    const progressPercentage = (video.currentTime / video.duration) * 100;
    progressIndicator.style.width = `${progressPercentage}%`;
    updateDuration();
};

const handleKeyUp = (e) => {
    if (e.code === "Space") {
        togglePlayPause();
    }
};

const handleVideoEnded = () => {
    mask.classList.remove('hidden');
    continueButton.classList.remove('hidden');
};

const handlePlay = () => {
    video.play();
    mask.classList.add('hidden');
    play.classList.add('hidden');
};

playButton.addEventListener('click', togglePlayPause);
replayButton.addEventListener('click', handleReplay);
volumeButton.addEventListener('click', toggleMute);
video.addEventListener('timeupdate', updateProgress);
window.addEventListener("keyup", handleKeyUp);
video.addEventListener('ended', handleVideoEnded);
play.addEventListener('click', handlePlay);
replay2.addEventListener('click', handleReplay);

export default {
    updateDuration
};