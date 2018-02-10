import {html, LitElement} from '../node_modules/@polymer/lit-element/lit-element.js';

export class SpenderTimer extends LitElement {
    constructor() {
        super();
        this.time = new Date(0);
    }

    static get properties() {
        return {
            time: Date
        }
    }

    blink() {
        this.shadowRoot.getElementById('coin').animate(
            [
                {transform: "scale(1)"},
                {transform: "scale(1.1)"},
                {transform: "scale(1)"},
            ],
            {
                duration: 300,
            }
        )
    }

    render({time}) {
        return html`
            <style>  
                :host {
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                }
                
                #coin {
                  width: 68px;
                  height: 68px;
                  background-position: 0 0;
                  transition: background-position 300ms steps(18,end);
                }
                
                #coin.piggy {
                  background-position: -1224px 0;
                }

                #time {
                    display: block;
                  text-align: center;
                }
            </style>
            
            <div id="coin" class$="${this.time.getTime() > 0 ? "piggy" : ""}" style="background-image: url(spender/sprite_60fps.svg)"></div>
            <div id="time">${this._formatTime(time)}</div>`
    }

    _formatTime(time) {
        const options = {minute: '2-digit', second: '2-digit'};
        return time.toLocaleTimeString('de-DE', options);
    }
}

customElements.define('spender-timer', SpenderTimer);