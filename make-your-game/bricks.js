import { PowerUp } from "./PowerUp.js"

export class Brick {
    constructor(_width, _height, _id, _x, _y){
        this.id = _id;
        this.type = "brick";
        this.bricktype = "normal";
        this.width = _width;
        this.height = _height;
        this.health = 1;
        this.points = 10;
        this.color = "white";
        this.pos = {"x": _x, "y": _y};
    }

    removeBrick() {
        this.type = "none";
        return this.points;
    }

    killBrick() {
        this.health = 0;
    }
    getItemType() {
        return this.type;
    }
    decreaseLife(){
        this.health -= 1;
    }
    getHealth() {
        return this.health;
    }
}

export class Double extends Brick {
    constructor(...args){
        super(...args);
        this.health = 2;
        this.points = 50;
        this.bricktype = "double";
    }
    decreaseLife(){
        this.health -= 1;
        this.color = "orange";
        let el = document.getElementById(this.id);
        el.style.backgroundColor = this.color;
    }
}

export class PowerUpBrick extends Brick {
       constructor(...args){
        super(...args);
        this.bricktype = "powerup";
    }

    spawn(){
        //return new PowerUp();
        console.log("spawning powerup")
    }

}

export class Steel extends Brick {
    constructor(...args){
        super(...args);
        this.points = 42069;
        this.health = 99;
        this.bricktype = "steel";
    }

}