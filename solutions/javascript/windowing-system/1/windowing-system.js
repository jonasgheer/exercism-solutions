// @ts-check

export function Size(width = 80, height = 60) {
    this.width = width;
    this.height = height;
}

Size.prototype.resize = function(width, height) {
    this.width = width;
    this.height = height;
}

export function Position(x = 0, y = 0) {
    this.x = x;
    this.y = y;
}

Position.prototype.move = function(x, y) {
    this.x = x;
    this.y = y;
}

export class ProgramWindow {
    screenSize = new Size(800, 600);
    size = new Size();
    position = new Position();

    resize(size) {
        let {width, height} = size;
        width = Math.max(1, width);
        height = Math.max(1, height);
        this.size = new Size(Math.min(this.screenSize.width - this.position.x, width), Math.min(this.screenSize.height - this.position.y, height));
    }

    move(position) {
        let {x, y} = position;
        x = Math.max(0, x);
        y = Math.max(0, y);
        this.position = new Position(Math.min(this.screenSize.width - this.size.width, x), Math.min(this.screenSize.height - this.size.height, y));
    }
}

export function changeWindow(programWindow) {
    programWindow.resize(new Size(400, 300));
    programWindow.move(new Position(100, 150));
    return programWindow;
}

