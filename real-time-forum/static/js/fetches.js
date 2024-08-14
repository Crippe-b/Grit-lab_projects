function getUsers() {
  fetch("http://localhost:8080/losers", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  })
    .then((response) => response.json())
    .then((json) => {
      if (json.status === "Failed") {
        location.assign("http://localhost:8080/");
        return;
      }
      let el = document.querySelector(".username-section");
      if (el !== null) el.remove();
      let usrs = onlineoroffline(json);
      rootDiv.appendChild(newUserList(usrs[0], usrs[1]));
    });
}

// Throttling Function
// const throttleFunction=(func, delay)=>{

//   // Previously called time of the function
//   let prev = 0;
//   return (...args) => {
//     // Current called time of the function
//     let now = new Date().getTime();

//     // Logging the difference between previously
//     // called and current called timings
//     console.log(now-prev, delay);

//     // If difference is greater than delay call
//     // the function again.
//     if(now - prev> delay){
//       prev = now;

//       // "..." is the spread operator here
//       // returning the function with the
//       // array of arguments
//       return func(...args);
//     }
//   }
// }

function throttle(cb, delay) {
  let wait = false;

  return (...args) => {
    if (wait) {
      return;
    }

    cb(...args);
    wait = true;
    setTimeout(() => {
      wait = false;
    }, delay);
  };
}

function throttleforScroll(cb, delay) {
  let wait = false;

  return (...args) => {
    if (wait) {
      return;
    }

    cb(...args);
    wait = true;
    setTimeout(() => {
      cb(...args)
      wait = false;
    }, delay);
  };
}
let throthle = throttleforScroll(function (event) {
  
  let el = document.querySelector(".message-box");
  if (el === null || el.scrollTop !== 0) return;
  let newMessages = l2CountIdiot();
  let oldScrollSize = el.scrollHeight
  el.prepend(...newMessages);
  el.scrollTop = (el.scrollHeight - oldScrollSize)
}, 500);

let trottle = throttle(function (event, user) {
  socket.send(
    JSON.stringify({
      recipient: user,
      content: "typing_update",
      type: "typing_update",
    })
  );
}, 500);

let allMessages;
let msgCounter = 0;

function l2CountIdiot() {
  let res = allMessages.slice(
    -msgCounter - 10,
    allMessages.length - msgCounter
  );
  msgCounter += 10;
  return res;
}

function getMessages(user) {
  fetch("http://localhost:8080/nodobylikesu", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Body: user,
    },
  })
    .then((response) => response.json())
    .then((json) => {
      if (json.status === "Failed") {
        location.assign("http://localhost:8080/");
        return;
      }
      let el = document.querySelector(".chat-section");
      if (el !== null) el.remove();

      allMessages = gatherMsg(json, user);
      msgCounter = 0;
      let messages = l2CountIdiot();
      console.log(messages);

      rootDiv.appendChild(msgBox(messages, user));
      let el2 = document.getElementById("user-" + user);
      el2.classList.remove("new_Message");
      document
        .querySelector(".message")
        .addEventListener("input", (event) => trottle(event, user));
      let el3 = document.querySelector(".message-box")
      el3.scrollTop = el3.scrollHeight
      el3.addEventListener("scroll", (event) => throthle(event))
    });
}
