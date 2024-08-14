class Animal {

    constructor(name) {
      this.speed = 5;
      this.name = name;
    }
    lole() {
        console.log("lole")
    }
    // ...
}
  
  class Rabbit extends Animal {
  
    constructor(name,earLength) {
      super();
      this.name = name
      this.earLength = earLength;
      this.speed = 300
    }
  
    // ...
}

// now fine
let an = new Animal("hello")
let rabbit = new Rabbit("White Rabbit", 10);
console.log(rabbit.name); // White Rabbit
console.log(rabbit.earLength); // 10
//console.log(rabbit.speed)
//console.log(an.name)
//console.log(an.speed)
//an.lole()
//rabbit.lole()