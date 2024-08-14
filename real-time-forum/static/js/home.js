let home = `<p class="front-p">
Register or Sign in to see posts and chat with fellow users
</p>
<section class="front-page">
<div class="register-window" id="register">
  <form action="/register" method="post" class="content" id="registration-form" onsubmit="return validateRegistration()"name="registration-form">
    <h1>Registeration</h1>
    <input type="text" value="" placeholder="First Name" class="validate" name="firstname" required >
    <input type="text" value="" placeholder="Last Name" class="validate" name="lastname" required >
    <input type="number" value="" placeholder="Age" class="validate" name="age" required >
    <input type="text" value="" placeholder="Gender" class="validate" name="gender" required >
    <input type="text" value="" placeholder="Email" class="validate" name="email" required >
    <input type="text" value="" placeholder="Username" class="validate" name="username" required >
    <input type="password" value="" placeholder="Password" class="validate" name="password" required >
    <input class="second-button" id="register" type="button" from="registration-form" value="Submit" onclick="validateRegistration();" >
  </form>
</div>
<p class="line-divider"></p>
<div class="popup" id="popup-1">
  <form class="content" id="login-form" name="login-form">
    <h1>Sign in</h1>
    <input type="text" value="" placeholder="Email/Username" class="validate" name="name" />
    <input type="password" value="" placeholder="Password" class="validate" name="password" />
    <input type="button" name="submit" class="second-button" id="popup-1" from="login-form" value="Submit" onclick="sendLoginData(this.form, 'http://localhost:8080/loggedin');" />
  </form>
</div>
</section>
`;

function sendLoginData(form, path) {
  let data = new FormData(form);
  fetch(path, {
    method: "POST",
    body: data,
  })
    .then((json) => json.json())
    .then((response) => {
      if (response.status === "Success") {
        onNavigate("/loggedin");
        openChatSocket();
      } else alert("you suck :)");
    });
}
