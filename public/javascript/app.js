function answerToggle() {
    const answerWrapper = document.querySelectorAll('.answer-wrapper');
    const toggleButtons = document.querySelectorAll('.answer-toggle');

    answerWrapper.forEach((element) => {
        element.style.display = 'none'
    });

    toggleButtons.forEach((button) => {
        button.addEventListener('click', (e) => {
            const answer = e.target.parentElement.nextElementSibling;
            answer.style.display =
                answer.style.display === 'none'
                    ? 'block'
                    : 'none';
        })
    });
}

answerToggle();