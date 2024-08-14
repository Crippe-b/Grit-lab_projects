

export function getNextPosition(position,direction,speed) {
    if (direction["x"] === "LEFT") position["x"] -= speed["x"];
    if (direction["x"] === "RIGHT") position["x"] += speed["x"];
    if (direction["y"] === "UP") position["y"] -= speed["y"];
    if (direction["y"] === "DOWN") position["y"] += speed["y"];

    return position;
}

// export function collisionHandler(item) {
//     if (item.type === "logo") logoCollision(item, brick);
//     if (item.type === "") "lol"
// }

export function getCollisionInfo(obj) {
    let info = {}
    info.HalfW = obj.width / 2
    info.HalfH = obj.height / 2
    info.CenterX = obj.pos.x + obj.width / 2
    info.CenterY = obj.pos.y + obj.height / 2
    return info
}

export function logoCollision(logo, brick) {
    let score = 0
    if (brick.type === "brick" && brick.getHealth() <= 0) return score;
    if (
        logo.pos["x"] > brick.pos["x"] + brick.width ||
        logo.pos["x"] + logo.width < brick.pos["x"] ||
        logo.pos["y"] > brick.pos["y"] + brick.height ||
        logo.pos["y"] + logo.height < brick.pos["y"]
      ) return score
      logo.c = getCollisionInfo(logo)
      brick.c = getCollisionInfo(brick)
      logo.updateColor()
    
    if (brick.type === "brick" && brick.bricktype === "powerup") {
        brick.spawn()
    }
    // let logoHalfW = logo.width / 2
    // let logoHalfH = logo.height / 2
    // let brickHalfW = brick.width / 2
    // let brickHalfH = brick.height / 2
    // let logoCenterX = logo.pos.x + logo.width / 2
    // let logoCenterY = logo.pos.y + logo.height / 2
    // let brickCenterX = brick.pos.x + brick.width / 2
    // let brickCenterY = brick.pos.y + brick.height / 2

// Calculate the distance between centers
    let diffX = logo.c.CenterX - brick.c.CenterX
    let diffY = logo.c.CenterY - brick.c.CenterY
// Calculate the minimum distance to separate along X and Y
    let minXDist = logo.c.HalfW + brick.c.HalfW
    let minYDist = logo.c.HalfH + brick.c.HalfH
// Calculate the depth of collision for both the X and Y axis
    let depthX = diffX > 0 ? minXDist - diffX : -minXDist - diffX
    let depthY = diffY > 0 ? minYDist - diffY : -minYDist - diffY
    
    if (depthX != 0 && depthY != 0) {
        if (Math.abs(depthX) < Math.abs(depthY)) {
            //"HORIZONTAL"
            logo.direction["x"] = (logo.direction["x"] === "LEFT") ? "RIGHT" : "LEFT"
            logo.pos["x"] += depthX
            if (brick.type === "paddle") {
                logo.c = getCollisionInfo(logo)
                logo.getNewLogoSpeed(brick); }
            else {
                brick.decreaseLife()
            }
        } else {
            //"VERTICAL"
            logo.direction["y"] = (logo.direction["y"] === "UP") ? "DOWN" : "UP"
            logo.pos["y"] += depthY
            if (brick.type === "brick") {
                brick.decreaseLife();
            }
            if (brick.type === "paddle") {
                logo.c = getCollisionInfo(logo)
                logo.getNewLogoSpeed(brick); }
        }
        if (brick.type === "brick" && brick.getHealth() <= 0) score += brick.points
    }
    return score
}

// export function CalculateCollisionInfo(w, h ) {
//     let cInfo = {}
//     cInfo.halfW
//     cInfo.halfH 
//     cInfo.centerX
//     cInfo.centerY
// }