const calculator = {
  displayValue: '0',
  firstOperand: null,
  waitingForSecondOperand: false,
  operator: null,
};

async function processData(val1, op, val2) {
  const values = {
    "val1": val1.toString(),
    "val2": val2.toString(),
    "op": op.toString(), 
  }

  try {
    const data = await fetch("/solve", {
      method: 'POST', 
      body: JSON.stringify(values), 
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      }
    })
    const jsonData = await data.json();
    console.log(jsonData)
    return jsonData;
  } catch(e) {
    return e;
  }
}


function inputDigit(digit) {
  const { displayValue, waitingForSecondOperand } = calculator;

  if (waitingForSecondOperand === true) {
    calculator.displayValue = digit;
    calculator.waitingForSecondOperand = false;
  } else {
    calculator.displayValue = displayValue === '0' ? digit : displayValue + digit;
  }
}

function inputDecimal(dot) {
  // If the `displayValue` does not contain a decimal point
  if (!calculator.displayValue.includes(dot)) {
    // Append the decimal point
    calculator.displayValue += dot;
  }
}

function handleOperator(nextOperator) {
  const { firstOperand, displayValue, operator } = calculator
  const inputValue = parseFloat(displayValue);

  if (operator && calculator.waitingForSecondOperand)  {
    calculator.operator = nextOperator;
    return;
  }

  if (firstOperand == null) {
    calculator.firstOperand = inputValue;
  } else if (operator) {
    const currentValue = firstOperand || 0;
    //const result = performCalculation[operator](currentValue, inputValue);
    res = processData(currentValue, operator, inputValue)
    res.then(data => {
      calculator.displayValue = data.result;
      calculator.firstOperand = data.result;
      console.log("data in fun", data.result)
      document.getElementById("showResult").value = data.result
    })
  }

  calculator.waitingForSecondOperand = true;
  calculator.operator = nextOperator;

}

function resetCalculator() {
  calculator.displayValue = '0';
  calculator.firstOperand = null;
  calculator.waitingForSecondOperand = false;
  calculator.operator = null;
}

function updateDisplay() {
  const display = document.querySelector('.calculator-screen');
  display.value = calculator.displayValue;
}


updateDisplay();

const keys = document.querySelector('.calculator-keys');
keys.addEventListener('click', (event) => {
  const { target } = event; // target.values is digits
  if (!target.matches('button')) { // ignore, if its not button
    return;
  }

  if (target.classList.contains('operator')) {
    handleOperator(target.value);
		updateDisplay();
    return;
  }

  if (target.classList.contains('decimal')) {
    inputDecimal(target.value);
		updateDisplay();
    return;
  }

  if (target.classList.contains('all-clear')) {
    resetCalculator();
		updateDisplay();
    return;
  }

  inputDigit(target.value); 
  updateDisplay();
});
