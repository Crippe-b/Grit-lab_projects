import { getCollisionInfo } from "./collisionHelper.js";
import { guid, updateMovementHtml } from "./helper.js"
export class Paddle { // removed default
    constructor(startX, startY) {
        this.type = "paddle";
        this.width = 130;
        this.height = 25;
        this.movDirection = "none";
        this.shift = false;
        this.speed = 8;
        this.bounceCounter = 0;
        this.pos = {"x": startX, "y": startY};
        this.x = startX;
        this.y = startY;
        this.id = guid();
        this.health = 1;
        this.c = getCollisionInfo(this);
    }
    move() {
        switch(this.movDirection) {
            case "right":
                if (this.shift) this.pos["x"] += (this.speed * 2);
                else this.pos["x"] += this.speed;
                break;
            case "left":
                if (this.shift) this.pos["x"] -= (this.speed * 2);
                else this.pos["x"] -= this.speed;
                break;
            default:
                break;
        }
        if (this.pos["x"] < 0) this.pos["x"] = 0;
        if (this.pos["x"] + this.width > 1024) this.pos["x"] = 1024 - this.width;
        updateMovementHtml(this.id, this.pos["x"], this.pos["y"]);
    }

    increaseSpeed(newSpeed) {
        this.speed = newSpeed;
    }

    getItemType() {
        return this.type;
    }

}