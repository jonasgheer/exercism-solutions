export function timeToMixJuice(name) {
  switch (name) {
    case "Pure Strawberry Joy":
      return 0.5;
    case "Energizer":
    case "Green Garden":
      return 1.5;
    case "Tropical Island":
      return 3;
    case "All or Nothing":
      return 5;
    default:
      return 2.5;
  }
}

export function limesToCut(wedgesNeeded, limes) {
  let count = 0;
  let wedges = 0;
  while (wedges < wedgesNeeded && limes.length > 0) {
    switch (limes.shift()) {
      case "small":
        wedges += 6;
        break;
      case "medium":
        wedges += 8;
        break;
      case "large":
        wedges += 10;
        break;
    }
    count++;
  }
  return count;
}

export function remainingOrders(timeLeft, orders) {
  while (timeLeft > 0)
    timeLeft -= timeToMixJuice(orders.shift());
  return orders;
}
