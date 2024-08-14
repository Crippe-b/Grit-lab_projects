import {levelCreator} from "/levels.js";
import {pause_clock, resume_clock} from "/timer.js";
import {powerUpWide, bonusPoints, powerUpSlowmotion, powerUpShoot} from "/powerups.js";
const section = document.querySelector("section");
const logo = document.querySelector(".logo");
const paddle = document.querySelector(".paddle")

//change this to change both framerate and game speed (🤮). be aware that request animation frame won't go over 60 though.
const FPS = 60;


// sets game size
// section.style.width = 1024 + "px";
// section.style.height = 768 + "px";

//initializes pause
let STOP = false;

//game over variable(...no shit) only really required so that you can't pause/resume once you get a game over
let gameOver, started, finishedLevel = false;




// Score and lives variables
export let playerObj = {
  lives: 3,
  score: 0,
  powerUp: []
}

// Logo velocity variables
let xPosLogo, yPosLogo;
let xSpeed, ySpeed = 0;


// Paddle velocity variables
const startPaddleSpeed = 8
let xPaddle;
let paddleSpeed = 8;
let moveLeft, moveRight = false
export let paddleWidth = 130;

// Logo speed multiplier
export let logoSpeed = 1.5;

// Updates the score
export const updateScore = () => {
  document.getElementById("score").innerHTML = "SCORE: " + playerObj.score;
  document.getElementById("player-lives").innerHTML = "LIVES: "+ playerObj.lives;
}


let currentLevel = 1;
function levelSetup() {

  finishedLevel ? currentLevel++ : currentLevel = currentLevel 
  levelCreator(currentLevel.toString());
  started = false;

  xPosLogo = 490; yPosLogo = 713;
  ySpeed = 0; xSpeed = 0;
  pause_clock();
  section.style.left = (window.innerWidth/2) - (section.clientWidth/2) + "px"
  paddle.style.left = (section.clientWidth/2) - (paddle.clientWidth/2) + "px"
  paddle.style.top = section.clientHeight - paddle.clientHeight + "px"
  xPaddle = (section.clientWidth/2) - (paddle.clientWidth/2)
}

function update() {
  logo.style.left = xPosLogo + "px";
  logo.style.top = yPosLogo + "px";
}

export function randomColor() {
  let color = "#";
  color += Math.random().toString(16).slice(2, 8).toUpperCase();
  return color;
}

let lastDrawTime, fpsInterval, lastSampleTime
let requestID

function startAnimating(fps) {
    fpsInterval = 1000 / fps
    lastDrawTime = performance.now()
    lastSampleTime = lastDrawTime
    animate()
}

function animate(now) {
    //request another frame
    requestID = requestAnimationFrame(animate);
    //variable that pauses animation when true
    if (!STOP) {

    // calc elapsed time since last loop
    let elapsed = now - lastDrawTime
    
    // if enough time has elapsed, draw the next frame
    if (elapsed > fpsInterval) {
        // get ready for next frame by setting lastDrawTime = now, but...
        // also adjust for fpsInterval not being multiple of 16.67
        lastDrawTime = now - (elapsed % fpsInterval)
        //draw
        drawNextFrame(now)
    }
  }
  
}




function drawNextFrame(now) {
  if (started === false) {
    xPosLogo = xPaddle +40
  }  
  if (xPosLogo + logo.clientWidth >= (section.clientWidth + 10) || xPosLogo <= 8) {
        xSpeed = -xSpeed;
        logo.style.fill = randomColor();
        //doTheFunny()
  }

  if (yPosLogo >= 744) {
    playerObj.lives--;
    updateScore();
    levelSetup();
  }

  if (yPosLogo >= 714 && started) {
    checkPaddleCollision();
  }
  if (yPosLogo + logo.clientHeight >= section.clientHeight || yPosLogo <= 60) {
    // console.log("y:", yPosLogo)
    ySpeed = -ySpeed;
    logo.style.fill = randomColor();
    //doTheFunny()
  }
  if (moveLeft === true) {
    xPaddle <= 10 ? xPaddle = 10 : xPaddle -= paddleSpeed;
    paddle.style.left = xPaddle + "px";
  }
  if (moveRight === true) {
    xPaddle >= 1030 - paddleWidth ? xPaddle = 1030 - paddleWidth : xPaddle += paddleSpeed;
    paddle.style.left = xPaddle + "px";
  }
  xPosLogo += xSpeed * logoSpeed;
  yPosLogo += ySpeed * logoSpeed;
  update();
}

