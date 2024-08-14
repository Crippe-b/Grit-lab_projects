export function guid() {
    return Math.floor((1 + Math.random()) * 0x10000).toString(16).substring(1);
}

export function updateMovementHtml(id, left, top) {
    let el = document.getElementById(id);
    el.style.top = top + "px";
    el.style.left = left + "px";
}

export function getRandomStartSpeeds(speed) {
    let result = {};
    let val2 = (Math.random() * (speed - ((speed/100) * 80))) + ((speed/100) * 80);
    let val1 = speed - val2;
    result["x"] = val1;
    result["y"] = val2;
    return result
}

export function getRandomXDirection() {
    let n = Math.random();
    if (n >= 0.5) return ("LEFT");
    return "RIGHT";
}

export function randomColor() {
    let color = "#";
    color += Math.random().toString(16).slice(2, 8).toUpperCase();
    return color;
}

// export function getNewSpeed(percentage) {

// }