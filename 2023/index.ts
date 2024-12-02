// get input.txt and put is in a string variable

const fs = require('fs');
const input = fs.readFileSync('days/01/part2/input.txt', 'utf8');
const lines = input.split(/\r?\n/);

const firstPart = () => {
    const DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
    const numbers = [];

    for (const line of lines as string[]) {
        // find first digit
        let firstDigit = '', secondDigit = '';
        for (const digit of line.split('')) {
            if (DIGITS.includes(digit)) {
                firstDigit = digit;
                break;
            }
        }

        for (const digit of line.split('').reverse()) {
            if (DIGITS.includes(digit)) {
                secondDigit = digit;
                break;
            }
        }

        let number = firstDigit + secondDigit;
        if (number.length === 1) {
            number += number;
        }

        const final = parseInt(number);
        numbers.push(final);
    }

    const sum = numbers.reduce((a, b) => a + b, 0);
    console.log(sum);
}

const secondPart = () => {
    const DIGITS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];

    const DIGITS_MAP: Record<string, number> = {
        '1': 1,
        '2': 2,
        '3': 3,
        '4': 4,
        '5': 5,
        '6': 6,
        '7': 7,
        '8': 8,
        '9': 9,
        'one': 1,
        'two': 2,
        'three': 3,
        'four': 4,
        'five': 5,
        'six': 6,
        'seven': 7,
        'eight': 8,
        'nine':9,
    };

    function indexOfAll(inputStr: string, value: string): number[] {
        let result: number[] = [];
        let i = 0;
        while (i < inputStr.length) {
            let foundIndex = inputStr.indexOf(value, i);
            if (foundIndex !== -1) {
                result.push(foundIndex);
                i = foundIndex + 1;
            } else {
                break;
            }
        }
        return result;
    }

    function flatten(input: number[][][]): number[][] {
        let flattened: number[][] = [];
        for (let i = 0; i < input.length; i++) {
            for (let j = 0; j < input[i].length; j++) {
                flattened.push(input[i][j]);
            }
        }
        return flattened;
    }



    const numbers = [];

    let i = 0

    for (const line of lines as string[]) {
        console.log(line);
        // find first digit
        let firstDigit = '', secondDigit = '';

        const foundDigits = flatten(DIGITS
            .map((digit) =>
                indexOfAll(line, digit).map((index) => [DIGITS_MAP[digit], index])));

        const firstDigits = foundDigits
            .filter(digit => digit[1] !== -1)
            .sort((a, b) => a[1] - b[1])

        firstDigit = firstDigits[0][0].toString();

        const secondDigits = foundDigits
            .filter(digit => digit[1] !== -1)
            .sort((a, b) => b[1] - a[1])

        secondDigit = secondDigits[0][0].toString();

        console.log(firstDigits, secondDigits);

        let number = firstDigit + secondDigit;
        console.log(number);
        if (number.length === 1) {
            number += number;
        }

        const final = parseInt(number);
        numbers.push(final);

        i++
    }

    const sum = numbers.reduce((a, b) => a + b, 0);
    console.log(sum);
}

secondPart();