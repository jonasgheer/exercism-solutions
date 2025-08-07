#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = vec![];

    for input in inputs {
        match input {
            CalculatorInput::Value(n) => stack.push(*n),
            CalculatorInput::Add => {
                let (b, a) = (stack.pop()?, stack.pop()?);
                stack.push(a + b);
            }
            CalculatorInput::Subtract => {
                let (b, a) = (stack.pop()?, stack.pop()?);
                stack.push(a - b);
            }
            CalculatorInput::Multiply => {
                let (b, a) = (stack.pop()?, stack.pop()?);
                stack.push(a * b);
            }
            CalculatorInput::Divide => {
                let (b, a) = (stack.pop()?, stack.pop()?);
                stack.push(a / b);
            }
        }
    }
    if stack.len() != 1 {
        return None;
    }
    return Some(stack[0]);
}
