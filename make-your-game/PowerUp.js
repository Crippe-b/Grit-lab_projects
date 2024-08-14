import { Logo } from "./Logo.js";
import { guid } from "./helper.js"

// ["shoot", "wide", "slowmotion", "points"]
const powers = ["shoot", "wide", "slowmotion", "points"]
const paddle = document.querySelector(".paddle");

export class PowerUp {
    constructor(_x, _y) {
        this.type = "powerup";
        this.width = 48;
        this.height = 48;
        this.speed = 2;
        this.botLeft = [_x, _y + this.height];
        this.botRight = [_x + this.width, _y + this.height];
        this.position = [_x, _y];
        this.timer = 30;
        this.id = guid()
    }
    move(_x) {
        this.position = [_x, _y - this.speed];
    }
    removePowerup() {
        this.type = "none";
    }

    getItemType() {
        return this.type;
    }

    countdown() {

    }
}

export class Width extends PowerUp {
    constructor(){
        super(_x, _y);
    }

    toggle(_width){
        return _width === 130 ? 200 : 130;
    }

}

export class Slowmo extends PowerUp {
    constructor(){
        super(_x, _y);
    }
    toggle(_speed) {
        return _speed === 8 ? 5 : 8;
    }

}

export class Points extends PowerUp {
    constructor(){
        super(_x, _y);
    }
    give(){
        return 100;
    }
}

export class LogoPowerUp extends PowerUp {
    constructor(){
        super(_x, _y);
    }
    toggle(_speed) {
        const logo2 = new Logo(_x, _y, 3, _speed);
        const logo3 = new Logo(_x, _y, -3, _speed);
    }
}