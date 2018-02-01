import {Element} from '../node_modules/@polymer/polymer/polymer-element.js';

export class SpenderToggle extends Element { // Define a string template instead of a `<template>` element.
    static get template() {
        return `
<style>
    #toggle {
  flex: 1;
  position: relative;
  background-color: var(--spender-toggle-color, grey);
  color: #f1f1f1;
  font-size: 32px;
  font-weight: bold;
  border: none;
  outline: none;
  user-select: none;
  margin: 0 8px;
  border-radius: 2px;
  box-shadow: 0 3px 6px rgba(0,0,0,0.16), 0 3px 6px rgba(0,0,0,0.23);
}


#toggle:active {
  filter: brightness(1.5);
}

#toggle:disabled {
  background-color: #3f3f3f;
  color: #e5e5e5;
}
</style>

<button id="toggle" disabled$="[[disabled]]">[[title]]</button>
        
`
    }

    constructor() {
        super();
        this.time = new Date(0);
    }

    // properties, observers, etc. are identical to 2.x
    static get properties() {
        return {
            disabled: {
                Type: Boolean
            },
            title: {
                Type: String
            }
        }
    }

}

customElements.define('spender-toggle', SpenderToggle);