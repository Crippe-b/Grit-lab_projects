let create = `
<div class="headbar">
<div class="indentifybar">
<img class="logo" src="../css/text-bubbles.png">
<h1><button class="ident" onclick="onNavigate('/loggedin')">Heads</button></h1>
</div>
 <div class="loginbar">
 <button onclick="gotoLogout()" class="first-button">Sign out</button>
 </div>
 </div>

 <section class="createPost-section">
      <form method="post" name="create-post" id="create-post">
        <div class="title-div">
          <label class="title">Title:</label>
          <div class="input-group">
            <input id="title-input" type="text" name="title" cols="45" maxlength="35" required placeholder="Title"/>
          </div>
        </div>
        <div class="content-div">
          <label id="content-label" for="content-lable">Content:</label>
          <textarea class="create-content" required type="text" name="content" cols="70" maxlength="500" placeholder="Post content"></textarea>
        </div>
        <label id="category-label" for="category"
          >Choose a category/categories</label
        >
        <div class="categories">
          <input class="check-cat" name="category" type="checkbox" value="sport"/>
          <label class="category-box"></label>Sport<br />
          <input class="check-cat" name="category" type="checkbox" value="music"/>
          <label class="category-box"></label>Music<br />
          <input class="check-cat" name="category" type="checkbox" value="news"/>
          <label class="category-box"></label>News<br />
          <input class="check-cat" name="category" type="checkbox" value="off-topic"/>
          <label class="category-box"></label>Off-topic<br />
          <input class="check-cat" name="category" type="checkbox" value="game"/>
          <label class="category-box"></label>Game<br />
          <input class="check-cat" name="category" type="checkbox" value="funny"/>
          <label class="category-box"></label>Funny<br />
          <input class="check-cat" name="category" type="checkbox" value="korn"/>
          <label class="category-box"></label>Korn<br />
        </div>
        <input id="createPost" type="button" from="create-post" value="Create" onclick="validatePost(this.form, 'http://localhost:8080/createpost');"/>
      </form>
    </section>
`;

function sendPostData(form, path) {
  let data = new FormData(form);
  fetch(path, {
    method: "POST",
    body: data,
  })
    .then((json) => json.json())
    .then((response) => {
      if (response.status === "Success") {
        onNavigate("/loggedin");
      } else if (response.status === "Failed")
        location.assign("http://localhost:8080/");
      else if (response.status === "No_Content") alert("missing content");
      else if (response.status === "No_Title") alert("missing title");
      else if (response.status === "No_Category") alert("missing category");
    });
}

function validatePost() {
  const r = /\s+[^a-zA-Z0-9]/g;
  let form = document.forms["create-post"];

  if (r.test(form["title"].value) === true) {
    alert("title can not be empty");
    return;
  }
  if (r.test(form["content"].value) === true) {
    alert("content can not be empty");
    return;
  }
  sendPostData(form, "http://localhost:8080/createpost");
}
