export class InvalidInputError extends Error {
    constructor(message: string) {
        super();
        this.message = message || "Invalid Input";
    }
}

const DIRECTION = ["north", "east", "south", "west"] as const;
type Direction = typeof DIRECTION[number];

type Turn = "R" | "L";

type ChangeDirection = `${Direction} + ${Turn}`;

type Coordinates = [number, number];

export class Robot {
    #bearing: Direction = "north";
    #x = 0;
    #y = 0;

    get bearing(): Direction {
        return this.#bearing;
    }

    get coordinates(): Coordinates {
        return [this.#x, this.#y];
    }

    place({ x, y, direction }: { x: number; y: number; direction: string }): void {
        this.#x = x;
        this.#y = y;
        if (!Robot.#isDirection(direction)) {
            throw new InvalidInputError("invalid direction");
        }
        this.#bearing = direction;
    }

    evaluate(instructions: string): void {
        Robot.#tryValidateInstructions(instructions);

        for (const instruction of instructions.split("")) {
            switch (instruction) {
                case "R":
                case "L":
                    this.#bearing = this.#changeDirection(instruction);
                    continue;
                case "A":
                    this.#advance();
            }
        }
    }

    #advance(): void {
        switch (this.#bearing) {
            case "north":
                this.#y++;
                break;
            case "east":
                this.#x++;
                break;
            case "south":
                this.#y--;
                break;
            case "west":
                this.#x--;
                break;
        }
    }

    #changeDirection(turn: Turn): Direction {
        const directionAndTurn: ChangeDirection = `${this.#bearing} + ${turn}`;
        switch (directionAndTurn) {
            case "north + R":
                return "east";
            case "north + L":
                return "west";
            case "east + R":
                return "south";
            case "east + L":
                return "north";
            case "south + R":
                return "west";
            case "south + L":
                return "east";
            case "west + R":
                return "north";
            case "west + L":
                return "south";
            default:
                throw new Error("Invalid case");
        }
    }

    static #isDirection(direction: string | Direction): direction is Direction {
        return DIRECTION.includes(direction as Direction);
    }

    static #tryValidateInstructions(instructions: string): boolean {
        if (instructions.match(/[RAL]+/)) return true;
        throw new InvalidInputError(`${instructions} is not valid`);
    }
}
