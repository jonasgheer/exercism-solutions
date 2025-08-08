export function twoSum(array1, array2) {
  return Number(array1.join("")) + Number(array2.join(""));
}

export function luckyNumber(value) {
  value = value.toString().split("");
  for (let i = 0; i < value.length / 2; i++) {
    if (value[i] !== value[value.length - i - 1]) return false;
  }
  return true;
}

export function errorMessage(input) {
  if (!input) 
    return "Required field";
  if (Number.isNaN(Number(input)) || Number(input) === 0) 
    return "Must be a number besides 0";
  else 
    return "";
}
