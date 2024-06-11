const QUESTIONS = {{ .Questions }};

const questionElement = document.querySelector("#question");
const optionsContainer = document.querySelector("#options");
const nextButton = document.querySelector("#next");
const submitButton = document.querySelector("#submit");
const questionCounter = document.querySelector("#qcounter");
const template = document.querySelector("#option-template");
const takeQuizButton = document.querySelector("#take-quiz");
const quiz = document.querySelector("#quiz");
const end = document.querySelector("#end");
const correct = document.querySelector("#correct");
const total = document.querySelector("#total");
const finish = document.querySelector("#finish");

let currentQuestionIndex = 0;
let score = 0;
let questionValue = 100 / QUESTIONS.length;
let correctAnswers = 0;

const loadQuestion = () => {
    const currentQuestion = QUESTIONS[currentQuestionIndex];
    questionElement.innerText = currentQuestion.question;
    questionCounter.innerText = `Question ${currentQuestionIndex + 1} / ${QUESTIONS.length}`;
    
    optionsContainer.innerHTML = '';  // Clear previous options

    currentQuestion.options.forEach((option, index) => {
        const clone = template.content.cloneNode(true);
        const span = clone.querySelector("span");
        const input = clone.querySelector("input");
        const label = clone.querySelector("label");

        span.innerText = `${index + 1}. ${option}`;
        span.classList.add("font-normal");
        input.setAttribute("value", index);
        input.setAttribute("name", "option");
        label.setAttribute("id", `option-${index}`);

        optionsContainer.appendChild(clone);
    });

    attachOptionListeners();
};

const attachOptionListeners = () => {
    const options = optionsContainer.querySelectorAll("label[id^='option-']");
    
    options.forEach(option => {
        option.addEventListener("click", handleOptionClick);
    });
};

const handleOptionClick = (event) => {
    const options = optionsContainer.querySelectorAll("label[id^='option-']");
    
    options.forEach(opt => opt.classList.remove("border-blue-500"));
    event.currentTarget.classList.add("border-blue-500");
    event.currentTarget.querySelector("input").checked = true;

    submitButton.classList.remove("hidden");
};

const removeOptionListeners = () => {
    const options = optionsContainer.querySelectorAll("label[id^='option-']");

    options.forEach(option => {
        option.removeEventListener("click", handleOptionClick);
        option.classList.add("pointer-events-none");
    });
};

nextButton.addEventListener("click", () => {
    nextButton.classList.add("hidden");

    if (currentQuestionIndex < QUESTIONS.length - 1) {
        currentQuestionIndex++;
        loadQuestion();

        return 
    }

    end.classList.remove("hidden");

    correct.innerText = correctAnswers;
    total.innerText = QUESTIONS.length;
});

submitButton.addEventListener("click", () => {
    removeOptionListeners();

    const selectedOption = optionsContainer.querySelector("input[name='option']:checked");

    if (!selectedOption) {
        alert("Please select an option before submitting.");
        return;
    }

    const currentQuestion = QUESTIONS[currentQuestionIndex];
    const correctAnswer = currentQuestion.answer;
    const selectedAnswer = selectedOption.value;

    correctAnswer == selectedAnswer ? handleCorrectAnswer(selectedOption) : handleIncorrectAnswer(selectedOption);

    submitButton.classList.add("hidden");
    nextButton.classList.remove("hidden");
});

takeQuizButton.addEventListener("click", () => {
    takeQuizButton.classList.add("hidden");
    loadQuestion();
    quiz.classList.remove("hidden");
});


finish.addEventListener("click", () => {
    if (window.API) {
        window.API.LMSSetValue("cmi.core.score.raw", score);
        window.API.LMSSetValue("cmi.core.lesson_status", "passed");
        window.API.LMSCommit("");
        window.API.LMSFinish("");
    }

    // If the window is not closed, display a message to the user.
    document.body.innerHTML = `<div class="flex flex-col text-center text-white"><h1 class='text-3xl pt-20 mb-2 font-semibold'>Thank you for taking the quiz.</h1> <p class="text-lg font-light">You can now close this window.</p></div>`;
});


const correctSVG = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 stroke-green-500 ml-auto">
<path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75 11.25 15 15 9.75M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
</svg>
`;

const incorrectSVG = `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6 stroke-red-500 ml-auto">
<path stroke-linecap="round" stroke-linejoin="round" d="m9.75 9.75 4.5 4.5m0-4.5-4.5 4.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
</svg>
`;

const handleCorrectAnswer = (answer) => {
    const label = answer.parentElement;
    label.innerHTML = label.innerHTML + correctSVG;
    label.classList.add("border-green-500");
    score += questionValue;
    correctAnswers++;
};

const handleIncorrectAnswer = (answer) => {
    const label = answer.parentElement;
    label.innerHTML = label.innerHTML + incorrectSVG;
    label.classList.add("border-red-500");

    const correctOption = optionsContainer.querySelector(`#option-${QUESTIONS[currentQuestionIndex].answer}`);
    correctOption.innerHTML = correctOption.innerHTML + correctSVG;
    correctOption.classList.add("border-green-500");
}


loadQuestion();
