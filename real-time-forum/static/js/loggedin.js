let loggedin = `
<div class="headbar">
<div class="indentifybar">
<img class="logo" src="../css/text-bubbles.png">
<h1><button class="ident" onclick="onNavigate('/loggedin')">Heads</button></h1>

<button onclick="onNavigate('/createpost')" class="first-button">Create Post</button>
</div>
 <div class="loginbar">
 <button onclick="gotoLogout()" class="first-button">Sign out</button>
 </div>
 </div>
 <select class="category-filter" onchange="filterCategory(this.value)">
 <option value="Show_all">Show all</option>
 <option value="funny">Funny</option>
 <option value="sport">Sport</option>
 <option value="off-topic">Off Topic</option>
 <option value="music">Music</option>
 <option value="game">Game</option>
 <option value="news">News</option>
 <option value="korn">Korn</option>
 </select>
 `;

function newSmallPost(title, author, timestamp, categories, postid) {
  let el = document.createElement("div");
  el.classList.add("post-info");
  el.classList.add.apply(el.classList, categories);
  el.innerHTML = `<button class="post-dir" onclick="onNavigate('/post?id=${postid}')"></button>
  <p class="author-link" id="author">${author}</p>
  <p class="timestamp">${timestamp}</p>
  <h5 class="category">${categories}</h5>`;
  el.getElementsByClassName("post-dir")[0].innerText = title;
  return el;
}

function newUserList(online, offline) {
  let el = document.createElement("section");
  el.classList.add("username-section");
  el.innerHTML = `<div class="usernames">
  <ul class="username-list" id="user-list">
      <p class="online">Online</p>
      ${online}
  </ul>
  <ul class="username-list">
      <p class="offline">Offline</p>
      ${offline}
  </ul>
</div>`;
  return el;
}

function onlineoroffline(users) {
  let offline = "";
  let online = "";

  users.sort((a, b) => {
    if (a.hassentmsg === true && b.hassentmsg === false) return -1;
    if (a.hassentmsg === false && b.hassentmsg === true) return 1;
    if (a.hassentmsg === true && b.hassentmsg === true) {
      let c = Date.parse(a.lastmsg);
      let d = Date.parse(b.lastmsg);
      if (c === NaN || d === NaN) console.log("fuck");
      return d - c;
    }
    if (a.username < b.username) {
      return -1;
    } else if (a.username > b.username) {
      return 1;
    } else {
      return 0;
    }
  });

  for (let i = 0; i < users.length; i++) {
    if (users[i].loggedin === true) {
      if (users[i].you === true)
        online += `<li class="names" id="user-${users[i].username}"><button class="chat-link you">${users[i].username}</button></li>`;
      else
        online += `<li class="names" id="user-${users[i].username}"><button onclick="getMessages('${users[i].username}')" class="chat-link">${users[i].username}</button></li>`;
    } else {
      offline += `<li class="names" id="user-${users[i].username}"><button onclick="getMessages('${users[i].username}')" class="chat-link">${users[i].username}</button></li>`;
    }
  }
  return [online, offline];
}

function closeMsgBox() {
  let el = document.querySelectorAll(".chat-section");
  el.forEach((chat) => {
    chat.remove();
  });
}

function msgBox(messages, name) {
  let temp = document.createElement("section");
  temp.classList.add("chat-section");
  temp.id = name;
  let msgBoxEl = document.createElement("div");
  msgBoxEl.classList.add("message-box");
  msgBoxEl.append(...messages);
  temp.innerHTML += `<button type="close" id="message-close" class="send-msg" onclick="closeMsgBox()" value="close">&times</button>`;
  temp.appendChild(msgBoxEl);
  temp.innerHTML += `
       <form class="message" onsubmit="return sendMessage()">
       <input name="message" id="messafge" class="msg-input" maxlength="400" placeholder="Write a msg" type="text">
       </form>
       `;
  return temp;
}

function gatherMsg(msg, otherUsr) {
  let res = [];
  for (let i = 0; i < msg.length; i++) {
    if (msg[i].sender === otherUsr) {
      let temp = document.createElement("div");
      temp.innerHTML = `<div class="container darker">
      <h4 class="username" id="nickname2">${msg[i].sender}</h4>
      <p class="text-content" id="textId2"></p>
      <span class="time-left">${msg[i].timestamp}</span>
      </div>`;
      temp.firstChild.getElementsByTagName("p")[0].textContent = msg[i].message;
      res.push(temp.firstChild);
    } else {
      let temp = document.createElement("div");
      temp.innerHTML = `<div class="container">
        <h4 class="username" id="nickname1">${msg[i].sender}</h4>
        <p class="text-content" id="textId1"></p>
        <span class="time-right">${msg[i].timestamp}</span>
        </div>`;
      temp.firstChild.getElementsByTagName("p")[0].textContent = msg[i].message;
      res.push(temp.firstChild);
    }
  }
  return res;
}

function filterCategory(selected) {
  let allPosts = document.querySelectorAll(".post-info");
  console.log(allPosts);
  allPosts.forEach((pst) => {
    if (selected === "Show_all") pst.style.display = "block";
    else if (pst.classList.contains(selected) === true)
      pst.style.display = "block";
    else pst.style.display = "none";
  });
}

function loggedinFetch() {
  getJson();
  getUsers();
}
