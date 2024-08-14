function getCookie(cname) {
  let name = cname + "=";
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(";");
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == " ") {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

const routes = {
  "/": home,
  "/loggedin": loggedin,
  "/createpost": create,
  "/post" : ""
};

const routesFuncs = {
  "/":resetCredentials,
  "/loggedin":loggedinFetch,
  "/post": getSinglePostJson,
}

function resetCredentials() {
  let name = "session_token";
  document.cookie = name + "=; Max-Age=-99999999;";
  if (socket !== undefined) socket = socket.close();
}

const rootDiv = document.getElementById("root");
rootDiv.innerHTML = routes[window.location.pathname];
document.body.append(rootDiv);

const onNavigate = (pathname) => {
  window.history.pushState({page: pathname}, pathname, window.location.origin + pathname);
  console.log(window.location.pathname)
  let popStateEvent = new PopStateEvent('popstate', { state: "lol" });
  dispatchEvent(popStateEvent);
};

window.onpopstate = () => {
  rootDiv.innerHTML = routes[window.location.pathname];
  if (window.location.pathname in routesFuncs) routesFuncs[window.location.pathname]();
  
};

onNavigate("/")