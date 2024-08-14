import { guid, updateMovementHtml, getRandomStartSpeeds, randomColor } from "./helper.js"
import {getCollisionInfo, getNextPosition, logoCollision} from "./collisionHelper.js"

const BrickAreaMargin = {
    left: 2,
    top: 2
}
const BrickInfo = {
    width: 96,
    height: 48
}
const BrickBorder = {
    horizontal: 3,
    vertical: 2
}
export class Logo {
    constructor(_x, _y, Dir, _speed) {
        this.type = "logo";
        this.width = 48;
        this.height = 24;
        this.speedNumber = _speed;
        this.direction = {"x": Dir[0], "y": Dir[1]};
        this.speed = getRandomStartSpeeds(_speed);
        this.pos = {"x": _x, "y": _y};
        // this.topLeft = []
        // this.topRight = []
        // this.botLeft = []
        // this.botRight = []
        this.attachedToPaddle = false;
        this.id = guid();
        //this.insideGrids = []
        this.health = 1;
        this.c = getCollisionInfo(this);
    }
    // updatePosition(_x,_y) {
    //     this.topLeft = [_x, _y]
    //     this.topRight = [_x + this.width, _y+this.height]
    //     this.botLeft = [_x, _y + this.height]
    //     this.botRight = [_x + this.width, _y + this.height]
    // }
    // updateGrids(positions) {
    //     let xMod = (positions[0] - BrickBorder.horizontal) % (BrickInfo.width + (BrickBorder.horizontal * 2))
    //     let yMod = (positions[1] - BrickBorder.vertical) % (BrickInfo.height + (BrickBorder.vertical * 2))
    //     let xCord = (positions[0] - BrickBorder.horizontal) / xMod
    //     let yCord = (positions[1] - BrickBorder.vertical) / yMod
    //     return xCord + (yCord * 10)

    // }
    
    updateSpeed(_newSpeed) {
        this.speed = _newSpeed;
    }

    move(bricks, paddle) {
        let score = 0
        if (this.attachedToPaddle === true) {
            let paddle = document.querySelector(".paddle");
            let paddleX = parseInt(paddle.style.left.slice(0,-2)) + 65 - (this.width/2);
            this.pos["x"] = paddleX;
        } else {
            this.pos = getNextPosition(this.pos, this.direction, this.speed)
            if (this.pos["y"] >= 759) {
                this.killLogo();
                return score
            }
            if (this.pos["y"] >= 600) {
                logoCollision(this, paddle);
            }
            if (this.pos["y"] <= 600) {
                for (let i = 0; i < bricks.length; i++) {
                    score += logoCollision(this, bricks[i]);
                }
            }
            if (this.pos["x"] <= 0) {
                this.pos["x"] = 0;
                this.swapDirectionX();
                this.updateColor();
            }
            if (this.pos["x"] + this.width >= 1024) {
                this.pos["x"] = 1024 - this.width;
                this.swapDirectionX();
                this.updateColor();
            }
            if (this.pos["y"] <= 0) {
                this.pos["y"] = 0;
                this.swapDirectionY();
                this.updateColor();
            }
            // this.updatePosition()
            // this.insideGrids[this.updateGrids(this.topLeft), this.updateGrids(this.topRight), this.updateGrids(this.botLeft), this.updateGrids(this.botRight)]
        }
        updateMovementHtml(this.id, this.pos["x"], this.pos["y"]);
        return score
    }
    // Used for start?
    detachFromPaddle() {
        this.attachedToPaddle = false;
    }
    getItemType() {
        return this.type;
    }

    swapDirectionX() {
        this.direction["x"] = (this.direction["x"] === "LEFT") ? "RIGHT" : "LEFT";
    }

    swapDirectionY() {
        this.direction["y"] = (this.direction["y"] === "UP") ? "DOWN" : "UP";
    }
    getNewLogoSpeed(paddle) {
        if (this.c.CenterX === paddle.c.CenterX) {
            this.speed = {"x": 0, "y": this.speedNumber};
            return
        }
        let xDiffFromCenter;
        let range = paddle.c.HalfW + (this.c.HalfW);
        let maxPercent = 80;
        if (this.c.CenterX > paddle.c.CenterX) {
            xDiffFromCenter = this.c.CenterX - (paddle.c.CenterX);
            //let range = paddle.c.HalfW + (this.c.HalfW);
            this.direction["x"] = "RIGHT";
        }
        if (this.c.CenterX < paddle.c.CenterX) {
            xDiffFromCenter = (paddle.c.CenterX) - this.c.CenterX;
            this.direction["x"] = "LEFT";
        }
        let percent = (maxPercent / range) * xDiffFromCenter;
        this.speed["x"] = (this.speedNumber / 100) * percent;
        this.speed["y"] = this.speedNumber - this.speed["x"];
    }

    updateColor() {
        document.getElementById(this.id).style.fill = randomColor();
    }

    killLogo() {
        this.health = 0;
    }
}