function handleAttri(d, at) {
  for (let i = 0; i < at.length; i += 2) {
    let k = at[i];
    let v = at[i + 1];

    d.setAttribute(k, v);
  }
}

function createChild(parent, child) {
  if (child instanceof Node) {
    parent.appendChild(child);
  } else {
    parent.appendChild(document.createTextNode(child));
  }
}

function removeChild(parent, child) {
  parent.removeChild(child);
}

function getParent(element, num = 2) {
  let ret = element;

  for (let i = 0; i < num; i++) {
    ret = ret.parentElement;
  }

  return ret;
}

function createStructure(st) {
  let parent = document.createElement(st.tag);

  if ("attri" in st) {
    handleAttri(parent, st.attri);
  }

  if ("children" in st) {
    if (Array.isArray(st.children)) {
      for (const child of st.children) {
        createChild(parent, child);
      }
    } else {
      createChild(parent, st.children);
    }
  }

  return parent;
}

function redirect(url) {
  window.history.pushState(null, null, url);
}

let state = {
  todo: [],
  selectedBtn: 1,
  page: "/#",
  filter: "",
  hideInfo: false,
  hideClear: false,
  allChecked: false,
};

function addEvent(eventType, element, callback) {
  element.addEventListener(eventType, callback);
}

function removeEvent(eventType, element, callback) {
  element.removeEventListener(eventType, callback);
}

function saveState(currentState) {
  localStorage.setItem("state", JSON.stringify(currentState));
}

function loadState() {
  let data = localStorage.getItem("state");

  if (data) {
    state = JSON.parse(data);
  }
}
