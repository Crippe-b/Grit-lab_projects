// ["shoot", "wide", "slowmotion", "points"]

import {randomColor} from "/main.js"
let paddle = document.querySelector(".paddle");
// let ammo = 20;

const shot = '<div class="shots"><div class="shot-left"></div><div class="shot-right"></div><div>'
export const powerUpShoot = () => {
        // SFX HERE
        let färg = randomColor();
        paddle.style.backgroundColor = färg;
        paddle.innerHTML += shot;
}

let count = 0;
export const powerUpWide = (pw) => {
    
    // SFX HERE
    if (count%2==0){
        paddle.style.width = 200 + "px";
        count++;
        return pw + 70;
    } else {
        paddle.style.width = 130 + "px";
        count++;
        return pw - 70;
    }
}

export const powerUpSlowmotion = (s) => {
    // SFX HERE
    return s * 0.7;
} 

export const bonusPoints = (s) => {
   return s + 100;
}