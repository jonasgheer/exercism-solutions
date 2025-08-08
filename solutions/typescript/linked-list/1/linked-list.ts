export class LinkedList<Element> {
    #list: Element[] = []

    push(element: Element): void {
        this.#list.push(element);
    }

    pop(): Element | undefined {
        return this.#list.pop();
    }

    shift(): Element | undefined {
        return this.#list.shift();
    }

    unshift(element: Element): void {
        this.#list.unshift(element);
    }

    delete(element: Element): void {
        this.#list = this.#list.filter(v => v !== element);
    }

    count(): number {
        return this.#list.length;
    }
}
