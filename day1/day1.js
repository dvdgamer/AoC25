const { error } = require('console');
const fs = require('fs');
const input = fs.readFileSync('./input.txt').toString()
// const input = fs.readFileSync('./test-input.txt').toString()

let i;
let safe = [];
let result = 0;
let number = '';
let currentSafePosition = 50;

function populateSafeDial() {
	i = 0;
	while (i < 100) {
		safe.push(i);
		i++;
	}
}

function extractNumber(input) {
	i++;
	while (input[i] != '\n') {
		number += input[i];
		i++;
	}

}
function moveSafeDial(direction, amount) {
	if (direction === 'R') {
		currentSafePosition += amount;
		if (currentSafePosition < 100) { //ERROR: Infiinite loop
			currentSafePosition = currentSafePosition % 100;
		}
		console.log(currentSafePosition);
	}
	if (direction === 'L') {
		currentSafePosition -= amount;
		if (currentSafePosition < 0)
			currentSafePosition = 100 + currentSafePosition;
		console.log(currentSafePosition);
	}
	if (direction !== 'R' && direction !== 'L') {
		error("choose a good direction")
	}
	if (currentSafePosition === 0)
		result++;
}

populateSafeDial();
i = 0;
while (i < input.length) {
	if (input[i] == 'L') {
		extractNumber(input);
		number = parseInt(number);
		moveSafeDial('L', number);
		number = 0;
	}
	if (input[i] == 'R') {
		extractNumber(input);
		number = parseInt(number);
		moveSafeDial('R', number);
		number = 0;
	}
	i++;
}
console.log("result: ", result);
