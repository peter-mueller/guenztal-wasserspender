/**
@license
Copyright (c) 2018 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
import { html, css } from "../../node_modules/lit-element/lit-element.js";
import { PageViewElement } from './page-view-element.js';
import "./water-toggle.js";
import "../../node_modules/@polymer/paper-toast/paper-toast.js";
import "./object-table.js";
import { WasserSpenderClient } from "../client.js";

const url = () => {
  if (location.host == "localhost:8081") {
    return "http://localhost:8080";
  }

  return "";
};

const refreshDuration = 500;
const client = new WasserSpenderClient(url());

class ViewControl extends PageViewElement {
  static get properties() {
    return {
      unlocked: {
        type: Boolean
      },
      // Timer number in Remaining Seconds
      timer: {
        type: Number
      },
      // Current state of valves
      valves: {
        type: Object
      }
    };
  }

  constructor() {
    super();
    this.valves = {
      Warm: {
        Opened: false
      },
      Osmose: {
        Opened: false
      },
      Cold: {
        Opened: false
      }
    };
    window.setInterval(this.updateInfo.bind(this), refreshDuration);
  }

  static get styles() {
    return [css`
      #control {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        padding: 8px;
      }

      water-toggle {
        padding: 8px;
      }

      json-table {
        width: 512px;
        margin: auto;
      }

    `];
  }

  updateInfo() {
    client.findTimer().then(timer => {
      if (timer < 0) {
        this.unlocked = false;
        this.timer = 0;
        return;
      }

      this.unlocked = true;
      this.timer = timer;
    });
    client.findAllValves().then(valves => {
      this.valves = valves;
    });
  }

  toggle(e) {
    const type = e.detail.type;

    for (var v in this.valves) {
      var valve = this.valves[v];

      if (valve.Name == type) {
        client.setValve(type, !valve.Opened);
      }
    }
  }

  render() {
    return html`
      <section id="control" @toggle=${this.toggle}>
        <water-toggle imagesrc="/images/fire.jpg" type="warm" ?active=${this.valves.Warm.Opened} ?locked=${!this.unlocked} label="Warmes Wasser"  @toggle=${this.toggle}></water-toggle>
        <water-toggle imagesrc="/images/waterfall.jpg"  type="cold" ?active=${this.valves.Cold.Opened} label="Kaltes Wasser"  @toggle=${this.toggle}></water-toggle>
        <water-toggle imagesrc="/images/osmose.jpg" ?active=${this.valves.Osmose.Opened} ?locked=${!this.unlocked} type="osmose" label="Osmose Wasser"  @toggle=${this.toggle}></water-toggle>

      </section>

      <paper-toast ?opened=${!this.unlocked} text="Für Warmes und Osmosewasser Geld einwerfen." 
        duration = "0" 
        ></paper-toast >
      <paper-toast ?opened=${this.unlocked} text=${this._helpText(this.timer)} 
      duration = "0" 
       ></paper-toast >
    `;
  }

  _helpText(timer) {
    return `Noch ${Math.floor(timer / 1000 / 60)} Minute(n) und ${Math.floor(timer / 1000 % 60)} Sekunde(n) übrig.`;
  }

}

window.customElements.define('view-control', ViewControl);