function checkPaddleCollision() {
  let xDir = -5;
  let rightCorner = xPaddle + paddleWidth + 2;
  let leftCorner = xPaddle - 46;
  let division = 10 / (paddleWidth + 48);

  // COLLISION CONFIRMED
    if(xPosLogo >= leftCorner && xPosLogo <= rightCorner) {
      console.log("calc: ",  division * (xPosLogo - leftCorner))
      xDir += division * (xPosLogo - leftCorner);
      console.log("xDir", xDir)
      console.log("ySpeed: ", ySpeed)
      if (xDir <= -4 || xDir >= 4) {
        ySpeed *= -0.95;
      } else {
        ySpeed < 6 ? ySpeed *= -1.01 : ySpeed *= -1;
        console.log("ySpeed: ", ySpeed)
      }
      xSpeed = xDir * logoSpeed;
    }
}

let canPause = true
const moveCheck = () =>  {document.body.addEventListener("keydown", (event)  => {
  
  if (event.key.toLowerCase() === " " && started === false) {
    //add speed to the ball
    resume_clock();
    started = true
    xSpeed = ((Math.random()*2) -1) * logoSpeed;
    ySpeed = -4 * logoSpeed;
    document.querySelector(".start-text").style.display = "none";
  }

  if (event.key.toLowerCase() === "a" || event.key === 'ArrowLeft') {
    moveLeft = true
  }
  if (event.key.toLowerCase() === "d" || event.key === 'ArrowRight') {
    moveRight = true
  }
  if (event.key === "Shift") {
    paddleSpeed = startPaddleSpeed * 2
  }

  
  // JUST TO TEST POWERUPS
  if (event.key.toLowerCase() === "u") {
    paddleWidth = powerUpWide(paddleWidth);
    logoSpeed = powerUpSlowmotion(logoSpeed);
    playerObj.score = bonusPoints(playerObj.score);
    powerUpShoot();
    updateScore();
  }

  if(started === false){
    return
  }
  if (gameOver) return
  if (event.key.toLowerCase() === "p" || event.key.toLowerCase() === "escape" && canPause === true) {
    canPause = false
    if(started){
      STOP ? resume_clock() : pause_clock();
    }
    STOP = !STOP
    //console.log(STOP)
  }

})
}

const stopCheck = () =>  {document.body.addEventListener("keyup", (event)  => {
  //console.log(event)
  if (event.key.toLowerCase() === "a" || event.key === 'ArrowLeft') {
    moveLeft = false
  }
  if (event.key.toLowerCase() === "d" || event.key === 'ArrowRight') {
    moveRight = false
  }

  if (event.key.toLowerCase() === "p" || event.key === 'escape') {
    canPause = true
  }
  if (event.key === "Shift") {
    paddleSpeed = startPaddleSpeed
  }
})
}


  // setInterval(bonusPoints, 1000);


// function vibeChecker() {
//   if (vibeCheck > 10) {
//     lives = 0
//     STOP = true
//     gameOver = true
//     youFuckedUp.style.width = window.innerWidth + "px"
//     youFuckedUp.style.height = window.innerHeight + "px"
//     youFuckedUp.style.top = 0 + "px"
//     //youFuckedUp.style.left = (window.innerWidth - section.clientWidth)/2 + "px"
//     youFuckedUp.style.display = "initial"
//     youFuckedUp.play()
//     console.log("window h", window.innerHeight)
//     console.log("window w", window.innerWidth)
//     console.log("video h", youFuckedUp.clientHeight)
//     console.log("video w", youFuckedUp.clientWidth)
//   }
//   return
// }








updateScore();
levelSetup();
stopCheck();
moveCheck();
startAnimating(FPS);
// console.log(window.innerWidth);
// console.log(paddle.clientWidth);
// console.log(xPaddle);
// console.log((window.innerWidth/2));