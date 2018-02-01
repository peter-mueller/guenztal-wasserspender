import {Element} from '../node_modules/@polymer/polymer/polymer-element.js';

export class SpenderTimer extends Element {

    // Define a string template instead of a `<template>` element.
    static get template() {
        return ` 
<style>
    #time {
        display: block;
      text-align: center;
    }
</style>

<div id="time">[[_formatTime(time)]]</div>`
    }

    constructor() {
        super();
        this.time = new Date(0);
    }

    // properties, observers, etc. are identical to 2.x
    static get properties() {
        return {
            time: {
                Type: Date
            }
        }
    }

    _formatTime(time) {
        const options = {minute: '2-digit', second: '2-digit'};
        return time.toLocaleTimeString('de-DE', options);
    }
}

customElements.define('spender-timer', SpenderTimer);