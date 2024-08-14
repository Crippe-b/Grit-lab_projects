import Player from './Player.js'
import {KeyboardController} from './keyboardcontroller.js'
import {PowerUp} from './PowerUp.js'
import ItemLoader from './ItemLoader.js'
 
export const GameState = {
    RUNNING: 'running',
    PAUSE: 'pause',
    STOPPED: 'stopped'
}
const ItemType = {
    LOGO: "logo",
    PADDLE: "paddle",
    Brick: "brick"
}

export class Arkanoid { // removed default
    constructor() {
        this.timer = 300;
        this.state = GameState.RUNNING;
        this.currentLevel = 0;
        this.keyboardController = new KeyboardController(this);
        this.movingItems;
        this.timeInterval;
        this.gameHasStarted = false;
        this.loadNextLevel();
        this.items = ItemLoader.getBricks();
    }

    update() {
        this.updateScore();
        this.constructor.updateLifeCounter();

        if (this.state === GameState.RUNNING) {
            this.moveItems();
            this.checkRemainingBricks();
            this.checkPlayerLife();
        }
    }

    moveItems() {
        // gets all moving items (ex: if there are multiple balls/dvdlogos
        // in play or powerup icons + lasers from powerups and the paddle) 
        // and then applies movement to them
        this.movingItems = ItemLoader.getMovingItems()
        for (let i = 0; i < this.movingItems.length; i++) {
            if (this.movingItems[i].getItemType() === "logo") {
                Player.increaseScore(this.movingItems[i].move(this.items, this.movingItems[0]));
            }
            if (this.movingItems[i].getItemType() === "paddle") this.movingItems[i].move();
        }
        ItemLoader.removeDeadBricks();
        ItemLoader.removeDeadMovingItems();
        this.items = ItemLoader.getBricks();
    }

    checkPlayerLife() {
        //checks if there is a single ball/logo in play or if there is a ball attatched to the paddle, otherwise it removes a life from Player class
        let allLogosDestroyed = true;
        for (let i = 0; i < this.movingItems.length; i++) {
            if (this.movingItems[i].getItemType() === ItemType.LOGO) {
                allLogosDestroyed = false;
                break;
            }
        }
        if (allLogosDestroyed) {
            if (Player.getLife() <= 1) {
                this.gameOver();
                Player.decreaseLife();
            } else {
                this.reloadItems();
                Player.decreaseLife(); 
            }
        }
    }

    checkRemainingBricks() {
        //ranges over all items and checks if there are any remaining bricks in play, otherwise it will load the next level
        let bricksRemaining = false;
        for (let i = 0; i < this.items.length; i++) {
            if (this.items[i].getItemType() === ItemType.Brick && this.items[i].bricktype !== "steel") {
                bricksRemaining = true;
                break;
            }
        }
        if (!bricksRemaining) {
            this.loadNextLevel();
        }
    }

    togglePause() {
        // just toggles a pause variable each time the function is run
        this.state = (this.state === GameState.RUNNING) ?
            GameState.PAUSE : GameState.RUNNING;
            document.querySelector(".pause-text").style.display = (this.state === GameState.RUNNING) ?
            "none" : "block";
    }

    toggleTimer() { 
        (this.state === GameState.PAUSE) ?
            this.pauseTimer() : this.runTimer();
    }
    updateClock() {
        this.timer -= 1;
        let el = document.getElementById("timer");
        el.innerText = "TIME: " + this.timer;
        if(this.timer <= 0) clearInterval(this.timeInterval);
    }
    pauseTimer() {
		this.timeInterval = clearInterval(this.timeInterval); // stop the clock
    }

    runTimer() {
        this.timeInterval = setInterval(this.updateClock.bind(this),1000);
    }


    loadNextLevel() {
        //add comment
        this.reloadItems();
        this.currentLevel += 1;
        if (this.currentLevel <= 5 /*|| this.currentLevel !== 666*/) {
            ItemLoader.loadLevel(this.currentLevel);
            this.updateLevelCounter();
        } else {
            this.endGame();
        }

    }

    gameOver() {
        //just toggles pause and shows the gameoverscreen. may need to change to something more permanent than just pausing though.
        this.togglePause();
        document.querySelector(".gameover-text").style.display = "block"
        document.querySelector(".pause-text").style.display = "none"
        this.pauseTimer()
        this.state = GameState.STOPPED;
        // const gameOverScreen = document.querySelector('.gameOverScreen')
        // gameOverScreen.className += ' show'
    }

    updateLevelCounter() {
        document.getElementById("level").innerText = "LEVEL: " + this.currentLevel;
    }

    static updateLifeCounter() {
        document.getElementById("player-lives").innerText = "LIVES: " + Player.getLife();
    }

    updateScore() {
        document.getElementById("score").innerText = "SCORE: " + Player.getScore();
    }

    reloadItems() {
        ItemLoader.clearAllMovingItems();
        ItemLoader.loadInitialItems();
        this.movingItems = ItemLoader.getMovingItems();
        this.keyboardController.rebindPaddle(this.movingItems[0]);

        this.gameHasStarted = false;
        clearInterval(this.timeInterval);
        this.timer = 300;
        let el = document.getElementById("timer");
        el.innerText = "TIME: " + this.timer;
    }

    endGame() {
        this.togglePause();
        document.querySelector(".end-text").style.display = "block"
        document.querySelector(".pause-text").style.display = "none"
        this.pauseTimer()
        //this.state = GameState.STOPPED;
    }
}


