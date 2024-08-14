import {
    Brick,
    Double,
    PowerUpBrick,
    Steel
} from "./bricks.js"

// export class Levels {

// }
const BrickAreaMargin = {
    top: 2,
    left: 2
}
const BrickBorder = {
    horizontal: 3,
    vertical: 2
}
const BrickInfo = {
    width: 96,
    height: 48
}
// "B" = BLOCK / WHITE
// "D" = DOUBLE / YELLOW
// "P" = POWERUP / LIGHT BLUE
// "S" = STEEL / GRAY
// " " = EMPTY / GUESS WHAT COLOR ?
//let i = 20
let level = {
    "1": [
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," ","B","B","B","B","B","B"," "," "],
        [" "," ","B","D","D","D","D","B"," "," "],
        [" "," ","B","D","D","D","D","B"," "," "],
        [" "," ","B","D","D","D","D","B"," "," "],
        [" "," ","B","B","B","B","B","B"," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
    ],

    "2": [
        [" "," "," "," "," "," "," "," "," "," "],
        [" ","B","B"," "," "," "," ","B","B"," "],
        [" ","B","D","B"," "," ","B","D","B"," "],
        [" "," ","B","D","B","B","D","B"," "," "],
        [" "," "," ","B","P","P","B"," "," "," "],
        [" "," ","B","D","B","B","D","B"," "," "],
        [" ","B","D","B"," "," ","B","D","B"," "],
        [" ","B","B"," "," "," "," ","B","B"," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
    ],

    "3": [
        ["B","P","B","S"," "," ","S","B","P","B"],
        ["D","D","D","S"," "," ","S","D","D","D"],
        ["B","B","B","S"," "," ","S","B","B","B"],
        ["B","B","B","S"," "," ","S","B","B","B"],
        ["B","B","B","S"," "," ","S","B","B","B"],
        ["B","B","B","S"," "," ","S","B","B","B"],
        ["B","P","B","S"," "," ","S","B","P","B"],
        ["D","D","D","S"," "," ","S","D","D","D"],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
    ],

    "4": [
        [" "," "," "," "," "," "," "," "," "," "],
        [" ","S","D","D","S","S","D","D","S"," "],
        [" ","D","B","B","B","B","B","B","D"," "],
        [" ","D","B","B","B","B","B","B","D"," "],
        [" ","S","B","P","B","B","P","B","S"," "],
        [" ","D","B","B","B","B","B","B","D"," "],
        [" ","D","B","B","B","B","B","B","D"," "],
        [" ","S","D","D","S","S","D","D","S"," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
    ],

/*     "5": [
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," ","B"," "," "," "," "],
    ], */
    "5": [
        ["D"," "," "," "," "," "," "," "," ","D"],
        ["S","S","S","S","B","B","S","S","S","S"],
        [" "," "," "," "," "," "," "," "," "," "],
        ["S","S","B","S","S","S","S","B","S","S"],
        [" "," "," "," "," "," "," "," "," "," "],
        ["B","S","S","S","S","S","S","S","S","B"],
        [" "," "," "," "," "," "," "," "," "," "],
        ["S","S","S","S","B","B","S","S","S","S"],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        ],
    
    /*"6": [
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i),Rand(i)],
        [" "," "," "," "," "," "," "," "," "," "],
        [" "," "," "," "," "," "," "," "," "," "],
        ],
 */   
}


export const levelCreator = (l) => {
    let bricks = []
    // let width = 96
    // let height = 48
    let idCounter = 0
    //let levelString = "";
    for(let row = 0; row < level[l].length; row++) {
        // let temp = []
        for(let i = 0; i < level[l][row].length; i++){
            let xpos = BrickAreaMargin.left + BrickBorder.horizontal + ((BrickInfo.width + (BrickBorder.horizontal * 2)) * i)
            let ypos = BrickAreaMargin.top + BrickBorder.vertical + ((BrickInfo.height + (BrickBorder.vertical * 2)) * row)
            switch(level[l][row][i]){
                case 'B':
                    //levelString += `<div class="brick normal-brick" id="${row.toString() + i.toString()}"></div>`;
                    bricks.push(new Brick(BrickInfo.width, BrickInfo.height, idCounter,xpos,ypos))
                    idCounter++
                    break;
                case 'D':
                    //levelString += `<div class="brick double-brick" id="${row.toString() + i.toString()}"></div>`;
                    bricks.push(new Double(BrickInfo.width, BrickInfo.height, idCounter,xpos,ypos))
                    idCounter++
                    break;
                case 'S':
                    //levelString += `<div class="brick steel-brick" id="${row.toString() + i.toString()}"></div>`;
                    bricks.push(new Steel(BrickInfo.width, BrickInfo.height, idCounter,xpos,ypos))
                    idCounter++
                    break;
                case 'P':
                    //levelString += `<div class="brick powerup-brick" id="${row.toString() + i.toString()}"></div>`;
                    bricks.push(new PowerUpBrick(BrickInfo.width, BrickInfo.height, idCounter,xpos,ypos))
                    idCounter++
                    break;
                default:
                    //levelString += `<div class="brick empty-brick" id="${row.toString() + i.toString(),xpos,ypos}"></div>`;
                    break;
            }
        }
        // gameArea.push(temp);
    }

   //document.querySelector(".bricks").innerHTML = levelString;
    return bricks;
}
/*
function Rand(n = 0) {
    n += Math.round(Math.random() * 100);
    if (n >= 0 && n <= 15) return " ";
    if (n >= 16 && n <= 50) return "B";
    if (n >= 51 && n <= 90) return "D";
    if (n >= 91 && n <= 99) return "P";
    if (n >= 100) return "S";
}
*/
/*
    EMPTY MODEL
" ": [
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
    [" "," "," "," "," "," "," "," "," "," "],
],

*/