export function createVisitor(name, age, ticketId) {
  return {
    name,
    age,
    ticketId
  }
}

export function revokeTicket(visitor) {
  visitor.ticketId = null;
  return visitor;
}

export function ticketStatus(tickets, ticketId) {
  if (tickets[ticketId] === undefined) 
    return "unknown ticket id";
  else if (tickets[ticketId] === null)
    return "not sold";
  else
    return `sold to ${tickets[ticketId]}`;
}

export function simpleTicketStatus(tickets, ticketId) {
  if (tickets[ticketId] == null) 
    return "invalid ticket !!!"
  return tickets[ticketId];
}

export function gtcVersion(visitor) {
  return visitor.gtc?.version;
}
