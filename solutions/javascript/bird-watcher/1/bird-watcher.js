export function totalBirdCount(birdsPerDay) {
  let total = 0;
  for (const count of birdsPerDay) 
    total += count;
  return total;
}

export function birdsInWeek(birdsPerDay, week) {
  let total = 0;
  for (let i = 7 * (week - 1); i < 7 * week; i++) 
    total += birdsPerDay[i];
  return total;
}

export function fixBirdCountLog(birdsPerDay) {
  for (let i = 0; i < birdsPerDay.length; i += 2) 
    birdsPerDay[i]++;
  return birdsPerDay;
}
