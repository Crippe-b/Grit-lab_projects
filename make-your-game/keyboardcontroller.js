export class KeyboardController { // removed default
    constructor(game) {
        this.game = game;
        //this.paddle = game.movingItems[0]
        this.canPause = true;
        this.shift = false;
        window.addEventListener('keydown',this.keydown.bind(this));
        window.addEventListener('keyup',this.keyup.bind(this));
    }
    keydown(event) {
        switch (event.key.toLowerCase()) {
            case "escape":
            case "p":
                if (this.canPause === true && this.game.gameHasStarted === true) {
                    this.canPause = false;
                    this.game.togglePause();
                    this.game.toggleTimer();
                }
                break;
            case "r":
                if (document.querySelector(".pause-text").style.display === "block" || document.querySelector(".gameover-text").style.display === "block" || document.querySelector(".end-text").style.display === "block") {
                    location.reload();
                }
                break;
            case "shift":
                this.paddle.shift = true;
                break;
            case "a":
            case "arrowleft":
                this.paddle.movDirection = "left";
                break;
            case "d":
            case "arrowright":
                this.paddle.movDirection = "right";
                break;
            case " ":
                if (this.game.gameHasStarted === false) {
                    this.game.gameHasStarted = true;
                    document.querySelector(".start-text").style.display = "none";
                    this.game.movingItems[1].detachFromPaddle();
                    this.game.runTimer();
                }
                break;
            default:
                break;
        }
    }
    keyup(event) {
        switch (event.key.toLowerCase()) {
            case "shift":
                this.paddle.shift = false;
                break
            case "escape":
            case "p":
                this.canPause = true;
                break;
            case "a":
            case "arrowleft":
                if (this.paddle.movDirection === "left") this.paddle.movDirection = "none";
                break;
            case "d":
            case "arrowright":
                if (this.paddle.movDirection === "right") this.paddle.movDirection = "none";
                break;
        }
    }
    rebindPaddle(newPaddle) {
        this.paddle = newPaddle;
    }
}