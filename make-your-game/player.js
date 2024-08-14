//player class that holds the current lives and score of a player. default export starts the players lives at 3 and score at 0 when it's imported.
import {Paddle} from "./Paddle.js";
import {Logo} from "./Logo.js";
// POWERUPS: "wide", "slowmo", "logo", "shoot" (bonus points will never go in array)
class Player {
    constructor(life, score) {
        this.life = life;
        this.score = score;
        this.powerup = [];
        //this.paddle = new Paddle();
        //this.logo = new Logo;
        this.logoAmount = 1;
    }

    logoCheck() {
        if (this.logoAmount <= 0){                  
            this.life -= 1;
            resetPlayer();
        }
    }
    
    decreaseLife() {
        this.life -= 1;
    }
    
    increaseLife() {
        this.life += 1;
    }

    getLife() {
        return this.life;
    }

    getScore() {
        return this.score;
    }

    increaseScore(increase) {
        this.score += increase;
    }

    addPowerup(p) {
        this.powerup.push(p);
    }

    removePowerup() {
        this.powerup.shift();
    }

    powerUpCheck() {
        for (let i = 0; i < this.powerup.length; i++){
            switch (this.powerup[i]) {
                case 'wide':
                    break;
                case 'slowmo':
                    break;
                case 'shoot':
                    break;
                case 'logo':
                    break;
            }
        }
    }
}

export default new Player(3, 0);