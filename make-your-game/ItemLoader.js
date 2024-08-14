import { levelCreator } from './levels.js'
import { Logo } from './Logo.js'
import { Paddle } from './Paddle.js'
//import { level } from './levels.js'
import { getRandomXDirection } from './helper.js'

class ItemLoader {
    constructor() {
        this.bricks = [];
        this.movingItems = [];
        this.gamewidth = document.querySelector(".game-wrapper").clientWidth;
    }

    loadInitialItems() {
        const paddle = new Paddle((this.gamewidth/2) - 65, 720);
        const logo = new Logo(paddle.x + (paddle.width/2) - 24, paddle.y - 30, [getRandomXDirection(),"UP"], 8);
        logo.attachedToPaddle = true;
        this.addMovingItems(paddle);
        this.addMovingItems(logo);
        this.attachItemToHtml(logo, logo.pos, "logo", logo.id, "svg");
        this.attachItemToHtml(paddle, [paddle.x, paddle.y], "paddle", paddle.id ,"div");
    }
    
    addMovingItems(item) {
        this.movingItems.push(item);
    }

    loadLevel(levelNbr) {
        this.clearAllBricks();
//        this.clearAllMovingItems()
//        this.loadInitialItems()
        this.bricks = levelCreator(levelNbr);
        for (let i = 0; i < this.bricks.length; i++) {
            this.attachItemToHtml(this.bricks[i],[this.bricks[i].pos["x"],this.bricks[i].pos["y"]],this.bricks[i].type,this.bricks[i].id, "div");
        }
    }

    getMovingItems() {
        return this.movingItems;
    }

    getBricks() {
        return this.bricks;
    }



    attachItemToHtml(item, position, clss, id, element) {
        let el = document.createElement(element);
        el.setAttribute("class", clss);
        el.setAttribute("id",id);
        el.style.top = position[1] + "px";
        el.style.left = position[0] + "px";
        switch (item.type) {
            case "brick":
                el.style.width = item.width + "px";
                el.style.height = item.height + "px";
                el.classList.add(item.bricktype);
                document.querySelector(".bricks").appendChild(el);
                break;
            case "paddle":
                el.classList.add("moving");
                document.querySelector(".game-wrapper").appendChild(el);
                break;
            case "logo":
                let ns = "http://www.w3.org/2000/svg";
                let el2 = document.createElementNS(ns,element);
                el2.setAttribute("class", clss);
                el2.classList.add("moving");
                el2.style.top = position[1] + "px";
                el2.style.left = position[0] + "px";
                el2.setAttribute("id",id);
                el2.setAttribute("viewBox","0 0 16 8");
                el2.setAttribute("fill","none");
                let p = document.createElementNS(ns,"path");
                p.setAttribute("d", "M7.4 4.959C3.268 4.959 0 5.509 0 6.186C0 6.864 3.268 7.413 7.4 7.413C11.532 7.413 14.943 6.864 14.943 6.186C14.944 5.508 11.533 4.959 7.4 4.959ZM7.263 6.51C6.306 6.51 5.53 6.273 5.53 5.98C5.53 5.687 6.306 5.45 7.263 5.45C8.22 5.45 8.995 5.687 8.995 5.98C8.995 6.273 8.219 6.51 7.263 6.51ZM13.319 0.052002H9.701L7.769 2.291L6.849 0.0830021H1.145L0.933 1.045H3.202C3.202 1.045 4.239 0.909002 4.273 1.739C4.273 3.177 1.897 3.055 1.897 3.055L2.341 1.555H0.869L0.194 3.988H2.862C2.862 3.988 5.683 3.738 5.683 1.77C5.683 1.77 5.797 1.196 5.749 0.943002L7.124 4.62L10.559 1.055H13.165C13.165 1.055 13.963 1.123 13.963 1.74C13.963 3.178 11.604 3.028 11.604 3.028L11.969 1.556H10.682L9.946 3.989H12.399C12.399 3.989 15.465 3.799 15.465 1.71C15.465 1.709 15.404 0.052002 13.319 0.052002Z");
                el2.appendChild(p);
                document.querySelector(".game-wrapper").appendChild(el2);
                break;
            default:
                el.classList.add("moving");
                document.querySelector(".game-wrapper").appendChild(el);
                break;
        }
    }
    removeDeadBricks() {
        for (let i = 0; i < this.bricks.length; i++) {
            if (this.bricks[i].health <= 0) {
                let el = document.getElementById(this.bricks[i].id);
                el.remove();
                this.bricks.splice(i, 1);
            }
        }
    }

    removeDeadMovingItems() {
        for (let i = 0; i < this.movingItems.length; i++) {
            if (this.movingItems[i].health <= 0) {
                let el = document.getElementById(this.movingItems[i].id);
                el.remove();
                this.movingItems.splice(i, 1);
            }
        }
    }

    clearAllBricks() {
        let elems = document.querySelectorAll(".brick");
        elems.forEach(element => {element.remove()});
        this.bricks = [];
    }

    clearAllMovingItems() {
        let elems = document.querySelectorAll(".moving");
        elems.forEach(element => {element.remove();});
        this.movingItems = [];
    }
}
export default new ItemLoader()