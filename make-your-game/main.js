import {Arkanoid} from "./arkanoid.js";

// point is to bring the entire game into a single class for readability's sake.
const game = new Arkanoid;

function startGameLooping() {
    game.update();
    window.requestAnimationFrame(startGameLooping);
}
function test() {
    //let el = document.createElement("div")
    //el.classList.add("scanlines")
    //el.setAttribute("id","lole")
    let bleh = document.getElementsByClassName("game-wrapper");
    //let bleh2 = document.getElementsByClassName("game")

    bleh[0].classList.add("crt");

    //bleh[0].appendChild(el)
}
startGameLooping();
test();