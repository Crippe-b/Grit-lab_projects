function validateRegistration() {
  const r = /^[a-z0-9]+$/i;
  let a = /^[0-9]+$/g;
  let e = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g;
  let form = document.forms["registration-form"];

  if (r.test(form["username"].value) === false) {
    alert("Username can not be empty");
    return
  }
  if (r.test(form["gender"].value) === false) {
    alert("Gender can't be empty");
    return
  }
  if (r.test(form["firstname"].value) === false) {
    alert("Firstname can't be empty");
    return
  }
  if (r.test(form["lastname"].value) === false) {
    alert("Lastname can't be empty");
    return
  }
  if (a.test(form["age"].value) === false) {
    alert("Age can't be empty and must be a number");
    return
  }
  if (e.test(form["email"].value) === false) {
    alert("Please enter a valid email adress");
    return
  }
  if (r.test(form["password"].value) === false) {
    alert("Password can not be empty");
    return
  }
  sendLoginData(form, "http://localhost:8080/register")
}

function gotoLogout() {
  onNavigate("/");
}
