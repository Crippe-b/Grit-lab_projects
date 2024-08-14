function getJson() {
  fetch("http://localhost:8080/hello", {
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

      let el = document.createElement("div");
      el.classList.add("post-section");
      for (let i = 0; i < json.posts.length; i++) {
        console.log(json.posts[i]);
        let author = json.posts[i].author;
        let timestamp = json.posts[i].created;
        let title = json.posts[i].title;
        let categories = json.posts[i].categories;
        let postid = json.posts[i].postid;
        el.appendChild(
          newSmallPost(title, author, timestamp, categories, postid)
        );
      }
      rootDiv.appendChild(el);
    });
}

function getSinglePostJson() {
  fetch(`http://localhost:8080/goodbye${window.location.search}`, {
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
      let author = json.author;
      let timestamp = json.created;
      let title = json.title;
      let categories = json.categories;
      let comments = json.comments;
      let content = json.body;
      let el = newBigPost(
        title,
        author,
        timestamp,
        categories,
        content,
        comments,
        commentForm
      );
      rootDiv.innerHTML = el;
      rootDiv.getElementsByClassName("post-title")[0].innerText = title;
      rootDiv.getElementsByClassName("postContent")[0].innerText = content;
    });
}
