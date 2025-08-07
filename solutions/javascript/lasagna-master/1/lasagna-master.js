export function cookingStatus(timer) {
  if (timer === undefined) 
    return "You forgot to set the timer.";
  else if (timer === 0) 
    return "Lasagna is done.";
  else 
    return "Not done, please wait.";
}

export function preparationTime(layers, timePerLayer = 2) {
  return layers.length * timePerLayer;
}

export function quantities(layers) {
  return {
    noodles: layers.filter(l => l === "noodles").length * 50,
    sauce: layers.filter(l => l === "sauce").length * 0.2
  }
}

export function addSecretIngredient(friendsList, myList) {
  myList.push(friendsList[friendsList.length - 1]);
}

export function scaleRecipe(recipe, portions) {
  const scaled = {...recipe};
  for (const p in scaled) {
    scaled[p] *= portions / 2;
  }
  return scaled;
}
