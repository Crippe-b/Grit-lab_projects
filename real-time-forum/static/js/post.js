function newBigPost(
  title,
  author,
  timestamp,
  categories,
  content,
  comments,
  commentform
) {
  return `
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

<div class="big-post">
<h2 class="post-title" id="title">${title}</h2>
<div class="single-post">

<div class="singlePost-info">
    <div id="author">
    <label class="post_author">Author: </label>
        <p class="author-link">${author}</p>
    </div>
    <label class="posted">Posted: </label>
    <p class=timestamp>${timestamp}</p>

    <div class="post-content" id="postId">
    <p class="postContent">${content}</p>
</div>
<label class="posted">Categories:</label>
<p class="category" id="categoryId">${categories}</p>
</div>

</div>
<section class="comments">
        ${this.newComments(comments)}
</section>
<section class="comment-section">
<p class="errorMsg"></p>
${commentform}
</section>
`;
}

let commentForm = `<div class="comment-form">
<p class="comment-p">Write a comment:</p>
<form class="text-input" id="comment-form" name="comment-form" onsubmit="return sendCommentData(this);">
    <input name="comment" id="comment" type="text" maxlength="500" cols="70" placeholder="Write a comment">
    <input type="button" name="submit" class="post-btn" from="comment-form" value="Submit" onclick="return sendCommentData(this.form);">
</form>
</div>
</div>
`;

function newComments(comments) {
  let res = "";
  for (let i = 0; i < comments.length; i++) {
    let el = document.createElement("div");
    el.innerHTML = `<div class="comment-content">
    <p class="comment">${comments[i].author}</p>
    <p class="comment">${comments[i].created}</p>
    <p class="comment dummy"></p>
</div>`;
    el.getElementsByClassName("dummy")[0].innerText = comments[i].commenttext;
    //   res += `<div>
    //       <div class="comment-content">
    //           <p class="comment">${comments[i].author}</p>
    //           <p class="comment">${comments[i].created}</p>
    //           <p class="comment">${comments[i].commenttext}</p>
    //       </div>
    // </div>`;
    console.log(el.innerHTML);
    res += `<div>
  ${el.innerHTML}
  </div`;
  }
  return res;
}

function sendCommentData(form) {
  let data = new FormData(form);

  const r = /\S/;

  if (r.test(data.get("comment")) === false) {
    alert("empty comment :(");
    return false;
  }

  params = new URLSearchParams(window.location.search);
  data.set("id", params.get("id"));
  fetch(window.location.pathname, {
    method: "POST",
    body: data,
  })
    .then((json) => json.json())
    .then((response) => {
      if (response.status === "Success") {
        //return false
        onNavigate(window.location.pathname + window.location.search);
      } else if (response.status === "Failed") {
        alert("you fucked up, idiot");
        //return false
      }
    });
  return false;
}
