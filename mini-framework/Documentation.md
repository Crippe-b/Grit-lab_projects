## MINI-FRAMEWORK PROJECT

## Explanation

The Mini-Framework project works by providing a set of functions that simplify the process of creating and manipulating elements and events in JavaScript, making it easier to build dynamic web applications.

This documentation file provides an explanation of the features along with code examples and explanations on :

- How to create an element.
- Create an event.
- Add attributes to an element.
- Nest elements.

## Features

The Mini-Framework project includes the following features:

- `handleAttri(d,at)`: This function handles attributes of an element by setting the attribute key-value pairs using `setAttribute()`.
- `removeChild(parent,child)`: This function removes a child element from a parent element.
- `getParent(element,num = 2)`: This function returns the parent element of a given element.
- `createStructure(st)`: This function creates an element structure based on the given tag, attributes, and children.
- `redirect(url)`: This function redirects the user to a new URL.
- `state`: This object stores the state of the application.

## Code Examples

#### Creating an Element

To create an element using the Mini-Framework project, use the `createStructure()` function. It takes an object with the following properties:
`tag`: The tag name of the element to be created.
`attri (optional)`: An array of attribute key-value pairs for the element.
`children (optional)`: An array of child elements for the element.

Here is an example of how to create a div element with a class of "example":

```js
const exampleDiv = createStructure({
  tag: "div",
  attri: ["class", "example"],
  children: [node / text],
});
```

#### Creating an Event

To create an event using the Mini-Framework project, use the `addEvent()` function.
`addEvent(eventType, element, callback)`: Attaches an event listener to the specified element using the `addEventListener()` method.

```js
function addEvent(eventType, element, callback) {
  element.addEventListener(eventType, callback);
}
```

This function takes three arguments:
`eventType`: The type of event to listen for (e.g. "click").
`element`: The element to attach the event listener to.
`callback`: The function to be called when the event is triggered.

Example usage:

```js
addEvent("keypress", this.input, this.handleNewTodo);
addEvent("click", this.clear, this.handleClearChecked);
addEvent("change", this.checkAllBox, this.handleCheckAll);
```

#### Removing an Event

`removeEvent(eventType, element, callback)`: This function removes an event listener from the specified element using the `removeEventListener()` method.

```js
function removeEvent(eventType, element, callback) {
  element.removeEventListener(eventType, callback);
}
```

It takes the same three parameters as the `addEvent()` function:
`eventType`: The type of event that was previously added.
`element`: The element from which the event listener will be removed.
`callback`: The function that was previously attached as the event listener.

Example Usage:

```js
removeEvent("blur", targetElement, this.handleChangeTodoText);
removeEvent("keypress", targetElement, this.handleChangeTodoText);
```

#### Adding Attributes to an Element & Nesting Elements

To add attributes to an element using the Mini-Framework project, use the `createStructure()` function with an attri property that is an array of attribute key-value pairs.

To nest elements using the Mini-Framework project, use the `createStructure()` function to create child elements and the `createChild()` function to append them to a parent element.

Here is an example of how to create a div element with two child p elements:

```js
const li = createStructure({
  tag: "li",
  attri: ["data-id", todo.id, "class", cl],
  children: [
    createStructure({
      tag: "div",
      attri: ["class", "view"],
    }),
  ],
});
const body = document.querySelector("body");
createChild(body, parentDiv);
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- []
- []
- []
- []
- []
