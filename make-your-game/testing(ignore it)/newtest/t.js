var c = document.getElementById("myCanvas");
let o = document.getElementById("loko")
var ctx = c.getContext("2d");
var imgData = ctx.createImageData(1024, 759);
let l = new ImageData(1024,759)
// let ll = new ArrayBuffer(l)
// let lll = new Uint32Array(ll)
let lll = new Uint32Array(l.data.buffer)
let b = new Uint32Array(imgData.data.buffer)
let len = lll.length

// for (i = 0; i < imgData.data.length; i += 4) {
//   imgData.data[i+0] = 255;
//   imgData.data[i+1] = 0;
//   imgData.data[i+2] = 0;
//   imgData.data[i+3] = 255;
// }
for (var i = 0; i < len; i++) {
    lll[i] = ((255 * Math.random()) | 0) << 24;
}
ctx.putImageData(l,0,0)

//ctx.putImageData(imgData, 10, 10);
console.log(imgData.data.buffer)
console.log(new ArrayBuffer(3108864))
console.log(b)
console.log(lll)


function putImageData(ctx,imageData,dx,dy,dirtyX,dirtyY,dirtyWidth,dirtyHeight) {
    const data = imageData.data;
    const height = imageData.height;
    const width = imageData.width;
    dirtyX = dirtyX || 0;
    dirtyY = dirtyY || 0;
    dirtyWidth = dirtyWidth !== undefined ? dirtyWidth : width;
    dirtyHeight = dirtyHeight !== undefined ? dirtyHeight : height;
    const limitBottom = dirtyY + dirtyHeight;
    const limitRight = dirtyX + dirtyWidth;
    for (let y = dirtyY; y < limitBottom; y++) {
      for (let x = dirtyX; x < limitRight; x++) {
        const pos = y * width + x;
        ctx.fillStyle = `rgba(${data[pos * 4 + 0]}, ${data[pos * 4 + 1]}, ${
          data[pos * 4 + 2]
        }, ${data[pos * 4 + 3] / 255})`;
        ctx.fillRect(x + dx, y + dy, 1, 1);
      }
    }
  }